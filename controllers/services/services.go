package services

import "github.com/capomanpc/go-blog-api/models"

// MyAppService構造体のメソッドを実装する全ての型をMyAppControllerに代入できるように
// 新たにMyAppServiceインターフェースを定義
// MyAppService構造体はMyAppServiceインターフェースを実装しているのでMyAppController構造体に代入できる

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
