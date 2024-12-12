package store

import (
	"backend/internals/models"
	"context"
	"database/sql"
)

type JobStore interface {
	CreateJob(ctx context.Context, job *models.Job) (*models.Job, error)
	DeleteJob(ctx context.Context, id int) error
	GetAllJobs(ctx context.Context) ([]models.Job, error)
	GetAllJobsByUserID(ctx context.Context, userID int) ([]models.Job, error)
	GetJobByID(ctx context.Context, id int) (*models.Job, error)
	UpdateJob(ctx context.Context, job *models.Job) (*models.Job, error)
}

type SQLiteJobStore struct {
	db *sql.DB
}

func NewJobStore(db *sql.DB) *SQLiteJobStore {
	return &SQLiteJobStore{db: db}
}

func (s *SQLiteJobStore) CreateJob(ctx context.Context, job *models.Job) (*models.Job, error) {
	result, err := s.db.Exec("INSERT INTO jobs (title, description, company, location, salary, user_id) VALUES (?, ?, ?, ?, ?, ?)", job.Title,
		job.Description, job.Company, job.Location, job.Salary, job.UserID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	job.ID = int(id)
	return job, nil
}

func (s *SQLiteJobStore) GetAllJobs(ctx context.Context) ([]models.Job, error) {
	rows, err := s.db.Query("SELECT * FROM jobs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Company, &job.Salary, &job.UserID, &job.CreatedAt); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (s *SQLiteJobStore) GetAllJobsByUserID(ctx context.Context, userID int) ([]models.Job, error) {
	rows, err := s.db.Query("SELECT * FROM jobs WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Company, &job.Salary, &job.UserID, &job.CreatedAt); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (s *SQLiteJobStore) GetJobByID(ctx context.Context, id int) (*models.Job, error) {
	job := &models.Job{}
	row := s.db.QueryRow("SELECT * FROM jobs WHERE id = ?", id)

	if err := row.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Company, &job.Salary, &job.UserID, &job.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return job, nil
}

func (s *SQLiteJobStore) UpdateJob(ctx context.Context, job *models.Job) (*models.Job, error) {
	_, err := s.db.Exec("UPDATE jobs SET title = ?, description = ?, company = ?, location = ?, salary = ? WHERE id = ?", job.Title,
		job.Description, job.Company, job.Location, job.Salary, job.ID)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *SQLiteJobStore) DeleteJob(ctx context.Context, id int) error {
	_, err := s.db.Exec("DELETE FROM jobs WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
