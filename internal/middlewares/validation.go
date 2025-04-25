package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ClothShop/product-service/internal/utils"
)

func ValidateProductMiddleware[T any](target *T) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "❌ Ошибка чтения тела запроса", http.StatusBadRequest)
				log.Println("❌ Ошибка чтения тела запроса:", err)
				return
			}

			if err := json.Unmarshal(body, target); err != nil {
				http.Error(w, "❌ Неверный формат JSON", http.StatusBadRequest)
				log.Println("❌ Неверный формат JSON:", err)
				return
			}

			if err := utils.ValidateStruct(target); err != nil {
				http.Error(w, "❌ Ошибка валидации: "+err.Error(), http.StatusBadRequest)
				log.Println("❌ Ошибка валидации:", err)
				return
			}

			r.Body = io.NopCloser(bytes.NewReader(body))

			next.ServeHTTP(w, r)
		})
	}
}
