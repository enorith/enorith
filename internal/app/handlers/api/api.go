package api

import (
	"github.com/enorith/enorith/internal/app/models"
	dba "github.com/enorith/enorith/internal/pkg/database"

	"github.com/enorith/framework/database"
	"github.com/enorith/http/content"
	"github.com/enorith/http/contracts"
	"github.com/enorith/http/router"
	"gorm.io/gorm"
)

type CURDListScope interface {
	ListScope(req contracts.RequestContract, user models.AuthUser) func(db *gorm.DB) *gorm.DB
	AggTableScope(req contracts.RequestContract) func(db *gorm.DB) *gorm.DB
}

type CURDHandler[T models.Model] struct {
}

func (CURDHandler[T]) List(builder *database.Builder[T], req contracts.RequestContract, user models.AuthUser) (*database.PageResult[T], error) {
	var model T
	var opt database.PaginateOptions

	if ls, ok := models.Model(model).(CURDListScope); ok {
		opt.AggTableScope = ls.AggTableScope(req)
	}
	return builder.Query(func(d *gorm.DB) *gorm.DB {
		if ls, ok := models.Model(model).(CURDListScope); ok {
			d = ls.ListScope(req, user)(d)
		}
		// d.Order("id desc")

		dba.WithSorts(d, req)
		return dba.WithFilters(d, req)
	}).Paginate(opt)
}

// Post, if id > 0, update, else create
func (CURDHandler[T]) Post(model T, tx *gorm.DB) (content.JsonMessage, error) {
	// update or create
	e := tx.Transaction(func(tx *gorm.DB) error {
		var e error

		var origin T

		if model.GetID() > 0 {
			tx.First(&origin, "id = ?", model.GetID())
			e = tx.Model(&model).Omit("id").Updates(model).Error
		} else {
			origin = model
			e = tx.Create(&model).Error
		}

		if e != nil {
			return e
		}

		if ap, ok := models.Model(model).(models.AfterPost[T]); ok {
			return ap.AfterPost(tx, origin)
		}

		return nil
	})

	if e != nil {
		return content.JsonMessage(422), e
	}

	return content.JsonMessage(200), nil
}

func (CURDHandler[T]) Detail(id content.ParamInt64, tx *gorm.DB) (T, error) {
	var model T
	e := tx.First(&model, "id = ?", id).Error
	return model, e
}

func (CURDHandler[T]) Delete(id content.ParamInt64, tx *gorm.DB) error {
	var model T
	return tx.Delete(&model, "id = ?", id).Error
}

func (CURDHandler[T]) Toggle(id content.ParamInt64, tx *gorm.DB) error {
	var model T

	return tx.Model(&model).Where("id = ?", id).Update("enabled", gorm.Expr(" !enabled")).Error
}

func (h CURDHandler[T]) RegisterRoutes(r *router.Wrapper, prefix string) {
	r.Post(prefix, h.List)
	r.Post(prefix+"/write", h.Post)
	r.Delete(prefix+"/:id", h.Delete)
	r.Get(prefix+"/:id", h.Detail)
	r.Post(prefix+"/toggle/:id", h.Toggle)
}
