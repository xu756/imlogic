package config

type Config struct {
	Addr        Addr
	Rpc         Rpc
	DbConfig    DbConfig
	RedisConfig RedisConfig
	JwtConfig   JwtConfig
	EmailConfig EmailConfig
	Etcd        Etcd
}

type Addr struct {
	ApiAddr    string
	PublicAddr string
	ImAddr     string
}

type Rpc struct {
	PublicRpc string
}

type DbConfig struct {
	DbType   string
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

type Etcd struct {
	Addrs []string
}
