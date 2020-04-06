package task

import (
	"context"
	"examples/globlezap"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"reflect"
	"sync"
)

var runningtaskmap cmap.ConcurrentMap
var once sync.Once
var taskpool *ants.Pool
var options = ants.Options{
	PreAlloc: true,
}
var l = &sync.RWMutex{}

var log = globlezap.GetLogger()

type TaskInMem struct {
	Task   *Task
	Cancel *context.CancelFunc
}

func NewTask() *Task {
	return &Task{
		TaskId:     uuid.NewV4().String(),
		TaskStatus: 0,
	}
}

func GetRunningTaskMap() *cmap.ConcurrentMap {
	once.Do(func() {
		runningtaskmap = cmap.New()
	})
	return &runningtaskmap
}

func GetTaskInMemById(taskid string) (*TaskInMem, bool) {
	taskmap := GetRunningTaskMap()
	task, ok := taskmap.Get(taskid)
	return task.(*TaskInMem), ok
}

func GetTaskPool() *ants.Pool {
	if taskpool == nil {
		l.Lock()
		defer l.Unlock()
		if taskpool == nil {
			InitTaskPool(100, ants.WithOptions(options))
		}
	}
	return taskpool
}

func InitTaskPool(size int, options ...ants.Option) {
	var err error
	taskpool, err = ants.NewPool(100, options...)
	if err != nil {
		fmt.Println("pool init error")
		panic(err)
	}
}

func (t *TaskInMem) PutToTaskMap() error {
	if t.Task.TaskId == "" {
		return errors.New("TaskId Missing")
	}
	taskmap := GetRunningTaskMap()
	taskmap.Set(t.Task.TaskId, t)
	return nil
}

func (t *TaskInMem) RemoveFromTaskMap() {
	taskmap := GetRunningTaskMap()
	taskmap.Remove(t.Task.TaskId)
}

func (t *TaskInMem) TaskIsRunning() bool {
	taskmap := GetRunningTaskMap()
	return taskmap.Has(t.Task.TaskId)
}

func (t *TaskInMem) StopTask() error {
	taskmap := GetRunningTaskMap()
	if !taskmap.Has(t.Task.TaskId) {
		return errors.New("task not exixts in runningmap")
	}
	cancel := *t.Cancel
	cancel()
	t.RemoveFromTaskMap()
	return nil
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
