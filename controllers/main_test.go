package controllers_test

import (
	"testing"

	"github.com/capomanpc/go-blog-api/controllers"
	"github.com/capomanpc/go-blog-api/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
