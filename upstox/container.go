package upstox

import "database/sql"

type UpstoxAppContainer struct {
	Repo      UpstoxRep
	dbService UpstoxService
	Api       UpstoxApiService
	Handler   *Handler
}

func NewUpstoxAppContainer(db *sql.DB) *UpstoxAppContainer {
	repo := NewUpstoxRep(db)
	dbService := NewUpstoxService(repo)
	api := NewUpstoxApiService()
	handler := NewHandler(dbService, api)
	return &UpstoxAppContainer{
		Repo:      repo,
		dbService: dbService,
		Api:       api,
		Handler:   handler,
	}
}


func (c *UpstoxAppContainer) GetUpstoxRepo() UpstoxRep {
	return c.Repo
}
func (c *UpstoxAppContainer) GetUpstoxService() UpstoxService {
	return c.dbService
}
func (c *UpstoxAppContainer) GetUpstoxAPI() UpstoxApiService {
	return c.Api
}
func (c *UpstoxAppContainer) GetUpstoxHandler() *Handler {
	return c.Handler
}
