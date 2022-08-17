package adapters

import (
	"MyAPI/adapters/database"
	"MyAPI/adapters/resty"
	"fmt"
)

// Adapters is wrapper for lib/driver that needed to be injected
type Adapters struct {
	Database *database.Database
	Resty    *resty.Resty
}

// Init initialization new adapters
func Init() (Adapters, error) {
	database, err := database.Init()
	if err != nil {
		fmt.Println(err)
	}

	resty := resty.Init()

	return Adapters{
		database,
		resty,
	}, nil
}
