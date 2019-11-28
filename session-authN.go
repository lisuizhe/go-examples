package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

const (
	sessionTokenKey = "session_token"
)

var (
	cache redis.Conn
	users map[string]string
)

func init() {
	users = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}

func main() {
	initCache()
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Credentials is
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Signin does
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionTokenUUID, _ := uuid.NewV4()
	sessionToken := sessionTokenUUID.String()
	_, err = cache.Do("SETEX", sessionToken, "120", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    sessionTokenKey,
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

// Welcome does
func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(sessionTokenKey)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", response)))
}

// Refresh does
func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(sessionTokenKey)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionTokenUUID, err := uuid.NewV4()
	newSessionToken := newSessionTokenUUID.String()
	_, err = cache.Do("SETEX", newSessionToken, "120", fmt.Sprintf("%s", response))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = cache.Do("DEL", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    sessionTokenKey,
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}
