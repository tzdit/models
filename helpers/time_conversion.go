package helpers

import (
	"encoding/base64"
	"time"
)

//PurseTime converts string time int golang time in a given format
func PurseTime(format string, timeString string) *time.Time {
	newTime := time.Time{}
	newTime, _ = time.Parse(format, timeString)
	return &newTime
}

//BytesToBase64String converts byte array to base64 string
func BytesToBase64String(data []byte) string {
	encodedString := base64.StdEncoding.EncodeToString(data)
	return encodedString
}

//Base64StringToBytes converts base64 string into byte array
func Base64StringToBytes(data string) ([]byte, error) {
	decodedString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decodedString, nil
}
