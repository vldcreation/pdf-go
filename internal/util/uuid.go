package util

import (
	"github.com/google/uuid"
)

func GenerateRandoomUUID() string {
	uuid := uuid.New()

	return uuid.String()
}
