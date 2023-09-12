variable "tenant_id" {
  description = "The AAD tenant id of service principal or user"
  type        = string
  default     = "***"
}

variable "client_id" {
  description = "The client ID of the of the service principal"
  default     = "***"
  type        = string

}
variable "client_secret" {
  description = "The client secret of the service principal"
  sensitive   = true
  type        = string
  default     = "..."
}

variable "subscription_id" {
  description = "The Azure subscription for ASLG2 deployment"
  default     = "***"
  sensitive   = true
  type        = string
}

variable "prefix" {
  default     = "terraformppp"
  description = "The Prefix used for all resources in this example"
}

variable "location" {
  default     = "westeurope"
  description = "The Azure Region in which all resources in this example should be created."
}

variable "dataverse_env_display_name" {
  default     = "synapse_link"
  description = "The display name for the dataverse environment."
}

variable "dataverse_location" {
  default     = "europe"
  description = "Region where the dataverse instance is provisioned."
}

variable "dataverse_language_code" {
  default     = "1033"
  description = "Language code for Dataverse."
}

variable "dataverse_currency_code" {
  default     = "USD"
  description = "Language code for Dataverse."
}

variable "dataverse_environment_type" {
  default     = "Production"
  description = "Environment type for the dataverse to reside in."
}

variable "tenant_domain" {
  default     = "mngenvmcap080290"
  description = "Lowercase domain name for the tenant."
}

variable "dataverse_security_group_id" {
  default     = "00000000-0000-0000-0000-000000000000"
  description = "Security group id to be granted access to the Dataverse."
}
