package entities

type TaskStatus string

const (
	InProgress TaskStatus = "ทำอยู่"
	Done       TaskStatus = "ทำเสร็จแล้ว"
)

type TaskModel struct {
	TaskID  string     `json:"task_id" bson:"task_id,omitempty"`
	Name    string     `json:"name" bson:"name,omitempty"`
	Details string     `json:"details" bson:"details,omitempty"`
	Status  TaskStatus `json:"status" bson:"status,omitempty"`
}
