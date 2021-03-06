package session

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (s *Session) BooleanToString(element bool) string {
	if element == true {
		return "true"
	} else {
		return "false"
	}
}

func (s *Session) StringToBoolean(element string) bool {
	if strings.TrimSpace(element) == "true" {
		return true
	} else {
		return false
	}
}

func (s *Session) IntegerToString(element int) string {
	value := strconv.Itoa(element)
	return value
}

func (s *Session) StringToInteger(element string) int {
	converted, _ := strconv.Atoi(element)
	return converted
}

func (s *Session) StringToInt64(element string) int64 {
	n, _ := strconv.ParseInt(element, 10, 64)
	return n
}

func (s *Session) IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
