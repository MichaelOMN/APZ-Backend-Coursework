package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	entity "sport_app"
	"sport_app/pkg/repository"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwt_go "github.com/golang-jwt/jwt"
)

const (
	salt       = "b09d60dd88305ed83c3e55e1dae03ee3bf550231"
	signingKey = "57d44c0c0169cc99ed18d1b24c312b0b219cde9c"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	EntityId int `json:"entity_id"`
}

// type tokenActivityClaims struct {
// 	jwt.StandardClaims
// 	EntityName string `json:"entity_name"`
// }

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateVisitor(user entity.Visitor) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateVisitor(user)
}

func (s *AuthService) CreateCoach(user entity.Coach) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateCoach(user)
}

func (s *AuthService) GenerateToken(username, password string, isCoach bool) (string, error) {
	var user any
	var err error

	if isCoach {
		user, err = s.repo.GetCoach(username, generatePasswordHash(password))
	} else {
		user, err = s.repo.GetVisitor(username, generatePasswordHash(password))
	}

	if err != nil {
		return "", err
	}

	var token *jwt.Token

	if isCoach {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.(entity.Coach).Id,
		})
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.(entity.Visitor).Id,
		})
	}

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateTokenForActivity(name string) (string, error) {
	// var user entity.Activity
	// var err error

	// if err != nil {
	// 	return "", err
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenActivityClaims{
	// 	jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
	// 		IssuedAt:  time.Now().Unix(),
	// 	},
	// 	user.Name,
	// })

	token := jwt_go.New(jwt_go.SigningMethodHS256)
	claims := token.Claims.(jwt_go.MapClaims)
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour)
	claims["iat"] = time.Now()
	claims["EntityName"] = name

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("token claims are neither of type *tokenClaims")

	}

	return claims.EntityId, nil
}

func decodeTokenPayload(tokenString string) (map[string]interface{}, error) {
	// Split the token into three parts: header, payload, and signature
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	// Decode the payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode token payload: %v", err)
	}

	// Unmarshal the JSON payload into a map
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token payload: %v", err)
	}

	return claims, nil
}

func verifySignatureHS256(token string, secretKey string) (bool, error) {
	// Split the token into three parts: header, payload, and signature
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid token format")
	}

	// Decode the signature
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return false, fmt.Errorf("failed to decode token signature: %v", err)
	}

	//logrus.Fatal("Compute the HMAC-SHA256 hash!")
	// Compute the HMAC-SHA256 hash of the token's header and payload using the secret key
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(parts[0] + "." + parts[1]))
	expectedSignature := mac.Sum(nil)

	//logrus.Fatal("Compare the computed signature!")
	// Compare the computed signature with the decoded signature
	return hmac.Equal(signature, expectedSignature), nil
}

func (s *AuthService) ParseActivityToken(accessToken string) (string, error) {

	payload, err := decodeTokenPayload(accessToken)
	if err != nil {
		return ``, err
	}

	//logrus.Fatal("In da ParseActToken1!")

	var valid bool
	valid, err = verifySignatureHS256(accessToken, signingKey)

	if err != nil {
		return ``, err
	}

	//logrus.Fatal("In da ParseActToken2!")

	if !valid {
		return ``, errors.New("signature is invalid")
	}

	name, ok := payload["EntityName"].(string)

	if !ok {
		return ``, errors.New("name is not a string type")
	}

	return name, nil
	//logrus.Fatal("In da ParseActToken1!")
	// this fucking crashed
	// token, err := jwt_go.Parse(accessToken, func(token *jwt_go.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt_go.SigningMethodHMAC); !ok {
	// 		return nil, errors.New("invalid signing method")
	// 	}

	// 	return []byte(signingKey), nil
	// })
	//logrus.Fatal("In da ParseActToken2!")

	// if err != nil {
	// 	return ``, err
	// }

	// claims, ok := token.Claims.(jwt_go.MapClaims)
	// if !ok {
	// 	return ``, errors.New("token claims are neither of type *MapClaims")

	// }

	// name, ok := claims["EntityName"].(string)
	// if !ok {
	// 	return ``, errors.New("there is no EntityName inside token")
	// }

	// return name, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
