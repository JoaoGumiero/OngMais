package config

import (
	"os"
	// this will automatically load your .env file:
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Firebase          FirebaseConfig
	FirebaseProjectID string
}

type FirebaseConfig struct {
	Type            string
	Project_id      string
	Private_key_id  string
	Private_key     string
	Client_email    string
	Client_id       string
	Auth_uri        string
	Token_uri       string
	Auth_provider   string
	Client_x509     string
	Universe_domain string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		FirebaseProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
		Firebase: FirebaseConfig{
			Type:            os.Getenv("type"),
			Project_id:      os.Getenv("project_id"),
			Private_key_id:  os.Getenv("private_key_id"),
			Private_key:     os.Getenv("private_key"),
			Client_email:    os.Getenv("client_email"),
			Client_id:       os.Getenv("client_id"),
			Auth_uri:        os.Getenv("auth_uri"),
			Token_uri:       os.Getenv("token_uri"),
			Auth_provider:   os.Getenv("auth_provider_x509_cert_url"),
			Client_x509:     os.Getenv("client_x509_cert_url"),
			Universe_domain: os.Getenv("universe_domain"),
		},
	}
	return cfg, nil
}
