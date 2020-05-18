package httpgin

import (
	"errors"
	"examples/globledatasource"
	task "examples/task"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type CreateRequest struct {
	// TaskName define your taskname
	TaskName  string
	TaskType  string
	AutoStart bool
}

// @Summary Create Task
// @Tags 创建任务
// @version 1.0
// @Param CreateRequest body httpgin.CreateRequest true  "json for createtask"
// @Accept application/x-json-stream
// @Success 200  {object} httpgin.Response
// @Router /api/v1/createtask [POST]
func CreateTask(c *gin.Context) {
	var ct CreateRequest
	c.Bind(&ct)

	err := ct.Check()
	if err != nil {
		Logger.Error("Error", zap.String("error", err.Error()))
		c.JSON(http.StatusOK, newAPIException(123, err.Error()))
		return
	}

	task := task.NewTask()
	task.TaskType = strings.ToLower(ct.TaskType)
	db := globledatasource.GetDb()
	task.CreateTask(db)
	response := Response{
		Code:    "2000",
		Message: task,
		Error:   nil,
	}

	c.JSON(http.StatusOK, response)

}

func (cr *CreateRequest) Check() error {
	if _, ok := task.TaskMap[strings.ToLower(cr.TaskType)]; !ok {
		return errors.New("TaskType not support")
	}
	return nil
}
