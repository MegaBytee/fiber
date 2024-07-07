package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateRandomTokenPlus(x int, c string) string {
	token := GenerateRandomToken(x) + c
	return token
}

func GenerateRandomToken(x int) string {
	key := make([]byte, x)
	_, err := rand.Read(key)
	if err != nil {
		// handle error here
		fmt.Println(err)
	}

	return hex.EncodeToString(key)
}

// recast a to b
func Recast(a, b any) error {
	js, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return json.Unmarshal(js, b)
}

func Serialize(s any) []byte {
	data, err := json.Marshal(s)
	if err != nil {
		print(err)
		return []byte(nil)
	}

	return data
}

func DeSerialize(data []byte) any {
	var a any
	err := json.Unmarshal(data, &a)
	if err != nil {
		print(err)
	}

	return a

}

func Hash(data string) string {

	//fmt.Println("hash:>data=", data)
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}

func StringToInt64(value string) int64 {
	p, _ := strconv.ParseInt(value, 10, 64)
	return p
}
func StringToInt32(value string) int32 {
	p, _ := strconv.ParseInt(value, 10, 32)
	return int32(p)
}
func StringToFloat64(value string) float64 {
	p, _ := strconv.ParseFloat(value, 64)
	return p
}

func StringToBool(value string) bool {
	p, _ := strconv.ParseBool(value)
	return p
}

func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 2)

}

func Int32ToString(value int32) string {
	return strconv.FormatInt(int64(value), 2)

}

func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func SlugToKeyword(slug string) string {
	//fmt.Println("slug:", slug)
	slug = strings.Replace(slug, "-", " ", -1)

	keyword := "\"" + slug + "\""
	//fmt.Println("keyword:", keyword)
	return keyword
}

func GetTimeFromUnixTime(i int64) time.Time {
	return time.Unix(i, 0).UTC()
}
