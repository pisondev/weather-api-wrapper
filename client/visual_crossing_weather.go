package client

type VisualCrossingWeather struct {
	ResolvedAddress  string      `json:"resolvedAddress"`
	CurentConditions CurrentCond `json:"currentConditions"`
}

type CurrentCond struct {
	Temp       float64 `json:"temp"`
	Conditions string  `json:"conditions"`
}
