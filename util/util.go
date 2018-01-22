package util

import (
	uuid "github.com/satori/go.uuid"
)

//NewUUID return a new UUIDv4 string
func NewUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
