package config

type Config struct {
	Addr        Addr
	Rpc         Rpc
	DbConfig    DbConfig
	RedisConfig RedisConfig
	JwtConfig   JwtConfig
	MqUrl       string
	EmailConfig EmailConfig
	MinioConfig MinioConfig
}

// Addr 服务地址 运行地址
type Addr struct {
	ApiAddr      string
	PublicAddr   string
	ImAddr       string // ws 地址
	ImServerAddr string // 调用ws rpc 地址
	ImRpcAddr    string //处理ws消息的rpc地址
}

// Rpc 服务发现 调用的地址
type Rpc struct {
	PublicRpc    string
	ImServerAddr string
	ImRpcAddr    string
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

// MinioConfig  配置
type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}
