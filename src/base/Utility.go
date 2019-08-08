package base

import (
	"crypto/sha1"
	"fmt"
	"log"
	"strings"
	"time"
)

func GetServerTime() int {
	return int(time.Now().Unix()) + serverTimeDelta
}

func BinaryToHex(data []byte) string {
	builder := strings.Builder{}

	builder.Grow(len(data) * 2)

	for _, b := range data {
		i := b & 0xff

		if i < 0x10 {
			builder.WriteString("0")
		}

		builder.WriteString(fmt.Sprintf("%x", i))
	}

	return strings.ToLower(builder.String())
}

func GetSHA1String(stringToHash string) string {
	hash := sha1.Sum([]byte(stringToHash))

	return BinaryToHex(hash[:])
}

func Debug(v ...interface{}) {
	log.Println("[DEBUG]", v)
}

func Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

func Warning(v ...interface{}) {
	log.Println("[WARN]", v)
}

func Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}
