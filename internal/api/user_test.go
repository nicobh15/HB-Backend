package api

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	mockdb "github.com/nicobh15/hb-backend/internal/db/mock"
	db "github.com/nicobh15/hb-backend/internal/db/sqlc"
	"github.com/nicobh15/hb-backend/internal/util"
	"go.uber.org/mock/gomock"
)

func TestFechUserAPI(t *testing.T) {
	user := randomUser()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		FetchUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
		Times(1).
		Return(user, nil)

	// server := NewServer(store)
}

func randomUser() db.User {
	return db.User{
		UserID:      util.RandomUUID(),
		Username:    util.RandomUserName(),
		Email:       util.RandomEmail(),
		FirstName:   util.RandomName(),
		Role:        util.RandomString(5),
		HouseholdID: util.RandomUUID(),
		CreatedAt:   pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now()},
	}
}
