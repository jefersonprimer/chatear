package application

import (
	"context"
	"fmt"
	"time"

	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

// SchedulePermanentDeletions is a use case for scheduling permanent user deletions.
type SchedulePermanentDeletions struct {
	UserDeletionRepository    repositories.UserDeletionRepository
	DeletionCapacityRepository repositories.DeletionCapacityRepository
	MaxDeletionsPerDay        int
}

// NewSchedulePermanentDeletions creates a new SchedulePermanentDeletions use case.
func NewSchedulePermanentDeletions(userDeletionRepo repositories.UserDeletionRepository, deletionCapacityRepo repositories.DeletionCapacityRepository, maxDeletionsPerDay int) *SchedulePermanentDeletions {
	return &SchedulePermanentDeletions{
		UserDeletionRepository:    userDeletionRepo,
		DeletionCapacityRepository: deletionCapacityRepo,
		MaxDeletionsPerDay:        maxDeletionsPerDay,
	}
}

// Execute schedules permanent deletions for users whose deletion due date has passed.
func (uc *SchedulePermanentDeletions) Execute(ctx context.Context) error {
	// Find users whose deletion due date has passed and are not yet scheduled for hard deletion
	usersToSchedule, err := uc.UserDeletionRepository.FindScheduledBefore(ctx, time.Now())
	if err != nil {
		return fmt.Errorf("failed to find users to schedule for permanent deletion: %w", err)
	}

	for _, userDeletion := range usersToSchedule {
		// Get today's deletion capacity
		now := time.Now().UTC()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		capacity, err := uc.DeletionCapacityRepository.GetDeletionCapacity(ctx, today)
		if err != nil {
			return fmt.Errorf("failed to get deletion capacity for %s: %w", today.Format("2006-01-02"), err)
		}

		if capacity == nil {
			// If no capacity entry for today, create one
			capacity = entities.NewDeletionCapacity(today, uc.MaxDeletionsPerDay)
		}

		if capacity.CanDelete() {
			// Schedule for today
			userDeletion.ScheduledDate = today
			userDeletion.Status = entities.UserDeletionStatusScheduled
			if err := uc.UserDeletionRepository.Update(ctx, userDeletion); err != nil {
				return fmt.Errorf("failed to update user deletion status for user %s: %w", userDeletion.UserID, err)
			}

			// Increment today's deletion count
			if err := uc.DeletionCapacityRepository.IncrementDeletionCapacity(ctx, today); err != nil {
				return fmt.Errorf("failed to increment deletion capacity for %s: %w", today.Format("2006-01-02"), err)
			}
		} else {
			// Find the next available day
			nextAvailableDay := today.Add(24 * time.Hour)
			for {
				nextCapacity, err := uc.DeletionCapacityRepository.GetDeletionCapacity(ctx, nextAvailableDay)
				if err != nil {
					return fmt.Errorf("failed to get deletion capacity for %s: %w", nextAvailableDay.Format("2006-01-02"), err)
				}

				if nextCapacity == nil {
					nextCapacity = entities.NewDeletionCapacity(nextAvailableDay, uc.MaxDeletionsPerDay)
				}

				if nextCapacity.CanDelete() {
					userDeletion.ScheduledDate = nextAvailableDay
					userDeletion.Status = entities.UserDeletionStatusScheduled
					if err := uc.UserDeletionRepository.Update(ctx, userDeletion); err != nil {
						return fmt.Errorf("failed to update user deletion status for user %s: %w", userDeletion.UserID, err)
					}
					// Increment the next available day's deletion count
					if err := uc.DeletionCapacityRepository.IncrementDeletionCapacity(ctx, nextAvailableDay); err != nil {
						return fmt.Errorf("failed to increment deletion capacity for %s: %w", nextAvailableDay.Format("2006-01-02"), err)
					}
					break
				} else {
					nextAvailableDay = nextAvailableDay.Add(24 * time.Hour)
				}
			}
		}
	}

	return nil
}
