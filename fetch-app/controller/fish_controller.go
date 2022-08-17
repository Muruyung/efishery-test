package controller

import (
	"MyAPI/entity"
	"sort"
)

// FishController is wrapper for fish controller
type FishController interface {
	GetFishList() (fish []entity.Fish, err error)
	GroupFish() (groupFish map[string]interface{}, err error)
}

// GetFishList get list of fish from fish usecase
func (controller Controller) GetFishList() (fish []entity.Fish, err error) {
	fish, err = controller.usecase.GetFishList()
	if err != nil {
		return
	}

	currencyKey := "usd"
	currency, err := controller.usecase.GetCurrency(currencyKey)
	if err != nil {
		return
	}

	for key := range fish {
		if fish[key].Price != nil {
			priceUSD := float64(*fish[key].Price) * currency[currencyKey].(float64)
			fish[key].PriceUSD = &priceUSD
		}
	}
	return
}

// GroupFish get list of fish and aggregate data by province and weekly
func (controller Controller) GroupFish() (groupFish map[string]interface{}, err error) {
	fish, err := controller.usecase.GetFishList()
	if err != nil {
		return
	}

	groupFish = map[string]interface{}{}
	for _, val := range fish {
		groupWeek := make(map[int]entity.Aggregate)
		if val.AreaProvinsi != nil {
			_, week := val.TglParsed.ISOWeek()
			if _, ok := groupFish[*val.AreaProvinsi]; ok {
				groupWeek = groupFish[*val.AreaProvinsi].(map[int]entity.Aggregate)

				if _, ok := groupWeek[week]; ok {
					priceData := groupWeek[week].Price
					sizeData := groupWeek[week].Size

					if priceData.Min > *val.Price {
						priceData.Min = *val.Price
					}
					if priceData.Max < *val.Price {
						priceData.Max = *val.Price
					}
					priceData.Avg = (priceData.Avg + float64(*val.Price)) / 2
					priceData.Data = append(priceData.Data, int(*val.Price))

					if sizeData.Min > *val.Size {
						sizeData.Min = *val.Size
					}
					if sizeData.Max < *val.Size {
						sizeData.Max = *val.Size
					}
					sizeData.Avg = (sizeData.Avg + float64(*val.Size)) / 2
					sizeData.Data = append(sizeData.Data, int(*val.Size))

					groupWeek[week] = entity.Aggregate{
						Price: priceData,
						Size:  sizeData,
					}
				} else {
					groupWeek[week] = entity.Aggregate{
						Price: entity.AggregateData{
							Min:    *val.Price,
							Max:    *val.Price,
							Median: float64(*val.Price),
							Avg:    float64(*val.Price),
							Data:   []int{int(*val.Price)},
						},
						Size: entity.AggregateData{
							Min:    *val.Size,
							Max:    *val.Size,
							Median: float64(*val.Size),
							Avg:    float64(*val.Size),
							Data:   []int{int(*val.Size)},
						},
					}
				}
			} else {
				groupWeek[week] = entity.Aggregate{
					Price: entity.AggregateData{
						Min:    *val.Price,
						Max:    *val.Price,
						Median: float64(*val.Price),
						Avg:    float64(*val.Price),
						Data:   []int{int(*val.Price)},
					},
					Size: entity.AggregateData{
						Min:    *val.Size,
						Max:    *val.Size,
						Median: float64(*val.Size),
						Avg:    float64(*val.Size),
						Data:   []int{int(*val.Size)},
					},
				}
			}
			groupFish[*val.AreaProvinsi] = groupWeek
		}
	}

	for keyFish, valFish := range groupFish {
		groupWeek := valFish.(map[int]entity.Aggregate)
		for keyWeek, valWeek := range groupWeek {
			priceData := valWeek.Price
			sizeData := valWeek.Size

			if len(priceData.Data) > 1 {
				sort.Ints(priceData.Data)
				index := len(priceData.Data) / 2
				if len(priceData.Data)%2 == 0 {
					priceData.Median = (float64(priceData.Data[index]) + float64(priceData.Data[index-1])) / 2
				} else {
					priceData.Median = float64(priceData.Data[index])
				}
			}

			if len(sizeData.Data) > 1 {
				sort.Ints(sizeData.Data)
				index := len(sizeData.Data) / 2
				if len(sizeData.Data)%2 == 0 {
					sizeData.Median = (float64(sizeData.Data[index]) + float64(sizeData.Data[index-1])) / 2
				} else {
					sizeData.Median = float64(sizeData.Data[index])
				}
			}

			groupWeek[keyWeek] = entity.Aggregate{
				Price: priceData,
				Size:  sizeData,
			}
		}

		groupFish[keyFish] = groupWeek
	}

	return
}
