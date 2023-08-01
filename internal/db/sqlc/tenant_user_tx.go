package db

import "context"

type CreateUserArgs struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TenantUserTxInput struct {
	TenantData CreateTenantParams `json:"tenant"`
	UserData   CreateUserArgs     `json:"user"`
}

type TenantUserTxOutput struct {
	Message string `json:"message"`
	Tenant  Tenant `json:"tenant"`
	User    User   `json:"user"`
}

func (s *Store) CreateTenantUserTx(ctx context.Context, input TenantUserTxInput) (TenantUserTxOutput, error) {
	var output TenantUserTxOutput
	err := s.execTx(ctx, func(q *Queries) error {

		// stripe payment first

		t, err := q.CreateTenant(ctx, input.TenantData)

		if err != nil {
			return err
		}

		output.User, err = q.CreateUser(ctx, CreateUserParams{
			Username: input.UserData.Username,
			Email:    input.UserData.Email,
			TenantID: t.ID,
			Password: input.UserData.Password,
		})

		if err != nil {
			return err
		}

		output.Tenant = t

		return nil
	})
	return output, err
}
