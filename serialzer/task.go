package serialzer

import "memorandum/model"

type Task struct {
	Id        uint   `json:"id" example:"1"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int    `json:"status"`
	CreateAt  int64  `json:"createAt"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

//DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
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

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}

//BulidListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
