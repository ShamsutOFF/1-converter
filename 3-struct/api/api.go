package api

import (
	"1-converter/3-struct/config"
	"fmt"
)

type Api struct {
	Config *config.Config // будем хранить указатель
}

// Конструктор (функция пакета, а не метод!)
func NewApi(cfg *config.Config) *Api {
	fmt.Println("Initializing...")
	fmt.Println("Config = ", cfg)
	return &Api{Config: cfg}
}
