// A simple stdlib http handler that validates a request inline - rather than using
// the validator.Validator interface.
//
// I would recommend using this method of validation in your service layer
// and transport agnostic to ensure portability, but this example shows
// a simple usage.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/theflyingcodr/govalidator"
)

// Request is a request, duh.
type Request struct {
	Name      string    `json:"name"`
	DOB       time.Time `json:"dob"`
	IsEnabled bool      `json:"isEnabled"`
	Count     int       `json:"count"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		switch r.Method {
		case http.MethodPost:
			var req Request
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to decode request: %s", err)
				return
			}
			if err := validator.New().
				Validate("name", validator.StrLength(req.Name, 4, 10)).
				Validate("dob", validator.NotEmpty(req.DOB), validator.DateBefore(req.DOB, time.Now().AddDate(-16, 0, 0))).
				Validate("isEnabled", validator.Bool(req.IsEnabled, false)).
				Validate("count", validator.PositiveInt(req.Count)).Err(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp := map[string]interface{}{
					"errors": err,
				}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not supported")
		}

	})
	log.Fatal(http.ListenAndServe(":1234", nil))
}
