package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"context"
	"database/sql"
	"errors"
)

type JobService struct {
	store store.JobStore
}

func NewJobService(db *sql.DB) *JobService {
	return &JobService{store: store.NewJobStore(db)}
}

func (s *JobService) CreateJob(ctx context.Context, job *models.Job) (*models.Job, error) {
	return s.store.CreateJob(ctx, job)
}

func (s *JobService) GetAllJobs(ctx context.Context) ([]models.Job, error) {
	return s.store.GetAllJobs(ctx)
}

func (s *JobService) GetAllJobsByUserID(ctx context.Context, userID int) ([]models.Job, error) {
	return s.store.GetAllJobsByUserID(ctx, userID)
}

func (s *JobService) GetJobByID(ctx context.Context, id int) (*models.Job, error) {
	return s.store.GetJobByID(ctx, id)
}

func (s *JobService) UpdateJob(ctx context.Context, job *models.Job, userID int, isAdmin bool) (*models.Job, error) {
	existingJob, err := s.store.GetJobByID(ctx, job.ID)

	if err != nil {
		return nil, err
	}

	if !isAdmin && existingJob.UserID != userID {
		return nil, errors.New("unauthorized to update this job")
	}

	return s.store.UpdateJob(ctx, job)
}

func (s *JobService) DeleteJob(ctx context.Context, id int, userID int, isAdmin bool) error {
	existingJob, err := s.store.GetJobByID(ctx, id)

	if err != nil {
		return err
	}

	if !isAdmin && existingJob.UserID != userID {
		return errors.New("unauthorized to delete this job")
	}

	return s.store.DeleteJob(ctx, id)
}
