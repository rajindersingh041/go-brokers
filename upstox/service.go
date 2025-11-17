package upstox

import "fmt"

type UpstoxService interface {
	StoreInstrumentData(res Response)
}

type upstoxservice struct {
	repo UpstoxRep
}

func NewUpstoxService(repo UpstoxRep) UpstoxService {
	return &upstoxservice{repo: repo}
}

func (s *upstoxservice) StoreInstrumentData(res Response) {
	fmt.Println("Storing data to DB")
	s.repo.EnsureIntrumentTable()
	for token, info := range res.Data.TokenData {
		info.SymbolID = token
		s.repo.StoreInstrumentData(info)
	}

}

// func (s *Service) Fetch() {

// }