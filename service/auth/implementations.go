package auth

import (
	"crypto/rsa"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	repository repository.RepositoryInterface
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type AuthServiceOpts struct {
	PrivateKeyFile string
	PublicKeyFile  string
	Repository     repository.RepositoryInterface
}

func NewAuthService(opts AuthServiceOpts) (Service, error) {
	publicKey, privateKey, err := initKey(opts.PublicKeyFile, opts.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	return &service{
		repository: opts.Repository,
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (s *service) AuthenticateUserPassword(u *repository.User, pwd string) error {
	if !utils.ComparePasswordHash(pwd, u.Password) {
		return errors.New("wrong password")
	}

	return nil
}

func (s *service) GenerateJWT() (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *service) ValidateJWT(token string) error {
	validToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid token")
		}

		return s.publicKey, nil
	})
	if err != nil {
		return err
	}

	claims, ok := validToken.Claims.(jwt.RegisteredClaims)
	if !ok {
		return errors.New("invalid token")
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return errors.New("token expired")
	}

	return nil
}

func initKey(publicKeyFile, privateKeyFile string) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	rawPublicKey, err := os.ReadFile(filepath.Join(dir, publicKeyFile))
	if err != nil {
		return nil, nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		return nil, nil, err
	}

	rawPrivateKey, err := os.ReadFile(filepath.Join(dir, privateKeyFile))
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(rawPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	return publicKey, privateKey, nil
}
