terraform {
  required_providers {
    powerplatform = {
      source = "microsoft/power-platform"
    }
  }
}

provider "powerplatform" {
  tenant_id = var.tenant_id
  client_id = var.client_id
  secret    = var.secret
}

resource "powerplatform_tenant_settings" "settings" {
  walk_me_opt_out                              = false
  disable_support_tickets_visible_by_all_users = true
}


