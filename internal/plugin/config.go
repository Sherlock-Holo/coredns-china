package plugin

type Config struct {
	Forest     *forest
	ChinaDns   string
	ForeignDns string
}

func NewConfig() Config {
	return Config{
		Forest: new(forest),
	}
}
