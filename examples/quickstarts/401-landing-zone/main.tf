terraform {
  required_version = ">= 1.5"
  required_providers {
    powerplatform = {
      source = "microsoft/power-platform"
    }
  }
}

provider "powerplatform" {
  client_id = var.client_id
  secret    = var.secret
  tenant_id = var.tenant_id
}

data "powerplatform_tenant_settings" "tenant_settings" {

}
