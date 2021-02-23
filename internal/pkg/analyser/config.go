package analyser

// NewConfig ...
func NewConfig(skills ...[]string) *Config {
	s := make([][]string, 0)
	for _, skill := range skills {
		s = append(s, skill)
	}

	return &Config{
		Skills: s,
	}
}

// Config ...
type Config struct {
	Skills [][]string
}
