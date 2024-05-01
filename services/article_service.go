package services

import (
	"github.com/capomanpc/go-blog-api/models"
	"github.com/capomanpc/go-blog-api/repositories"
)

// ArticleDetailHandlerの処理をサービス化
func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// '...'はスライスの展開
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}
