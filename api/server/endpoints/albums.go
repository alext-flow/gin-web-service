package endpoints

import (
	"common/models"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"api/server/db"
	"gorm.io/gorm"
	"errors"
)


func CreateAlbum(c *gin.Context) {
    fmt.Println("request received to create a new Album")
    var newAlbum models.Album
    if err := c.BindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // start connection
    tx := db.DbConn.Begin()

    var albumType models.AlbumType
    var album models.Album

    // check if there's an AlbumType stored in the db by name, if so, get it
    if err := tx.Where("name = ?", newAlbum.AlbumType.Name).First(&albumType).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // create new AlbumType if it does not exist
            albumType = models.AlbumType{
                ID:   uuid.New(),
                Name: newAlbum.AlbumType.Name,
                Desc: newAlbum.AlbumType.Desc,
            }

            fmt.Println("creating a new AlbumType: ", albumType)
            if err := tx.Create(&albumType).Error; err != nil {
                tx.Rollback()
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create album type"})
                return
            }
        } else {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the album type"})
            return
        }
    }

    // Set the created/found AlbumType to newAlbum.AlbumType
    newAlbum.AlbumType = albumType
    newAlbum.TypeID = albumType.ID

    // check if the combination of albums.title and album_types.id exists, if so, return the object
    if err := tx.Where("title = ? AND type_id = ?", newAlbum.Title, albumType.ID).First(&album).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // create the Album
            newAlbum.ID = uuid.New().String()

            fmt.Println("Creating a new Album: ", newAlbum)
            if err := tx.Create(&newAlbum).Error; err != nil {
                tx.Rollback()
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create album"})
                return
            }

            // commit the transaction
            tx.Commit()
            c.JSON(http.StatusCreated, newAlbum)
            return
        } else {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the album"})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Album with that title and Album Type already exists"})
        return
    }
}

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	// get all albums from the database using GORM
	if err := db.DbConn.Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
    var album models.Album

    if err := db.DbConn.Preload("Type").First(&album, "id = ?", c.Param("id")).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the album"})
        }
        return
    }

    c.JSON(http.StatusOK, album)
}

func DeleteAlbumById(c *gin.Context) {
	var album models.Album
	albumID := c.Param("id")

	// delete the album with the given ID from the database using GORM
	if err := db.DbConn.Delete(&album, "id = ?", albumID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func GetAlbumsByType(c *gin.Context) {
    typeName := c.Param("type")

	var albumsOfType []models.Album

	// Find all albums of the specified type
	result := db.DbConn.Preload("Type", "name = ?", typeName).Find(&albumsOfType)

	// Check if any error occurred during the DB operation
	if result.Error != nil {
		// If error is due to not finding any records, return a 404
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No albums found of specified type"})
		} else {
			// If error is due to something else, return a 500 (Internal Server Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// If no albums of the specified type were found, return a 404
	if len(albumsOfType) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No albums found of specified type"})
		return
	}

	// Return the albums of the specified type
	c.JSON(http.StatusOK, albumsOfType)
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	inputId := c.Param("id")

	// Get the album with the given ID from the database using GORM
	if err := db.DbConn.First(&album, "id = ?", inputId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	var updateData models.AlbumUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateData.Title != "" {
		album.Title = updateData.Title
	}
	if updateData.Artist != "" {
		album.Artist = updateData.Artist
	}
	if updateData.Price != 0 {
		album.Price = updateData.Price
	}

	// Save the updated album back to the database using GORM
	if err := db.DbConn.Save(&album).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}
