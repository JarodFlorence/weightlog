package main

import (
	"./sessions"
	"./users"
	"./util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = ":8080"
	defaultResp = `{"login":"/login"}`
	contentType = "Content-Type"
	applicationJSON = "application/json"
)

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type LoginResponse struct {
	Errors    []string `json:"errors"`
	SessionID string   `json:"session_id,omitempty"`
}

type RegisterResponse struct {
	Errors []string      `json:"errors"`
	User   *users.User `json:"user,omitempty"`
}

func writeJSON(rw http.ResponseWriter, v interface{}) {
	rw.Header().Set(contentType, applicationJSON)
	json.NewEncoder(rw).Encode(v)
}

func handleError(rw http.ResponseWriter, err error, statusCode int) {
	resp := ErrorResponse{Errors: []string{err.Error()}}
	rw.WriteHeader(statusCode)
	writeJSON(rw, resp)
}

func rootHandler(rw http.ResponseWriter, _ *http.Request) {
	writeJSON(rw, defaultResp)
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	email, pw := req.FormValue("email"), req.FormValue("password")
	if len(email) == 0 || len(pw) == 0 {
		handleError(rw, errors.New("no_username_password"), http.StatusBadRequest)
		return
	}

	db := util.GetDB()
	defer db.Close()

	u, err := users.GetByEmail(db, email)
	if err != nil {
		log.Printf("error while looking up user: %v", err)
		handleError(rw, errors.New("bad_login"), http.StatusBadRequest)
		return
	}

	if !u.Verify(pw) {
		handleError(rw, errors.New("bad_login"), http.StatusBadRequest)
		return
	}

	s, _ := sessions.New(u.Id)
	if err = s.Save(db); err != nil {
		log.Printf("error while saving session in database: %v", err)
		handleError(rw, errors.New("internal_error"), http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{Errors: []string{}, SessionID: s.Id}
	writeJSON(rw, resp)
}

func registerHandler(rw http.ResponseWriter, req *http.Request) {
	email, pw, pwConf := req.FormValue("email"), req.FormValue("password"), req.FormValue("password_confirmation")
	if email == "" || pw == "" || pwConf == "" {
		handleError(rw, errors.New("no_email_password"), http.StatusBadRequest)
		return
	}

	if pw != pwConf {
		handleError(rw, errors.New("pw_conf_bad_match"), http.StatusBadRequest)
		return
	}

	user, err := users.New(email, pw)
	if err != nil {
		handleError(rw, errors.New("hash_error"), http.StatusInternalServerError)
		return
	}

	db := util.GetDB()
	defer db.Close()

	if err = user.Save(db); err != nil {
		handleError(rw, errors.New("db_error"), http.StatusInternalServerError)
		return
	}

	resp := RegisterResponse{Errors: []string{}, User: user}
	writeJSON(rw, resp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", registerHandler).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	} else {
		port = ":" + port
	}

	log.Printf("starting api server on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
