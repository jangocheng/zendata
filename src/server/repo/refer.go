package serverRepo

import (
	"github.com/easysoft/zendata/src/model"
	"github.com/jinzhu/gorm"
)

type ReferRepo struct {
	db *gorm.DB
}

func (r *ReferRepo) CreateDefault(fieldId uint) (err error) {
	refer := &model.Refer{FieldID: fieldId}
	err = r.db.Create(&refer).Error

	return
}

func (r *ReferRepo) Create(refer *model.Refer) (err error) {
	err = r.db.Create(&refer).Error
	return
}

func (r *ReferRepo) Get(fieldId uint) (refer model.Refer, err error) {
	err = r.db.Where("fieldID=?", fieldId).First(&refer).Error
	return
}

func (r *ReferRepo) Save(ref *model.Refer) (err error) {
	err = r.db.Save(ref).Error
	return
}

func NewReferRepo(db *gorm.DB) *ReferRepo {
	return &ReferRepo{db: db}
}