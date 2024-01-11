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
	if c.Title != "" || c.IsDone != nil {
		return nil
	}
	fmt.Errorf("at least one property 'title' or 'isDone' must be provided")
	return nil
}
