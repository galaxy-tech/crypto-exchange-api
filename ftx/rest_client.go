package ftx

import (
	"github.com/galaxy-tech/go-ftx/auth"
	"github.com/galaxy-tech/go-ftx/rest"
)

type RestClient struct {
	client *rest.Client
}

func NewRestClient(key, secret, subaccount string) *RestClient {

	clientWithSubAccounts := rest.New(
		auth.New(
			key,
			secret,
			auth.SubAccount{
				UUID:     1,
				Nickname: subaccount,
			},
		))

	// switch subaccount
	clientWithSubAccounts.Auth.UseSubAccountID(1)

	return &RestClient{
		client: clientWithSubAccounts,
	}
}