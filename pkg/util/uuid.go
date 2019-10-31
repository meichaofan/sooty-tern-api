package util

import (
	"github.com/google/uuid"
)

// MustUUID 创建UUID，如果发生错误则抛出panic
func MustUUID() string {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}

// NewUUID 创建UUID
func NewUUID() (string, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// NewRecordId
func NewRecordId() string {
	recordId, err := NewUUID()
	if err != nil {
		return ""
	}
	return recordId[:8]
}

// NewSalt
func NewSalt() string {
	recordId, err := NewUUID()
	if err != nil {
		return ""
	}
	return recordId[:6]
}