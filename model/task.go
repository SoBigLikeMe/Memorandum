package model

import "github.com/jinzhu/gorm"

//建表Task
type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"defalut:'0'"` //0为未完成，1为已完成
	Content   string `gorm:"type:longtext"`
	StratTime int64  //备忘录开始时间
	EndTime   int64  // 备忘录完成时间
}
