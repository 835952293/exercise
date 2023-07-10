package management

import (
	"exercise/model"

	"gorm.io/gorm"
)

type ArticleTagManager struct {
	db *gorm.DB
}

func (atm *ArticleTagManager) Select(at *model.ArticleTag) {
	return
}

func (atm *ArticleTagManager) Insert(at *model.ArticleTag) {
	return
}

func (atm *ArticleTagManager) Update(at *model.ArticleTag) {
	return
}

func (atm *ArticleTagManager) Delete(at *model.ArticleTag) {
	return
}
