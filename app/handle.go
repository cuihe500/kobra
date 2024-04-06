package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log/slog"
)

type Handle struct {
	Context *gin.Context
	Orm     *gorm.DB
	Error   error
}

func (handle *Handle) AddError(err error) {
	if handle.Error == nil {
		handle.Error = err
	} else if handle.Error != nil {
		slog.Error("Something went wrong but not get!", "reason", handle.Error)
		slog.Error("Something get error!", "reason", err)
		handle.Error = err
	}
}

func (handle *Handle) MakeContext(ctx *gin.Context) *Handle {
	handle.Context = ctx
	return handle
}

func (handle *Handle) MakeOrm() *Handle {
	db, exists := handle.Context.Get("db")
	if !exists {
		handle.AddError(errors.New("Can not get db connect in context!Reason:Not Found."))
		return nil
	}
	handle.Orm = db.(*gorm.DB)
	return handle
}
