package handlers

import (
	"net/http"

	"github.com/sant470/gocommon/errors"
)

type Handler func(rw http.ResponseWriter, r *http.Request) *errors.AppError

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := h(rw, r); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Message))
	}
}
