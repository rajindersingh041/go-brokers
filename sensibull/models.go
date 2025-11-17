package sensibull

import (
	"time"
)

type Price struct {
    Future            float64 `json:"future"`
    FutureChangePct   float64 `json:"future_change_percent"`
}

type Volume struct {
    Future float64 `json:"future"`
}

type OiOptions struct {
    CallOi float64 `json:"call_oi"`
    PutOi  float64 `json:"put_oi"`
}

type OiFutures struct {
    FuturesOi float64 `json:"futures_oi"`
}

type OiChangeOptions struct {
    CallOiChange float64 `json:"call_oi_change"`
    PutOiChange  float64 `json:"put_oi_change"`
}

type OiChangeFutures struct {
    FutureOiChange float64 `json:"future_oi_change"`
}

type PcrData struct {
    Pcr              float64   `json:"pcr"`
    AutomaticExpiry  string    `json:"automatic_expiry"`
}

type MaxPainData struct {
    MaxPain          float64   `json:"max_pain"`
    AutomaticExpiry  string    `json:"automatic_expiry"`
}

type IV struct {
    AtmIV        float64 `json:"atm_iv"`
    AtmIVChange  float64 `json:"atm_iv_change"`
    AtmStrike    float64 `json:"atm_strike"`
    AtmIVExpiry  string  `json:"atm_iv_expiry"`
}

type IndiaVix struct {
    IndiaVixPrice       float64 `json:"indiavix_price"`
    IndiaVixPriceChange float64 `json:"indiavix_price_change"`
}

type IVP struct {
    IVP       float64 `json:"ivp"`
    AtmStrike float64 `json:"atm_strike"`
    Expiry    string  `json:"expiry"`
}

type RollingAtmStraddle struct {
    AtmStrike float64 `json:"atm_strike"`
    Ltp       float64 `json:"ltp"`
    LtpChange float64 `json:"ltp_change"`
}

type SensibullData struct {
    Timestamp         time.Time                       `json:"-"`
    Nifty             float64                         `json:"nifty"`
    NiftyChangePct    float64                         `json:"nifty_change_percent"`
    Spot              float64                         `json:"spot"`
    SpotChangePct     float64                         `json:"spot_change_percent"`
    Price             Price                           `json:"price"`
    Volume            Volume                          `json:"volume"`
    OiOptions         OiOptions                       `json:"oi_options"`
    OiFutures         OiFutures                       `json:"oi_futures"`
    OiChangeOptions   OiChangeOptions                 `json:"oi_change_options"`
    OiChangeFutures   OiChangeFutures                 `json:"oi_change_futures"`
    PcrData           PcrData                         `json:"pcr_data"`
    MaxPainData       MaxPainData                     `json:"max_pain_data"`
    IV                IV                              `json:"iv"`
    IndiaVix          IndiaVix                        `json:"indiavix"`
    IVP               IVP                             `json:"ivp"`
    RollingAtmStraddle map[string]RollingAtmStraddle  `json:"rolling_atm_straddle"`
}

type ChartData map[string]SensibullData

type Root struct {
    Payload struct {
        ChartData ChartData `json:"chart_data"`
    } `json:"payload"`
}
