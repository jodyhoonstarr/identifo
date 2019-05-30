package fake

import (
	"fmt"
	"path"

	"github.com/madappgang/identifo/jwt"
	jwtService "github.com/madappgang/identifo/jwt/service"
	"github.com/madappgang/identifo/model"
	"github.com/madappgang/identifo/storage/mem"
)

// NewComposer creates new database composer.
func NewComposer(settings model.ServerSettings, options ...func(*DatabaseComposer) error) (*DatabaseComposer, error) {
	c := DatabaseComposer{
		settings:                   settings,
		newAppStorage:              mem.NewAppStorage,
		newUserStorage:             mem.NewUserStorage,
		newTokenStorage:            mem.NewTokenStorage,
		newVerificationCodeStorage: mem.NewVerificationCodeStorage,
	}

	for _, option := range options {
		if err := option(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

// DatabaseComposer composes in-memory services.
type DatabaseComposer struct {
	settings                   model.ServerSettings
	newAppStorage              func() (model.AppStorage, error)
	newUserStorage             func() (model.UserStorage, error)
	newTokenStorage            func() (model.TokenStorage, error)
	newVerificationCodeStorage func() (model.VerificationCodeStorage, error)
}

// Compose composes all services with in-memory storage support.
func (dc *DatabaseComposer) Compose() (
	model.AppStorage,
	model.UserStorage,
	model.TokenStorage,
	model.VerificationCodeStorage,
	jwtService.TokenService,
	error,
) {

	appStorage, err := dc.newAppStorage()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	userStorage, err := dc.newUserStorage()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	tokenStorage, err := dc.newTokenStorage()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	verificationCodeStorage, err := dc.newVerificationCodeStorage()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	tokenServiceAlg, ok := jwt.StrToTokenSignAlg[dc.settings.Algorithm]
	if !ok {
		return nil, nil, nil, nil, nil, fmt.Errorf("Unknow token service algoritm %s", dc.settings.Algorithm)
	}

	tokenService, err := jwtService.NewJWTokenService(
		path.Join(dc.settings.PEMFolderPath, dc.settings.PrivateKey),
		path.Join(dc.settings.PEMFolderPath, dc.settings.PublicKey),
		dc.settings.Issuer,
		tokenServiceAlg,
		tokenStorage,
		appStorage,
		userStorage,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return appStorage, userStorage, tokenStorage, verificationCodeStorage, tokenService, nil
}
