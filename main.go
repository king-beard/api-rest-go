package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{Id: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{Id: "2", Title: "21", Artist: "Adele", Year: 2011},
	{Id: "3", Title: "The Eminem show", Artist: "Eminem", Year: 2022},
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found!"})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbum)

	router.Run("localhost:91")
}
