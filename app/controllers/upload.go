package controllers

import (
  "github.com/robfig/revel"
  "fmt"
  "errors"
  "share/lib/uploaded_image"
)

type Upload struct {
  *revel.Controller
}

func (c Upload) Do() revel.Result {

  files := c.Params.Files["file"]
  f_len := len(files)
  if f_len < 1 {
    err := errors.New("Binary file data expected.")
    c.RenderError(err)
  } else if f_len > 1 {
    err := errors.New(fmt.Sprintf("Single file handler per request expected, %d given", f_len))
    c.RenderError(err)
  }

  file := files[0]
  u_img, err := uploaded_image.New(file)

  revel.INFO.Println(u_img)

  if u_img == nil {
  
  } else if err != nil {
  
  }

//  for i, _ := range files {
//    file, err := files[i].Open()
//    defer file.Close()
//    if err != nil {
//      return c.RenderError(err)
//    }
//    revel.INFO.Println(files[i].Filename)
//  }

  return c.RenderText("hello upload")
}
