package apiV1_test

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/handler/apiV1"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/service/mocks"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/authentication"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortURLHandler_GetUser(t *testing.T) {
	userService := new(mocks.MockUserService)
	user := domain.User{Model: gorm.Model{ID: 1}, Email: "test@test.com", Links: []domain.Link{
		{Model: gorm.Model{ID: 1}, LongUrl: "http://test.com", ShortUrl: "http://test.com/1"},
		{Model: gorm.Model{ID: 2}, LongUrl: "http://test.com", ShortUrl: "http://test.com/1"},
	}}
	authRequest := httptest.NewRequest(http.MethodGet, "/v1/user/1", nil)
	token, err := authentication.GenerateToken(user.Email)
	authRequest.Header.Add("Authorization", "Bearer "+token)
	assert.NoError(t, err)
	userService.On("GetUser", "1").Return(user, nil)
	linkService := new(mocks.MockLinkService)
	logger := logrus.New()
	testCases := []struct {
		name           string
		responseWriter *httptest.ResponseRecorder
		request        *http.Request
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GetUser",
			responseWriter: httptest.NewRecorder(),
			request:        authRequest,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"email":"test@test.com","links":[{"id":1,"url":"http://test.com","short_url":"http://test.com/1"},{"id":2,"url":"http://test.com","short_url":"http://test.com/1"}]}`,
		},
		{
			name:           "GetUser_NotFound",
			responseWriter: httptest.NewRecorder(),
			request:        &http.Request{},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `"Error getting user"`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			userService.On("GetUserByEmail", mock.Anything).Return(&user, nil)
			shortURLHandler := apiV1.NewShortURLHander(linkService, userService, nil, logger)
			shortURLHandler.GetUser(testCase.responseWriter, testCase.request)
			assert.Equal(t, testCase.expectedStatus, testCase.responseWriter.Code)
			assert.Equal(t, testCase.expectedBody, testCase.responseWriter.Body.String())
		})
	}
}
