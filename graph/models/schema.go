// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type NewTodo struct {
	Text string `json:"text"`
	Name string `json:"name"`
}

type Todo struct {
	ID   uint64 `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
