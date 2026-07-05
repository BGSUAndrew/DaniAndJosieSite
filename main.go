package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//describes the data in the call

type dogphoto struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// getPhotos responds with list of all the photos as JSON.
func getPhotos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, photos)
}

var photos = []dogphoto{
	{Name: "IMG_7507.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7507.JPG"},
	{Name: "IMG_7821.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7821.JPG"},
	{Name: "IMG_7873.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7873.JPG"},
	{Name: "IMG_7914.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7914.JPG"},
	{Name: "IMG_7924.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7924.JPG"},
	{Name: "IMG_7960.JPG", URL: "https://joisethedog.sfo3.cdn.digitaloceanspaces.com/IMG_7960.JPG"},
}

func main() {
	router := gin.Default()
	router.GET("/photos", getPhotos)
	router.Run("localhost:8080")
}
