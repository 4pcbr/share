package controllers

import (
  "github.com/robfig/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  //revel.INFO.Println(revel.Config.StringDefault("app.name", "unknown"))
  return c.Render()
}
