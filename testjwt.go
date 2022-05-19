package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret []byte

func testjwtinit() {
	// Load sample key data
	if keyData, e := ioutil.ReadFile("jwtkey"); e == nil {
		hmacSampleSecret = keyData
		fmt.Println(string(keyData))
	} else {
		panic(e)
	}
}

func testjwt() {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)

}

func testjwtdecode() {
	// Token from another example.  This token is expired
	var tokenString1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.QTPHyqQr8lQEUjZ7KphcpUIt7zx7Z0EDzDCJ0GybDQk"

	token, err := jwt.Parse(tokenString1, func(token *jwt.Token) (interface{}, error) {
		return []byte("sampekey"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	}

}
