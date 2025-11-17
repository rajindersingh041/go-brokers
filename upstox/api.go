package upstox

import (
	"errors"
	"fmt"
	"net/http"
)

type UpstoxApiService struct {
}

func NewUpstoxApiService() UpstoxApiService {
	return UpstoxApiService{}
}

func (a *UpstoxApiService) SetHeaders(req *http.Request) {
	req.Header.Add("x-request-id", "abc")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")

}

func (a *UpstoxApiService) FetchOHLC(Instruments string) (*http.Response, error) {
	url := fmt.Sprintf("https://service.upstox.com/market-data-api/v2/open/quote?i=%s", Instruments)
	fmt.Println("URL is ", url)
	req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Add("x-request-id", "abc")
	// req.Header.Add("content-type", "application/json")
	// req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")
	a.SetHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// fmt.Errorf("Failed to fetch OHLC data")
		return nil, errors.New("failed to fetch OHLC data")
	}
	return resp, nil

}