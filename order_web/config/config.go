package config

type SrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name             string       `mapstructure:"name" json:"name"`
	Host             string       `mapstructure:"host" json:"host"`
	Tags             []string     `mapstructure:"tags" json:"tags"`
	Port             int          `mapstructure:"port" json:"port"`
	GoodsSrvInfo     SrvConfig    `mapstructure:"goods_srv" json:"goods_srv"`
	OrderSrvInfo     SrvConfig    `mapstructure:"order_srv" json:"order_srv"`
	InventorySrvInfo SrvConfig    `mapstructure:"inventory_srv" json:"inventory_srv"`
	JWTInfo          JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo       ConsulConfig `mapstructure:"consul" json:"consul"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}