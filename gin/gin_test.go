package gin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func TestGinServer(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbumsWrap)

	err := router.Run("localhost:8080")
	if err != nil {
		t.Errorf("run gin server failed, error: %v", err)
		panic(err)
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// postAlbumsWrap adds an album from JSON received in the request body.
func postAlbumsWrap(c *gin.Context) {
	type NewAlbumWrap struct {
		Data album `json:"data"`
	}

	req := NewAlbumWrap{Data: album{}}

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, req.Data)
	c.IndentedJSON(http.StatusCreated, req)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
