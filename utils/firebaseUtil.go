package utils

import (
	"bytes"
	"drinks-api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ValidateUser(user models.User) interface{} {

	firebaseKey := os.Getenv("FIREBASE_KEY")

	body, _ := json.Marshal(user)
	resp, err := http.Post("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key="+firebaseKey, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	corpo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var p2 interface{}
	json.Unmarshal(corpo, &p2)
	m := p2.(map[string]interface{})
	return m
}
