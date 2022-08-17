package entity

import "time"

// Fish entity of fish
type Fish struct {
	Uuid         *string    `xorm:"uuid" json:"uuid"`
	Komoditas    *string    `xorm:"komoditas" json:"komoditas"`
	AreaProvinsi *string    `xorm:"area_provinsi" json:"area_provinsi"`
	AreaKota     *string    `xorm:"area_kota" json:"area_kota"`
	Size         *int32     `xorm:"size" json:"size,string"`
	Price        *int32     `xorm:"price" json:"price,string"`
	PriceUSD     *float64   `xorm:"price_usd" json:"price_usd"`
	TglParsed    *time.Time `xorm:"tgl_parsed" json:"tgl_parsed"`
	Timestamp    *string    `xorm:"timestamp" json:"timestamp"`
}

// AggregateData entity of agregate data
type AggregateData struct {
	Min    int32   `xorm:"min" json:"min"`
	Max    int32   `xorm:"max" json:"max"`
	Median float64 `xorm:"median" json:"median"`
	Avg    float64 `xorm:"avg" json:"avg"`
	Data   []int   `json:"-"`
}

// Aggregate entity of aggregate
type Aggregate struct {
	Price AggregateData `xorm:"price" json:"price"`
	Size  AggregateData `xorm:"size" json:"size"`
}
