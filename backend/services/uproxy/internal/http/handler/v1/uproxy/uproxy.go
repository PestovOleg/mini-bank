// TODO: сделать  описание
package uproxy

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/uproxy/internal/http/mapper"
	"go.uber.org/zap"
)

type UnleashProxyHandler struct {
	logger *zap.Logger
}

func NewUnleashProxyHandler() *UnleashProxyHandler {
	return &UnleashProxyHandler{
		logger: logger.GetLogger("UProxyAPI"),
	}
}

// GetToggles godoc
// @Version 1.0
// @title GetToggles
// @Summary Unleash Proxy for Web
// @Description Unleash Proxy for Web
// @Tags uProxy
// @Accept  json
// @Produce  json
// @Success 200 {object} mapper.ToggleList "Successfully retrieved feature toggles"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /uproxy [get]
func (u *UnleashProxyHandler) ListToggles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{
			Timeout: time.Second * 2,
		}
		unleashServer := os.Getenv("UNLEASH_TOGGLES_URL")
		token := os.Getenv("UNLEASH_ADMIN_TOKEN")
		if unleashServer == "" || token == "" {
			u.logger.Error(
				"sysvar UNLEASH_TOGGLES or UNLEASH_ADMIN_TOKEN is not enabled, URL and token to Unleash Server cannot be found",
			)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, unleashServer, nil)

		if err != nil {
			u.logger.Sugar().Error("Failed to get features: %s", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", token)

		resp, err := client.Do(req)

		// Если статус отличен от успешного - возвращаем ошибку
		if err != nil {
			u.logger.Debug("Error from unleash: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		var Toggles mapper.ToggleList

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&Toggles); err != nil {
			u.logger.Sugar().Error("Failed to decode JSON response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		resp.Body.Close()

		if err := json.NewEncoder(w).Encode(Toggles); err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading Toggles"))
			if err != nil {
				u.logger.Error(err.Error())
			}
		}
	})
}
