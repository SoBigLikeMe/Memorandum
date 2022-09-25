package serialzer

import "memorandum/model"

type Task struct {
	Id        uint   `json:"id" example:"1"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      uint64 `json:"view"`
	Status    int    `json:"status"`
	CreateAt  int64  `json:"createAt"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

// BuildTask 序列化备忘录
func BuildTask(item model.Task) Task {
	return Task{
		Id:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreateAt:  item.CreatedAt.Unix(),
		StartTime: item.StratTime,
		EndTime:   item.EndTime,
	}
}
