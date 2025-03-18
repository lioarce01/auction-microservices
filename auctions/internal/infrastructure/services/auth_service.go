package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthService struct {
	AuthURL string
}

func NewAuthService(authURL string) *AuthService {
	return &AuthService{AuthURL: authURL}
}

func (s *AuthService) GetCreatorID(sub string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s", s.AuthURL, sub))
	if err != nil {
		return "", fmt.Errorf("failed to fetch creator ID from auth service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch creator ID, status code: %d", resp.StatusCode)
	}

	var user struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	return user.ID, nil
}
