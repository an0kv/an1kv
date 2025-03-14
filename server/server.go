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
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image,omitempty"`
}

type Server struct {
	users    []User
	ImageDir string
	mutex    sync.Mutex
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	Image       string  `json:"image,omitempty"`
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
		ImageDir: imageDir,
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
	filePath := filepath.Join(s.ImageDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", err
	}
	return filePath, nil
}
func (s *Server) CategoryExists(id int) bool {
	for _, cat := range Categories {
		if cat.ID == id {
			return true
		}
	}
	return false
}

func (s *Server) ProductExists(id int) bool {
	for _, prod := range Products {
		if prod.ID == id {
			return true
		}
	}
	return false
}
