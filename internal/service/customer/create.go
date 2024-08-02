package customer

import (
	"context"
	"log"

	"github.com/google/uuid"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

func (s *service) Create(ctx context.Context, info *model.CustomerInfo) (uint32, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Failed to generate uuid: %v\n", err)
		return 0, err
	}
	err = s.repo.Create(ctx, userUUID.ID(), info)
	if err != nil {
		log.Printf("Failed to create user: %v\n", err)
		return 0, err
	}

	return userUUID.ID(), nil
}
