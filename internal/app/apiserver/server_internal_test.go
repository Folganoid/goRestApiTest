package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"goRestApi/internal/app/model"
	"goRestApi/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {

	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "valid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionCreate(t *testing.T) {

	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)

	s := newServer(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string {
				"email": u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload: "lkjlkj",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string {
				"email": "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid pass",
			payload: map[string]string {
				"email": u.Email,
				"password": "kjkjk",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}