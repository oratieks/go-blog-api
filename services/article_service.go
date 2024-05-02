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
	defer db.Close()

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

// PostArticleHandlerの処理をサービス化
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		// 構造体は値型なので初期値である空の構造体を返す
		return models.Article{}, err
	}

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

// GetArticleListHandlerの処理をサービス化
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		// スライスは参照型なので初期値であるnilを返す
		return nil, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

// PostNiceHandlerの処理をサービス化
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	newArticle := article
	newArticle.NiceNum = newArticle.NiceNum + 1

	return newArticle, nil
}
