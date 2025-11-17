package sensibull

import "database/sql"

type SensibullAppContainer struct {
	Repo      *SensibullDBRepo
	dbService *SensibullDBService
	Api       *SensibullApiService
	Handler   *SensibullHandler
}

func NewSensibullAppContainer(db *sql.DB) *SensibullAppContainer {
	repo := NewSensibullDBRepo(db)
	dbService := NewSensibullDBService(repo)
	api := NewSensibullApiService()
	handler := NewSensibullHandler(dbService, api)

	return &SensibullAppContainer{
		Repo:      repo,
		dbService: dbService,
		Api:       api,
		Handler:   handler,
	}
}


func (c *SensibullAppContainer) GetSensibullRepo() *SensibullDBRepo {
	return c.Repo
}
func (c *SensibullAppContainer) GetSensibullService() *SensibullDBService {
	return c.dbService
}
// func (c *SensibullAppContainer) GetSensibullAPI() SensibullApiService {
// 	return c.Api
// }
func (c *SensibullAppContainer) GetSensibullHandler() *SensibullHandler {
	return c.Handler
}
