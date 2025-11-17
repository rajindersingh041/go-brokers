package main

import (
	"fmt"
	"net/http"

	"github.com/rajindersingh041/go-brokers/database"
	"github.com/rajindersingh041/go-brokers/sensibull"
)

func main() {
    // Read JSON from local file
    mux := http.NewServeMux()
    fmt.Print("Listen and Serve to 8080 port")
    db, _ := database.InitDB()

    sensibullAppContainer := sensibull.NewSensibullAppContainer(db)
    sensibullAppContainer.GetSensibullHandler().RegisterRoutes(mux)
    http.ListenAndServe(":8080", mux)

}


// func main() {
//     // Read JSON from local file
// 	fmt.Print("Reading JSON file")
//     jsonBytes, err := os.ReadFile("testdata.json")
//     if err != nil {
// 		fmt.Print("e1")

//         panic(err)
//     }

//     // Unmarshal and extract data
//     dataSlice, err := sensibull.UnmarshalChartData(jsonBytes)
//     if err != nil {
//         panic(err)
//     }

// 	fmt.Print("Printing data")
//     // Print results
//     for _, d := range dataSlice {
//         for key, ras := range d.RollingAtmStraddle {
//             fmt.Printf("Timestamp: %v, Nifty: %v, Key: %v, AtmStrike: %v\n", d.Timestamp, d.Nifty, key, ras.AtmStrike)
//         }
//     }

//     mux := http.NewServeMux()
//     fmt.Print("Listen and Serve to 8080 port")
//     db, _ := database.InitDB()
//     // repo := sensibull.NewSensibullDBRepo(db)
//     // SenService := sensibull.NewSensibullDBService(repo)

//     // sensibull.SensibullDBService.ProcessAndStoreData(SenService, dataSlice, repo, "TEST_SYMBOL")
//     sensibullAppContainer := sensibull.NewSensibullAppContainer(db)
//     sensibullAppContainer.GetUpstoxHandler().RegisterRoutes(mux)
//     http.ListenAndServe(":8080", mux)

// }
// func main(){
// 	fmt.Print("Setting up mux")

// 	mux := http.NewServeMux()
// 	fmt.Print("Listen and Serve to 8080 port")

// 	db, err := database.InitDB()
// 	if err != nil {

// 		panic("Failed to connect to database")
		
// 	}

// 	defer db.Close()
// 	NewUpstoxAppContainer := upstox.NewUpstoxAppContainer(db)
// 	NewUpstoxAppContainer.GetUpstoxHandler().RegisterRoutes(mux)
// 	// repo := upstox.NewUpstoxRep(db)
// 	// Service := upstox.NewUpstoxService(repo)
// 	// Api := upstox.NewUpstoxApiService()
// 	// UpstoxHandler := upstox.NewHandler(Service, Api)
// 	// UpstoxHandler.RegisterRoutes(mux)
// 	http.ListenAndServe(":8080", mux)

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )


// var rawJSON = []byte(`
// {
//     "data": {
//         "request_id": "WSIM-35aab43d-a839-499f-acfd-5f9471e8f69a",
//         "time_in_millis": 1763070026523,
//         "token_data": {
//             "NSE_FO|48599": {
//                 "timestamp": "2025-11-14 03:10:26",
//                 "lastTradeTime": "2025-11-13 03:29:57",
//                 "lastPrice": 2775.5,
//                 "closePrice": 2788.7,
//                 "lastQuantity": 375,
//                 "buyQuantity": 109875.0,
//                 "sellQuantity": 307875.0,
//                 "volume": 3264750,
//                 "averagePrice": 2795.9,
//                 "oi": 1.1172E7,
//                 "poi": 1.134975E7,
//                 "oiDayHigh": 1.150275E7,
//                 "oiDayLow": 1.1121E7,
//                 "netChange": -13.2,
//                 "lowerCircuitLimit": 2509.9,
//                 "upperCircuitLimit": 3067.5,
//                 "yl": 0.0,
//                 "yh": 0.0,
//                 "ohlc": {
//                     "interval": "1d",
//                     "open": 2788.7,
//                     "high": 2844.5,
//                     "low": 2752.3,
//                     "close": 2788.7,
//                     "volume": 3264750,
//                     "ts": 1762972200000
//                 },
//                 "depth": {
//                     "sell": [
//                         {
//                             "quantity": 375,
//                             "price": 2776.9,
//                             "orders": 1
//                         },
//                         {
//                             "quantity": 375,
//                             "price": 2777.0,
//                             "orders": 1
//                         },
//                         {
//                             "quantity": 375,
//                             "price": 2777.7,
//                             "orders": 1
//                         },
//                         {
//                             "quantity": 3750,
//                             "price": 2777.9,
//                             "orders": 2
//                         },
//                         {
//                             "quantity": 750,
//                             "price": 2778.0,
//                             "orders": 2
//                         }
//                     ],
//                     "buy": [
//                         {
//                             "quantity": 3000,
//                             "price": 2775.0,
//                             "orders": 5
//                         },
//                         {
//                             "quantity": 375,
//                             "price": 2774.3,
//                             "orders": 1
//                         },
//                         {
//                             "quantity": 375,
//                             "price": 2774.1,
//                             "orders": 1
//                         },
//                         {
//                             "quantity": 375,
//                             "price": 2774.0,
//                             "orders": 1
//                         },
//                         {
//                             "quantity": 375,
//                             "price": 2773.7,
//                             "orders": 1
//                         }
//                     ]
//                 }
//             }
//         }
//     },
//     "success": true
// }
// `)


// func main() {
//     var resp Response

//     if err := json.Unmarshal(rawJSON, &resp); err != nil {
//         panic(err)
//     }

//     fmt.Println("Request ID:", resp.Data.RequestID)
//     fmt.Println("Time:", resp.Data.TimeInMillis)

//     fmt.Println("\n--- Token Data ---")
//     for token, info := range resp.Data.TokenData {
// 		info.SymbolID = token
//         fmt.Println("Token:", token)
//         // fmt.Println("  Timestamp:", info.Timestamp)
//         // fmt.Println("  LastTradeTime:", info.LastTradeTime)
//         // fmt.Println("  LastPrice:", info.LastPrice)
//         // fmt.Println(" Store SymbolID:", info.SymbolID)
// 		fmt.Println(info.Depth)
//     }

//     // Access a specific token safely
//     if info, ok := resp.Data.TokenData["BSE_FO|1127618"]; ok {
//         fmt.Println("Specific LastPrice:", info.LastPrice)
//     }
// }
