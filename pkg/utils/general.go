package utils

import "github.com/lithammer/shortuuid/v3"

func GenerateUUID() string {
	return "app-" + shortuuid.New()
}
