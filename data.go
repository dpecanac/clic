package main

const (
	DEFAULT_DATA_FILE = "clic.yaml"
)

type Config struct {
	File    string
	Name    string
	Tags    string
	Verbose bool
}

// Data holds global params with list of all requests
type Data struct {
	BaseURL  *string   `yaml:"base_url"`
	Headers  *[]string `yaml:"headers"`
	Requests []SR      `yaml:"requests"`
}

// SR holds single request params
type SR struct {
	Name        *string   `yaml:"name"`
	Description *string   `yaml:"description"`
	Tags        *[]string `yaml:"tags"`
	Method      string    `yaml:"method"`
	Endpoint    *string   `yaml:"endpoint"`
	Body        *string   `yaml:"body"`
	Headers     *[]string `yaml:"headers"`
}

type RQ struct {
	Name    *string
	URL     string
	Headers map[string]string
	Method  string
	Body    *string
}
