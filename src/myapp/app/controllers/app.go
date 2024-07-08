package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	indexVariable := "Index Variable"
	c.isValueValid(indexVariable)
	return c.Render(indexVariable)
}

func (c App) isValueValid(value string) bool {
	if value == "valid" {
		return true
	}

	return false
}
