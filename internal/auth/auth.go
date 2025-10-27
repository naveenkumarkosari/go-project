package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	fmt.Println(val, "==val==")
	if val == " " {
		return "", errors.New("authorization is not provided")
	}
	vals := strings.Split(val, " ")
	if vals[0] != "ApiKey" {
		return "", errors.New("authorization required")
	}
	return vals[1], nil
}
