package endpoints

// import (
// 	"api/mocks"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"fmt"
// )

// func GetAlbumType(c *gin.Context) {
// 	typeParam := c.Param("type")
// 	typeName, ok := mocks.AlbumTypes[typeParam]
// 	if ok {
// 		c.JSON(http.StatusOK, typeName)
// 		return
// 	}
// 	c.AbortWithError(http.StatusNotFound, fmt.Errorf("invalid album type"))
// }