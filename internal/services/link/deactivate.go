package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/entities"
<<<<<<< HEAD
=======
	"github.com/h-varmazyar/p3o/internal/errors"
>>>>>>> 292128d (feat: add link creation)
)

func (s Service)Deactivate(ctx context.Context, userId uint, key string) error {
	link, err:=s.linkRepo.ReturnByKey(ctx, key)
	if err!=nil{
		return err
	}

	if link.OwnerId != userId {
<<<<<<< HEAD
		return ErrLinkOwnerMismatch
=======
		return errors.ErrLinkOwnerMismatch
>>>>>>> 292128d (feat: add link creation)
	}

	switch link.Status {
	case entities.LinkStatusDeactivatedByUser:
<<<<<<< HEAD
		return ErrLinkDeactivatedBefore
	case entities.LinkStatusDeactivatedByAdmin:
		return ErrLinkActivationBanned
=======
		return errors.ErrLinkDeactivatedBefore
	case entities.LinkStatusDeactivatedByAdmin:
		return errors.ErrLinkActivationBanned
>>>>>>> 292128d (feat: add link creation)
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