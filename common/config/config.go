package config

type Config struct {
	Addr        Addr        `yaml:"Addr"`
	Rpc         Rpc         `yaml:"Rpc"`
	DbConfig    DbConfig    `yaml:"DbConfig"`
	RedisConfig RedisConfig `yaml:"RedisConfig"`
	JwtConfig   JwtConfig   `yaml:"JwtConfig"`
	MqUrl       string      `yaml:"MqUrl"`
	Version     string      `yaml:"Version"`
	CosUrl      string      `yaml:"cosUrl"`
	UploadPath  string      `yaml:"uploadPath"`
	Minio       Minio       `yaml:"Minio"`
	MongodbUrl  string      `yaml:"MongodbUrl"`
}

type Addr struct {
	ApiAddr       string `yaml:"ApiAddr"`
	UserAddr      string `yaml:"UserAddr"`
	ImAddr        string `yaml:"ImAddr"`        // ws 地址
	ImServerAddr  string `yaml:"ImServerAddr"`  // 调用ws rpc 地址
	ImHandlerAddr string `yaml:"ImHandlerAddr"` //处理ws消息的rpc地址
}

type Rpc struct {
	UserRpc       string `yaml:"UserRpc"`
	ImServerAddr  string `yaml:"ImServerAddr"`
	ImHandlerAddr string `yaml:"ImHandlerAddr"`
}

type DbConfig struct {
	DbType   string `yaml:"DbType"`
	Addr     string `yaml:"Addr"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
}

type RedisConfig struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
	Prefix   string `yaml:"Prefix"`
}

type JwtConfig struct {
	SignKey string `yaml:"SignKey"`
	Expire  int64  `yaml:"Expire"`
}

type Minio struct {
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyID     string `yaml:"AccessKeyID"`
	SecretAccessKey string `yaml:"SecretAccessKey"`
	UseSSL          bool   `yaml:"UseSSL"`
	Bucket          string `yaml:"Bucket"`
}

func GetVersion() string {
	return RunData.Version

}
