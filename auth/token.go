package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
func TokenValid(r *http.Request) error {
	tokenStirng := ExtractToken(r)
	token, err := jwt.Parse(tokenStirng, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
