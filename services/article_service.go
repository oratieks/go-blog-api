package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/capomanpc/go-blog-api/apperrors"
	"github.com/capomanpc/go-blog-api/models"
	"github.com/capomanpc/go-blog-api/repositories"
)

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		// ErrCode型のInsertDataFailed定数のWrapメソッドを使用してラップされたエラーを返す
		// InsertDataFailedの部分は通常ErrCode型の変数だが、今回の場合はErrCode型の定数を使用している
		// 普段は変数だけど今回は定数が使用されている
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// ArticleDetailHandlerで使うことを想定したサービス
// 指定IDの記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// ゴルーチン内部で宣言された変数は、そのゴルーチン内でのみ有効なため、値を受け取る変数を定義している
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error // 非同期で実行されるのでエラー変数も二つ作る必要がある

	var wg sync.WaitGroup
	wg.Add(2)

	var amu sync.Mutex
	var cmu sync.Mutex

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		newArticle, err := repositories.SelectArticleDetail(db, articleID)
		amu.Lock() // 異なるインスタンスを使用する
		article, articleGetErr = newArticle, err
		amu.Unlock()
	}(s.db, articleID) // 引数を渡している

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		newCommentList, err := repositories.SelectCommentList(db, articleID)
		cmu.Lock() // 異なるインスタンスを使用する
		commentList, commentGetErr = newCommentList, err
		cmu.Unlock()
	}(s.db, articleID)

	wg.Wait()

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定IDの記事のいいね数を+1して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
