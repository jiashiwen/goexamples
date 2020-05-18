package task

import (
	"examples/commons"
	"github.com/jinzhu/gorm"
	"strconv"
	"testing"
)

type TaskTests struct{ Test *testing.T }

func (t *TaskTests) TestInitLocalStatusDB() {
	log.Info("Start TestInitLocalStatusDB")
	if commons.FileExists("../status.db") {
		commons.RemoveFile("../status.db")
	}
	InitLocalStatusDB("../")

}

func (t *TaskTests) TestCRUD() {

	var taskarry []*Task
	db, err := gorm.Open("sqlite3", "../status.db")
	if err != nil {
		t.Test.Error(err)
	}
	for i := 0; i < 10; i++ {
		task := NewTask()
		msg := "Insert:" + task.TaskId + "|" + strconv.Itoa(task.TaskStatus)
		log.Debug(msg)
		task.CreateTask(db)
		taskarry = append(taskarry, task)
	}

	for _, task := range taskarry {

		task2 := GetTaskFromDbById(db, task.TaskId)
		if len(task2) != 1 {
			t.Test.Error("Task not in db")
		}
	}

	for _, task := range taskarry {
		task.RemoveTaskById(db)
	}

	for _, task := range taskarry {
		task3 := GetTaskFromDbById(db, task.TaskId)
		if len(task3) != 0 {
			t.Test.Error("Task not be reomved")
		}
	}

}

func TestTask(t *testing.T) {
	t.Run("Task=initdb", func(t *testing.T) {
		test := TaskTests{Test: t}
		test.TestInitLocalStatusDB()
		test.TestCRUD()

	})
}

//func TestMain(m *testing.M) {
//	defer log.Sync()
//	log.Info("Start TestMain")
//	m.Run()
//
//}
