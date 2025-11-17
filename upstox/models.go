package upstox

type Response struct {
	Data struct {
		RequestID    string                    `json:"request_id"`
		TimeInMillis int64                     `json:"time_in_millis"`
		TokenData    map[string]InstrumentData `json:"token_data"`
	} `json:"data"`
}
type OHLC struct {
	SymbolID  string
	Timestamp string
	Interval  string  `json:"interval"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    int     `json:"volume"`
}

type DepthLevel struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Orders   int     `json:"orders"`
}

type Depth struct {
	SymbolID  string
	Timestamp string
	Sell      []DepthLevel `json:"sell"`
	Buy       []DepthLevel `json:"buy"`
}

type Greeks struct {
	SymbolID  string
	Timestamp string
	Op        float64 `json:"op"`
	Up        float64 `json:"up"`
	Iv        float64 `json:"iv"`
	Delta     float64 `json:"delta"`
	Theta     float64 `json:"theta"`
	Gamma     float64 `json:"gamma"`
	Vega      float64 `json:"vega"`
}

type InstrumentData struct {
	SymbolID          string  // `json:"symbolid"`
	Timestamp         int64   `json:"ts"`
	TimeString        string  `json:"timestamp"`
	LastTradeTime     string  `json:"lastTradeTime"`
	LastPrice         float64 `json:"lastPrice"`
	ClosePrice        float64 `json:"closePrice"`
	LastQuantity      int     `json:"lastQuantity"`
	BuyQuantity       float64 `json:"buyQuantity"`
	SellQuantity      float64 `json:"sellQuantity"`
	Volume            int     `json:"volume"`
	AveragePrice      float64 `json:"averagePrice"`
	Oi                float64 `json:"oi"`
	Poi               float64 `json:"poi"`
	OiDayHigh         float64 `json:"oiDayHigh"`
	OiDayLow          float64 `json:"oiDayLow"`
	NetChange         float64 `json:"netChange"`
	LowerCircuitLimit float64 `json:"lowerCircuitLimit"`
	UpperCircuitLimit float64 `json:"upperCircuitLimit"`
	Yl                float64 `json:"yl"`
	Yh                float64 `json:"yh"`
	Iv                float64 `json:"iv"`
	OHLC              OHLC    `json:"ohlc"`
	Depth             Depth   `json:"depth"`
	OptionGreeks      Greeks  `json:"optionGreeks"`
}