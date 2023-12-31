package memdb

import (
	"GoNews/pkg/storage"
	"errors"
)

// Хранилище данных.
type Store struct{}

// Конструктор объекта хранилища.
func New() *Store {
	return new(Store)
}

func (s *Store) Posts() ([]storage.Post, error) {
	return posts, nil
}

func (s *Store) AddPost(post storage.Post) error {
	posts = append(posts, post)
	return nil
}
func (s *Store) UpdatePost(upost storage.Post) error {
	for i, post := range posts {
		if post.ID == upost.ID {
			posts[i] = post
			return nil
		}
	}
	return errors.New("post not found")
}
func (s *Store) DeletePost(post storage.Post) error {
	for i, p := range posts {
		if p.ID == post.ID {
			posts = append(posts[:i], posts[i+1:]...)
			return nil
		}
	}
	return errors.New("post not found")
}

var posts = []storage.Post{
	{
		ID:      1,
		Title:   "Effective Go",
		Content: "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
	},
	{
		ID:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
	},
}
