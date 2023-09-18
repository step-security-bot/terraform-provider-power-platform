package powerplatform_bapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	models "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/bapi/models"
)

func (client *ApiClient) GetTenantSettings(ctx context.Context) (*models.TenantSettingsDto, error) {
	apiUrl := &url.URL{
		Scheme: "https",
		Host:   "api.bap.microsoft.com",
		Path:   "/providers/Microsoft.BusinessAppPlatform/listTenantSettings",
	}

	values := url.Values{}
	values.Add("api-version", "2023-06-01")
	apiUrl.RawQuery = values.Encode()

	request, err := http.NewRequestWithContext(ctx, "POST", apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+client.Token)
	body, err := client.doRequest(request)
	if err != nil {
		return nil, err
	}

	tenantSettings := models.TenantSettingsDto{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&tenantSettings)
	if err != nil {
		return nil, err
	}

	return &tenantSettings, nil
}
