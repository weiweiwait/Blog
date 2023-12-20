package model

// 创建文章model
type Article struct {
	*Model
	Title         string `json:"title" gorm:"title"`
	Desc          string `json:"desc" gorm:"desc"`
	Content       string `json:"content" gorm:"content"`
	CoverImageUrl string `json:"cover_image_url" gorm:"cover_image_url"`
	State         uint8  `json:"state" gorm:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
