package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockdb "github.com/lmnzx/gobank/db/mock"
	db "github.com/lmnzx/gobank/db/sqlc"
	"github.com/lmnzx/gobank/util"
	"github.com/stretchr/testify/require"
)

func randomAccount(id uuid.UUID) db.Account {
	return db.Account{
		ID:        id,
		Owner:     util.RandOwner(),
		Balance:   util.RandMoney(),
		Currency:  util.RandCurrency(),
		CreatedAt: time.Now(),
	}
}

// write test for getAccount api endpoint
func TestGetAccountAPI(t *testing.T) {
	account := randomAccount(uuid.New())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// build stub
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start test server and send request
	server := NewServer(store)

	// create a request
	url := fmt.Sprintf("/account/%s", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	// create a response recorder
	recorder := httptest.NewRecorder()

	// send request to server
	server.router.ServeHTTP(recorder, request)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}
