package task

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

const (
	Type_Sampletask string = "SampleTask"
)

type SampleTask struct {
	TaskInMem *TaskInMem
	StatusDb  *gorm.DB
}

func NewSampleTask() *SampleTask {

	taskinmem := TaskInMem{
		Task: NewTask(),
	}
	taskinmem.Task.TaskType = Type_Sampletask
	return &SampleTask{
		TaskInMem: &taskinmem,
	}
}

func (sample *SampleTask) Create() {
	sample.TaskInMem.Task.CreateTask(sample.StatusDb)
}

func (sample *SampleTask) Start() error {
	log.Info("startSamplTask...")

	pool := GetTaskPool()
	pool.Free()
	if pool.Free() < 1 {
		return errors.New("Task pool full")
	}
	ctx, cancel := context.WithCancel(context.Background())
	sample.TaskInMem.Cancel = &cancel
	sample.TaskInMem.Task.TaskStatus = int(Task_Running)
	submit := func() {
		execSampleTask(ctx, sample)
	}
	err := pool.Submit(submit)
	if err != nil {
		return err
	}

	sample.TaskInMem.PutToTaskMap()
	sample.TaskInMem.Task.UpdateTaskStatusById(sample.StatusDb)

	return nil
}

func (sample *SampleTask) Stop() {
	sample.TaskInMem.Task.TaskStatus = int(Task_Stop)
	sample.TaskInMem.Task.UpdateTaskStatusById(sample.StatusDb)
	sample.TaskInMem.StopTask()
}

func execSampleTask(ctx context.Context, t *SampleTask) {

	defer t.TaskInMem.Task.UpdateTaskStatusById(t.StatusDb)
	defer t.TaskInMem.RemoveFromTaskMap()
	for {
		select {
		case <-ctx.Done():
			t.TaskInMem.Task.TaskStatus = int(Task_Stop)
			return
		default:
			log.Info(t.TaskInMem.Task.TaskId + ":" + time.Now().String())
			time.Sleep(1 * time.Second)
		}
	}

}
