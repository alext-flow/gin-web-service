package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type album struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
	Type   albumType `json:"type"`
}

type albumType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
		Type: albumType{
			ID:   "vinyl-1",
			Name: "vinyl",
			Desc: "oldest and round",
		},
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
		Type: albumType{
			ID:   "cd-1",
			Name: "cd",
			Desc: "old, round and relatively small",
		},
	},
	{
		ID:     "3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
		Type: albumType{
			ID:   "minidisc-1",
			Name: "minidisc",
			Desc: "round, small and the biggest flop of all",
		},
	},
}

var albumTypes map[string]albumType = map[string]albumType{
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

func postAlbums(c *gin.Context) {
    var newAlbum album

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
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}


// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getType(c *gin.Context){
	albumType := c.Query("type")
	typeName, ok := albumTypes[albumType]
	if !ok{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid album type",
		})
		return
	}
	c.JSON(http.StatusOK, typeName)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

	// traverse through albums, find the album by id
	for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteById(c *gin.Context){
	id := c.Param("id")

	for index, album := range albums {
		if album.ID == id {
			deletedAlbum := album

			// delete from the "db" by getting all albums before all albums after
			albums = append(albums[:index], albums[index+1:]...)

			c.JSON(http.StatusOK, deletedAlbum)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error" : "album not found"})
}

func getAlbumsByType(c *gin.Context) {
    typeName := c.Param("type")
	
	// create a new slice to store albums of the specified type
	var albumsOfType []album

    // iterate over the list of albums
	for _, a := range albums {
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

func patchAlbum(c *gin.Context) {
	id := c.Param("id")
	var patchData map[string]interface{}

	if err := c.ShouldBindJSON(&patchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the album to patch
	for index, album := range albums {
		if album.ID == id {
			patchedAlbum := album
			patchDataBytes, _ := json.Marshal(patchData)
			if err := json.Unmarshal(patchDataBytes, &patchedAlbum); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patch data"})
				return
			}

			albums[index] = patchedAlbum
			c.JSON(http.StatusOK, patchedAlbum)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
}

func main() {
	router := gin.Default()
  
	router.GET("/type", getType);
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteById)
	router.GET("/albums/type/:type", getAlbumsByType)
	router.PATCH("/albums/:id", patchAlbum)
	router.Run()
  }