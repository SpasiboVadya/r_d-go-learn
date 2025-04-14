package main

type BaseConfig struct {
	Timeout int
	Retries int
	BaseURL string
}

type HTTPClientConfig struct {
	BaseConfig        // Embedded struct
	BaseURL    string // Custom field
}

func NewHTTPClientConfig() HTTPClientConfig {
	return HTTPClientConfig{
		BaseConfig: BaseConfig{
			Timeout: 10,
			Retries: 3,
			BaseURL: "https://www.amazon.com",
		},
		BaseURL: "https://www.google.com",
	}
}

//func (h *HTTPClientConfig) GetTimeout() time.Duration {
//	return time.Duration(h.BaseConfig.BaseURL) * time.Second
//}
