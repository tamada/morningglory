package points

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamada/morningglory/common"
	"github.com/tamada/morningglory/users"
)

func readBody(context *gin.Context) *common.Point {
	var point = common.Point{}
	context.Bind(&point)
	return &point
}

func RegisterPoints(context *gin.Context) error {
	if err := users.Authenticate(context); err != nil {
		return err
	}
	var point = readBody(context)
	point.User = context.Param("userName")
	point.Date = time.Now()
	return common.RegisterPoint(point)
}
