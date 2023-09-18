package powerplatform

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	powerplatform_mock "github.com/microsoft/terraform-provider-power-platform/internal/mocks"
	models "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/bapi/models"
)

func TestUnitTenantSettingsDataSource_Validate_Read(t *testing.T) {
	clientMock := powerplatform_mock.NewUnitTestsMockApiClientInterface(t)

	tenantSettings := models.TenantSettingsDto{
		WalkMeOptOut: true,
	}

	clientMock.EXPECT().GetTenantSettings(gomock.Any()).Return(&tenantSettings, nil)

	dataSource := NewTenantSettingsDataSource()
	dataSource.(*TenantSettingsDataSource).BapiApiClient = clientMock

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"powerplatform": powerPlatformProviderServerApiMock(clientMock),
		},
		Steps: []resource.TestStep{
			{
				Config: uniTestsProviderConfig + `
				data "powerplatform_tenant_settings" "all" {
					walk_me_opt_out = "true"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "walk_me_opt_out", "true"),
				),
			},
		},
	})
}
