package routers

import (
	"../repo"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetChainEvolution(c echo.Context) error {
	Id, _ := strconv.Atoi(c.QueryParam("id"))
	ChainEvolution, _ := repo.GetChainEvolution(Id)

	return c.JSON(http.StatusOK, ChainEvolution)
}
