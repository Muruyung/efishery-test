package usecase

import (
	"MyAPI/entity"
	"MyAPI/helper"
	"encoding/json"
	"fmt"
)

// FishUseCase is wrapper for fish use case
type FishUseCase interface {
	GetFishList() (user []entity.Fish, err error)
	GetCurrency(code string) (currency map[string]interface{}, err error)
}

// GetFishList query for select all fish data from efishery API
func (resty RestyUseCase) GetFishList() (fish []entity.Fish, err error) {
	resp, err := helper.RestyRequest(resty.Client, "https://stein.efishery.com").Get("/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(resp.Body()), &fish)
	return
}

// GetCurrency query for get currency data from currency API
func (resty RestyUseCase) GetCurrency(code string) (currency map[string]interface{}, err error) {
	resp, err := helper.RestyRequest(resty.Client, "https://cdn.jsdelivr.net").Get(fmt.Sprintf("/gh/fawazahmed0/currency-api@1/latest/currencies/idr/%s.json", code))
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(resp.Body()), &currency)
	return
}
