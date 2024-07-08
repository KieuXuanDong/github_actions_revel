package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.set("index_variable", "Index Variable")
	return c.Render()
}
