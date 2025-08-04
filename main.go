package main

import (
	"log"
	"net/http"
	"obs-go-backend/handler"
	"obs-go-backend/youtube"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, reading from environment variables")
	}

	// Environment'tan değişkenleri oku
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	channelID := os.Getenv("YOUTUBE_CHANNEL_ID")
	if apiKey == "" || channelID == "" {
		log.Fatal("YOUTUBE_API_KEY and YOUTUBE_CHANNEL_ID must be set")
	}

	// Bağımlılıkları oluştur
	ytClient := youtube.NewClient(apiKey, channelID)
	apiHandler := &handler.APIHandler{YTClient: ytClient}
	viewHandler := &handler.ViewHandler{}

	// Router'ı başlat
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Loglama için middleware

	// Route Grupları
	// 1. Grup: JSON dönen API endpoint'leri
	r.Route("/api", func(r chi.Router) {
		r.Get("/latest-subscriber", apiHandler.LatestSubscriberHandler)
		// Diğer YouTube API fonksiyonları buraya eklenecek...
	})

	// 2. Grup: OBS'te görünecek HTML sayfaları
	r.Route("/obs", func(r chi.Router) {
		r.Get("/", viewHandler.IndexPageHandler)
		// Diğer OBS sayfaları buraya eklenecek...
	})

	log.Println("Server starting on port :8080...")
	http.ListenAndServe(":8080", r)
}
