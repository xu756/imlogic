package config

type Config struct {
	Addr        Addr
	DbConfig    DbConfig
	RedisConfig RedisConfig
	JwtConfig   JwtConfig
	EmailConfig EmailConfig
}

type Addr struct {
	ApiAddr    string
	PublicAddr string
}

type DbConfig struct {
	Addr     string
	Port     int
	Username string
	Password string
	DbName   string
}

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

type JwtConfig struct {
	SignKey string
	Expire  int64
}

type EmailConfig struct {
	SMTPAddr string
	Password string
	From     string
	Host     string
}
