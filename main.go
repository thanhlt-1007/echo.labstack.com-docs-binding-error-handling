package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Query struct {
	IDS    []int64
	Active bool
}

func getSearch(context echo.Context) error {
	var query Query

	err := echo.QueryParamsBinder(context).
		MustInt64s("ids", &query.IDS).
		MustBool("active", &query.Active).
		BindError()
	if err != nil {
		return context.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	return context.JSON(
		http.StatusOK,
		query,
	)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/search", getSearch)
	e.Logger.Fatal(e.Start(":1323"))
}
