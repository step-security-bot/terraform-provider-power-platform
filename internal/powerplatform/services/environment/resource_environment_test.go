package environment

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"
	"github.com/microsoft/terraform-provider-power-platform/internal/powerplatform"
)

func TestUnitHttpClient(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	environment_resp := httpmock.NewStringResponder(http.StatusAccepted, httpmock.File("testdata/create_environment_response.json")).HeaderSet("Content-Type", "application/json")
	httpmock.RegisterResponder("POST", "https://api.bap.microsoft.com/providers/Microsoft.BusinessAppPlatform/environments", environment_resp)


	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"powerplatform": nil // func() { return providerserver.NewProtocol6WithError(powerplatform.NewPowerPlatformProvider()).(tfprotov6.ProviderServer), nil },
		},
		Steps: []resource.TestStep{},
	})

}

// func TestUnitEnvironmentsResource_Validate_Create_And_Force_Recreate(t *testing.T) {
// 	steps := []resource.TestStep{
// 		{
// 			Config: uniTestsProviderConfig + `
// 			resource "powerplatform_environment" "development" {
// 				display_name                              = "Example1"
// 				location                                  = "europe"
// 				language_code                             = "1033"
// 				currency_code                             = "USD"
// 				environment_type                          = "Sandbox"
// 				domain									  = "domain"
// 				security_group_id 						  = "security1"

// 			}`,
// 			Check: resource.ComposeTestCheckFunc(
// 				resource.TestCheckResourceAttr("powerplatform_environment.development", "environment_name", envIdBeforeChanges),
// 				resource.TestCheckResourceAttr("powerplatform_environment.development", "location", "europe"),
// 			),
// 		},

// 	ctrl := gomock.NewController(t)

// 	// Create a mock client
// 	mockClient := NewMockEnvironmentClientInterface(ctrl)

// 	resource.Test(t, resource.TestCase{
// 		IsUnitTest: true,
// 		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
// 			"powerplatform": powerPlatformProviderServerApiMock(clientMock),
// 		},
// 		Steps: nil,
// 	})
// }
