package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/entities"
)

func (s Service)Deactivate(ctx context.Context, userId uint, key string) error {
	link, err:=s.linkRepo.ReturnByKey(ctx, key)
	if err!=nil{
		return err
	}

	if link.OwnerId != userId {
		return ErrLinkOwnerMismatch
	}

	switch link.Status {
	case entities.LinkStatusDeactivatedByUser:
		return ErrLinkDeactivatedBefore
	case entities.LinkStatusDeactivatedByAdmin:
		return ErrLinkActivationBanned
	case entities.LinkStatusActive:
		fallthrough
	default:
		link.Status = entities.LinkStatusDeactivatedByUser
		if err=s.linkRepo.Update(ctx, link); err!=nil{
			return err
		}
		return nil
	}
}