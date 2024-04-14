package controllers

import (
	"az/app/models"
	"az/config"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, filename string, data interface{}, c echo.Context) error {

	err := t.templates.ExecuteTemplate(w, filename, data)
	if err != nil {
		// エラーハンドリング
		return err
	}
	return nil
}

type CommonData struct {
	Title string
}

func ViewTopPage(c echo.Context) error {
	data := "Hello"
	/* Renderでhtmlを表示 */
	path := "layout.html"
	return c.Render(http.StatusOK, path, data)
}

func ViewHomePage(c echo.Context) error {
	data := "good morning"
	path := "home.html"
	return c.Render(http.StatusOK, path, data)
}

func ViewSelection(c echo.Context) error {
	method := c.Request().Method
	if method == http.MethodPost {
		// POSTリクエストの処理
		selection := c.FormValue("selection")
		msg := "あなたが選んだのは" + selection + "です。"
		return c.Render(http.StatusOK, "home.html", msg)
	} else if method == http.MethodGet {
		// GETリクエストの処理
		msg := "あなたが選んだサイトは見つかりませんでした。"
		return c.Render(http.StatusOK, "home.html", msg)
	}
	return nil
}

func ViewOpinion(c echo.Context) error {
	opinion := c.FormValue("kansou")
	name := "murata"

	// Userモデルを作成
	opi := models.Opinion{
		Name:    name,
		Opinion: opinion,
	}

	// ユーザーを作成
	if err := opi.CreateOpinion(); err != nil {
		// エラーが発生した場合はログに記録
		log.Fatalln(err)
		return err
	}

	// 成功した場合はステータスコード200を返す
	return c.String(http.StatusOK, "User created successfully")
}

func StartMainServer() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("app/views/templates/*.html")),
	}
	e.Renderer = t

	e.Static("/static", config.Config.Static)

	e.GET("/top", ViewTopPage)
	e.GET("/home", ViewHomePage)
	e.POST("/select", ViewSelection)
	e.POST("/opinion", ViewOpinion)
	e.Start(":8080")
}
