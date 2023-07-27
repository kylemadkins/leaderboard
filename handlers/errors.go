package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrResponse{StatusCode: 405}
	ErrNotFound         = &ErrResponse{StatusCode: 404}
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func BadRequest(err error) *ErrResponse {
	return &ErrResponse{
		Err:        err,
		StatusCode: 400,
		Message:    err.Error(),
	}
}

func ServerError(err error) *ErrResponse {
	return &ErrResponse{
		Err:        err,
		StatusCode: 500,
		Message:    err.Error(),
	}
}
