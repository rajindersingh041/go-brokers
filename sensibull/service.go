package sensibull

type SensibullDBService struct {
	repo *SensibullDBRepo
}

func NewSensibullDBService(repo *SensibullDBRepo) *SensibullDBService {
	return &SensibullDBService{repo}
}

func (s SensibullDBService) ProcessAndStoreData(data []SensibullData, symbolID string) error {
	// s.repo.InsertSensibullData(symbolID, data)
	s.repo.InsertManySensibullData(symbolID, data)
	return nil
}

// func (s SensibullDBService) StorePrice(data []SensibullData, symbolID string) error {
// 	s.repo.InsertSensibullData(symbolID, data)
// 	s.repo.InsertManySensibullData(symbolID, data)
// 	return nil
// }

func (s SensibullDBService) StorePrice(data []SensibullData, symbolID string) error {
	s.repo.StorePrice(symbolID, data)
	return nil
}

func (s SensibullDBService) StoreOI(data []SensibullData, symbolID string) error {
	s.repo.StoreOI(symbolID, data)
	return nil
}

func (s SensibullDBService) StoreIV(data []SensibullData, symbolID string) error {
	s.repo.StoreIV(symbolID, data)
	return nil
}
func (s SensibullDBService) StoreRollingStraddle(data []SensibullData, symbolID string) error {
	s.repo.StoreRollingStraddle(symbolID, data)
	return nil
}