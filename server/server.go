package server

import (
	"os"
	"path/filepath"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	users    []User
	imageDir string
	mutex    sync.Mutex
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
}

var Categories = []Category{
	{ID: 1, Name: "Кофе"},
	{ID: 2, Name: "Чай"},
	{ID: 3, Name: "Десерты"},
}

var Products = []Product{
	{ID: 1, Name: "Эспрессо", Description: "Крепкий кофе", Price: 2.50, CategoryID: 1},
	{ID: 2, Name: "Латте", Description: "Кофе с молоком", Price: 3.50, CategoryID: 1},
	{ID: 3, Name: "Зеленый чай", Description: "Освежающий чай", Price: 2.00, CategoryID: 2},
	{ID: 4, Name: "Чизкейк", Description: "Классический десерт", Price: 4.00, CategoryID: 3},
}

func NewServer(imageDir string) (*Server, error) {
	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
		return nil, err
	}

	return &Server{
		users:    []User{},
		imageDir: imageDir,
	}, nil
}

func (s *Server) GetUsers() []User {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.users
}

func (s *Server) AddUser(user User) User {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	user.ID = len(s.users) + 1
	s.users = append(s.users, user)
	return user
}
func (s *Server) GetImagePath(filename string) (string, error) {
	filePath := filepath.Join(s.imageDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", err
	}
	return filePath, nil
}
