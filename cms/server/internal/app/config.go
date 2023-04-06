package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/gcp"
	"github.com/reearth/reearthx/appx"
	"github.com/reearth/reearthx/log"
	"github.com/samber/lo"
)

const configPrefix = "REEARTH_CMS"

type Config struct {
	Port         string `default:"8080" envconfig:"PORT"`
	ServerHost   string
	Host         string `default:"http://localhost:8080"`
	Dev          bool
	Host_Web     string
	GraphQL      GraphQLConfig
	Origins      []string
	DB           string `default:"mongodb://localhost"`
	Mailer       string
	SMTP         SMTPConfig
	SendGrid     SendGridConfig
	SignupSecret string
	GCS          GCSConfig
	Task         gcp.TaskConfig
	AssetBaseURL string
	Web          WebConfig
	Web_Disabled bool
	// auth
	Auth          AuthConfigs
	Auth0         Auth0Config
	Auth_ISS      string
	Auth_AUD      string
	Auth_ALG      *string
	Auth_TTL      *int
	Auth_ClientID *string
	// auth for m2m
	AuthM2M AuthM2MConfig
}

type AuthConfig struct {
	ISS      string
	AUD      []string
	ALG      *string
	TTL      *int
	ClientID *string
}

type GraphQLConfig struct {
	ComplexityLimit int `default:"6000"`
}

type AuthConfigs []AuthConfig

type Auth0Config struct {
	Domain       string
	Audience     string
	ClientID     string
	ClientSecret string
	WebClientID  string
}

type SendGridConfig struct {
	Email string
	Name  string
	API   string
}

type SMTPConfig struct {
	Host         string
	Port         string
	SMTPUsername string
	Email        string
	Password     string
}

type GCSConfig struct {
	BucketName              string
	PublicationCacheControl string
}

type AuthM2MConfig struct {
	ISS   string
	AUD   []string
	ALG   *string
	TTL   *int
	Email string
}

func (c Config) Auths() (res AuthConfigs) {
	if ac := c.Auth0.AuthConfig(); ac != nil {
		res = append(res, *ac)
	}
	if c.Auth_ISS != "" {
		var aud []string
		if len(c.Auth_AUD) > 0 {
			aud = append(aud, c.Auth_AUD)
		}
		res = append(res, AuthConfig{
			ISS:      c.Auth_ISS,
			AUD:      aud,
			ALG:      c.Auth_ALG,
			TTL:      c.Auth_TTL,
			ClientID: c.Auth_ClientID,
		})
	}

	return append(res, c.Auth...)
}

func (c Config) JWTProviders() (res []appx.JWTProvider) {
	return c.Auths().JWTProviders()
}

func (c Config) AuthForWeb() *AuthConfig {
	if ac := c.Auth0.AuthConfigForWeb(); ac != nil {
		return ac
	}
	if c.Auth_ISS != "" {
		var aud []string
		if len(c.Auth_AUD) > 0 {
			aud = append(aud, c.Auth_AUD)
		}
		return &AuthConfig{
			ISS:      c.Auth_ISS,
			AUD:      aud,
			ALG:      c.Auth_ALG,
			TTL:      c.Auth_TTL,
			ClientID: c.Auth_ClientID,
		}
	}
	// if ac := c.AuthSrv.AuthConfig(c.Dev, c.Host); ac != nil {
	// 	return ac
	// }
	return nil
}

func (c Auth0Config) AuthConfig() *AuthConfig {
	domain := c.Domain
	if c.Domain == "" {
		return nil
	}
	if !strings.HasPrefix(domain, "https://") && !strings.HasPrefix(domain, "http://") {
		domain = "https://" + domain
	}
	if !strings.HasSuffix(domain, "/") {
		domain = domain + "/"
	}
	aud := []string{}
	if c.Audience != "" {
		aud = append(aud, c.Audience)
	}
	return &AuthConfig{
		ISS: domain,
		AUD: aud,
	}
}

func (c Auth0Config) AuthConfigForWeb() *AuthConfig {
	if c.Domain == "" || c.WebClientID == "" {
		return nil
	}
	domain := prepareUrl(c.Domain)
	var aud []string
	if len(c.Audience) > 0 {
		aud = []string{c.Audience}
	}
	return &AuthConfig{
		ISS:      domain,
		AUD:      aud,
		ClientID: &c.WebClientID,
	}
}

func (a AuthConfig) JWTProvider() appx.JWTProvider {
	return appx.JWTProvider{
		ISS: a.ISS,
		AUD: a.AUD,
		ALG: a.ALG,
		TTL: a.TTL,
	}
}

func (a AuthM2MConfig) JWTProvider() []appx.JWTProvider {
	domain := a.ISS
	if a.ISS == "" {
		return nil
	}
	if !strings.HasPrefix(domain, "https://") && !strings.HasPrefix(domain, "http://") {
		domain = "https://" + domain
	}

	return []appx.JWTProvider{{
		ISS: domain,
		AUD: a.AUD,
		ALG: a.ALG,
		TTL: a.TTL,
	}}
}

// Decode is a custom decoder for AuthConfigs
func (ipd *AuthConfigs) Decode(value string) error {
	if value == "" {
		return nil
	}

	var providers []AuthConfig

	err := json.Unmarshal([]byte(value), &providers)
	if err != nil {
		return fmt.Errorf("invalid identity providers json: %w", err)
	}

	*ipd = providers
	return nil
}

func (a AuthConfigs) JWTProviders() []appx.JWTProvider {
	return lo.Map(a, func(a AuthConfig, _ int) appx.JWTProvider { return a.JWTProvider() })
}

func ReadConfig(debug bool) (*Config, error) {
	// load .env
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, err
	} else if err == nil {
		log.Infof("config: .env loaded")
	}

	var c Config
	err := envconfig.Process(configPrefix, &c)

	if debug {
		c.Dev = true
	}

	return &c, err
}

func (c Config) Print() string {
	s := fmt.Sprintf("%+v", c)
	for _, secret := range []string{c.DB, c.Auth0.ClientSecret} {
		if secret == "" {
			continue
		}
		s = strings.ReplaceAll(s, secret, "***")
	}
	return s
}

func prepareUrl(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}
	url = strings.TrimSuffix(url, "/")
	return url
}
