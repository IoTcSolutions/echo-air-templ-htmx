package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"insights/eath/v2/views/components"
	"insights/eath/v2/views/pages"
	"net/http"
)

var valid = 0

func LoginPage(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	fmt.Println(sess)

	sess.Save(c.Request(), c.Response())
	component := pages.Login()
	err := render(c, http.StatusOK, component)
	return err
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println(email, password)
	values := []int{200, 150, 350, 225, 300, 145, 500, 1200}
	str := "line"
	component := pages.Dashboard(values, str)
	err := render(c, http.StatusOK, component)
	return err
}

func LoginValidate(c echo.Context) error {
	name := c.Param("name")
	f, ok := components.LoginFields[name]
	if !ok {
		c.Response().WriteHeader(500)
		return errors.New("Could not extract from LoginFields")
	}
	if err := c.Request().ParseForm(); err != nil {
		c.Response().WriteHeader(500)
		return errors.New("Could not parse login form")
	}
	val := c.Request().FormValue(name)
	err := f.Validator(val)
	if err == nil {
		valid++
	}
	if valid >= 2 {
		c.Response().Header().Set("HX-Trigger", "validLogin")
	}
	component := components.Inp(f, name, val, f.Validator(val))
	err = render(c, http.StatusOK, component)
	return err
}
