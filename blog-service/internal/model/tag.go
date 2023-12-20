package model

// 创建标签model
type Tag struct {
	*Model
	Name  string `json:"name" gorm:"name"`
	State uint8  `json:"state" gorm:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
