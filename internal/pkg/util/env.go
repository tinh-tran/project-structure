package util


import (
	"fmt"
	"os"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func VerifyListEnv(keys []string) {
	isPanic := false
	stringEnvs := ""
	for _, s := range keys {
		if GetEnv(s, "") == "" {
			stringEnvs = stringEnvs + " " + s + " "
			isPanic = true
		}
	}
	if isPanic {
		panic(fmt.Sprintf("Please input env  \n %v", stringEnvs))
	}
}
