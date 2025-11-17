package upstox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// this is client will use

// interface which client will work with
type Handler struct {
	service UpstoxService
	api UpstoxApiService
}


// function which will return pointer to Handler struct
func NewHandler(Service UpstoxService, api UpstoxApiService) *Handler {
	return &Handler{service: Service, api:api}
}


// this function add the functionality to Handler to register the Routes
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /upstox", h.GetHomePage)
	mux.HandleFunc("GET /upstox/ohlc", h.GetOHLC)
	mux.HandleFunc("GET /upstox/clients/{id}", h.UpstoxClients)
}

func (h *Handler) GetOHLC(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "OHLC is xxxxxxxx")
	w.Header().Set("content-type","application/json")
	// w.Write([]byte("OHCL is xxxxxxxxxxxxx"))

	// instrument :=  r.PathValue("i")
	instrument := r.URL.Query().Get("i")
	fmt.Println("Instrument is ", instrument)
	resp, err := h.api.FetchOHLC(instrument)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// repo := NewUpstoxRep()
	// service := NewUpstoxService(repo)
	var upstoxResponse Response
	err = json.NewDecoder(resp.Body).Decode(&upstoxResponse)
	if err != nil {
		http.Error(w, "Failed to decode OHLC data", http.StatusInternalServerError)
		return

	}

	h.service.StoreInstrumentData(upstoxResponse)
	fmt.Println("OHLC Data fetched successfully")

	w.WriteHeader(resp.StatusCode)
	jsonBytes, err := json.Marshal(upstoxResponse)
	if err != nil {
		http.Error(w, "Failed to encode OHLC data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)
       // defer resp.Body.Close()
}




func (h *Handler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Welcome to the home page of upstox")
	w.Header().Set("content","application/json")
	// w.Write([]byte("Welcome to Upstox Home page"))


	user := r.URL.Query().Get("user")
	if user != "" {
		h.AuthenticateUser(w,r, user)
		return
	}


	response := map[string]interface{}{
		"message": "Welcome to home page",
		"status":  http.StatusOK,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)
	}


func (h *Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request, user string) {

	w.Header().Set("content", "application/json")
	json_data := map[string]interface{} {
		"message": fmt.Sprintf("Welcome %s to the Upstox homepage", user),
		"status":http.StatusOK,
	}
	json_encoded, _ := json.Marshal(json_data) //json.NewEncoder(w).Encode(json_data)

	w.Write(json_encoded)

	
}

func (h *Handler) UpstoxClients(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("content", "application/json")
	json_data := map[string]interface{} {
		"message": fmt.Sprintf("you have access to /upstox/clients/%s", id),
		"status":http.StatusOK,
	}
	json_encoded, _ := json.Marshal(json_data)

	w.Write(json_encoded)

}

