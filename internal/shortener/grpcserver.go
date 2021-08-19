package shortener

import (
	"context"
	//"fmt"

	//"link_shortener/internal/store"
	"github.com/Kernunn/linkShortener/internal/store"

	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
const lenShortLink = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

type LinkShortener struct {
	store *store.Store
}

func (sh *LinkShortener) mustEmbedUnimplementedLinkShortenerServer() { }

func New() (*LinkShortener, error) {
	ls := &LinkShortener{}
	err := ls.store.Open()
	if err != nil {
		return nil, err
	}
	return ls, nil
}

func (sh *LinkShortener) Create(ctx context.Context, url *LongURL) (*ShortURL, error) {
	//if shLink, ok := sh.longToShort[url.GetUrl()]; ok {
	//	return &ShortURL{Url: shLink}, nil
	//}
	//var shLink string
	//for i := 0; i < 100; i++ {
	//	shLink = getRandomString(lenShortLink)
	//	if _, ok := sh.shortToLong[shLink]; !ok {
	//		break
	//	}
	//}
	//if _, ok := sh.shortToLong[shLink]; ok {
	//	return nil, fmt.Errorf("unable to create short url")
	//}
	//sh.longToShort[url.GetUrl()] = shLink
	//sh.shortToLong[shLink] = url.GetUrl()
	//return &ShortURL{Url: shLink}, nil
	return nil, nil
}

func (sh *LinkShortener) Get(ctx context.Context, url *ShortURL) (*LongURL, error) {
	//shURL, ok := sh.shortToLong[url.GetUrl()]
	//if !ok {
	//	return nil, fmt.Errorf("not found")
	//}
	//return &LongURL{Url: shURL}, nil
	return nil, nil
}

func getRandomString(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Int63n(int64(len(alphabet)))])
	}
	return sb.String()
}
