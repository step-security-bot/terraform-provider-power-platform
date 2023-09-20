package powerplatform

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	powerplatform_bapi "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/bapi"
	models "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/bapi/models"
)

var _ resource.Resource = &TenantSettingsResource{}

//var _ resource.ResourceWithImportState = &TenantSettingsResource{}

func NewTenantSettingsResource() resource.Resource {
	return &TenantSettingsResource{
		ProviderTypeName: "powerplatform",
		TypeName:         "_tenant_settings",
	}
}

type TenantSettingsResource struct {
	BapiApiClient    powerplatform_bapi.ApiClientInterface
	ProviderTypeName string
	TypeName         string
}

type TenantSettingsResourceModel struct {
	//	Id                                             types.String `tfsdk:"id"`
	WalkMeOptOut                                   types.Bool `tfsdk:"walk_me_opt_out"`
	DisableNPSCommentsReachout                     types.Bool `tfsdk:"disable_nps_comments_reachout"`
	DisableNewsletterSendout                       types.Bool `tfsdk:"disable_newsletter_sendout"`
	DisableEnvironmentCreationByNonAdminUsers      types.Bool `tfsdk:"disable_environment_creation_by_non_admin_users"`
	DisablePortalsCreationByNonAdminUsers          types.Bool `tfsdk:"disable_portals_creation_by_non_admin_users"`
	DisableSurveyFeedback                          types.Bool `tfsdk:"disable_survey_feedback"`
	DisableTrialEnvironmentCreationByNonAdminUsers types.Bool `tfsdk:"disable_trial_environment_creation_by_non_admin_users"`
	DisableCapacityAllocationByEnvironmentAdmins   types.Bool `tfsdk:"disable_capacity_allocation_by_environment_admins"`
	DisableSupportTicketsVisibleByAllUsers         types.Bool `tfsdk:"disable_support_tickets_visible_by_all_users"`
	// PowerPlatform                                  PowerPlatformSettings `tfsdk:"power_platform"`
}

// type PowerPlatformSettings struct {
// 	Search               SearchSettings               `tfsdk:"search"`
// 	TeamsIntegration     TeamsIntegrationSettings     `tfsdk:"teams_integration"`
// 	PowerApps            PowerAppsSettings            `tfsdk:"power_apps"`
// 	PowerAutomate        PowerAutomateSettings        `tfsdk:"power_automate"`
// 	Environments         EnvironmentsSettings         `tfsdk:"environments"`
// 	Governance           GovernanceSettings           `tfsdk:"governance"`
// 	Licensing            LicensingSettings            `tfsdk:"licensing"`
// 	PowerPages           PowerPagesSettings           `tfsdk:"power_pages"`
// 	Champions            ChampionsSettings            `tfsdk:"champions"`
// 	Intelligence         IntelligenceSettings         `tfsdk:"intelligence"`
// 	ModelExperimentation ModelExperimentationSettings `tfsdk:"model_experimentation"`
// 	CatalogSettings      CatalogSettingsSettings      `tfsdk:"catalog_settings"`
// }

// type SearchSettings struct {
// 	DisableDocsSearch      types.Bool `tfsdk:"disable_docs_search"`
// 	DisableCommunitySearch types.Bool `tfsdk:"disable_community_search"`
// 	DisableBingVideoSearch types.Bool `tfsdk:"disable_bing_video_search"`
// }

// type TeamsIntegrationSettings struct {
// 	ShareWithColleaguesUserLimit types.Int64 `tfsdk:"share_with_colleagues_user_limit"`
// }

// type PowerAppsSettings struct {
// 	DisableShareWithEveryone             types.Bool `tfsdk:"disable_share_with_everyone"`
// 	EnableGuestsToMake                   types.Bool `tfsdk:"enable_guests_to_make"`
// 	DisableMembersIndicator              types.Bool `tfsdk:"disable_members_indicator"`
// 	DisableMakerMatch                    types.Bool `tfsdk:"disable_maker_match"`
// 	DisableUnusedLicenseAssignment       types.Bool `tfsdk:"disable_unused_license_assignment"`
// 	DisableCreateFromImage               types.Bool `tfsdk:"disable_create_from_image"`
// 	DisableCreateFromFigma               types.Bool `tfsdk:"disable_create_from_figma"`
// 	DisableConnectionSharingWithEveryone types.Bool `tfsdk:"disable_connection_sharing_with_everyone"`
// }

// type PowerAutomateSettings struct {
// 	DisableCopilot types.Bool `tfsdk:"disable_copilot"`
// }

// type EnvironmentsSettings struct {
// 	DisablePreferredDataLocationForTeamsEnvironment types.Bool `tfsdk:"disable_preferred_data_location_for_teams_environment"`
// }

// type GovernanceSettings struct {
// 	DisableAdminDigest                                 types.Bool     `tfsdk:"disable_admin_digest"`
// 	DisableDeveloperEnvironmentCreationByNonAdminUsers types.Bool     `tfsdk:"disable_developer_environment_creation_by_non_admin_users"`
// 	EnableDefaultEnvironmentRouting                    types.Bool     `tfsdk:"enable_default_environment_routing"`
// 	Policy                                             PolicySettings `tfsdk:"policy"`
// }

// type PolicySettings struct {
// 	EnableDesktopFlowDataPolicyManagement types.Bool `tfsdk:"enable_desktop_flow_data_policy_management"`
// }

// type LicensingSettings struct {
// 	DisableBillingPolicyCreationByNonAdminUsers     types.Bool  `tfsdk:"disable_billing_policy_creation_by_non_admin_users"`
// 	EnableTenantCapacityReportForEnvironmentAdmins  types.Bool  `tfsdk:"enable_tenant_capacity_report_for_environment_admins"`
// 	StorageCapacityConsumptionWarningThreshold      types.Int64 `tfsdk:"storage_capacity_consumption_warning_threshold"`
// 	EnableTenantLicensingReportForEnvironmentAdmins types.Bool  `tfsdk:"enable_tenant_licensing_report_for_environment_admins"`
// 	DisableUseOfUnassignedAIBuilderCredits          types.Bool  `tfsdk:"disable_use_of_unassigned_ai_builder_credits"`
// }

// type PowerPagesSettings struct {
// }

// type ChampionsSettings struct {
// 	DisableChampionsInvitationReachout   types.Bool `tfsdk:"disable_champions_invitation_reachout"`
// 	DisableSkillsMatchInvitationReachout types.Bool `tfsdk:"disable_skills_match_invitation_reachout"`
// }

// type IntelligenceSettings struct {
// 	DisableCopilot            types.Bool `tfsdk:"disable_copilot"`
// 	EnableOpenAiBotPublishing types.Bool `tfsdk:"enable_open_ai_bot_publishing"`
// }

// type ModelExperimentationSettings struct {
// 	EnableModelDataSharing types.Bool `tfsdk:"enable_model_data_sharing"`
// 	DisableDataLogging     types.Bool `tfsdk:"disable_data_logging"`
// }

// type CatalogSettingsSettings struct {
// 	PowerCatalogAudienceSetting types.String `tfsdk:"power_catalog_audience_setting"`
// }

// metadata
func (r *TenantSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + r.TypeName
}

// schema
func (r *TenantSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Power Platform Tenant Settings Resource",
		MarkdownDescription: "Power Platform Tenant Settings Resource",
		Attributes: map[string]schema.Attribute{
			// "id": schema.StringAttribute{
			// 	Description: "Id",
			// 	Computed:    true,
			// },
			"walk_me_opt_out": schema.BoolAttribute{
				Description: "Walk Me Opt Out",
				Optional:    true, Computed: true,
			},
			"disable_nps_comments_reachout": schema.BoolAttribute{
				Description: "Disable NPS Comments Reachout",
				Optional:    true, Computed: true,
			},
			"disable_newsletter_sendout": schema.BoolAttribute{
				Description: "Disable Newsletter Sendout",
				Optional:    true, Computed: true,
			},
			"disable_environment_creation_by_non_admin_users": schema.BoolAttribute{
				Description: "Disable Environment Creation By Non Admin Users",
				Optional:    true, Computed: true,
			},
			"disable_portals_creation_by_non_admin_users": schema.BoolAttribute{
				Description: "Disable Portals Creation By Non Admin Users",
				Optional:    true, Computed: true,
			},
			"disable_survey_feedback": schema.BoolAttribute{
				Description: "Disable Survey Feedback",
				Optional:    true, Computed: true,
			},
			"disable_trial_environment_creation_by_non_admin_users": schema.BoolAttribute{
				Description: "Disable Trial Environment Creation By Non Admin Users",
				Optional:    true, Computed: true,
			},
			"disable_capacity_allocation_by_environment_admins": schema.BoolAttribute{
				Description: "Disable Capacity Allocation By Environment Admins",
				Optional:    true, Computed: true,
			},
			"disable_support_tickets_visible_by_all_users": schema.BoolAttribute{
				Description: "Disable Support Tickets Visible By All Users",
				Optional:    true, Computed: true,
			},
			// "power_platform": schema.SingleNestedAttribute{
			// 	Description: "Power Platform",
			// 	Optional:    true, Computed: true,
			// 	Attributes: map[string]schema.Attribute{
			// 		"search": schema.SingleNestedAttribute{
			// 			Description: "Search",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_docs_search": schema.BoolAttribute{
			// 					Description: "Disable Docs Search",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_community_search": schema.BoolAttribute{
			// 					Description: "Disable Community Search",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_bing_video_search": schema.BoolAttribute{
			// 					Description: "Disable Bing Video Search",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"teams_integration": schema.SingleNestedAttribute{
			// 			Description: "Teams Integration",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"share_with_colleagues_user_limit": schema.Int64Attribute{
			// 					Description: "Share With Colleagues User Limit",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"power_apps": schema.SingleNestedAttribute{
			// 			Description: "Power Apps",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_share_with_everyone": schema.BoolAttribute{
			// 					Description: "Disable Share With Everyone",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"enable_guests_to_make": schema.BoolAttribute{
			// 					Description: "Enable Guests To Make",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_members_indicator": schema.BoolAttribute{
			// 					Description: "Disable Members Indicator",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_maker_match": schema.BoolAttribute{
			// 					Description: "Disable Maker Match",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_unused_license_assignment": schema.BoolAttribute{
			// 					Description: "Disable Unused License Assignment",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_create_from_image": schema.BoolAttribute{
			// 					Description: "Disable Create From Image",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_create_from_figma": schema.BoolAttribute{
			// 					Description: "Disable Create From Figma",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_connection_sharing_with_everyone": schema.BoolAttribute{
			// 					Description: "Disable Connection Sharing With Everyone",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"power_automate": schema.SingleNestedAttribute{
			// 			Description: "Power Automate",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_copilot": schema.BoolAttribute{
			// 					Description: "Disable Copilot",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"environments": schema.SingleNestedAttribute{
			// 			Description: "Environments",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_preferred_data_location_for_teams_environment": schema.BoolAttribute{
			// 					Description: "Disable Preferred Data Location For Teams Environment",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"governance": schema.SingleNestedAttribute{
			// 			Description: "Governance",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_admin_digest": schema.BoolAttribute{
			// 					Description: "Disable Admin Digest",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_developer_environment_creation_by_non_admin_users": schema.BoolAttribute{
			// 					Description: "Disable Developer Environment Creation By Non Admin Users",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"enable_default_environment_routing": schema.BoolAttribute{
			// 					Description: "Enable Default Environment Routing",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"policy": schema.SingleNestedAttribute{
			// 					Description: "Policy",
			// 					Optional:    true, Computed: true,
			// 					Attributes: map[string]schema.Attribute{
			// 						"enable_desktop_flow_data_policy_management": schema.BoolAttribute{
			// 							Description: "Enable Desktop Flow Data Policy Management",
			// 							Optional:    true, Computed: true,
			// 						},
			// 					},
			// 				},
			// 			},
			// 		},
			// 		"licensing": schema.SingleNestedAttribute{
			// 			Description: "Licensing",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_billing_policy_creation_by_non_admin_users": schema.BoolAttribute{
			// 					Description: "Disable Billing Policy Creation By Non Admin Users",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"enable_tenant_capacity_report_for_environment_admins": schema.BoolAttribute{
			// 					Description: "Enable Tenant Capacity Report For Environment Admins",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"storage_capacity_consumption_warning_threshold": schema.Int64Attribute{
			// 					Description: "Storage Capacity Consumption Warning Threshold",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"enable_tenant_licensing_report_for_environment_admins": schema.BoolAttribute{
			// 					Description: "Enable Tenant Licensing Report For Environment Admins",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_use_of_unassigned_ai_builder_credits": schema.BoolAttribute{
			// 					Description: "Disable Use Of Unassigned AI Builder Credits",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"power_pages": schema.SingleNestedAttribute{
			// 			Description: "Power Pages",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{},
			// 		},
			// 		"champions": schema.SingleNestedAttribute{
			// 			Description: "Champions",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_champions_invitation_reachout": schema.BoolAttribute{
			// 					Description: "Disable Champions Invitation Reachout",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_skills_match_invitation_reachout": schema.BoolAttribute{
			// 					Description: "Disable Skills Match Invitation Reachout",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"intelligence": schema.SingleNestedAttribute{
			// 			Description: "Intelligence",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"disable_copilot": schema.BoolAttribute{
			// 					Description: "Disable Copilot",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"enable_open_ai_bot_publishing": schema.BoolAttribute{
			// 					Description: "Enable Open AI Bot Publishing",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"model_experimentation": schema.SingleNestedAttribute{
			// 			Description: "Model Experimentation",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"enable_model_data_sharing": schema.BoolAttribute{
			// 					Description: "Enable Model Data Sharing",
			// 					Optional:    true, Computed: true,
			// 				},
			// 				"disable_data_logging": schema.BoolAttribute{
			// 					Description: "Disable Data Logging",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 		"catalog_settings": schema.SingleNestedAttribute{
			// 			Description: "Catalog Settings",
			// 			Optional:    true, Computed: true,
			// 			Attributes: map[string]schema.Attribute{
			// 				"power_catalog_audience_setting": schema.StringAttribute{
			// 					Description: "Power Catalog Audience Setting",
			// 					Optional:    true, Computed: true,
			// 				},
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func ConvertToTenantSettingsDto(plan *TenantSettingsResourceModel) models.TenantSettingsDto {
	dto := models.TenantSettingsDto{}
	if !plan.WalkMeOptOut.IsUnknown() {
		dto.WalkMeOptOut = plan.WalkMeOptOut.ValueBool()
	}

	return dto
}

// todo: fix func name
func ConvertResourceFromTenantSettingsDto(tenantSettingsDto models.TenantSettingsDto) TenantSettingsResourceModel {
	ts := TenantSettingsResourceModel{
		//		Id:                         types.StringValue(""),
		WalkMeOptOut:                                   types.BoolValue(tenantSettingsDto.WalkMeOptOut),
		DisableNPSCommentsReachout:                     types.BoolValue(tenantSettingsDto.DisableNPSCommentsReachout),
		DisableNewsletterSendout:                       types.BoolValue(tenantSettingsDto.DisableNewsletterSendout),
		DisableEnvironmentCreationByNonAdminUsers:      types.BoolValue(tenantSettingsDto.DisableEnvironmentCreationByNonAdminUsers),
		DisablePortalsCreationByNonAdminUsers:          types.BoolValue(tenantSettingsDto.DisablePortalsCreationByNonAdminUsers),
		DisableSurveyFeedback:                          types.BoolValue(tenantSettingsDto.DisableSurveyFeedback),
		DisableTrialEnvironmentCreationByNonAdminUsers: types.BoolValue(tenantSettingsDto.DisableTrialEnvironmentCreationByNonAdminUsers),
		DisableCapacityAllocationByEnvironmentAdmins:   types.BoolValue(tenantSettingsDto.DisableCapacityAllocationByEnvironmentAdmins),
		//DisableSupportTicketsVisibleByAllUsers:         types.BoolValue(tenantSettingsDto.DisableSupportTicketsVisibleByAllUsers),
		// PowerPlatform: PowerPlatformSettings{
		// 	Search: SearchSettings{
		// 		DisableDocsSearch:      types.BoolValue(tenantSettingsDto.PowerPlatform.Search.DisableDocsSearch),
		// 		DisableCommunitySearch: types.BoolValue(tenantSettingsDto.PowerPlatform.Search.DisableCommunitySearch),
		// 		DisableBingVideoSearch: types.BoolValue(tenantSettingsDto.PowerPlatform.Search.DisableBingVideoSearch),
		// 	},
		// 	TeamsIntegration: TeamsIntegrationSettings{
		// 		ShareWithColleaguesUserLimit: types.Int64Value(tenantSettingsDto.PowerPlatform.TeamsIntegration.ShareWithColleaguesUserLimit),
		// 	},
		// 	PowerApps: PowerAppsSettings{
		// 		DisableShareWithEveryone:             types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableShareWithEveryone),
		// 		EnableGuestsToMake:                   types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.EnableGuestsToMake),
		// 		DisableMembersIndicator:              types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableMembersIndicator),
		// 		DisableMakerMatch:                    types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableMakerMatch),
		// 		DisableUnusedLicenseAssignment:       types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableUnusedLicenseAssignment),
		// 		DisableCreateFromImage:               types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableCreateFromImage),
		// 		DisableCreateFromFigma:               types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableCreateFromFigma),
		// 		DisableConnectionSharingWithEveryone: types.BoolValue(tenantSettingsDto.PowerPlatform.PowerApps.DisableConnectionSharingWithEveryone),
		// 	},
		// 	PowerAutomate: PowerAutomateSettings{
		// 		DisableCopilot: types.BoolValue(tenantSettingsDto.PowerPlatform.PowerAutomate.DisableCopilot),
		// 	},
		// 	Environments: EnvironmentsSettings{
		// 		DisablePreferredDataLocationForTeamsEnvironment: types.BoolValue(tenantSettingsDto.PowerPlatform.Environments.DisablePreferredDataLocationForTeamsEnvironment),
		// 	},
		// 	Governance: GovernanceSettings{
		// 		DisableAdminDigest: types.BoolValue(tenantSettingsDto.PowerPlatform.Governance.DisableAdminDigest),
		// 		DisableDeveloperEnvironmentCreationByNonAdminUsers: types.BoolValue(tenantSettingsDto.PowerPlatform.Governance.DisableDeveloperEnvironmentCreationByNonAdminUsers),
		// 		EnableDefaultEnvironmentRouting:                    types.BoolValue(tenantSettingsDto.PowerPlatform.Governance.EnableDefaultEnvironmentRouting),
		// 		Policy: PolicySettings{
		// 			EnableDesktopFlowDataPolicyManagement: types.BoolValue(tenantSettingsDto.PowerPlatform.Governance.Policy.EnableDesktopFlowDataPolicyManagement),
		// 		},
		// 	},
		// 	Licensing: LicensingSettings{
		// 		DisableBillingPolicyCreationByNonAdminUsers:     types.BoolValue(tenantSettingsDto.PowerPlatform.Licensing.DisableBillingPolicyCreationByNonAdminUsers),
		// 		EnableTenantCapacityReportForEnvironmentAdmins:  types.BoolValue(tenantSettingsDto.PowerPlatform.Licensing.EnableTenantCapacityReportForEnvironmentAdmins),
		// 		StorageCapacityConsumptionWarningThreshold:      types.Int64Value(tenantSettingsDto.PowerPlatform.Licensing.StorageCapacityConsumptionWarningThreshold),
		// 		EnableTenantLicensingReportForEnvironmentAdmins: types.BoolValue(tenantSettingsDto.PowerPlatform.Licensing.EnableTenantLicensingReportForEnvironmentAdmins),
		// 		DisableUseOfUnassignedAIBuilderCredits:          types.BoolValue(tenantSettingsDto.PowerPlatform.Licensing.DisableUseOfUnassignedAIBuilderCredits),
		// 	},
		// 	PowerPages: PowerPagesSettings{},
		// 	Champions: ChampionsSettings{
		// 		DisableChampionsInvitationReachout:   types.BoolValue(tenantSettingsDto.PowerPlatform.Champions.DisableChampionsInvitationReachout),
		// 		DisableSkillsMatchInvitationReachout: types.BoolValue(tenantSettingsDto.PowerPlatform.Champions.DisableSkillsMatchInvitationReachout),
		// 	},
		// 	Intelligence: IntelligenceSettings{
		// 		DisableCopilot:            types.BoolValue(tenantSettingsDto.PowerPlatform.Intelligence.DisableCopilot),
		// 		EnableOpenAiBotPublishing: types.BoolValue(tenantSettingsDto.PowerPlatform.Intelligence.EnableOpenAiBotPublishing),
		// 	},
		// 	ModelExperimentation: ModelExperimentationSettings{
		// 		EnableModelDataSharing: types.BoolValue(tenantSettingsDto.PowerPlatform.ModelExperimentation.EnableModelDataSharing),
		// 		DisableDataLogging:     types.BoolValue(tenantSettingsDto.PowerPlatform.ModelExperimentation.DisableDataLogging),
		// 	},
		// 	CatalogSettings: CatalogSettingsSettings{
		// 		PowerCatalogAudienceSetting: types.StringValue(tenantSettingsDto.PowerPlatform.CatalogSettings.PowerCatalogAudienceSetting),
		// 	},
		// },
	}

	if !ts.DisableSupportTicketsVisibleByAllUsers.IsUnknown() {
		ts.DisableSupportTicketsVisibleByAllUsers = types.BoolValue(tenantSettingsDto.DisableSupportTicketsVisibleByAllUsers)
	}

	return ts
}

// configure
func (r *TenantSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*PowerPlatformProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.BapiApiClient = provider.bapiClient
}

// create
func (r *TenantSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//tflog.Debug(ctx, fmt.Sprintf("CREATE RESOURCE START: %s", r.ProviderTypeName))
	var plan TenantSettingsResourceModel

	// read the terraform plan into the plan variable
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert to dto
	dto := ConvertToTenantSettingsDto(&plan)

	ts, err := r.BapiApiClient.UpdateTenantSettings(ctx, &dto)
	if err != nil {
		//resp.Diagnostics.AddError(fmt.Sprintf("Client error when reading %s", d.ProviderTypeName), err.Error())
		return
	}

	state := ConvertResourceFromTenantSettingsDto(*ts)
	// set to empty guid
	//state.Id = types.StringValue("00000000-0000-0000-0000-000000000000")

	resp.State.Set(ctx, &state)
}

// read
func (r *TenantSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//tflog.Debug(ctx, fmt.Sprintf("READ RESOURCE START: %s", r.ProviderTypeName))
	var state TenantSettingsResourceModel
	ts, err := r.BapiApiClient.GetTenantSettings(ctx)
	if err != nil {
		//resp.Diagnostics.AddError(fmt.Sprintf("Client error when reading %s", d.ProviderTypeName), err.Error())
		return
	}

	state = ConvertResourceFromTenantSettingsDto(*ts)
	//state.Id = types.StringValue("00000000-0000-0000-0000-000000000000")
	state.WalkMeOptOut = types.BoolValue(ts.WalkMeOptOut)

	diags := resp.State.Set(ctx, &state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// update
func (r *TenantSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("UPDATE RESOURCE START: %s", r.ProviderTypeName))
	var plan TenantSettingsResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state *TenantSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert to dto
	ts := ConvertToTenantSettingsDto(&plan)

	// update
	uts, err := r.BapiApiClient.UpdateTenantSettings(ctx, &ts)
	if err != nil {
		//resp.Diagnostics.AddError(fmt.Sprintf("Client error when reading %s", d.ProviderTypeName), err.Error())
		return
	}

	updatedState := ConvertResourceFromTenantSettingsDto(*uts)
	//updatedState.Id = types.StringValue("00000000-0000-0000-0000-000000000000")

	diags := resp.State.Set(ctx, &updatedState)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// delete
func (r *TenantSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// no operation - tenant settings cannot be deleted
}
