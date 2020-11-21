package random

import (
	gu "github.com/google/uuid"
)

// GenUUIDV4 means
func GenUUIDV4() (uuid string, err error) {
	uuidByte, err := gu.NewRandom()
	if err != nil {
		return "", err
	}
	uuid = uuidByte.String()
	return uuid, nil
}
