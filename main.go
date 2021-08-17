package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var longToShort = make(map[string]string)
var shortToLong = make(map[string]string)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
const lenShortLink = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomString(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Int63n(int64(length))])
	}
	return sb.String()
}

func Create(url string) (string, bool) {
	if shLink, ok := longToShort[url]; ok {
		return shLink, true
	}
	var shortLink string
	for i := 0; i < 100; i++ {
		shortLink = getRandomString(lenShortLink)
		if _, ok := shortToLong[shortLink]; !ok {
			break
		}
	}
	if _, ok := shortToLong[shortLink]; ok {
		return "", false
	}
	longToShort[url] = shortLink
	shortToLong[shortLink] = url
	return shortLink, true
}

func Get(shortLink string) (string, bool) {
	url, ok := shortToLong[shortLink]
	return url, ok
}

func main() {
	shLink, _ := Create("longLink1")
	fmt.Println(Get(shLink + "f"))
}
