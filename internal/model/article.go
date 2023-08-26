package model

import "blog-service/pkg/app"

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// 文章
type Article struct {
	*Model               //引用公共欄位
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
