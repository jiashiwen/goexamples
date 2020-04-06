package task

import (
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

type TestSampleTasks struct {
	Test       *testing.T
	SampleTask *SampleTask
}

func (t *TestSampleTasks) TestSampleTask_Create() {
	t.SampleTask.Create()
	tasks := GetTaskFromDbById(t.SampleTask.StatusDb, t.SampleTask.TaskInMem.Task.TaskId)
	if t.SampleTask.TaskInMem.Task.TaskId != (tasks[0].TaskId) {
		t.Test.Error("not get created task")
	}

	t.SampleTask.TaskInMem.Task.RemoveTaskById(t.SampleTask.StatusDb)
}

func (t *TestSampleTasks) TestSampleTask_Start() {
	t.SampleTask.Create()
	tasks := GetTaskFromDbById(t.SampleTask.StatusDb, t.SampleTask.TaskInMem.Task.TaskId)
	if t.SampleTask.TaskInMem.Task.TaskId != (tasks[0].TaskId) {
		t.Test.Error("not get created task")
	}
}

func (t *TestSampleTasks) TestSampleTask_Stop() {
	t.SampleTask.Stop()
	time.Sleep(2 * time.Second)

	tasks := GetTaskFromDbById(t.SampleTask.StatusDb, t.SampleTask.TaskInMem.Task.TaskId)
	if tasks[0].TaskStatus != 0 {
		t.Test.Error("Task status wrong")
	}

}

func TestSampleTask(t *testing.T) {
	t.Run("TestSampleTask=alltest", func(t *testing.T) {
		db, err := gorm.Open("sqlite3", "../status.db")
		defer db.Close()
		if err != nil {
			t.Error(err)
		}
		testsample := NewSampleTask()
		testsample.StatusDb = db

		ts := TestSampleTasks{
			Test:       t,
			SampleTask: testsample,
		}
		ts.TestSampleTask_Create()
		ts.TestSampleTask_Start()
		ts.TestSampleTask_Stop()

	})
}

//func TestSampleTask_Start(t *testing.T) {
//	db, err := gorm.Open("sqlite3", "../status.db")
//	defer db.Close()
//	if err != nil {
//		t.Error(err)
//	}
//
//	sampletask := NewSampleTask()
//	sampletask.StatusDb = db
//	//创建任务
//	sampletask.Create()
//	tasks := GetTaskFromDbById(sampletask.StatusDb, sampletask.TaskInMem.Task.TaskId)
//	if sampletask.TaskInMem.Task.TaskId != (tasks[0].TaskId) {
//		t.Error("not get created task")
//	}
//
//	//启动任务
//	sampletask.Start()
//	time.Sleep(2 * time.Second)
//	task, _ := GetTaskInMemById(sampletask.TaskInMem.Task.TaskId)
//	tasks = GetTaskFromDbById(sampletask.StatusDb, sampletask.TaskInMem.Task.TaskId)
//	if task.Task.TaskStatus != 1 {
//		t.Error("Task status wrong")
//	}
//	//停止任务
//	sampletask.Stop()
//	time.Sleep(2 * time.Second)
//
//	tasks = GetTaskFromDbById(sampletask.StatusDb, sampletask.TaskInMem.Task.TaskId)
//	if sampletask.TaskInMem.Task.TaskStatus != 0 {
//		t.Error("Task status wrong")
//	}
//
//}
