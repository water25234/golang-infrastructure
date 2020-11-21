package logger

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/common/random"
	"github.com/water25234/golang-infrastructure/common/time"
	log "github.com/water25234/golang-infrastructure/core/logger"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type responseBodyStruct struct {
	Metadata interface{} `json:"metadata"`
	Data     interface{} `json:"data"`
}

// Logger means
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start execute time
		startTime := time.GetCurrentMilliUnix()

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		requestUUID := c.MustGet("requestUUID").(string)
		if len(requestUUID) == 0 {
			uuid, err := random.GenUUIDV4()
			if err != nil {
				logrus.WithField("err", err).Panic("middleware logger gen uuid failure")
				c.Abort()
			}
			requestUUID = uuid
		}

		// put here for new.
		log.SetLoggerField(requestUUID)

		c.Next()

		responseBody := bodyLogWriter.body.String()

		response := make(map[string]interface{})

		if responseBody != "" {
			res := responseBodyStruct{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				response["json"] = res
			}
		}

		// end execute time
		endTime := time.GetCurrentMilliUnix()

		responseMap := make(map[string]interface{})

		responseMap["request_time"] = startTime
		responseMap["request_method"] = c.Request.Method
		responseMap["request_uri"] = c.Request.RequestURI
		responseMap["request_proto"] = c.Request.Proto
		responseMap["request_ua"] = c.Request.UserAgent()
		responseMap["request_referer"] = c.Request.Referer()
		responseMap["request_post_data"] = c.Request.PostForm.Encode()
		responseMap["request_client_ip"] = c.ClientIP()
		responseMap["response_time"] = endTime
		responseMap["response"] = response
		responseMap["speed_time"] = fmt.Sprintf("%vms", endTime-startTime)

		log.Record.WithFields(responseMap).Info("execute api done")
	}
}
