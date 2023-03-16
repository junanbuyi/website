package db

import (
	"test/config"
)

// 定义文章模型
type Article struct {
	Id      int    `DB:"id"`
	Title   string `DB:"Title"`
	Content string `DB:"Content"`
}

// 获取所有文章
func GetAllArticles() ([]Article, error) {
	var articles []Article
	rows, err := config.DB.Queryx("SELECT id, title, content FROM bbs")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Title, &article.Content)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// 获取单个文章
func GetArticleById(id uint) (*Article, error) {
	var article Article
	_, err := config.DB.Query("select  limit 1", &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 创建文章
func CreateArticle(article *Article) error {
	_, err := config.DB.Exec("")
	if err != nil {
		return err
	}
	return nil
}

// 更新文章
func UpdateArticle(article *Article) error {
	_, err := config.DB.Exec("")
	if err != nil {
		return err
	}
	return nil
}
