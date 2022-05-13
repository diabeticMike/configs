package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/internship/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestController_Users(t *testing.T) {
	url := "/users"
	t.Run("success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockReader := mock.NewMockReader(mockCtrl)
		mockReader.EXPECT().Read("users.json").Return("hello", nil)

		c := &controller{reader: mockReader}
		r := *mux.NewRouter()
		rr := httptest.NewRecorder()
		r.HandleFunc(url, c.users).Methods(http.MethodGet)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, `hello`, rr.Body.String())
	})

	t.Run("failed_read_error", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockReader := mock.NewMockReader(mockCtrl)
		mockReader.EXPECT().Read("users.json").Return("", errors.New("error"))

		c := &controller{reader: mockReader}
		r := *mux.NewRouter()
		rr := httptest.NewRecorder()
		r.HandleFunc(url, c.users).Methods(http.MethodGet)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, `users upload error`, rr.Body.String())
	})
}
