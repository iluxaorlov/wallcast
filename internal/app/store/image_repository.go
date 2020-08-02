package store

import (
	"database/sql"
	"github.com/iluxaorlov/wallcast/internal/app/model"
)

type ImageRepository struct {
	store *Store
}

func (r *ImageRepository) Insert(img *model.Image) error {
	return r.store.db.QueryRow("INSERT INTO images (url, date) VALUES ($1, $2) RETURNING id", &img.Url, &img.Date).Scan(&img.Id)
}

func (r *ImageRepository) FindByUrl(url string) (*model.Image, error) {
	img := new(model.Image)

	err := r.store.db.QueryRow("SELECT id, url, date FROM images WHERE url = $1", url).Scan(&img.Id, &img.Url, &img.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return img, nil
}
