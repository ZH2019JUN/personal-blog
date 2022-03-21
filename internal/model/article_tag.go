package model

import (
	"github.com/jinzhu/gorm"
)

//文章标签的关联模型

type ArticleTag struct {
	*Model
	TagID uint32  `json:"tag_id"`
	ArticleID uint32  `json:"article_id"`
}

func (a ArticleTag)TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag)GetByAID(db *gorm.DB) (ArticleTag,error) {
	var articleTag ArticleTag
	err := db.Where("article_id=? AND is_del=?",a.ArticleID,0).First(&articleTag).Error
	if err != nil&&err != gorm.ErrRecordNotFound{
		return articleTag,err
	}
	return articleTag, err
}

func (a ArticleTag)ListByTID(db *gorm.DB) ([]*ArticleTag,error) {
	var articleTags []*ArticleTag
	err := db.Where("tag_id=? AND is_del=?",a.TagID,0).Find(&articleTags).Error
	if err != nil{
		return nil, err
	}
	return articleTags,nil
}

func (a ArticleTag)ListByAIDs(db *gorm.DB,articleIDs []uint32) ([]*ArticleTag,error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND is_del=?",articleIDs,0).Find(articleTags).Error
	if err != nil{
		return nil, err
	}
	return articleTags,nil
}

func (a ArticleTag)Create(db *gorm.DB) error {
	err := db.Create(&a).Error
	if err != nil{
		return err
	}
	return nil
}

func (a ArticleTag)UpdateOne(db *gorm.DB,values interface{}) error {
	err := db.Model(&a).Where("article_id=? AND is_del=?",a.ArticleID,0).Limit(1).Updates(values).Error
	if err != nil{
		return err
	}
	return nil
}

func (a ArticleTag)Delete(db *gorm.DB) error {
	err := db.Where("id=? AND is_del=?",a.Model.ID,0).Delete(&a).Error
	if err != nil{
		return err
	}
	return nil
}

func (a ArticleTag)DeleteOne(db *gorm.DB) error {
	err := db.Where("article_id=? AND is_del=?",a.ArticleID,0).Delete(&a).Limit(1).Error
	if err != nil{
		return err
	}
	return nil
}
