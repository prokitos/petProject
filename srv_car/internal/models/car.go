package models

// структура таблицы машин
type Car struct {
	Id     string `json:"id" example:"12"`
	RegNum string `json:"regNum" example:""`
	Mark   string `json:"mark" example:""`
	Model  string `json:"model" example:"tesla"`
	Year   string `json:"year" example:""`
	Owner  People `json:"owner"`
}

type People struct {
	Id         string `json:"id" example:""`
	Name       string `json:"name" example:"jamson"`
	Surname    string `json:"surname" example:""`
	Patronymic string `json:"patronymic" example:""`
}

type CarToRM struct {
	Car
	Types string
}
