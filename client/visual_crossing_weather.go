package client

type VisualCrossingWeather struct {
	ResolvedAddress string `json:"resolved_address"`
	Days            Days   `json:"days"`
}

type Days struct {
	Temp       float64 `json:"temp"`
	Conditions string  `json:"conditions"`
}
