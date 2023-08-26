package service

//用ID查指定文章
type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=01"`
}

//用Tag查詢文章清單
type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=01"`
}

//創建文章
type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	State         uint8  `form:"state,default=1" binding:"oneof=01"`
}

//更新文章
type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	State         uint8  `form:"state,default=1" binding:"oneof=01"`
}

//刪除文章
type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
