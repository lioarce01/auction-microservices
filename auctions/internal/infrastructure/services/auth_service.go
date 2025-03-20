package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AuthService struct {
	AuthURL string
}

func NewAuthService(authURL string) *AuthService {
	return &AuthService{AuthURL: authURL}
}

func (s *AuthService) GetCreatorID(token string) (string, error) {
	fmt.Println("🔍 [AuthService] Solicitando ID del creador con token:", token)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/me", s.AuthURL), nil)
	if err != nil {
		fmt.Println("❌ [AuthService] Error creando request:", err)
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	fmt.Println("📡 [AuthService] Enviando solicitud a:", req.URL)
	client := &http.Client{
		Timeout: 5 * time.Second, // ⚠️ Evita que la solicitud se quede colgada.
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ [AuthService] Error en la solicitud HTTP:", err)
		return "", fmt.Errorf("failed to fetch creator ID from auth service: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("✅ [AuthService] Respuesta recibida, código:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Reemplazado ioutil.ReadAll por io.ReadAll
		fmt.Printf("⚠️ [AuthService] Respuesta no OK: %d, Body: %s\n", resp.StatusCode, string(body))
		return "", fmt.Errorf("failed to fetch creator ID, status code: %d", resp.StatusCode)
	}

	var response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("❌ [AuthService] Error parseando JSON:", err)
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	fmt.Println("✅ [AuthService] ID del creador obtenido:", response.Data.ID)
	return response.Data.ID, nil
}
