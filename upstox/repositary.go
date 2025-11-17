package upstox

import (
	"database/sql"
	"fmt"
)

type UpstoxRep struct {
	db *sql.DB
}

func NewUpstoxRep(db *sql.DB) UpstoxRep {
	return UpstoxRep{db:db}
}

// type UpstoxRep struct {
	// EnsureTable()
	// StoreInstrumentData(InstrumentData)
	// StoreOHLCData(OHLC)
	// StoreDepthData(Depth)
	// StoreGreeksData(Greeks)

	
// }

// func NewUpstoxRep() UpstoxRep {
// 	return UpstoxRep{}
// }


func (r *UpstoxRep) EnsureIntrumentTable() {
	fmt.Println("Ensuring Table exists..Table created if not exists")
	query := `
	CREATE TABLE IF NOT EXISTS instruments (
			symbol_id           TEXT    NOT NULL,
			timestamp           TIMESTAMP  NOT NULL,
			last_trade_time     TIMESTAMP,
			last_price          DOUBLE PRECISION,
			close_price         DOUBLE PRECISION,
			last_quantity       INTEGER,
			buy_quantity        DOUBLE PRECISION,
			sell_quantity       DOUBLE PRECISION,
			volume              INTEGER,
			average_price       DOUBLE PRECISION,
			oi                  DOUBLE PRECISION,
			poi                 DOUBLE PRECISION,
			oi_day_high         DOUBLE PRECISION,
			oi_day_low          DOUBLE PRECISION,
			net_change          DOUBLE PRECISION,
			lower_circuit_limit DOUBLE PRECISION,
			upper_circuit_limit DOUBLE PRECISION,
			yearly_low                  DOUBLE PRECISION,
			yearly_high                  DOUBLE PRECISION,
			iv                  DOUBLE PRECISION,
			PRIMARY KEY (symbol_id, timestamp)
	);`
	
	_, err := r.db.Exec(query)
	if err != nil {
		// handle error
		fmt.Println("Error creating table:", err)
	}
}

func (r *UpstoxRep) EnsureOHLCTable() {
	fmt.Println("Ensuring Table exists..Table created if not exists")
	query := `
	CREATE TABLE IF NOT EXISTS ohlc (
    symbol_id  TEXT NOT NULL,
    timestamp  TIMESTAMP NOT NULL,
    interval   TEXT,
    open       DOUBLE PRECISION,
    high       DOUBLE PRECISION,
    low        DOUBLE PRECISION,
    close      DOUBLE PRECISION,
    volume     BIGINT,
    PRIMARY KEY (symbol_id, timestamp)
	);`
	
	_, err := r.db.Exec(query)
	if err != nil {
		// handle error
		fmt.Println("Error creating table:", err)
	}
}

func (r *UpstoxRep) EnsureDepthTable() {
	fmt.Println("Ensuring Table exists..Table created if not exists")
}

func (r *UpstoxRep) EnsureGreeksTable() {
	fmt.Println("Ensuring Table exists..Table created if not exists")
	query := `
	CREATE TABLE IF NOT EXISTS option_greeks (
    symbol_id  TEXT NOT NULL,
    timestamp  TIMESTAMP NOT NULL,
	option_price        	DOUBLE PRECISION,
	underlying_price        DOUBLE PRECISION,
	iv        				DOUBLE PRECISION,
	delta     				DOUBLE PRECISION,
	theta     				DOUBLE PRECISION,
	gamma     				DOUBLE PRECISION,
	vega      				DOUBLE PRECISION,
    PRIMARY KEY (symbol_id, timestamp)
	);`
	
	_, err := r.db.Exec(query)
	if err != nil {
		// handle error
		fmt.Println("Error creating table:", err)
	}

}

func (r *UpstoxRep) StoreInstrumentData(data InstrumentData) {
	fmt.Printf("\nStoring Instrument Data for SymbolID: %s with %f \n", data.SymbolID, data.AveragePrice)

	query := `
		INSERT INTO instruments (
			symbol_id, timestamp, last_trade_time, last_price, close_price, last_quantity,
			buy_quantity, sell_quantity, volume, average_price, oi, poi, oi_day_high, oi_day_low,
			net_change, lower_circuit_limit, upper_circuit_limit, yearly_low, yearly_high, iv
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
		)
		ON CONFLICT (symbol_id, timestamp) DO UPDATE SET
			timestamp = EXCLUDED.timestamp,
			last_trade_time = EXCLUDED.last_trade_time,
			last_price = EXCLUDED.last_price,
			close_price = EXCLUDED.close_price,
			last_quantity = EXCLUDED.last_quantity,
			buy_quantity = EXCLUDED.buy_quantity,
			sell_quantity = EXCLUDED.sell_quantity,
			volume = EXCLUDED.volume,
			average_price = EXCLUDED.average_price,
			oi = EXCLUDED.oi,
			poi = EXCLUDED.poi,
			oi_day_high = EXCLUDED.oi_day_high,
			oi_day_low = EXCLUDED.oi_day_low,
			net_change = EXCLUDED.net_change,
			lower_circuit_limit = EXCLUDED.lower_circuit_limit,
			upper_circuit_limit = EXCLUDED.upper_circuit_limit,
			yearly_low = EXCLUDED.yearly_low,
			yearly_high = EXCLUDED.yearly_high,
			iv = EXCLUDED.iv;
`

		_, err := r.db.Exec(query,
			data.SymbolID, data.TimeString, data.LastTradeTime, data.LastPrice, data.ClosePrice, data.LastQuantity,
			data.BuyQuantity, data.SellQuantity, data.Volume, data.AveragePrice, data.Oi, data.Poi, data.OiDayHigh, data.OiDayLow,
			data.NetChange, data.LowerCircuitLimit, data.UpperCircuitLimit, data.Yl, data.Yh, data.Iv,
		)
		if err != nil {
			fmt.Println("Error storing instrument data:", err)
			// handle error
		}

	// fmt.Println("Storing the ohlc data...")
	// fmt.Println("Storing the Greeks...")
	// fmt.Println("Storing the Price info...")
	data.OHLC.SymbolID = data.SymbolID
	data.OHLC.Timestamp = data.TimeString
	data.OptionGreeks.SymbolID = data.SymbolID
	data.OptionGreeks.Timestamp = data.TimeString

	data.Depth.SymbolID = data.SymbolID
	data.Depth.Timestamp = data.TimeString

	r.StoreGreeksData(data.OptionGreeks)
	r.StoreOHLCData(data.OHLC)
}

func (r *UpstoxRep) StoreOHLCData(ohlc OHLC) {
	fmt.Printf("Storing OHLC Data for SymbolID: %s with Close Price: %f\n", ohlc.SymbolID, ohlc.Close)
	r.EnsureOHLCTable()

	query := `
		INSERT INTO ohlc (
			symbol_id, timestamp, interval, open, high, low, close, volume
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
		ON CONFLICT (symbol_id, timestamp) DO UPDATE SET
			interval = EXCLUDED.interval,
			open = EXCLUDED.open,
			high = EXCLUDED.high,
			low = EXCLUDED.low,
			close = EXCLUDED.close,
			volume = EXCLUDED.volume;
		`
		_, err := r.db.Exec(query,
			ohlc.SymbolID, ohlc.Timestamp, ohlc.Interval, ohlc.Open, ohlc.High, ohlc.Low, ohlc.Close, ohlc.Volume,
		)
		if err != nil {
			fmt.Println("Error storing OHLC data:", err)
			// handle error
		}

}

func (r *UpstoxRep) StoreDepthData(depth Depth) {
	fmt.Printf("Storing Depth Data for SymbolID: %s at Timestamp: %s\n", depth.SymbolID, depth.Timestamp)
	r.EnsureDepthTable()
	// Implement depth data storage logic
}

func (r *UpstoxRep) StoreGreeksData(greeks Greeks) {
	fmt.Printf("Storing Greeks Data for SymbolID: %s with IV: %f\n", greeks.SymbolID, greeks.Iv)

	r.EnsureGreeksTable()
	query := `
		INSERT INTO option_greeks (
			symbol_id, timestamp, option_price, underlying_price, iv, delta, theta, gamma, vega
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
		ON CONFLICT (symbol_id, timestamp) DO UPDATE SET
			option_price = EXCLUDED.option_price,
			underlying_price = EXCLUDED.underlying_price,
			iv = EXCLUDED.iv,
			delta = EXCLUDED.delta,
			theta = EXCLUDED.theta,
			gamma = EXCLUDED.gamma,
			vega = EXCLUDED.vega;
		`
		_, err := r.db.Exec(query,
			greeks.SymbolID, greeks.Timestamp, greeks.Op, greeks.Up, greeks.Iv, greeks.Delta, greeks.Theta, greeks.Gamma, greeks.Vega,
		)
		if err != nil {
			fmt.Println("Error storing Greeks data:", err)
			// handle error
		}
}