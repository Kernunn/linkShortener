package store

import (
	"database/sql"
	//"link_shortener/internal/model"
	"github.com/Kernunn/linkShortener/internal/model"
)

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) Create(u *model.Link) error {
	_, err := r.store.db.Exec(
		"INSERT INTO link "+
			"VALUES ($1, $2)", u.Url, u.ShortLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *LinkRepository) GetByUrl(u *model.Link) (*model.Link, error) {
	var shortLink string
	err := r.store.db.QueryRow("select shortlink "+
		"from link "+
		"where url = $1;", u.Url).Scan(&shortLink)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &model.Link{
		Url:       u.Url,
		ShortLink: shortLink,
	}, nil
}

func (r *LinkRepository) GetByShortLink(u *model.Link) (*model.Link, error) {
	var url string
	err := r.store.db.QueryRow("select url "+
		"from link "+
		"where shortlink = $1;", u.ShortLink).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &model.Link{
		Url:       url,
		ShortLink: u.ShortLink,
	}, nil
}
