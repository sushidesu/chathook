package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime+", "+os.Getenv("VERCEL_ENV"))
}
