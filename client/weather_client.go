package client

type WeatherClient struct {
	Location string `json:"location"`
	Days     Days   `json:"days"`
}

type Days struct {
	Temp       float64 `json:"temp"`
	Conditions string  `json:"conditions"`
}
