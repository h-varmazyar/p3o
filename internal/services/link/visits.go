package link

import "context"

func (s Service) TotalVisits(ctx context.Context, userId uint) (uint, error) {
	visits, err := s.linkRepo.Visits(ctx, userId)
	if err != nil {
		return 0, err
	}

	return visits, nil
}
