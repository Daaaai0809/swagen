package constant

import (
	"fmt"
	"strings"
)

const (
	API_KEY_AUTH         = "ApiKey"
	BASIC_AUTH           = "BasicAuth"
	BEARER_AUTH          = "Bearer"
	OAUTH2_AUTH          = "OAuth2"
	OPEN_ID_CONNECT_AUTH = "OpenId"

	API_KEY_AUTH_LOWER         = "apikey"
	BASIC_AUTH_LOWER           = "basicauth"
	BEARER_AUTH_LOWER          = "bearer"
	OAUTH2_AUTH_LOWER          = "oauth2"
	OPEN_ID_CONNECT_AUTH_LOWER = "openid"
)

func GetCamelCaseSecurityType(securityType string) string {
	switch strings.ToLower(securityType) {
	case BASIC_AUTH_LOWER:
		return BASIC_AUTH
	case BEARER_AUTH_LOWER:
		return BEARER_AUTH
	case OAUTH2_AUTH_LOWER:
		return OAUTH2_AUTH
	case OPEN_ID_CONNECT_AUTH_LOWER:
		return OPEN_ID_CONNECT_AUTH
	default:
		if strings.Contains(securityType, API_KEY_AUTH_LOWER) {
			splitted := strings.Split(securityType, API_KEY_AUTH_LOWER)
			return fmt.Sprintf("%s%s", API_KEY_AUTH, splitted[1])
		}
	}

	return ""
}

var SecurityTypes = []string{
	API_KEY_AUTH,
	BASIC_AUTH,
	BEARER_AUTH,
	OAUTH2_AUTH,
	OPEN_ID_CONNECT_AUTH,
}
