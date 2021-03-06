package provider

import (
	"encoding/base64"
	"os"
)

func getConfigString(configFile map[string]interface{}, key string, envName string) string {
	// load a value from config file
	if v, ok := configFile[key]; ok && v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}

	// load from the environment if there is no value in config file
	return os.Getenv(envName)
}

// base64Decode decodes the Base64url encoded string
// steel from https://github.com/golang/oauth2/blob/master/jws/jws.go
func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	switch len(s) % 4 {
	case 1:
		s += "==="
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}
