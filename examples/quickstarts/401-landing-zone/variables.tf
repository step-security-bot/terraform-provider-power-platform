variable "client_id" {
  description = "The username of the Power Platform API in user@domain format"
  type        = string
}

variable "secret" {
  description = "The password of the Power Platform API user"
  sensitive   = true
  type        = string
}

variable "tenant_id" {
  description = "The AAD tenant id of service principal or user"
  type        = string
}

variable "PPGuestMakerSetting" {
  type    = string
  default = null
}

variable "PPAppSharingSetting" {
  type    = string
  default = null
}

variable "PPEnvCreationSetting" {
  type    = string
  default = null

}

variable "PPTrialEnvCreationSetting" {
  type    = string
  default = null
}

variable "PPEnvCapacitySetting" {
  type    = string
  default = null
}

variable "PPTenantIsolationSetting" {
  type    = string
  default = null
}

variable "PPTenantDLP" {
  type    = string
  default = null
}

variable "PPTenantIsolationDomains" {
  type    = string
  default = null
}

variable "PPAdminEnvNaming" {
  description = "value for the environment naming convention"
  type        = string
  default     = null
}

variable "PPAdminRegion" {
  description = "The region where the environment will be deployed"
  type        = string
  validation {
    condition     = contains(["unitedstates", "europe", "asia", "australia", "india", "japan", "canada", "unitedkingdom", "unitedstatesfirstrelease", "southamerica", "france", "switzerland", "germany", "unitedarabemirates", "norway"], var.PPAdminRegion)
    error_message = "The region must be a valid region"
  }
  default = "unitedstates"
}

# param (
#     #Security, govarnance and compliance
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPGuestMakerSetting,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPAppSharingSetting,
#     #Admin environment and settings
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPEnvCreationSetting,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPTrialEnvCreationSetting,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPEnvCapacitySetting,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPTenantIsolationSetting,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPTenantDLP,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPTenantIsolationDomains,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPAdminEnvNaming,
#     [ValidateSet('unitedstates', 'europe', 'asia', 'australia', 'india', 'japan', 'canada', 'unitedkingdom', 'unitedstatesfirstrelease', 'southamerica', 'france', 'switzerland', 'germany', 'unitedarabemirates', 'norway')][Parameter(Mandatory = $false)][string]$PPAdminRegion,
#     [Parameter(Mandatory = $false)][string]$PPAdminBilling,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPAdminCoeSetting,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPAdminDlp,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPAdminEnvEnablement,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPAdminManagedEnv,
#     #Landing Zones
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPDefaultRenameText,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPDefaultDLP,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPDefaultManagedEnv,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPDefaultManagedSharing,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizen,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenCount,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenNaming,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenRegion,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenDlp,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenBilling,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenManagedEnv,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenAlm,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenDescription,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenCurrency,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPCitizenLanguage,
#     [Parameter(Mandatory = $false)]$PPCitizenConfiguration,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPPro,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProCount,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProNaming,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProRegion,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProDlp,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProBilling,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProManagedEnv,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProAlm,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProDescription,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProCurrency,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPProLanguage,
#     [Parameter(Mandatory = $false)]$PPProConfiguration,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPSelectIndustry,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPIndustryNaming,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPIndustryRegion,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPIndustryBilling,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPIndustryAlm,
#     [Parameter(Mandatory = $false)][string][AllowEmptyString()][AllowNull()]$PPIndustryManagedEnv
# )
