package models

type Users struct {
	User_id     int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Login       string `json:"login" example:""`
	Password    string `json:"password" example:""`
	AccessLevel int    `json:"access_level" example:""`
}
