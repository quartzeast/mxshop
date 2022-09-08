package config

type Mysql struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type Consul struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type Server struct {
	Name       string   `mapstructure:"name" json:"name"`
	Tags       []string `mapstructure:"tags" json:"tags"`
	MysqlInfo  Mysql    `mapstructure:"mysql" json:"mysql"`
	ConsulInfo Consul   `mapstructure:"consul" json:"consul"`
}

type Nacos struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
