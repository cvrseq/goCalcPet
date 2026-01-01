package service

type RequestBody struct {
	Task   string `gorm:"column:task" json:"task"`
	ID     uint   `gorm:"primaryKey" json:"id"`
	IsDone bool   `gorm:"column:is_done" json:"is_done"`
}
