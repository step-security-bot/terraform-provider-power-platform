// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package dlp_policy_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"
	"github.com/microsoft/terraform-provider-power-platform/internal/mocks"
)

func TestUnitDataLossPreventionPolicyResource_Validate_Update(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	policyId := "00000000-0000-0000-0000-000000000000"
	policyResponse1 := fmt.Sprintf(`{
		"policyDefinition": {
			"name": "%s",
			"displayName": "Block All Policy",
			"defaultConnectorsClassification": "Blocked",
			"connectorGroups": [
				{
					"classification": "Confidential",
					"connectors": []
				},
				{
					"classification": "General",
					"connectors": []
				},
				{
					"classification": "Blocked",
					"connectors": []
				}
			],
			"environmentType": "AllEnvironments",
			"environments": [],
			"createdBy": {
				"displayName": "admin"
			},
			"createdTime": "2023-10-09T09:38:40.3125874Z",
			"lastModifiedBy": {
				"displayName": "admin"
			},
			"lastModifiedTime": "2023-10-09T09:38:40.3125874Z",
			"etag": "14d0df1a-418e-47f5-972e-0430ec36ef47",
			"isLegacySchemaVersion": false
		},
		"customConnectorUrlPatternsDefinition": {
			"rules": [
				{
					"order": 1,
					"customConnectorRuleClassification": "Ignore",
					"pattern": "*"
				}
			]
		}
	}`, policyId)

	policyResponse2 := fmt.Sprintf(`{
		"policyDefinition": {
			"name": "%s",
			"displayName": "Block All Policy_1",
			"defaultConnectorsClassification": "Blocked",
			"connectorGroups": [
				{
					"classification": "Confidential",
					"connectors": []
				},
				{
					"classification": "General",
					"connectors": []
				},
				{
					"classification": "Blocked",
					"connectors": []
				}
			],
			"environmentType": "AllEnvironments",
			"environments": [],
			"createdBy": {
				"displayName": "admin"
			},
			"createdTime": "2023-10-09T09:38:40.3125874Z",
			"lastModifiedBy": {
				"displayName": "admin"
			},
			"lastModifiedTime": "2023-10-09T09:38:40.3125874Z",
			"etag": "14d0df1a-418e-47f5-972e-0430ec36ef47",
			"isLegacySchemaVersion": false
		},
		"customConnectorUrlPatternsDefinition": {
			"rules": [
				{
					"order": 1,
					"customConnectorRuleClassification": "Ignore",
					"pattern": "*"
				}
			]
		}
	}`, policyId)

	policyResponse3 := fmt.Sprintf(`{
		"policyDefinition": {
			"name": "%s",
			"displayName": "Block All Policy_1",
			"defaultConnectorsClassification": "General",
			"connectorGroups": [
				{
					"classification": "Confidential",
					"connectors": []
				},
				{
					"classification": "General",
					"connectors": []
				},
				{
					"classification": "Blocked",
					"connectors": []
				}
			],
			"environmentType": "OnlyEnvironments",
			"environments": [
				{
					"id": "/providers/Microsoft.BusinessAppPlatform/scopes/admin/environments/00000000-0000-0000-0000-000000000000",
					"name": "00000000-0000-0000-0000-000000000000",
					"type": "Microsoft.BusinessAppPlatform/scopes/environments"
				}
        	],
			"createdBy": {
				"displayName": "admin"
			},
			"createdTime": "2023-10-09T09:38:40.3125874Z",
			"lastModifiedBy": {
				"displayName": "admin"
			},
			"lastModifiedTime": "2023-10-09T09:38:40.3125874Z",
			"etag": "14d0df1a-418e-47f5-972e-0430ec36ef47",
			"isLegacySchemaVersion": false
		},
		"customConnectorUrlPatternsDefinition": {
			"rules": [
				{
					"order": 1,
					"customConnectorRuleClassification": "Ignore",
					"pattern": "*"
				}
			]
		}
	}`, policyId)

	policyResponse4 := fmt.Sprintf(`{
		"policyDefinition": {
			"name": "%s",
			"displayName": "Block All Policy_1",
			"defaultConnectorsClassification": "General",
			"connectorGroups": [
				{
					"classification": "General",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_sql",
							"name": "shared_sql",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				},
				{
					"classification": "Confidential",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
							"name": "shared_sharepointonline",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				},
				{
					"classification": "Blocked",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_azureblob",
							"name": "shared_azureblob",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				}
			],
			"environmentType": "OnlyEnvironments",
			"environments": [
				{
					"id": "/providers/Microsoft.BusinessAppPlatform/scopes/admin/environments/00000000-0000-0000-0000-000000000000",
					"name": "00000000-0000-0000-0000-000000000000",
					"type": "Microsoft.BusinessAppPlatform/scopes/environments"
				}
        	],
			"createdBy": {
				"displayName": "admin"
			},
			"createdTime": "2023-10-09T09:38:40.3125874Z",
			"lastModifiedBy": {
				"displayName": "admin"
			},
			"lastModifiedTime": "2023-10-09T09:38:40.3125874Z",
			"etag": "14d0df1a-418e-47f5-972e-0430ec36ef47",
			"isLegacySchemaVersion": false
		},
		"customConnectorUrlPatternsDefinition": {
			"rules": [
				{
					"order": 1,
					"customConnectorRuleClassification": "Ignore",
					"pattern": "*"
				}
			]
		}
	}`, policyId)

	policyResponse5 := fmt.Sprintf(`{
		"policyDefinition": {
			"name": "%s",
			"displayName": "Block All Policy_1",
			"defaultConnectorsClassification": "General",
			"connectorGroups": [
				{
					"classification": "General",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_sql",
							"name": "shared_sql",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				},
				{
					"classification": "Confidential",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
							"name": "shared_sharepointonline",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				},
				{
					"classification": "Blocked",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_azureblob",
							"name": "shared_azureblob",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				}
			],
			"environmentType": "OnlyEnvironments",
			"environments": [
				{
					"id": "/providers/Microsoft.BusinessAppPlatform/scopes/admin/environments/00000000-0000-0000-0000-000000000000",
					"name": "00000000-0000-0000-0000-000000000000",
					"type": "Microsoft.BusinessAppPlatform/scopes/environments"
				}
        	],
			"createdBy": {
				"displayName": "admin"
			},
			"createdTime": "2023-10-09T09:38:40.3125874Z",
			"lastModifiedBy": {
				"displayName": "admin"
			},
			"lastModifiedTime": "2023-10-09T09:38:40.3125874Z",
			"etag": "14d0df1a-418e-47f5-972e-0430ec36ef47",
			"isLegacySchemaVersion": false
		},
		"connectorConfigurationsDefinition": {
			"connectorActionConfigurations": [
				{
					"connectorId": "/providers/Microsoft.PowerApps/apis/shared_sql",
					"defaultConnectorActionRuleBehavior": "Allow",
					"actionRules": [
						{
							"actionId": "DeleteItem_V2",
							"behavior": "Block"
						},
						{
							"actionId": "ExecutePassThroughNativeQuery_V2",
							"behavior": "Block"
						}
					]
				}
			],
			"endpointConfigurations": [
				{
					"connectorId": "/providers/Microsoft.PowerApps/apis/shared_sql",
					"endpointRules": [
						{
							"order": 1,
							"behavior": "Allow",
							"endpoint": "contoso.com"
						},
						{
							"order": 2,
							"behavior": "Deny",
							"endpoint": "*"
						}
					]
				}
			]
		},
		"customConnectorUrlPatternsDefinition": {
			"rules": [
				{
					"order": 1,
					"customConnectorRuleClassification": "Ignore",
					"pattern": "*"
				}
			]
		}
	}`, policyId)

	policyResponse6 := fmt.Sprintf(`{
		"policyDefinition": {
			"name": "%s",
			"displayName": "Block All Policy_1",
			"defaultConnectorsClassification": "General",
			"connectorGroups": [
				{
					"classification": "General",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_sql",
							"name": "shared_sql",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				},
				{
					"classification": "Confidential",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
							"name": "shared_sharepointonline",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				},
				{
					"classification": "Blocked",
					"connectors": [
						{
							"id": "/providers/Microsoft.PowerApps/apis/shared_azureblob",
							"name": "shared_azureblob",
							"type": "Microsoft.PowerApps/apis"
						}
					]
				}
			],
			"environmentType": "OnlyEnvironments",
			"environments": [
				{
					"id": "/providers/Microsoft.BusinessAppPlatform/scopes/admin/environments/00000000-0000-0000-0000-000000000000",
					"name": "00000000-0000-0000-0000-000000000000",
					"type": "Microsoft.BusinessAppPlatform/scopes/environments"
				}
        	],
			"createdBy": {
				"displayName": "admin"
			},
			"createdTime": "2023-10-09T09:38:40.3125874Z",
			"lastModifiedBy": {
				"displayName": "admin"
			},
			"lastModifiedTime": "2023-10-09T09:38:40.3125874Z",
			"etag": "14d0df1a-418e-47f5-972e-0430ec36ef47",
			"isLegacySchemaVersion": false
		},
		"connectorConfigurationsDefinition": {
			"connectorActionConfigurations": [
				{
					"connectorId": "/providers/Microsoft.PowerApps/apis/shared_sql",
					"defaultConnectorActionRuleBehavior": "Allow",
					"actionRules": [
						{
							"actionId": "DeleteItem_V2",
							"behavior": "Block"
						},
						{
							"actionId": "ExecutePassThroughNativeQuery_V2",
							"behavior": "Block"
						}
					]
				}
			],
			"endpointConfigurations": [
				{
					"connectorId": "/providers/Microsoft.PowerApps/apis/shared_sql",
					"endpointRules": [
						{
							"order": 1,
							"behavior": "Allow",
							"endpoint": "contoso.com"
						},
						{
							"order": 2,
							"behavior": "Deny",
							"endpoint": "*"
						}
					]
				}
			]
		},
		"customConnectorUrlPatternsDefinition": {
			"rules": [
				{
					"order": 1,
					"customConnectorRuleClassification": "Blocked",
					"pattern": "https://*.contoso.com"
				},
				{
					"order": 2,
					"customConnectorRuleClassification": "Ignore",
					"pattern": "*"
				}
			]
		}
	}`, policyId)

	getResponsesInx := -1
	getResponsesArray := make([]string, 0)
	getResponsesArray = append(getResponsesArray, policyResponse1)
	getResponsesArray = append(getResponsesArray, policyResponse1)
	getResponsesArray = append(getResponsesArray, policyResponse2)
	getResponsesArray = append(getResponsesArray, policyResponse2)
	getResponsesArray = append(getResponsesArray, policyResponse3)
	getResponsesArray = append(getResponsesArray, policyResponse3)
	getResponsesArray = append(getResponsesArray, policyResponse4)
	getResponsesArray = append(getResponsesArray, policyResponse4)
	getResponsesArray = append(getResponsesArray, policyResponse5)
	getResponsesArray = append(getResponsesArray, policyResponse5)
	getResponsesArray = append(getResponsesArray, policyResponse6)
	getResponsesArray = append(getResponsesArray, policyResponse6)

	patchResponsesInx := -1
	patchResponsesArray := make([]string, 0)
	patchResponsesArray = append(patchResponsesArray, policyResponse2)
	patchResponsesArray = append(patchResponsesArray, policyResponse3)
	patchResponsesArray = append(patchResponsesArray, policyResponse4)
	patchResponsesArray = append(patchResponsesArray, policyResponse5)
	patchResponsesArray = append(patchResponsesArray, policyResponse6)

	httpmock.RegisterResponder("PATCH", fmt.Sprintf(`https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies/%s`, policyId),
		func(req *http.Request) (*http.Response, error) {
			patchResponsesInx++
			return httpmock.NewStringResponse(http.StatusOK, patchResponsesArray[patchResponsesInx]), nil
		})

	httpmock.RegisterResponder("POST", `https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusCreated, policyResponse1), nil
		})

	httpmock.RegisterResponder("GET", fmt.Sprintf(`https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies/%s`, policyId),
		func(req *http.Request) (*http.Response, error) {
			getResponsesInx++
			return httpmock.NewStringResponse(http.StatusOK, getResponsesArray[getResponsesInx]), nil
		})

	httpmock.RegisterResponder("DELETE", fmt.Sprintf(`https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies/%s`, policyId),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, ""), nil
		})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy"
					default_connectors_classification = "Blocked"
					environment_type                  = "AllEnvironments"
					environments = []

					business_connectors = []
					non_business_connectors = []
					blocked_connectors = []

					custom_connectors_patterns = toset([
						{
							order            = 1
							host_url_pattern = "*"
							data_group       = "Ignore"
						  }
					])
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "id", policyId),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "display_name", "Block All Policy"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "default_connectors_classification", "Blocked"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environment_type", "AllEnvironments"),
				),
			},
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy_1"
					default_connectors_classification = "Blocked"
					environment_type                  = "AllEnvironments"
					environments = []

					business_connectors = []
					non_business_connectors = []
					blocked_connectors = []
					custom_connectors_patterns = toset([
						{
							order            = 1
							host_url_pattern = "*"
							data_group       = "Ignore"
							}
					])
				}`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "display_name", "Block All Policy_1"),
				),
			},
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy_1"
					default_connectors_classification = "General"
					environment_type                  = "OnlyEnvironments"
					environments = [ "00000000-0000-0000-0000-000000000000" ]

					business_connectors = []
					non_business_connectors = []
					blocked_connectors = []
					custom_connectors_patterns = toset([
						{
							order            = 1
							host_url_pattern = "*"
							data_group       = "Ignore"
							}
					])
				}`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "default_connectors_classification", "General"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environment_type", "OnlyEnvironments"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environments.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environments.0", "00000000-0000-0000-0000-000000000000"),
				),
			},
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy_1"
					default_connectors_classification = "General"
					environment_type                  = "OnlyEnvironments"
					environments = [ "00000000-0000-0000-0000-000000000000" ]

					non_business_connectors = toset([
							{
								id                           = "/providers/Microsoft.PowerApps/apis/shared_sql"
								default_action_rule_behavior = "",
								action_rules = [],
								endpoint_rules = [],
							}
						])
					business_connectors = toset([
							{
								id                           = "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
								default_action_rule_behavior = "",
								action_rules = [],
								endpoint_rules = [],
							}
						])
					blocked_connectors = toset([
							{
								id                           = "/providers/Microsoft.PowerApps/apis/shared_azureblob",
								default_action_rule_behavior = "",
								action_rules = [],
								endpoint_rules = [],
							}
						])
					custom_connectors_patterns = toset([
						{
							order            = 1
							host_url_pattern = "*"
							data_group       = "Ignore"
							}
					])
					}`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_sql"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_sharepointonline"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_azureblob"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.#", "1"),
				),
			},
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy_1"
					default_connectors_classification = "General"
					environment_type                  = "OnlyEnvironments"
					environments = [ "00000000-0000-0000-0000-000000000000" ]

					non_business_connectors = toset([
					{
						id                           = "/providers/Microsoft.PowerApps/apis/shared_sql"
						default_action_rule_behavior = "Allow",
						action_rules = [
							{
							action_id = "DeleteItem_V2",
							behavior  = "Block",
							},
							{
							action_id = "ExecutePassThroughNativeQuery_V2",
							behavior  = "Block",
							}
						],
						endpoint_rules = [
							{
							order    = 1,
							behavior = "Allow",
							endpoint = "contoso.com"
							},
							{
							order    = 2,
							behavior = "Deny",
							endpoint = "*"
							}
						]
					}
				])
				business_connectors = toset([
						{
							id                           = "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
							default_action_rule_behavior = "",
							action_rules = [],
							endpoint_rules = [],
						}
					])
				blocked_connectors = toset([
						{
							id                           = "/providers/Microsoft.PowerApps/apis/shared_azureblob",
							default_action_rule_behavior = "",
							action_rules = [],
							endpoint_rules = [],
						}
				])
				custom_connectors_patterns = toset([
					{
						order            = 1
						host_url_pattern = "*"
						data_group       = "Ignore"
					}
				])
				}`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_sql"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.default_action_rule_behavior", "Allow"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.0.action_id", "DeleteItem_V2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.0.behavior", "Block"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.1.action_id", "ExecutePassThroughNativeQuery_V2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.1.behavior", "Block"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.0.order", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.0.behavior", "Allow"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.0.endpoint", "contoso.com"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.1.order", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.1.behavior", "Deny"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.1.endpoint", "*"),
				),
			},
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy_1"
					default_connectors_classification = "General"
					environment_type                  = "OnlyEnvironments"
					environments = [ "00000000-0000-0000-0000-000000000000" ]

					non_business_connectors = toset([
							{
								id                           = "/providers/Microsoft.PowerApps/apis/shared_sql"
								default_action_rule_behavior = "Allow",
								action_rules = [
								{
									action_id = "DeleteItem_V2",
									behavior  = "Block",
								},
								{
									action_id = "ExecutePassThroughNativeQuery_V2",
									behavior  = "Block",
								}
								],
								endpoint_rules = [
								{
									order    = 1,
									behavior = "Allow",
									endpoint = "contoso.com"
								},
								{
									order    = 2,
									behavior = "Deny",
									endpoint = "*"
								}
								]
							}
						])
					business_connectors = toset([
							{
								id                           = "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
								default_action_rule_behavior = "",
								action_rules = [],
								endpoint_rules = [],
							}
						])
					blocked_connectors = toset([
							{
								id                           = "/providers/Microsoft.PowerApps/apis/shared_azureblob",
								default_action_rule_behavior = "",
								action_rules = [],
								endpoint_rules = [],
							}
						])
					custom_connectors_patterns = toset([
						{
							order            = 1
							host_url_pattern = "https://*.contoso.com"
							data_group       = "Blocked"
						},
						{
							order            = 2
							host_url_pattern = "*"
							data_group       = "Ignore"
						}
					])
				}`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.order", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.host_url_pattern", "https://*.contoso.com"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.data_group", "Blocked"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.order", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.host_url_pattern", "*"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.data_group", "Ignore"),
				),
			},
		},
	})
}

func TestUnitDataLossPreventionPolicyResource_Validate_Create(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", `https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusCreated, httpmock.File("tests/resource/Validate_Create/get_policy_00000000-0000-0000-0000-000000000001.json").String()), nil
		})

	httpmock.RegisterResponder("GET", `https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies/00000000-0000-0000-0000-000000000001`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, httpmock.File("tests/resource/Validate_Create/get_policy_00000000-0000-0000-0000-000000000001.json").String()), nil
		})

	httpmock.RegisterResponder("DELETE", `https://api.bap.microsoft.com/providers/PowerPlatform.Governance/v2/policies/00000000-0000-0000-0000-000000000001`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, ""), nil
		})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "Block All Policy"
					default_connectors_classification = "Blocked"
					environment_type                  = "OnlyEnvironments"
					environments = [ "00000000-0000-0000-0000-000000000000" ]

					non_business_connectors = toset([
						{
							id                           = "/providers/Microsoft.PowerApps/apis/shared_sql"
							default_action_rule_behavior = "Allow",
							action_rules = [
							  {
								action_id = "DeleteItem_V2",
								behavior  = "Block",
							  },
							  {
								action_id = "ExecutePassThroughNativeQuery_V2",
								behavior  = "Block",
							  }
							],
							endpoint_rules = [
							  {
								order    = 1,
								behavior = "Allow",
								endpoint = "contoso.com"
							  },
							  {
								order    = 2,
								behavior = "Deny",
								endpoint = "*"
							  }
							]
						  }
					])
					business_connectors = toset([
						{
							id                           = "/providers/Microsoft.PowerApps/apis/shared_sharepointonline",
							default_action_rule_behavior = "",
							action_rules                 = [],
							endpoint_rules               = []
						},
					])
					blocked_connectors      = toset([
						{
							id                           = "/providers/Microsoft.PowerApps/apis/shared_azureblob",
							default_action_rule_behavior = "",
							action_rules                 = []
							endpoint_rules               = []
						  },
					])
					custom_connectors_patterns = toset([
					  {
						order            = 1
						host_url_pattern = "https://*.contoso.com"
						data_group       = "Blocked"
					  },
					  {
						order            = 2
						host_url_pattern = "*"
						data_group       = "Ignore"
					  }
					])
				  }`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "id", "00000000-0000-0000-0000-000000000001"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "display_name", "Block All Policy"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "default_connectors_classification", "Blocked"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environment_type", "OnlyEnvironments"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environments.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environments.0", "00000000-0000-0000-0000-000000000000"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_sql"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.default_action_rule_behavior", "Allow"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.0.action_id", "DeleteItem_V2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.0.behavior", "Block"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.1.action_id", "ExecutePassThroughNativeQuery_V2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.1.behavior", "Block"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.0.order", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.0.behavior", "Allow"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.0.endpoint", "contoso.com"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.1.order", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.1.behavior", "Deny"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.1.endpoint", "*"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_sharepointonline"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.#", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_azureblob"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.order", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.host_url_pattern", "https://*.contoso.com"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.data_group", "Blocked"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.order", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.host_url_pattern", "*"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.data_group", "Ignore"),
				),
			},
		},
	})
}

func TestAccDataLossPreventionPolicyResource_Validate_Create(t *testing.T) {
	t.Skip("Skipping as there is inconsistency in API in connectors returned")

	resource.Test(t, resource.TestCase{
		IsUnitTest:               false,
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "powerplatform_environment" "env" {
					display_name     = "` + mocks.TestName() + `"
					location         = "unitedstates"
					environment_type = "Sandbox"
				}

				data "powerplatform_connectors" "all_connectors" {}

				locals {
					business_connectors = toset([
					  {
						action_rules = [
						  {
							action_id = "DeleteItem_V2"
							behavior  = "Block"
						  },
						  {
							action_id = "ExecutePassThroughNativeQuery_V2"
							behavior  = "Block"
						  },
						]
						default_action_rule_behavior = "Allow"
						endpoint_rules = [
						  {
							behavior = "Allow"
							endpoint = "contoso.com"
							order    = 1
						  },
						  {
							behavior = "Deny"
							endpoint = "*"
							order    = 2
						  },
						]
						id = "/providers/Microsoft.PowerApps/apis/shared_sql"
					  },
					  {
						action_rules                 = []
						default_action_rule_behavior = ""
						endpoint_rules               = []
						id                           = "/providers/Microsoft.PowerApps/apis/shared_approvals"
					  },
					  {
						action_rules                 = []
						default_action_rule_behavior = ""
						endpoint_rules               = []
						id                           = "/providers/Microsoft.PowerApps/apis/shared_cloudappsecurity"
					  }
					])
				  
					non_business_connectors = toset([for conn
					  in data.powerplatform_connectors.all_connectors.connectors :
					  {
						id                           = conn.id
						name                         = conn.name
						default_action_rule_behavior = ""
						action_rules                 = [],
						endpoint_rules               = []
					  }
					  if conn.unblockable == true && !contains([for bus_conn in local.business_connectors : bus_conn.id], conn.id)
					])
				  
					blocked_connectors = toset([for conn
					  in data.powerplatform_connectors.all_connectors.connectors :
					  {
						id                           = conn.id
						default_action_rule_behavior = ""
						action_rules                 = [],
						endpoint_rules               = []
					  }
					if conn.unblockable == false && !contains([for bus_conn in local.business_connectors : bus_conn.id], conn.id)])
				  }

				  resource "powerplatform_data_loss_prevention_policy" "my_policy" {
					display_name                      = "` + mocks.TestName() + `"
					default_connectors_classification = "Blocked"
					environment_type                  = "OnlyEnvironments"
					environments                      = [powerplatform_environment.env.id]
				  
					business_connectors     = local.business_connectors
					non_business_connectors = local.non_business_connectors
					blocked_connectors      = local.blocked_connectors
				  
					custom_connectors_patterns = toset([
					  {
						order            = 1
						host_url_pattern = "https://*.contoso.com"
						data_group       = "Blocked"
					  },
					  {
						order            = 2
						host_url_pattern = "*"
						data_group       = "Ignore"
					  }
					])
				  }`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "display_name", mocks.TestName()),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "default_connectors_classification", "Blocked"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environment_type", "OnlyEnvironments"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "environments.#", "1"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.#", "3"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_sql"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.1.id", "/providers/Microsoft.PowerApps/apis/shared_approvals"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.2.id", "/providers/Microsoft.PowerApps/apis/shared_cloudappsecurity"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.default_action_rule_behavior", "Allow"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.0.action_id", "DeleteItem_V2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.0.behavior", "Block"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.1.action_id", "ExecutePassThroughNativeQuery_V2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.action_rules.1.behavior", "Block"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.0.order", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.0.behavior", "Allow"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.0.endpoint", "contoso.com"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.1.order", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.1.behavior", "Deny"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "business_connectors.0.endpoint_rules.1.endpoint", "*"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.#", "22"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_commondataservice"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "non_business_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.id", "/providers/Microsoft.PowerApps/apis/shared_10to8"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.default_action_rule_behavior", ""),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.action_rules.#", "0"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "blocked_connectors.0.endpoint_rules.#", "0"),

					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.#", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.order", "1"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.host_url_pattern", "https://*.contoso.com"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.0.data_group", "Blocked"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.order", "2"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.host_url_pattern", "*"),
					resource.TestCheckResourceAttr("powerplatform_data_loss_prevention_policy.my_policy", "custom_connectors_patterns.1.data_group", "Ignore"),
				),
			},
		},
	})
}
