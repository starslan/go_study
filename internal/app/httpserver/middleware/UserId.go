package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go_study/internal/app/httpserver/handlers"
	"net/http"
	"strconv"
	"strings"
)

var secretKey = []byte("secret key")

func GetUserId(r *http.Request) (uint32, error) {

	cookie, err := r.Cookie("user-id")
	if err != nil {
		return 0, err
	}

	fmt.Println(cookie)

	return 55, nil
}

func nextId(shortURLList map[string]handlers.URLUserItem) uint32 {

	var i uint32 = 0
	for _, element := range shortURLList {
		if element.UserId > i {
			i = element.UserId
		}
	}
	return i + 1
}

func CheckUserIdMiddleware(shortURLList map[string]handlers.URLUserItem) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var idStr = ""
			id, err := GetUserId(r)
			if err != nil {

				idStr := string(nextId(shortURLList))

				cookie := http.Cookie{Name: "user-id", Value: GenerateSignedValue(secretKey, idStr)}
				http.SetCookie(w, &cookie)
			} else {
				idStr = strconv.Itoa(int(id))

			}

			r.Header.Set("X-USER-ID", idStr)
			next.ServeHTTP(w, r)
		})
	}
}

//func CheckUserIdMiddleware(shortURLList map[string]handlers.URLUserItem) http.Handler
//func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//			var idStr = ""
//			id, err := GetUserId(r)
//			if err != nil {
//
//				GenerateSignedValue(secretKey, value)
//
//			} else {
//				idStr = strconv.Itoa(int(id))
//
//			}
//
//			r.Header.Set("X-USER-ID", idStr)
//			next.ServeHTTP(w, r)
//		})
//	}
//}

func GenerateSignedValue(secretKey []byte, value string) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(value))
	signature := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s|%s", value, signature)
}

func VerifySignedValue(secretKey []byte, signedValue string) (bool, string) {
	parts := strings.Split(signedValue, "|")
	if len(parts) != 2 {
		return false, ""
	}

	value := parts[0]
	expectedSig := GenerateSignedValue(secretKey, value)

	return hmac.Equal([]byte(signedValue), []byte(expectedSig)), value
}
