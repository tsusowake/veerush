package process

import (
	"context"
)

func (p *Process) CreateUser(
	ctx context.Context,
) (*uint64, error) {
	userID, err := p.User.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	return userID, nil
}
