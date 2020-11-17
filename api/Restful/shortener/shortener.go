package apiv1shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-infrastructure/business/shortener"
)

// Handler means
func Handler(shortenerBiz shortener.Business) *ShortenerAPI {
	return &ShortenerAPI{
		shortenerBiz: shortenerBiz,
	}
}

// ShortenerAPI means
type ShortenerAPI struct {
	shortenerBiz shortener.Business
}

// GetShortenerURL means
func (impl *ShortenerAPI) GetShortenerURL(ctx *gin.Context) {
	shortenerID := ctx.Param("shortenerID")

	if shortenerID == "" {
		ctx.JSON(http.StatusUnauthorized, GetSuccessResponse("request parameter is failure"))
	}

	shortenerURL, _ := impl.shortenerBiz.GetShortenerURL(shortenerID)

	ctx.JSON(http.StatusOK, GetSuccessResponse(gin.H{"shortenerURL": shortenerURL}))
}

// GetSuccessResponse means
func GetSuccessResponse(data interface{}) map[string]interface{} {
	return gin.H{
		"metadata": gin.H{
			"status": "0000",
			"desc":   "success",
		},
		"data": data,
	}
}
