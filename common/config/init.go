package config

import "github.com/BurntSushi/toml"

var RunData Config

func Init(path string) {
	if _, err := toml.DecodeFile(path, &RunData); err != nil {
		panic(err)
	}

	//RunData = Config{
	//	Addr: Addr{
	//		ApiAddr:    "0.0.0.0:7080",
	//		PublicAddr: "0.0.0.0:7081",
	//	},
	//	DbConfig: DbConfig{
	//		Addr:     "localhost",
	//		Port:     3306,
	//		Username: "root",
	//		Password: "fyU0bj2KLndYN6CKYubb",
	//		DbName:   "xynu_grade",
	//	},
	//	RedisConfig: RedisConfig{
	//		Addr:     "localhost:6379",
	//		Password: "Y5dg5tg8050oigInC30sf",
	//		Db:       2,
	//	},
	//	JwtConfig: JwtConfig{
	//		SignKey: "Tqede8rNvZvbKmjqYpQRvwLxsGR",
	//		Expire:  3600,
	//	},
	//	EmailConfig: EmailConfig{
	//		SMTPAddr: "smtp.feishu.cn:587",
	//		Password: "smcfUO20MyiAEijy",
	//		From:     "grade@xynu.email",
	//		Host:     "smtp.feishu.cn",
	//	},
	//}
}
