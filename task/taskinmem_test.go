package task

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask()
	fmt.Println(task.TaskId)
}

func TestGetTaskMap(t *testing.T) {
	taskmap := GetRunningTaskMap()
	task := NewTask()
	taskinmem := TaskInMem{
		Task: task,
	}
	//(*taskmap)[task.TaskId] = task
	taskmap.Set(task.TaskId, &taskinmem)

	taskprt, ok := GetTaskInMemById(task.TaskId)
	if !ok {
		t.Error("task not exists")
	}
	fmt.Println(taskprt.Task.TaskId)
}

func TestGetTaskPool(t *testing.T) {
	var wg sync.WaitGroup
	pool := GetTaskPool()
	//defer pool.Release()

	for i := 0; i < 10; i++ {
		num := i
		submite := func() {
			fmt.Println("submet one task", num)
			wg.Done()
		}
		wg.Add(1)
		err := pool.Submit(submite)
		if err != nil {
			t.Error(err)
		}
	}

	wg.Wait()

}
