package handler

import "fmt"

type CreateTodoRequest struct {
	Title string `json:"title"`
}

func (c *CreateTodoRequest) Validate() error {
	if c.Title == "" {
		return fmt.Errorf("param: title is required")
	}
	return nil
}

type UpdateTodoRequest struct {
	Title  string `json:"title"`
	IsDone *bool  `json:"isDone"`
}

func (c *UpdateTodoRequest) Validate() error {
	if c.Title == "" {
		return fmt.Errorf("param title can not be blank")
	}
	return nil
}
