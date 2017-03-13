package config

import (
	"io"
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Errors 			errors
}

type errors struct {
	MaxStepsPerScenario 			int 	`toml:"max_steps_per_scenario"`
	MaxLenStep 				int	`toml:"max_len_step"`
	EmptyFeature 				bool	`toml:"empty_feature"`
	EmptyScenario 				bool 	`toml:"empty_scenario"`
	AllowedEmptyLinesBetweenSteps		bool	`toml:"allowed_empy_lines_between_steps"`
	Strict					bool
}

func GetConfig(reader io.Reader)(config *Config){
	if _, err := toml.DecodeReader(reader, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

