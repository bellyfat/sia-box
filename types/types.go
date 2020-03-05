package types

type (
	Config struct {
		Path     string `yaml:"path"`
		Password string `yaml:"password"`
	}

	StoreState struct {
		Files map[string]string
	}
)
