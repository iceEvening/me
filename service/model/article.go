package model

import (
	"github.com/jinzhu/gorm"

	"github.com/iceEvening/me/service/model/tables"
)

//NewArticle model func
func (m *Model) NewArticle(userID uint, nickname string, title string, markdown string, html string) error {
	article := tables.Article{
		UserID:   userID,
		NickName: nickname,
		Title:    title,
		Markdown: markdown,
		HTML:     html,
	}
	err := m.DB.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

//EditArticle model func
func (m *Model) EditArticle(userID uint, articleID uint, nickname string, title string, markdown string, html string) error {
	article := tables.Article{
		Model:    gorm.Model{ID: articleID},
		UserID:   userID,
		NickName: nickname,
		Title:    title,
		Markdown: markdown,
		HTML:     html,
	}

	err := m.DB.Save(&article).Error
	if err != nil {
		return err
	}
	return nil
}

//DelArticle model func
func (m *Model) DelArticle(userID uint, articleID uint) error {
	article := tables.Article{
		Model:  gorm.Model{ID: articleID},
		UserID: userID,
	}
	err := m.DB.Delete(&article).Error
	if err != nil {
		return err
	}
	return nil
}

//Article model func
func (m *Model) Article(ID uint) (*tables.Article, error) {
	article := tables.Article{
		Model: gorm.Model{ID: ID},
	}
	err := m.DB.Where(&article).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

//Articles model func
func (m *Model) Articles(ID uint) (*[]tables.Article, error) {
	var articles []tables.Article
	err := m.DB.Where("user_id = ?", ID).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	for i := range articles {
		articles[i].HTML = ""
		articles[i].Markdown = ""
	}
	return &articles, nil
}
