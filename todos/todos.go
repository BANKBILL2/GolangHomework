package todos

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/pallat/todos/logger"
)

func NewNewTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		var todo struct {
			Task string `json:"task"`
		}

		logger := logger.Extract(c)
		logger.Info("new task todo........")

		if err := c.Bind(&todo); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.Wrap(err, "new task").Error(),
			})
		}

		if err := db.Create(&Task{
			Task: todo.Task,
		}).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "create task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}

//Homework
func NewGetTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("get task todo")
		var todo []Task
		if err := db.Find(&todo).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "get task").Error(),
			})
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func NewPutTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("put task todo%")
		id := c.Param("id")

		var todo Task

		if err := db.Model(&todo).Where("id = ?", id).Update("processed", true).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "put task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}

type Task struct {
	gorm.Model
	Task      string
	Processed bool
}

func (Task) TableName() string {
	return "todos"
}
