package endpoints

import (
	"api/mocks"
	"common/models"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"api/server/db"
)

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	// bind the received JSON to newAlbum
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		// if error, abort the request with status code 400 and the error message
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newAlbum.ID = uuid.New().String()

	// insert the new album into the database
	insertAlbumQuery := `INSERT INTO Album (id, title, artist, price) VALUES (?, ?, ?, ?)`
	_, err := db.DbConn.Exec(insertAlbumQuery, newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		// if error, abort the request with status code 500 and the error message
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}

func GetAlbums(c *gin.Context) {
    c.JSON(http.StatusOK, mocks.Albums)
}


func GetAlbumByID(c *gin.Context) {
    inputId := c.Param("id")
	_, err := uuid.Parse(inputId)
	if err != nil {
		// abort with error if the id is not a valid UUID
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// traverse through albums, find the album by id
	for _, album := range mocks.Albums {
        if album.ID == inputId {
            c.JSON(http.StatusOK, album)
            return
        }
    }

	// if no album was found, abort with a custom error
	c.AbortWithError(http.StatusNotFound, fmt.Errorf("album not found"))
}

func DeleteAlbumById(c *gin.Context){
	albumID := c.Param("id")
	_, err := uuid.Parse(albumID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

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

// func GetAlbumsByType(c *gin.Context) {
//     typeName := c.Param("type")
	
// 	// create a new slice to store albums of the specified type
// 	var albumsOfType []models.Album

//     // iterate over the list of albums
// 	for _, a := range mocks.Albums {
// 		// If the album type matches the specified type, add it to the slice
//         if a.Type.Name == typeName {
// 			albumsOfType = append(albumsOfType, a)
//         }
//     }
	
// 	// if no albums of the specified type were found, return a 404 error
// 	if len(albumsOfType) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "No albums found of specified type"})
// 		return
// 	}

// 	// otherwise, return the albums of the specified type
// 	c.JSON(http.StatusOK, albumsOfType)
// }

func UpdateAlbum(c *gin.Context) {
    inputId := c.Param("id")
    _, err := uuid.Parse(inputId)
    if err != nil {
        // abort with error if the id is not a valid UUID
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }
    
	// this type does not have thr Type and ID field
    var updateData models.AlbumUpdate
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // find the album to patch
    for index, album := range mocks.Albums {
        if album.ID == inputId {
            if updateData.Title != "" {
                album.Title = updateData.Title
            }
            if updateData.Artist != "" {
                album.Artist = updateData.Artist
            }
            if updateData.Price != 0 {
                album.Price = updateData.Price
            }

            mocks.Albums[index] = album
            c.JSON(http.StatusOK, album)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
}



