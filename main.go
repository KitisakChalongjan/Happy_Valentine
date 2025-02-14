package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	fmt.Println("Happy Valentine Service Started!")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "🦉")
	})

	e.GET("/what_is_it", func(c echo.Context) error {
		heart := `
         ******       ******
       **      **   **      **
     **          ** **        **
    **            ***          	**
    **             *            **
     **                        **
       **                    **
         **                **
           **            **
             **        **
               **   **
                 **
		`
		return c.String(http.StatusOK, heart)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
