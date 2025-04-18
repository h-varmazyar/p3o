package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Activate(ctx context.Context, userId uint, key string) error {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return err
	}

	if link.OwnerId != userId {
		return errors.ErrLinkOwnerMismatch
	}

	switch link.Status {
	case entities.LinkStatusActive:
		return errors.ErrLinkActivatedBefore
	case entities.LinkStatusDeactivatedByAdmin:
		return errors.ErrLinkActivationBanned
	case entities.LinkStatusDeactivatedByUser:
		fallthrough
	default:
		link.Status = entities.LinkStatusActive
		if err = s.linkRepo.Update(ctx, link); err != nil {
			return err
		}
		return nil
	}
}
