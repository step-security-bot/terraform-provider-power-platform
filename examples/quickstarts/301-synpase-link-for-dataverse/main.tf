terraform {
  required_providers {
    powerplatform = {
      version = ">= 0.1"
      source  = "microsoft/power-platform"
    }
    azurerm = {
      version = ">= 1.5"
      source  = "hashicorp/azurerm"
    }
  }
}

provider "powerplatform" {
  client_id = var.client_id
  secret    = var.client_secret
  tenant_id = var.tenant_id
}

resource "powerplatform_environment" "development" {
  display_name      = var.dataverse_env_display_name
  location          = var.dataverse_location
  language_code     = var.dataverse_language_code
  currency_code     = var.dataverse_currency_code
  environment_type  = var.dataverse_environment_type
  domain            = var.tenant_domain
  security_group_id = var.dataverse_security_group_id
}

# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {}

  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
  subscription_id = var.subscription_id
}

resource "azurerm_resource_group" "example" {
  name     = "${var.prefix}-resources"
  location = var.location
}

resource "azurerm_storage_account" "example" {
  name                     = "${var.prefix}storageacct"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  network_rules {
    default_action = "Deny"
    ip_rules       = ["23.45.1.0/30"]
  }
}
