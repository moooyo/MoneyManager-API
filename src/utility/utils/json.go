package utils

import (
	"encoding/json"
	"log"
)

func PrettyPrintJson(args map[string]string) string {
	marshaled, err := json.MarshalIndent(args, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	return string(marshaled)
}
