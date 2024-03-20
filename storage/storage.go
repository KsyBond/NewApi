package storage

import (
	"awesomeProject/server/models"
	"database/sql"
)

type Storage struct {
	DB *sql.DB
}

func New() *Storage {
	var s Storage
	return &s
}

func (s *Storage) Create(new *models.Post) error {
	stmt, err := s.DB.Prepare("INSERT INTO News (Title, Description) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(new.Title, new.Description)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Read() error {
	return nil
}

func (s *Storage) Update() error {
	return nil
}

func (s *Storage) Delete() error {
	return nil
}
