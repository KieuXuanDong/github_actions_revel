package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	indexVariable := "Index Variable"
	isValueValid(indexVariable)
	return c.Render(indexVariable)
}

func (c App) isValueValid(value string) boolean {
	if value == "valid" {
		return true
	}

	return false
}
