package powerplatform

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	powerplatform_mock "github.com/microsoft/terraform-provider-power-platform/internal/mocks"
	models "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/bapi/models"
)

func TestUnitTenantSettingsDataSource_Validate_Read(t *testing.T) {
	clientMock := powerplatform_mock.NewUnitTestsMockApiClientInterface(t)

	tenantJson := `{
		"walkMeOptOut": true,
		"disableNPSCommentsReachout": true,
		"disableNewsletterSendout": true,
		"disableEnvironmentCreationByNonAdminUsers": true,
		"disablePortalsCreationByNonAdminUsers": true,
		"disableSurveyFeedback": true,
		"disableTrialEnvironmentCreationByNonAdminUsers": true,
		"disableCapacityAllocationByEnvironmentAdmins": true,
		"disableSupportTicketsVisibleByAllUsers": true,
		"powerPlatform": {
			"search": {
				"disableDocsSearch": true,
				"disableCommunitySearch": true,
				"disableBingVideoSearch": true
			},
			"teamsIntegration": {
				"shareWithColleaguesUserLimit": 10000
			},
			"powerApps": {
				"disableShareWithEveryone": true,
				"enableGuestsToMake": true,
				"disableMembersIndicator": true,
				"disableMakerMatch": true,
				"disableUnusedLicenseAssignment": true,
				"disableCreateFromImage": true,
				"disableCreateFromFigma": true,
				"disableConnectionSharingWithEveryone": true
			},
			"powerAutomate": {
				"disableCopilot": true
			},
			"environments": {
				"disablePreferredDataLocationForTeamsEnvironment": true
			},
			"governance": {
				"disableAdminDigest": true,
				"disableDeveloperEnvironmentCreationByNonAdminUsers": true,
				"enableDefaultEnvironmentRouting": true,
				"policy": {
					"enableDesktopFlowDataPolicyManagement": true
				}
			},
			"licensing": {
				"disableBillingPolicyCreationByNonAdminUsers": true,
				"enableTenantCapacityReportForEnvironmentAdmins": true,
				"storageCapacityConsumptionWarningThreshold": 85,
				"enableTenantLicensingReportForEnvironmentAdmins": true,
				"disableUseOfUnassignedAIBuilderCredits": true
			},
			"powerPages": {},
			"champions": {
				"disableChampionsInvitationReachout": true,
				"disableSkillsMatchInvitationReachout": true
			},
			"intelligence": {
				"disableCopilot": true,
				"enableOpenAiBotPublishing": true
			},
			"modelExperimentation": {
				"enableModelDataSharing": true,
				"disableDataLogging": true
			},
			"catalogSettings": {
				"powerCatalogAudienceSetting": "All"
			}
		}
	}
	`
	var tenantSettings models.TenantSettingsDto
	json.Unmarshal([]byte(tenantJson), &tenantSettings)

	// 	WalkMeOptOut:                                   true,
	// 	DisableNPSCommentsReachout:                     true,
	// 	DisableNewsletterSendout:                       true,
	// 	DisableEnvironmentCreationByNonAdminUsers:      true,
	// 	DisablePortalsCreationByNonAdminUsers:          true,
	// 	DisableSurveyFeedback:                          true,
	// 	DisableTrialEnvironmentCreationByNonAdminUsers: true,
	// 	DisableCapacityAllocationByEnvironmentAdmins:   true,
	// 	DisableSupportTicketsVisibleByAllUsers:         true,
	// 	PowerPlatform: //what goes here?},
	// }
	// tenantSettings.PowerPlatform.PowerApps.DisableShareWithEveryone = true

	clientMock.EXPECT().GetTenantSettings(gomock.Any()).Return(&tenantSettings, nil).AnyTimes()

	dataSource := NewTenantSettingsDataSource()
	dataSource.(*TenantSettingsDataSource).BapiApiClient = clientMock

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"powerplatform": powerPlatformProviderServerApiMock(clientMock),
		},
		Steps: []resource.TestStep{
			{
				Config: uniTestsProviderConfig + `data "powerplatform_tenant_settings" "all" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "walk_me_opt_out", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_nps_comments_reachout", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_newsletter_sendout", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_environment_creation_by_non_admin_users", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_portals_creation_by_non_admin_users", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_survey_feedback", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_trial_environment_creation_by_non_admin_users", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_capacity_allocation_by_environment_admins", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "disable_support_tickets_visible_by_all_users", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.search.disable_docs_search", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.search.disable_community_search", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.search.disable_bing_video_search", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.teams_integration.share_with_colleagues_user_limit", "10000"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_share_with_everyone", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.enable_guests_to_make", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_members_indicator", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_maker_match", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_unused_license_assignment", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_create_from_image", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_create_from_figma", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_apps.disable_connection_sharing_with_everyone", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.power_automate.disable_copilot", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.environments.disable_preferred_data_location_for_teams_environment", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.governance.disable_admin_digest", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.governance.disable_developer_environment_creation_by_non_admin_users", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.governance.enable_default_environment_routing", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.governance.policy.enable_desktop_flow_data_policy_management", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.licensing.disable_billing_policy_creation_by_non_admin_users", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.licensing.enable_tenant_capacity_report_for_environment_admins", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.licensing.storage_capacity_consumption_warning_threshold", "85"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.licensing.enable_tenant_licensing_report_for_environment_admins", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.licensing.disable_use_of_unassigned_ai_builder_credits", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.champions.disable_champions_invitation_reachout", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.champions.disable_skills_match_invitation_reachout", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.intelligence.disable_copilot", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.intelligence.enable_open_ai_bot_publishing", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.model_experimentation.enable_model_data_sharing", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.model_experimentation.disable_data_logging", "true"),
					resource.TestCheckResourceAttr("data.powerplatform_tenant_settings.all", "power_platform.catalog_settings.power_catalog_audience_setting", "All"),
				),
			},
		},
	})
}
