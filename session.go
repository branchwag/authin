package main

import (
	"fmt"
	"log"
	"net/http"
)

func Authorize(r *http.Request) error {
	// Get session token from cookies
	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil || sessionTokenCookie == nil {
		log.Println("Session token missing or invalid")
		return fmt.Errorf("session token missing or invalid")
	}
	fmt.Println("Session Token from Cookie:", sessionTokenCookie.Value)

	// Get CSRF token from request header
	csrfToken := r.Header.Get("X-CSRF-Token")
	if csrfToken == "" {
		log.Println("CSRF token missing")
		return fmt.Errorf("CSRF token missing")
	}
	fmt.Println("CSRF Token from Header:", csrfToken)

	// Get the username from the form value or body
	username := r.FormValue("username")
	if username == "" {
		log.Println("Username missing in form")
		return fmt.Errorf("username missing")
	}
	fmt.Println("Username from Form:", username)

	// Find user by username
	user, ok := users[username]
	if !ok {
		log.Println("User not found")
		return fmt.Errorf("user not found")
	}

	// Validate session token
	if sessionTokenCookie.Value != user.SessionToken {
		log.Println("Invalid session token")
		return fmt.Errorf("invalid session token")
	}

	// Validate CSRF token
	if csrfToken != user.CSRFToken {
		log.Println("CSRF token mismatch")
		return fmt.Errorf("CSRF token mismatch")
	}

	// If all checks pass, return nil (authorization successful)
	return nil
}
