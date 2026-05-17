package configs

type Config struct {
	Mail     MailConfig
	Session  SessionConfig
	Db       DBConfig
	Platform string
}

type MailConfig struct {
	Provider string
	Username string
	Password string
	Email    string
	Port     int
}

type SessionConfig struct {
	SessionAuthKey       []byte
	SessionEncryptionKey []byte
}

type DBConfig struct {
	Driver string
	Url    string
}
