package store

import "database/sql"

type Store struct {
	db              *sql.DB
	imageRepository *ImageRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Image() *ImageRepository {
	if s.imageRepository == nil {
		s.imageRepository = &ImageRepository{
			store: s,
		}
	}

	return s.imageRepository
}
