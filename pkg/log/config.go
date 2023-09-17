package log

type Config struct {
	Code              string     `yaml:"code"`
	Level             string     `yaml:"level"`
	EnableCaller      bool       `yaml:"enableCaller"`
	EnablePrettyPrint bool       `yaml:"enablePrettyPrint"`
	SampleRate        uint       `yaml:"sampleRate"`
	FileConfig        FileConfig `yaml:"fileConfig"`
}

type FileConfig struct {
	Enabled      bool   `json:"enabled"`
	Directory    string `yaml:"directory"`
	Filename     string `yaml:"filename"`
	MaxSizeInMb  int    `yaml:"maxSizeInMb"`
	MaxBackups   int    `yaml:"maxBackups"`
	MaxAgeInDays int    `yaml:"maxAgeInDays"`
}
