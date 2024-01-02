package main

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"insights/eath/v2/handlers"
	example "insights/eath/v2/views/components/examples"
	"insights/eath/v2/views/pages"
	"net/http"
)

func main() {

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Static("/static", "views/static")

	e.GET("/", handlers.LoginPage)
	e.POST("/login", handlers.Login)
	e.POST("/login/validate/:name", handlers.LoginValidate)
	e.POST("/dashboard", dashboard)
	e.POST("/newPage", echo.WrapHandler(http.HandlerFunc(newPage)))
	e.POST("/htmxexamples", echo.WrapHandler(http.HandlerFunc(examples)))

	exampleGroup := e.Group("/examples")
	exampleGroup.GET("/clicktoedit/:id", func(c echo.Context) error {
		// id := c.Param("id")
		component := example.ClickToEditResponse()
		err := render(c, http.StatusOK, component)
		return err
		// return c.String(http.StatusOK, id)
	})
	exampleGroup.PUT("/clicktoedit/:id", func(c echo.Context) error {
		// id := c.Param("id")
		component := example.ClickToEdit()
		err := render(c, http.StatusOK, component)
		return err
		// return c.String(http.StatusOK, id)
	})
	e.Logger.Fatal(e.Start(":3000"))
}

func examples(w http.ResponseWriter, r *http.Request) {
	component := pages.HtmxExamples()
	component.Render(context.Background(), w)
}

func dashboard(c echo.Context) error {

	values := []int{200, 150, 350, 225, 300, 145, 500, 1200}
	str := "line"
	//Rendering the page with email and values
	component := pages.Dashboard(values, str)
	return render(c, http.StatusOK, component)
}

func newPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Print(r)
	values := []int{200, 150, 350, 225, 300, 145, 500, 1200}
	email := r.FormValue("email")
	// If email and password are correct
	if email == "" {
		email = "line"
	}
	//Rendering the page with email and values
	component := pages.NewPage(email, values, email)
	component.Render(context.Background(), w)
}

func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}
