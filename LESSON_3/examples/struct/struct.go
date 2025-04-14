package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/push"
	"net/http"
	"time"
)

var jsonConfig = `{
  "baseURL": "https://api.example.com",
  "apiKey": "your-api-key-here",
  "timeout": 0,
  "isProd": false,
  "retryDelay": 500000000,
  "cacheSize": 1000,
  "isHttps": true
}`

type Config struct {
	http.Client
	BaseURL    string        `json:"baseURL,omitempty" yaml:"baseURL,omitempty" env:"baseURL" envDefault:"https://api.example.com"`
	APIKey     string        `json:"apiKey,omitempty" yaml:"apiKey,omitempty" env:"apiKey" envDefault:"your-api-key-here"`
	Timeout    time.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty" env:"timeout" envDefault:"5s"`
	RetryDelay time.Duration `json:"retryDelay,omitempty" yaml:"retryDelay,omitempty" env:"retryDelay" envDefault:"5s"`
	CacheSize  int           `json:"cacheSize,omitempty" yaml:"cacheSize,omitempty" env:"cacheSize" envDefault:"1000"`
	IsHttps    bool          `json:"isHttps,omitempty" yaml:"isHttps,omitempty" env:"isHttps" envDefault:"true"`
	IsProd     bool          `json:"isProd,omitempty" yaml:"isProd,omitempty" env:"isProd" envDefault:"true"`
}

func NewConfig(h push.HTTPDoer) int {
	return 0
}

func loadConfig() *Config {
	return &Config{
		BaseURL: "https://aoi.example.com",
		APIKey:  "123456",
	}
}

func doWithConfig(config *Config) {
	fmt.Printf("Config: %+v", config)
}

type Point struct {
	Z *int8
}

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

func main() {
	d := DocumentField{
		Type:  DocumentFieldTypeBool,
		Value: true,
	}
	if bVal, ok := d.Value.(bool); ok {
		fmt.Println(bVal)
	}

	//c := &Config{}
	//NewConfig(c)
	//p := Point{}
	//x := int8(1)
	//p.Z = &x
	////fmt.Println(&p.X)
	////fmt.Println(&p.Y)
	////fmt.Println(&p.Axis)
	//fmt.Println(p.Z)
	//fmt.Println(unsafe.Sizeof(p))
	////config := loadConfig()
	////doWithConfig(config)
	//
	//var config Config
	//err := json.Unmarshal([]byte(jsonConfig), &config)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Config: %+v\n", config)
	//_, err = json.Marshal(config)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(jsConfig))

	//var cfg Config

	//fmt.Printf("%+v\n", NewHTTPClientConfig().BaseConfig.BaseURL)
}
