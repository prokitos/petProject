package models

import "time"

type Car struct {
	Id        int                 `json:"car_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	Mark      string              `json:"mark" example:""`
	Year      string              `json:"year" example:""`
	Price     int                 `json:"price" example:""`
	Color     string              `json:"color" example:""`
	MaxSpeed  int                 `json:"speed" example:""`
	SeatsNum  int                 `json:"seats" example:""`
	Engine    CarEngine           `json:"engine" example:"" gorm:"foreignKey:Owner;references:Id"`
	Devices   []AdditionalDevices `json:"devices" example:"" gorm:"foreignKey:Owner;references:Id"`                                                                               // один ко многим
	OwnerList []People            `json:"all_owners" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;many2many:owners_bounds;joinForeignKey:car_uid;JoinReferences:owner_uid"` // много ко многим
	Status    string              `json:"status"`                                                                                                                                 // (sold / for sale)
}

type CarEngine struct {
	Id             int     `json:"engine_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	EngineCapacity float64 `json:"capacity" example:""`
	EnginePower    float64 `json:"power" example:""`
	Owner          int64   `json:"owner" example:""`
}

type People struct {
	Id      int    `json:"people_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	Name    string `json:"name" example:""`
	Surname string `json:"surname" example:""`
	Email   string `json:"email" example:""`
}

type Selling struct {
	Id     int       `json:"sell_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	Car    Car       `json:"car" example:"" gorm:"ForeignKey:Id"`
	People People    `json:"buyer" example:"" gorm:"ForeignKey:Id"`
	Time   time.Time `json:"time" example:""`
}

// ParkMaster, alarm system, trailer coupling, window lifter, power steering, nitro, DVR
type AdditionalDevices struct {
	Id         int    `json:"device_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	DeviceName string `json:"device_name" example:""`
	Owner      int64  `json:"owner" example:""`
}

type CarToRM struct {
	Car
	Types string
}

type SellingToRM struct {
	Id       int
	CarId    int
	PeopleId int
	Types    string
}
