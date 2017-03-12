package config

import (
	"log"
	"github.com/BurntSushi/toml"
	"io"
)

type Config struct {
	Errors 			errors
	Warnings 		warnings
}

type errors struct {
	MaxStepsPerScenario 			int 	`toml:"max_steps_per_scenario"`
	MaxLenStep 				int	`toml:"max_len_step"`
	EmptyFeature 				bool	`toml:"empty_feature"`
	EmptyScenario 				bool 	`toml:"empty_scenario"`
	Strict					bool
}

type warnings struct {
	AllowedEmptyLinesBetweenSteps		bool	`toml:"allowed_empy_lines_between_steps"`
}

func GetConfig(reader io.Reader)(config *Config){
	if _, err := toml.DecodeReader(reader, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

