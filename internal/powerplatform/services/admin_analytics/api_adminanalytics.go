package powerplatform

import (
	"context"
	"net/http"
	"net/url"

	api "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api"
)

func NewAdminAnalyticsClient(aaApi *api.AdminAnalyticsClientApi) AdminAnalyticsClient {
	return AdminAnalyticsClient{
		AdminAnalytics: aaApi,
	}
}

type AdminAnalyticsClient struct {
	AdminAnalytics *api.AdminAnalyticsClientApi
}

func (client *AdminAnalyticsClient) GetAppInsightConnections(ctx context.Context) (*[]AppInsightConnectionDto, error) {
	apiUrl := &url.URL{
		Scheme: "https",
		//todo don't use hardcoded url going to single region
		Host: client.AdminAnalytics.GetConfig().Urls.AdminAnalyticsUrl,
		Path: "/api/v1/sinks/appinsights/connections",
	}

	appInsightsArray := AppInsightConnectionDtoArray{}
	_, err := client.AdminAnalytics.Execute(ctx, "GET", apiUrl.String(), nil, nil, []int{http.StatusOK}, &appInsightsArray)
	if err != nil {
		return nil, err
	}

	return &appInsightsArray.Value, nil

}
