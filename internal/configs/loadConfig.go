package configs

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfigs() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	var cfg = Config{}

	// DB
	cfg.Db = DBConfig{
		Driver: os.Getenv("DB_DRIVER"),
		Url:    os.Getenv("DB_URL"),
	}

	// MAILING CONFIGS
	port, err := strconv.Atoi(os.Getenv("MAILER_PORT"))
	if err != nil {
		log.Fatalf("invalid MAILER_PORT: %v", err)
	}
	cfg.Mail = MailConfig{
		Provider: os.Getenv("MAILER_PROVIDER"),
		Username: os.Getenv("MAILER_USER"),
		Password: os.Getenv("MAILER_SECRET"),
		Email:    os.Getenv("MAILER_MAIL"),
		Port:     port,
	}

	// SESSIONS
	sessionEncryptKey, err := parseSessionKey(os.Getenv("SESSION_ENCRYPT_KEY"))
	if err != nil {
		log.Fatalf("Invalid SESSION_ENCRYPT_KEY: %v", err)
	}

	sessionAuthKey, err := parseSessionKey(os.Getenv("SESSION_AUTH_KEY"))
	if err != nil {
		log.Fatalf("Invalid SESSION_AUTH_KEY: %v", err)
	}

	cfg.Session = SessionConfig{
		SessionAuthKey:       sessionAuthKey,
		SessionEncryptionKey: sessionEncryptKey,
	}

	return cfg
}

func parseSessionKey(value string) ([]byte, error) {
	if len(value) == 32 || len(value) == 24 || len(value) == 16 {
		return []byte(value), nil
	}

	key, err := hex.DecodeString(value)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 && len(key) != 24 && len(key) != 16 {
		return nil, fmt.Errorf("must be 16, 24, or 32 bytes after hex decoding; got %d bytes", len(key))
	}

	return key, nil
}
