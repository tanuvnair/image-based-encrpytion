package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tanuvnair/image-based-encryption/internal/handlers"
	"github.com/tanuvnair/image-based-encryption/internal/service"
	"github.com/tanuvnair/image-based-encryption/internal/utils"
)

func main() {
	port := flag.Int("port", 8080, "HTTP server port")
	imageDir := flag.String("image-dir", "./images", "directory containing entropy source images")
	flag.Parse()

	if err := utils.ValidateImageDir(*imageDir); err != nil {
		log.Fatalf("invalid image dir: %v", err)
	}

	svc, err := service.NewEntropyService(*imageDir)
	if err != nil {
		log.Fatalf("failed to create entropy service: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/random", handlers.HandleRandom(svc))

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("starting image based encryption server on %s (images: %s)", addr, *imageDir)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
