package main

import (
	"api/mocks"
	"common/models"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

var albumTypes map[string]models.AlbumType = map[string]models.AlbumType{
	"vinyl": {
		ID:     "vinyl-1",
		Name:  "vinyl",
		Desc:  "oldest and round",
	},
	"cd": {
		ID:     "cd-1",
		Name:  "cd",
		Desc:  "old, round and relatively small",
	},
	"minidisc": {
		ID:     "minidisc-1",
		Name:  "minidisc",
		Desc:  "round, small and the biggest flop of all",
	},
}

func createAlbum(c *gin.Context) {
    var newAlbum models.Album

    // bind the received JSON to newAlbum
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
        // if error, send a response with
        // the status code 400 and the error message
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // add the new album to the slice
    mocks.Albums = append(mocks.Albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}


// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, mocks.Albums)
}

func getAlbumType(c *gin.Context){
	albumType := c.Query("type")
	//typeName, ok := albumTypes[albumType]
	if typeName, ok := albumTypes[albumType]; ok{
		c.JSON(http.StatusOK, typeName)
		return
	}
	c.AbortWithError(http.StatusNotFound, fmt.Errorf("invalid album type"))
	// c.JSON(http.StatusNotFound, gin.H{
	// 	"error": "Invalid album type",
	// })
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

	// traverse through albums, find the album by id
	for _, album := range mocks.Albums {
        if album.ID == id {
            c.IndentedJSON(http.StatusOK, album)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumById(c *gin.Context){
	albumID := c.Param("id")

	for index, album := range mocks.Albums {
		if album.ID == albumID {

			// delete from the "db" by getting all albums before all albums after
			mocks.Albums = append(mocks.Albums[:index], mocks.Albums[index+1:]...)

			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error" : "album not found"})
}

func getAlbumsByType(c *gin.Context) {
    typeName := c.Param("type")
	
	// create a new slice to store albums of the specified type
	var albumsOfType []models.Album

    // iterate over the list of albums
	for _, a := range mocks.Albums {
		// If the album type matches the specified type, add it to the slice
        if a.Type.Name == typeName {
			albumsOfType = append(albumsOfType, a)
        }
    }
	
	// if no albums of the specified type were found, return a 404 error
	if len(albumsOfType) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No albums found of specified type"})
		return
	}

	// otherwise, return the albums of the specified type
	c.JSON(http.StatusOK, albumsOfType)
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var patchData map[string]interface{}

	if err := c.ShouldBindJSON(&patchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the album to patch
	for index, album := range mocks.Albums {
		if album.ID == id {
			patchedAlbum := album
			patchDataBytes, _ := json.Marshal(patchData)
			if err := json.Unmarshal(patchDataBytes, &patchedAlbum); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patch data"})
				return
			}

			mocks.Albums[index] = patchedAlbum
			c.JSON(http.StatusOK, patchedAlbum)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
}

func main() {
	router := gin.Default()
  
	router.GET("/type/:type", getAlbumType);
	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.GET("/albums/type/:type", getAlbumsByType)
	router.PATCH("/albums/:id", updateAlbum)
	router.Run()
  }