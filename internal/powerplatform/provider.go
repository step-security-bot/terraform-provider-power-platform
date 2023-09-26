package powerplatform

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	api "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api"
	bapi "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api/bapi"
	dvapi "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api/dataverse"
	ppapi "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api/ppapi"
	common "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/common"
)

var _ provider.Provider = &PowerPlatformProvider{}

type PowerPlatformProvider struct {
	Config           *common.ProviderConfig
	BapiApi          *BapiClient
	DataverseApi     *DataverseClient
	PowerPlatformApi *PowerPlatoformApiClient
}

type BapiClient struct {
	Auth   bapi.BapiAuthInterface
	Client bapi.BapiClientInterface
}

type DataverseClient struct {
	Auth   dvapi.DataverseAuthInterface
	Client dvapi.DataverseClientInterface
}

type PowerPlatoformApiClient struct {
	Auth   ppapi.PowerPlatformAuthInterface
	Client ppapi.PowerPlatformClientInterface
}

func NewPowerPlatformProvider() func() provider.Provider {
	return func() provider.Provider {

		cred := common.ProviderCredentials{}
		config := common.ProviderConfig{
			Credentials: &cred,
			Urls: common.ProviderConfigUrls{
				BapiUrl:          "api.bap.microsoft.com",
				PowerAppsUrl:     "api.powerapps.com",
				PowerPlatformUrl: "api.powerplatform.com",
			},
		}

		//bapi
		baseAuthBapi := &api.AuthImplementation{
			Config: config,
		}
		bapiAuth := &bapi.BapiAuthImplementation{
			BaseAuth: baseAuthBapi,
		}
		baseApiForBapi := &api.ApiClientImplementation{
			Config:   config,
			BaseAuth: baseAuthBapi,
		}
		bapiClient := &BapiClient{
			Auth: bapiAuth,
			Client: &bapi.BapiClientImplementation{
				BaseApi: baseApiForBapi,
				Auth:    bapiAuth,
			},
		}
		bapiClient.Client.GetBase().SetAuth(bapiAuth)
		//

		//powerplatform
		baseAuthPowerPlatform := &api.AuthImplementation{
			Config: config,
		}
		powerplatformAuth := &ppapi.PowerPlatformAuthImplementation{
			BaseAuth: baseAuthPowerPlatform,
		}

		baseApiForPpApi := &api.ApiClientImplementation{
			Config:   config,
			BaseAuth: baseAuthPowerPlatform,
		}
		powerplatformClient := &PowerPlatoformApiClient{
			Auth: powerplatformAuth,
			Client: &ppapi.PowerPlatformClientImplementation{
				BaseApi: baseApiForPpApi,
				Auth:    powerplatformAuth,
			},
		}
		powerplatformClient.Client.GetBase().SetAuth(powerplatformAuth)
		//

		//dataverse
		baseAuthDataverse := &api.AuthImplementation{
			Config: config,
		}
		dataverseAuth := &dvapi.DataverseAuthImplementation{
			BaseAuth: baseAuthDataverse,
		}
		baseApiForDataverse := &api.ApiClientImplementation{
			Config:   config,
			BaseAuth: baseAuthDataverse,
		}
		dataverseClient := &DataverseClient{
			Auth: dataverseAuth,
			Client: &dvapi.DataverseClientImplementation{
				BaseApi:    baseApiForDataverse,
				Auth:       dataverseAuth,
				BapiClient: bapiClient.Client,
			},
		}
		//

		p := &PowerPlatformProvider{
			Config:           &config,
			BapiApi:          bapiClient,
			DataverseApi:     dataverseClient,
			PowerPlatformApi: powerplatformClient,
		}
		return p
	}
}

func (p *PowerPlatformProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "powerplatform"
}

func (p *PowerPlatformProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	tflog.Debug(ctx, "Schema request received")

	resp.Schema = schema.Schema{

		Description:         "The Power Platform Terraform Provider allows managing environments and other resources within Power Platform",
		MarkdownDescription: "The Power Platform Provider allows managing environments and other resources within [Power Platform](https://powerplatform.microsoft.com/)",
		Attributes: map[string]schema.Attribute{
			"tenant_id": schema.StringAttribute{
				Description:         "The id of the AAD tenant that Power Platform API uses to authenticate with",
				MarkdownDescription: "The id of the AAD tenant that Power Platform API uses to authenticate with",
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				Description:         "The client id of the Power Platform API app registration",
				MarkdownDescription: "The client id of the Power Platform API app registration",
				Optional:            true,
			},
			"secret": schema.StringAttribute{
				Description:         "The secret of the Power Platform API app registration",
				MarkdownDescription: "The secret of the Power Platform API app registration",
				Optional:            true,
				Sensitive:           true,
			},

			"username": schema.StringAttribute{
				Description:         "The username of the Power Platform API in user@domain format",
				MarkdownDescription: "The username of the Power Platform API in user@domain format",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				Description:         "The password of the Power Platform API use",
				MarkdownDescription: "The password of the Power Platform API use",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *PowerPlatformProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config common.ProviderCredentialsModel

	tflog.Debug(ctx, "Configure request received")

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tenantId := ""
	envTenantId := os.Getenv("POWER_PLATFORM_TENANT_ID")
	if config.TenantId.IsNull() {
		tenantId = envTenantId
	} else {
		tenantId = config.TenantId.ValueString()
	}

	username := ""
	envUsername := os.Getenv("POWER_PLATFORM_USERNAME")
	if config.Username.IsNull() {
		username = envUsername
	} else {
		username = config.Username.ValueString()
	}

	password := ""
	envPassword := os.Getenv("POWER_PLATFORM_PASSWORD")
	if config.Password.IsNull() {
		password = envPassword
	} else {
		password = config.Password.ValueString()
	}

	clientId := ""
	envClientId := os.Getenv("POWER_PLATFORM_CLIENT_ID")
	if config.ClientId.IsNull() {
		clientId = envClientId
	} else {
		clientId = config.ClientId.ValueString()
	}

	secret := ""
	envSecret := os.Getenv("POWER_PLATFORM_SECRET")
	if config.Secret.IsNull() {
		secret = envSecret
	} else {
		secret = config.Secret.ValueString()
	}

	ctx = tflog.SetField(ctx, "power_platform_tenant_id", tenantId)
	ctx = tflog.SetField(ctx, "power_platform_username", username)
	ctx = tflog.SetField(ctx, "power_platform_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "power_platform_password")
	ctx = tflog.SetField(ctx, "power_platform_client_id", clientId)
	ctx = tflog.SetField(ctx, "power_platform_secret", secret)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "power_platform_secret")

	if clientId != "" && secret != "" && tenantId != "" {
		p.Config.Credentials.TenantId = tenantId
		p.Config.Credentials.ClientId = clientId
		p.Config.Credentials.Secret = secret
	} else if username != "" && password != "" && tenantId != "" {
		p.Config.Credentials.TenantId = tenantId
		p.Config.Credentials.Username = username
		p.Config.Credentials.Password = password
	} else {
		if tenantId == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("tenant_id"),
				"Unknown API tenant id",
				"The provider cannot create the API client as there is an unknown configuration value for the tenant id. "+
					"Either target apply the source of the value first, set the value statically in the configuration, or use the POWER_PLATFORM_TENANT_ID environment variable.",
			)
		}
		if username == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("username"),
				"Unknown username",
				"The provider cannot create the API client as there is an unknown configuration value for the username. "+
					"Either target apply the source of the value first, set the value statically in the configuration, or use the POWER_PLATFORM_USERNAME environment variable.",
			)
		}
		if password == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("password"),
				"Unknown password",
				"The provider cannot create the API client as there is an unknown configuration value for the password. "+
					"Either target apply the source of the value first, set the value statically in the configuration, or use the POWER_PLATFORM_PASSWORD environment variable.",
			)
		}
		if clientId == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("client_id"),
				"Unknown client id",
				"The provider cannot create the API client as there is an unknown configuration value for the client id. "+
					"Either target apply the source of the value first, set the value statically in the configuration, or use the POWER_PLATFORM_CLIENT_ID environment variable.",
			)
		}
		if secret == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("secret"),
				"Unknown secret",
				"The provider cannot create the API client as there is an unknown configuration value for the secret. "+
					"Either target apply the source of the value first, set the value statically in the configuration, or use the POWER_PLATFORM_SECRET environment variable.",
			)
		}
	}

	resp.DataSourceData = p
	resp.ResourceData = p

	tflog.Info(ctx, "Configured API client", map[string]any{"success": true})
}

func (p *PowerPlatformProvider) Resources(ctx context.Context) []func() resource.Resource {
	return append(resources,
		//		func() resource.Resource { return NewEnvironmentResource() },
		NewDataLossPreventionPolicyResource,
		NewSolutionResource,
	)
}

func (p *PowerPlatformProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		func() datasource.DataSource { return NewConnectorsDataSource() },
		func() datasource.DataSource { return NewPowerAppsDataSource() },
		func() datasource.DataSource { return NewEnvironmentsDataSource() },
		func() datasource.DataSource { return NewSolutionsDataSource() },
	}
}

var dataSources []func() datasource.DataSource
var resources []func() resource.Resource

func RegisterDataSource(dataSource func() datasource.DataSource) error {
	dataSources = append(dataSources, dataSource)
	return nil
}

func RegisterResource(resource func() resource.Resource) error {
	resources = append(resources, resource)
	return nil
}
