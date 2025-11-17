package sensibull

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type SensibullDBRepo struct {
	db *sql.DB
}

func NewSensibullDBRepo(db *sql.DB) *SensibullDBRepo {
	return &SensibullDBRepo{db}
}


func (s *SensibullDBRepo) EnsureSensibullPriceTable() {
	query := ` 
	CREATE TABLE IF NOT EXISTS sensibull_price (
    symbol_id text,
    timestamp TIMESTAMPTZ NOT NULL,
    nifty DOUBLE PRECISION,
    nifty_change_pct DOUBLE PRECISION,
    spot DOUBLE PRECISION,
    spot_change_pct DOUBLE PRECISION,

    -- Price
    price_future DOUBLE PRECISION,
    price_future_change_pct DOUBLE PRECISION,

    -- Volume
    volume_future DOUBLE PRECISION,

    PRIMARY KEY (symbol_id, timestamp)
    )
    ;`

	_, err := s.db.Exec(query)
	if err != nil {
	
	fmt.Println("Error creating table:", err)
	}

}



func (s *SensibullDBRepo) InsertSensibullPriceTable(DataList []SensibullData, symbolID string) error {
	query := ` 
	INSERT INTO sensibull_price (
	symbol_id,
	timestamp,
	
	nifty,
	nifty_change_pct,
	spot,
	spot_change_pct,

	-- Price
	price_future,
	price_future_change_pct,

	-- Volume
	volume_future
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9)

	ON CONFLICT (symbol_id, timestamp) DO NOTHING 

    ;`

	// Batch insert for PostgreSQL: much faster than per-row inserts because it reduces round-trips and transaction overhead.
	// We build a single INSERT statement with multiple VALUES rows.
	if len(DataList) == 0 {
		return nil
	}
	totalStart := time.Now()
	baseQuery := `INSERT INTO sensibull_price (
		symbol_id, timestamp, nifty, nifty_change_pct, spot, spot_change_pct, price_future, price_future_change_pct, volume_future
	) VALUES `
	valueStrings := make([]string, 0, len(DataList))
	valueArgs := make([]interface{}, 0, len(DataList)*9)
	for i, data := range DataList {
		n := i*9
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			n+1, n+2, n+3, n+4, n+5, n+6, n+7, n+8, n+9))
		valueArgs = append(valueArgs,
			symbolID,
			data.Timestamp,
			data.Nifty,
			data.NiftyChangePct,
			data.Spot,
			data.SpotChangePct,
			data.Price.Future,
			data.Price.FutureChangePct,
			data.Volume.Future,
		)
	}
	query = baseQuery + strings.Join(valueStrings, ",") + ` ON CONFLICT (symbol_id, timestamp) DO NOTHING;`
	_, err := s.db.Exec(query, valueArgs...)
	if err != nil {
		fmt.Println("Error in batch insert:", err)
		return err
	}
	totalDuration := time.Since(totalStart)
	fmt.Printf("[TIMING] InsertSensibullPriceTable batch total: %v for %d rows\n", totalDuration, len(DataList))
	return nil
}


func (s *SensibullDBRepo) EnsureSensibullOITable() {
	query := ` 
	CREATE TABLE IF NOT EXISTS sensibull_oi (
    symbol_id text,
    timestamp TIMESTAMPTZ NOT NULL,
    call_oi DOUBLE PRECISION,
    put_oi DOUBLE PRECISION,
    futures_oi DOUBLE PRECISION,
    call_oi_change DOUBLE PRECISION,
    put_oi_change DOUBLE PRECISION,
    future_oi_change DOUBLE PRECISION,
    PRIMARY KEY (symbol_id, timestamp)
    )
    ;`

	_, err := s.db.Exec(query)
	if err != nil {
	
	fmt.Println("Error creating table:", err)
	}

}



func (s *SensibullDBRepo) InsertSensibullOITable(DataList []SensibullData, symbolID string) error {
	query := ` 
	INSERT INTO sensibull_oi (
    symbol_id,
    timestamp,
	call_oi,
	put_oi,
	futures_oi,
	call_oi_change,
	put_oi_change,
	future_oi_change)
	vALUES (
	$1, $2, $3, $4, $5, $6, $7, $8
    )
	ON CONFLICT (symbol_id, timestamp) DO NOTHING 

    ;`

	if len(DataList) == 0 {
		return nil
	}
	totalStart := time.Now()
	// Build the batch insert statement
	baseQuery := `INSERT INTO sensibull_oi (
		symbol_id, timestamp, call_oi, put_oi, futures_oi, call_oi_change, put_oi_change, future_oi_change
	) VALUES `
	valueStrings := make([]string, 0, len(DataList))
	valueArgs := make([]interface{}, 0, len(DataList)*8)
	for i, data := range DataList {
		n := i*8
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			n+1, n+2, n+3, n+4, n+5, n+6, n+7, n+8))
		valueArgs = append(valueArgs,
			symbolID,
			data.Timestamp,
			data.OiOptions.CallOi,
			data.OiOptions.PutOi,
			data.OiFutures.FuturesOi,
			data.OiChangeOptions.CallOiChange,
			data.OiChangeOptions.PutOiChange,
			data.OiChangeFutures.FutureOiChange,
		)
	}
	query = baseQuery + strings.Join(valueStrings, ",") + ` ON CONFLICT (symbol_id, timestamp) DO NOTHING;`
	_, err := s.db.Exec(query, valueArgs...)
	if err != nil {
		fmt.Println("Error in batch insert:", err)
		return err
	}
	totalDuration := time.Since(totalStart)
	fmt.Printf("[TIMING] InsertSensibullOITable batch total: %v for %d rows\n", totalDuration, len(DataList))
	return nil
}




func (s *SensibullDBRepo) EnsureSensibullMiscTable() {
	query := ` 
	CREATE TABLE IF NOT EXISTS sensibull_misc (
    symbol_id text,
    timestamp TIMESTAMPTZ NOT NULL,
    pcr DOUBLE PRECISION,
    pcr_automatic_expiry DATE,
    max_pain DOUBLE PRECISION,
    max_pain_automatic_expiry DATE,

    PRIMARY KEY (symbol_id, timestamp)
    )
    ;`

	_, err := s.db.Exec(query)
	if err != nil {
	
	fmt.Println("Error creating table:", err)
	}
}


func (s *SensibullDBRepo) EnsureSensibullIVTable() {
	query := ` 
	CREATE TABLE IF NOT EXISTS sensibull_iv (
    symbol_id text,
    timestamp TIMESTAMPTZ NOT NULL,
    atm DOUBLE PRECISION,
    atm_change DOUBLE PRECISION,
    atm_strike DOUBLE PRECISION,
    atm_expiry DATE,
    ivp_atm DOUBLE PRECISION,
    ivp_atm_strike DOUBLE PRECISION,
    ivp_atm_expiry DATE,

    PRIMARY KEY (symbol_id, timestamp)
    )
    ;`

	_, err := s.db.Exec(query)
	if err != nil {
	
	fmt.Println("Error creating table:", err)
	}

}



func (s *SensibullDBRepo) InsertSensibullIVTable(DataList []SensibullData, symbolID string) error {
	query := ` 
	INSERT INTO sensibull_iv (
    symbol_id,
    timestamp,
    atm,
    atm_change,
    atm_strike,
    atm_expiry,
    ivp_atm,
    ivp_atm_strike,
    ivp_atm_expiry
	
	)
	vALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9
    )
	ON CONFLICT (symbol_id, timestamp) DO NOTHING 

    ;`
	// Batch insert for PostgreSQL: much faster than per-row inserts because it reduces round-trips and transaction overhead.
	// We build a single INSERT statement with multiple VALUES rows.
	if len(DataList) == 0 {
		return nil
	}
	totalStart := time.Now()
	baseQuery := `INSERT INTO sensibull_iv (
		symbol_id, timestamp, atm, atm_change, atm_strike, atm_expiry, ivp_atm, ivp_atm_strike, ivp_atm_expiry
	) VALUES `
	valueStrings := make([]string, 0, len(DataList))
	valueArgs := make([]interface{}, 0, len(DataList)*9)
	for i, data := range DataList {
		n := i*9
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			n+1, n+2, n+3, n+4, n+5, n+6, n+7, n+8, n+9))
		valueArgs = append(valueArgs,
			symbolID,
			data.Timestamp,
			data.IV.AtmIV,
			data.IV.AtmIVChange,
			data.IV.AtmStrike,
			data.IV.AtmIVExpiry,
			data.IVP.IVP,
			data.IVP.AtmStrike,
			data.IVP.Expiry,
		)
	}
	query = baseQuery + strings.Join(valueStrings, ",") + ` ON CONFLICT (symbol_id, timestamp) DO NOTHING;`
	_, err := s.db.Exec(query, valueArgs...)
	if err != nil {
		fmt.Println("Error in batch insert:", err)
		return err
	}
	totalDuration := time.Since(totalStart)
	fmt.Printf("[TIMING] InsertSensibullIVTable batch total: %v for %d rows\n", totalDuration, len(DataList))
	return nil
}




func (s *SensibullDBRepo) EnsureSensibullRollingStraddleTable() {
	query := ` 
	CREATE TABLE IF NOT EXISTS sensibull_rolling_straddle (
    symbol_id text,
    timestamp TIMESTAMPTZ NOT NULL,
    strike DOUBLE PRECISION,
    ltp DOUBLE PRECISION,
    ltp_change DOUBLE PRECISION,

    PRIMARY KEY (symbol_id, timestamp)
    )
    ;`

	_, err := s.db.Exec(query)
	if err != nil {
	
	fmt.Println("Error creating table:", err)
	}

}




func (s *SensibullDBRepo) InsertSensibullRollingStraddleTable(DataList []SensibullData, symbolID string) error {
	query := ` 
	INSERT INTO sensibull_rolling_straddle (
    symbol_id,
    timestamp,
    strike,
    ltp,
    ltp_change)
	vALUES (
	$1, $2, $3, $4, $5
	)
	ON CONFLICT (symbol_id, timestamp) DO NOTHING 

	;`

    tx, err := s.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

	for _, data := range DataList {
		for _, ras := range data.RollingAtmStraddle {
		_, err := tx.Exec(query,symbolID,data.Timestamp,ras.AtmStrike, ras.Ltp, ras.LtpChange)
		if err != nil {
			fmt.Println("Error creating table:", err)
			}
		}
		}
	return tx.Commit()
}


func (s *SensibullDBRepo) InsertManySensibullData(symbolID string, dataList []SensibullData) error {
	// s.EnsureSensibullIVTable()
	// s.InsertSensibullIVTable(dataList, symbolID)

	s.EnsureSensibullOITable()
	s.InsertSensibullOITable(dataList, symbolID)

	// s.EnsureSensibullPriceTable()
	// s.InsertSensibullPriceTable(dataList, symbolID)

	// s.EnsureSensibullRollingStraddleTable()
	// s.InsertSensibullRollingStraddleTable(dataList, symbolID)

	return nil
}

func (s *SensibullDBRepo) StorePrice(symbolID string, dataList []SensibullData) error {
	s.EnsureSensibullPriceTable()
	s.InsertSensibullPriceTable(dataList, symbolID)
	return nil
}


func (s *SensibullDBRepo) StoreOI(symbolID string, dataList []SensibullData) error {
	s.EnsureSensibullOITable()
	s.InsertSensibullOITable(dataList, symbolID)
	return nil
}

func (s *SensibullDBRepo) StoreRollingStraddle(symbolID string, dataList []SensibullData) error {
	s.EnsureSensibullRollingStraddleTable()
	s.InsertSensibullRollingStraddleTable(dataList, symbolID)
	return nil
}

func (s *SensibullDBRepo) StoreIV(symbolID string, dataList []SensibullData) error {
	s.EnsureSensibullIVTable()
	s.InsertSensibullIVTable(dataList, symbolID)
	return nil
}