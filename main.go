package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

// describes the data in the call
type dogphoto struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Dog  string `json:"dog"`
}

// getPhotos responds with list of all the photos as JSON.
func getPhotos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, photos)
}

var photos = []dogphoto{
	{Name: "IMG_7507.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7507.JPG", Dog: "Dani"},
	{Name: "IMG_7512.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7512.JPG", Dog: "Josie"},
	{Name: "IMG_7821.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7821.JPG", Dog: "Josie"},
	{Name: "IMG_7356.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7356.JPG", Dog: "Josie"},
	{Name: "IMG_9999.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_9999.jpeg", Dog: "Dani"},
	{Name: "IMG_7873.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7873.JPG", Dog: "Dani_And_Josie"},
	{Name: "IMG_7914.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7914.JPG", Dog: "Josie"},
	{Name: "IMG_7924.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7924.JPG", Dog: "Dani"},
	{Name: "IMG_7960.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7960.JPG", Dog: "Josie"},
}

func main() {
	router := gin.Default()

	router.Static("/js", "./web/js")
	router.Static("/css", "./web/css")
	router.Static("/images", "./web/images")

	router.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})
	router.GET("/api/photos", getPhotos)
	router.Run("localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
}
