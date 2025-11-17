package sensibull

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type SensibullApiService struct {
}

func NewSensibullApiService() *SensibullApiService {
	return &SensibullApiService{}
}
func (s *SensibullApiService) FetchSensibullData(instrument string) (*http.Response, error) {
	
	url := "https://oxide.sensibull.com/v1/compute/compute_intraday"
	payload := fmt.Sprintf(`{"underlying":"%s","interval":"1M","chart_keys": ["iv", "ivp", "rolling_atm_straddle","price","oi_futures","oi_change_options","pcr","max_pain","indiavix","oi_options"]}`, instrument)

	fmt.Println("Payload is ", payload)
	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))
	fmt.Println("URL is ", url)
	s.SetHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("failed to fetch sensibull data")
	}
	return resp, nil
}

func (a *SensibullApiService) SetHeaders(req *http.Request) {
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")

}

