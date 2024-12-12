package handlers

import (
	"backend/internals/models"
	"backend/internals/services"
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	service *services.JobService
}

func NewJobHandler(db *sql.DB) *JobHandler {
	return &JobHandler{service: services.NewJobService(db)}
}

func (s *JobHandler) CreateJobHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var job models.Job

		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		job.UserID = userID

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		cratedJob, err := s.service.CreateJob(ctx, &job)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, cratedJob)
	}
}

func (s *JobHandler) GetAllJobsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		jobs, err := s.service.GetAllJobs(ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, jobs)
	}
}

func (s *JobHandler) GetAllJobsByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		jobs, err := s.service.GetAllJobsByUserID(ctx, c.GetInt("userID"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, jobs)
	}
}

func (s *JobHandler) GetJobByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
			return
		}

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		job, err := s.service.GetJobByID(ctx, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, job)
	}
}

func (s *JobHandler) UpdateJobByHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
			return
		}

		var job models.Job
		job.ID = id

		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		updateJob, err := s.service.UpdateJob(ctx, &job, userID, isAdmin)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updateJob)

	}
}

func (s *JobHandler) DeleteJobByHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		err = s.service.DeleteJob(ctx, id, userID, isAdmin)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
	}
}
