package responseutils

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"

type UserResponse struct {
	ID    uint           `json:"id"`
	Email string         `json:"email"`
	Links []LinkResponse `json:"links"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func MapUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:    user.Model.ID,
		Email: user.Email,
		Links: MapLinkListResponse(user.Links),
	}
}
