package server

import (
	"encoding/json"
	"fooservice/internal/model"
	"net/http"
	"strconv"
)

type DataService interface {
	QueryData(id int64) (model.Account, error)
}

type Service struct {
	service DataService
}

func NewService(service DataService) *Service {
	return &Service{service: service}
}

func (s *Service) Start() {
	http.Handle("GET /account/{id}", s.getAccount())
	_ = http.ListenAndServe(":8080", nil)
}

func (s *Service) getAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i int64
		id := r.PathValue("id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		a, err := s.service.QueryData(i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		_ = encoder.Encode(a)
	}
}
