package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// 数据库
	MySql struct {
		DataSource string
	}
}
