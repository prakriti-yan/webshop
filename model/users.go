package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

const passwordSalt = "L?fE*gpel)PCWWsLTeER"

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	Lastlogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString((hasher.Sum(nil)))
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname
		FROM public.users
		WHERE email = $1
		 AND password = $2`, email, pwd)
	fmt.Println(row)
	fmt.Println(pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("user not found!")
		return nil, fmt.Errorf("User not found")
	case err != nil:
		fmt.Println(err)
		return nil, err
	}
	// update the lastlogin
	t := time.Now()
	_, err = db.Exec(`
	UPDATE public.users
	SET lastlogin = $1
	WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	}
	return result, nil
}
