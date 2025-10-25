package application

import (
	"context"
	"fmt"
	"time"

	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

// HardDeleteUsers is a use case for permanently deleting users.
type HardDeleteUsers struct {
	UserDeletionRepository repositories.UserDeletionRepository
	UserRepository         repositories.UserRepository
}

// NewHardDeleteUsers creates a new HardDeleteUsers use case.
func NewHardDeleteUsers(userDeletionRepo repositories.UserDeletionRepository, userRepo repositories.UserRepository) *HardDeleteUsers {
	return &HardDeleteUsers{
		UserDeletionRepository: userDeletionRepo,
		UserRepository:         userRepo,
	}
}

// Execute permanently deletes users whose deletion is scheduled for today.
func (uc *HardDeleteUsers) Execute(ctx context.Context) error {
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// Find user deletions scheduled for today
	deletionsToExecute, err := uc.UserDeletionRepository.GetScheduledDeletions(ctx, today)
	if err != nil {
		return fmt.Errorf("failed to get scheduled deletions for today: %w", err)
	}

	for _, userDeletion := range deletionsToExecute {
		// Perform hard deletion of the user
		if err := uc.UserRepository.HardDelete(ctx, userDeletion.UserID); err != nil {
			// Log the error but continue with other deletions
			fmt.Printf("Error hard deleting user %s: %v\n", userDeletion.UserID, err)
			continue
		}

		// Mark the user deletion as executed
		userDeletion.MarkAsExecuted()
		if err := uc.UserDeletionRepository.Update(ctx, userDeletion); err != nil {
			fmt.Printf("Error marking user deletion %s as executed: %v\n", userDeletion.ID, err)
		}
	}

	return nil
}
