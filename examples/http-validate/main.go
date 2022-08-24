// This has the same logic as the http-inline example but it
// has implemented validator.Validator and as such requests can be checked
// in a single place. In this case, in a ParseRequest function
// will decode the json responses and check that the request is
// a validator.
//
// If it is, the validator will be evaluated and if a failure is found
// ie if Err() returns a non nil response, an error is returned.
//
// As is shown, this can then be checked in an error handler and a response
// created and returned.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/theflyingcodr/govalidator/v2"
)

// Request is a request, duh.
type Request struct {
	Name      string    `json:"name"`
	DOB       time.Time `json:"dob"`
	IsEnabled bool      `json:"isEnabled"`
	Count     int       `json:"count"`
}

// Validate implements validator.Validator and evaluates Request.
func (r *Request) Validate() validator.ErrValidation {
	return validator.New().
		Validate("name", validator.StrLength(r.Name, 4, 10)).
		Validate("dob", validator.NotEmpty(r.DOB), validator.DateBefore(r.DOB, time.Now().AddDate(-16, 0, 0))).
		Validate("isEnabled", validator.Equal(r.IsEnabled, false)).
		Validate("count", validator.PositiveNumber(r.Count))
}

func main() {
	http.Handle("/", errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Add("content-type", "application/json")
		switch r.Method {
		case http.MethodPost:
			// attempt to parse our request
			var req Request
			if err := parseRequest(r, &req); err != nil {
				// fire the error back up the chain
				return err
			}
			// Do your awesome stuff here
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not supported")
		}
		return nil
	}))
	log.Fatal(http.ListenAndServe(":1234", nil))
}

// parseRequest will attempt to decode the request as json into a struct, out.
//
// It will then check to see if the type implements Validator and if so checks
// for validation errors. It returns an error if validation fails.
func parseRequest(r *http.Request, out interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		return fmt.Errorf("failed to parse request")
	}
	if v, ok := out.(validator.Validator); ok {
		if err := v.Validate().Err(); err != nil {
			return err
		}
	}
	return nil
}

// errorHandler is a simple error handler for demonstration
// purposes only.
func errorHandler(h httpHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}
		if e, ok := err.(validator.ErrValidation); ok {
			resp := map[string]interface{}{
				"errors": e,
			}
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(&resp)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// httpHandlerFunc is a custom handler func that returns an error.
type httpHandlerFunc func(w http.ResponseWriter, r *http.Request) error
