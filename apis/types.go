package apis

import "fmt"

type openapiDesc struct {
	ConfigUrl         string    `json:"configUrl"`
	Oauth2RedirectUrl string    `json:"oauth2RedirectUrl"`
	ValidatorUrl      string    `json:"validatorUrl"`
	Urls              []urlInfo `json:"urls"`
}

type urlInfo struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type routes struct {
	GlobalEndpoint string `json:"global-endpoint" yaml:"global-endpoint" mapstructure:"global-endpoint"`
	Urls           []struct {
		Url      string `json:"url" yaml:"url" mapstructure:"url"`
		Name     string `json:"name" yaml:"name" mapstructure:"name"`
		Endpoint string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
	} `json:"urls" yaml:"urls" mapstructure:"urls"`
}

func (Self *routes) getUrlInfo() []urlInfo {
	if len(Self.Urls) == 0 {
		return []urlInfo{}
	}
	urls := make([]urlInfo, 0)
	for _, value := range Self.Urls {
		if value.Endpoint == "" {
			value.Endpoint = Self.GlobalEndpoint
		}
		info := urlInfo{
			Name: value.Name,
			Url:  fmt.Sprintf("%s%s", value.Url, value.Endpoint),
		}
		urls = append(urls, info)
	}
	return urls
}
