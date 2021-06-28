package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max of 32 bytes
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var str strings.Builder
	sz := len(alpha)

	for i := 0; i < n; i++ {
		c := alpha[rand.Intn(sz)]
		str.WriteByte(c)
	}

	return str.String()
}

// RandomStringList generates a list of random strings
func RandomStringList() []string {
	var res []string
	for i := 0; i < int(RandomInt(5, 10)); i++ {
		res = append(res, RandomString(int(RandomInt(10, 20))))
	}
	return res
}

// RandomName generates a random Anime Name
func RandomName() string {
	return RandomString(15)
}

func RandomSummary() string {
	return RandomString(100)
}

func RandomEpisodes() int32 {
	return RandomInt(11, 1000)
}

func RandomNames() []string {
	return RandomStringList()
}

func RandomGenre() []string {
	return RandomStringList()
}
