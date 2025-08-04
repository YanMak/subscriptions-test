package conv

import (
	"os"
	"strconv"
)

func ToInt(key string) int {
	raw := os.Getenv(key)
	res, err := strconv.Atoi(raw)
	if err != nil {
		println("reading .env variable (", key, ") : ", err)
		panic(err)
	}
	return res
}
