package model

import (
	"crypto/sha512"
	"encoding/base64"
	"testing"
)

func TestLoginSendsCorrectPasswordHash(t *testing.T) {
	testDB := new(mockDB)
	testDB.returnedRow = &mockRow{} // because login funcion has implemented scan() for the row!
	db = testDB                     // assign the testDB to the db object!!!

	password := "password"
	email := "email"
	Login(email, password)

	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if testDB.lastArgs[1] != pwd {
		t.Errorf("Login function failed to send correct password hash to database")
	}
}
