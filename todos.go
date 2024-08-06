package main

type Todo struct {
	ID        int    `json:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
