package handlers

import (
	"fmt"
	"net/http"
)

func Sniff(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Sniff")
}
