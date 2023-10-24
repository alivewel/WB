package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"service/pkg/memorycache"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Title     string
	Total     string
	Name      string
	PageNum   int
	TotalPage int
	PrevPage  int
	NextPage  int
	LastPage  int
}

func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, Wildberries!")
}

func GetDataHandler(c *gin.Context, cache *memorycache.Cache, htmlTemplate string) {
	id := c.Param("id")
	tmpl := template.Must(template.New("htmlTemplate").Parse(htmlTemplate))
	data, _ := cache.Get(id)
	if str, ok := data.(string); ok {
		pageNum := id
		pageNumInt, err := strconv.Atoi(pageNum)
		if err != nil {
			// Error handling if "page" is not an integer
			c.String(http.StatusBadRequest, "Invalid 'page' value: "+pageNum)
			return
		}

		totalRecord := cache.Count()
		if pageNumInt < 1 || pageNumInt > totalRecord {
			c.String(http.StatusBadRequest, "Invalid page. Out of range.")
			return
		}
		pageData := Page{
			Title:     "Places",
			Total:     "Total: " + strconv.Itoa(totalRecord),
			Name:      str,
			PageNum:   pageNumInt,
			TotalPage: totalRecord,
			PrevPage:  max(pageNumInt-1, 1),
		}

		if pageNumInt < totalRecord {
			pageData.NextPage = pageNumInt + 1
			pageData.LastPage = totalRecord
		}
		tmpl.Execute(c.Writer, pageData)
	} else {
		// Обработка случая, когда id выходит за пределы
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: id %s is not in the database", id))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
