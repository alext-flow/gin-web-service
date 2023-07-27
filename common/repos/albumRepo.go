package repos

// struct that holds all methods
import (
    "common/models"
    "gorm.io/gorm"
	"fmt"
	"errors"
	"github.com/google/uuid"
)

type AlbumRepo struct {
    DbConn *gorm.DB
}

func NewAlbumRepo(db *gorm.DB) *AlbumRepo {
    return &AlbumRepo{
        DbConn: db,
    }
}

func (ar *AlbumRepo) CreateAlbum(newAlbum models.Album, albumTypeName string, albumTypeDesc string) (*models.Album, error) {
    var albumType models.AlbumType
    var album models.Album

    tx := ar.DbConn.Begin()

    if err := tx.Where("name = ?", albumTypeName).First(&albumType).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            albumType = models.AlbumType{
                ID:   uuid.New(),
                Name: albumTypeName,
                Desc: albumTypeDesc,
            }

            if err := tx.Create(&albumType).Error; err != nil {
                tx.Rollback()
                return nil, err
            }
        } else {
            tx.Rollback()
            return nil, err
        }
    }

    newAlbum.AlbumType = albumType
    newAlbum.TypeID = albumType.ID

    if err := tx.Where("title = ? AND type_id = ?", newAlbum.Title, albumType.ID).First(&album).Error; err != nil {
		// combination of the given title and type id does not exist, create it
        if errors.Is(err, gorm.ErrRecordNotFound) {
            newAlbum.ID = uuid.New().String()

            if err := tx.Create(&newAlbum).Error; err != nil {
                tx.Rollback()
                return nil, err
            }

            tx.Commit()
            return &newAlbum, nil
        } else {
            tx.Rollback()
            return nil, err
        }
    } else {
        return nil, fmt.Errorf("Album with that title and Album Type already exists")
    }
}

func (ar *AlbumRepo) Albums() ([]*models.Album, error) {
    var albums []*models.Album
    result := ar.DbConn.Preload("AlbumType").Find(&albums)
    if result.Error != nil {
        return nil, result.Error
    }
    return albums, nil
}

func (ar *AlbumRepo) Album(id string) (*models.Album, error){
	var album models.Album
    result := ar.DbConn.Preload("AlbumType").Where("id = ?", id).Find(&album)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, result.Error
    }
    return &album, nil
}

func (ar *AlbumRepo) DeleteAlbum(id string) (bool, error){
	var album models.Album
	// check if the album exists
	if err := ar.DbConn.First(&album, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil
        }
		return false, err
	}

	// if the album exists terminate it
	result := ar.DbConn.Delete(&album)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}