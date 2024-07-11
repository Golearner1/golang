package secure

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "supersecret"

func Generatetoken(email string, authid int64) (string, error) {
	fmt.Print(authid)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"authid": authid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretkey))

}
func Verifytoken(token string) (int64, error) {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	tokenisvalid := parsedtoken.Valid
	if !tokenisvalid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// email:=claims["Email"].(string)
	authid := int64(claims["authid"].(float64))
	return authid, nil
}
