package sensibull

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type SensibullHandler struct {
	DbService *SensibullDBService
	ApiService *SensibullApiService

	// Cache for option chain data
	cachedData []SensibullData
	cacheTime time.Time
	cacheTTL time.Duration
}

func NewSensibullHandler(DbService *SensibullDBService, ApiService *SensibullApiService) *SensibullHandler {
	h := &SensibullHandler{
		DbService:  DbService,
		ApiService: ApiService,
		cacheTTL:   10 * time.Second, // Set TTL to 10 seconds
	}
	return h
}

func (s *SensibullHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /sensibull/price", s.StorePrice)
	mux.HandleFunc("GET /sensibull/oi", s.StoreOI)
	mux.HandleFunc("GET /sensibull/iv", s.StoreIV)

}

func (s *SensibullHandler) SensibullHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to Sensibull home page")

	resp, err := s.ApiService.FetchSensibullData("NIFTY")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to fetch sensibull data"))
		return
	}


	// Read all data from resp.Body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read response body"))
		return
	}

	// Unmarshal and extract data
	dataSlice, err := UnmarshalChartData(bodyBytes)
	if err != nil {
		panic(err)
	}
	
 	// fmt.Print("Printing data")
	s.DbService.ProcessAndStoreData(dataSlice, "NIFTY")

	w.Header().Set("content","application/json")
	w.Write([]byte("Sensibull Data Processed and stored successfully for"))


}

func (s *SensibullHandler) FetchOptionChain(w http.ResponseWriter, r *http.Request) []SensibullData {
	// Check cache validity
	if time.Since(s.cacheTime) < s.cacheTTL && s.cachedData != nil {
		return s.cachedData
	}

	resp, err := s.ApiService.FetchSensibullData("NIFTY")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to fetch sensibull data"))
		return nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read response body"))
		return nil
	}

	dataSlice, err := UnmarshalChartData(bodyBytes)
	if err != nil {
		panic(err)
	}

	// Update cache
	s.cachedData = dataSlice
	s.cacheTime = time.Now()

	return dataSlice
}

func (s *SensibullHandler) StorePrice(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Println("Welcome to Sensibull home page")

	fetchStart := time.Now()
	dataSlice := s.FetchOptionChain(w, r)
	fetchDuration := time.Since(fetchStart)
	if dataSlice == nil {
		return
	}
	dbStart := time.Now()
	s.DbService.StorePrice(dataSlice, "NIFTY")
	dbDuration := time.Since(dbStart)
	totalDuration := time.Since(start)
	fmt.Printf("[TIMING] FetchOptionChain: %v, StorePrice: %v, Total: %v\n", fetchDuration, dbDuration, totalDuration)

	w.Header().Set("content","application/json")
	w.Write([]byte("Sensibull Data Processed and stored successfully for"))


}

func (s *SensibullHandler) StoreOI(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fetchStart := time.Now()
	dataSlice := s.FetchOptionChain(w, r)
	fetchDuration := time.Since(fetchStart)
	if dataSlice == nil {
		return
	}
	dbStart := time.Now()
	s.DbService.StoreOI(dataSlice, "NIFTY")
	dbDuration := time.Since(dbStart)
	totalDuration := time.Since(start)
	fmt.Printf("[TIMING] FetchOptionChain: %v, StoreOI: %v, Total: %v\n", fetchDuration, dbDuration, totalDuration)

	w.Header().Set("content","application/json")
	w.Write([]byte("Sensibull Data Processed and stored successfully for"))


}

func (s *SensibullHandler) StoreIV(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fetchStart := time.Now()
	dataSlice := s.FetchOptionChain(w, r)
	fetchDuration := time.Since(fetchStart)
	if dataSlice == nil {
		return
	}
	dbStart := time.Now()
	s.DbService.StoreIV(dataSlice, "NIFTY")
	dbDuration := time.Since(dbStart)
	totalDuration := time.Since(start)
	fmt.Printf("[TIMING] FetchOptionChain: %v, StoreIV: %v, Total: %v\n", fetchDuration, dbDuration, totalDuration)

	w.Header().Set("content","application/json")
	w.Write([]byte("Sensibull Data Processed and stored successfully for"))


}