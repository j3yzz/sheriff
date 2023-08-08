package kavenegarsvc

type Config struct {
	ApiKey   string `koanf:"api_key" json:"api_key"`
	Template string `koanf:"template" json:"template"`
}
