package shortener

import (
	"context"
	"fmt"
	//"link_shortener/internal/model"
	"github.com/Kernunn/linkShortener/internal/model"

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

func (sh *LinkShortener) mustEmbedUnimplementedLinkShortenerServer() {}

func New() (*LinkShortener, error) {
	ls := &LinkShortener{}
	ls.store = store.New()
	err := ls.store.Open()
	if err != nil {
		return nil, err
	}
	return ls, nil
}

func (sh *LinkShortener) Create(ctx context.Context, url *LongURL) (*ShortURL, error) {
	link, err := sh.store.Link().GetByUrl(&model.Link{Url: url.GetUrl()})
	if err != nil {
		return nil, err
	}
	if link != nil {
		return &ShortURL{Url: link.ShortLink}, nil
	}

	var shLink string
	for i := 0; i < 100; i++ {
		shLink = getRandomString(lenShortLink)
		link, err := sh.store.Link().GetByShortLink(&model.Link{ShortLink: shLink})
		if err != nil {
			return nil, err
		}
		if link == nil {
			break
		}
	}
	link, err = sh.store.Link().GetByShortLink(&model.Link{ShortLink: shLink})
	if err != nil {
		return nil, err
	}
	if link != nil {
		return nil, fmt.Errorf("unable to create short url")
	}

	err = sh.store.Link().Create(&model.Link{Url: url.GetUrl(), ShortLink: shLink})
	if err != nil {
		return nil, err
	}
	return &ShortURL{Url: shLink}, nil
}

func (sh *LinkShortener) Get(ctx context.Context, url *ShortURL) (*LongURL, error) {
	link, err := sh.store.Link().GetByShortLink(&model.Link{ShortLink: url.GetUrl()})
	if err != nil {
		return nil, err
	}
	if link == nil {
		return nil, fmt.Errorf("not found")
	}
	return &LongURL{Url: link.Url}, nil
}

func (sh *LinkShortener) Close() error {
	return sh.store.Close()
}

func getRandomString(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Int63n(int64(len(alphabet)))])
	}
	return sb.String()
}
