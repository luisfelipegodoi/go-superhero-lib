package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// InterfaceToString Convert Interce to string and return
func InterfaceToString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

// CheckError - Verify if exists any error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//GetKeysFromMap return all map keys from a map
func GetKeysFromMap(object map[string]string) []string {
	var result []string
	for key := range object {
		result = append(result, key)
	}
	return result
}

//ArrayToString convert a array to string.
func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

//NullableString validate null for string
func NullableString(value string) sql.NullString {
	if len(value) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{String: value, Valid: true}
}

//NullableInt validate null for int
func NullableInt(value int) sql.NullInt64 {
	if value == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: int64(value), Valid: true}
}

//NullableFloat validate null for float
func NullableFloat(value float64) sql.NullFloat64 {
	if value == 0 {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{Float64: value, Valid: true}
}

//ValueInArray validate the existence of a value in Array.
func ValueInArray(array []interface{}, value interface{}) bool {

	for _, val := range array {
		if val == value {
			return true
		}
	}
	return false

}

//GetIntFromEnv get a int value from a Environment Variable, if error or environment variable not exist, the default value will be returned.
func GetIntFromEnv(key string, defaultValue int) (int, error) {
	env := os.Getenv(key)

	if env != "" {
		intEnv, err := strconv.Atoi(env)
		if err != nil {
			return defaultValue, err
		}
		return intEnv, err
	}

	return defaultValue, nil

}

//GetBoolFromEnv get a bool value from a Environment Variable, if error or environment variable not exist, the default value will be returned.
func GetBoolFromEnv(key string, defaultValue bool) (bool, error) {
	env := os.Getenv(key)

	if env != "" {
		boolEnv, err := strconv.ParseBool(env)
		if err != nil {
			return defaultValue, err
		}
		return boolEnv, err
	}

	return defaultValue, nil

}

// PrintJSON print in terminal console, object in json
func PrintJSON(obj interface{}) {
	fmt.Println(ObjectToJSON(obj))
}

// ObjectToJSON return string object json
func ObjectToJSON(obj interface{}) string {
	bytesObj, err := json.Marshal(obj)
	if err != nil {
		return "Erro ao converter obj em json. :("
	}

	return string(bytesObj)
}
