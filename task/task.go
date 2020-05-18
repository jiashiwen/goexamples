package task

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"examples/commons"

	uuid "github.com/satori/go.uuid"
)

type TaskStatus int32

const (
	Task_Stop       TaskStatus = 0
	Task_Running    TaskStatus = 1
	Type_Sampletask string     = "sampletask"
)

var TaskMap = map[string]string{
	Type_Sampletask: Type_Sampletask,
}

type Task struct {
	TaskId     string `gorm:"column:taskid"`
	TaskStatus int    `gorm:"column:status"`
	TaskType   string `gorm:"column:tasktype"`
}

func InitLocalStatusDB(path string) {
	dbfile := path + "status.db"
	if !commons.FileExists(dbfile) {
		_, err := commons.CreateNewFile(dbfile)
		if err != nil {
			panic(err)
		}

		db, err := sql.Open("sqlite3", dbfile)
		defer db.Close()
		if err != nil {
			panic(err)
		}

		dot, err := dotsql.LoadFromFile(path + "initdb.sql")
		if err != nil {
			panic(err)
		}

		_, err = dot.Exec(db, "create-tasks-table")
		if err != nil {
			panic(err)
		}

	}
}

func NewTask() *Task {
	return &Task{
		TaskId:     uuid.NewV4().String(),
		TaskStatus: 0,
	}
}

func (t *Task) CreateTask(db *gorm.DB) {
	db.Create(t)
}

func GetTaskFromDbById(db *gorm.DB, id string) []Task {
	var task []Task
	db.Where("taskid = ?", id).Find(&task)

	return task
}

func (t *Task) UpdateTaskStatusById(db *gorm.DB) {
	db.Model(t).Where("taskid = ?", t.TaskId).Update("status", t.TaskStatus)
}

func (t *Task) RemoveTaskById(db *gorm.DB) *Task {
	db.Delete(Task{}, "taskid = ?", t.TaskId)
	return t
}
