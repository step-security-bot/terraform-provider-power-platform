package powerplatform

import (
	"context"
	"net/http"

	common "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/common"
)

type AdminAnalyticsClientApi struct {
	baseApi *ApiClientBase
	auth    *AdminAnalyticsAuth
}

func NewAdminAnalyticsClientApi(baseApi *ApiClientBase, auth *AdminAnalyticsAuth) *AdminAnalyticsClientApi {
	return &AdminAnalyticsClientApi{
		baseApi: baseApi,
		auth:    auth,
	}
}

func (client *AdminAnalyticsClientApi) GetConfig() *common.ProviderConfig {
	return client.baseApi.Config
}

func (client *AdminAnalyticsClientApi) Execute(ctx context.Context, method string, url string, headers http.Header, body interface{}, acceptableStatusCodes []int, responseObj interface{}) (*ApiHttpResponse, error) {
	token, err := client.baseApi.InitializeBase(ctx, client.auth)
	if err != nil {
		return nil, err
	}
	return client.baseApi.ExecuteBase(ctx, token, method, url, headers, body, acceptableStatusCodes, responseObj)
}
