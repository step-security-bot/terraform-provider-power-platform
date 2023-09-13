terraform {
  required_providers {
    powerplatform = {
      version = "0.2"
      source  = "microsoft/power-platform"
    }
  }
}

provider "powerplatform" {
  client_id = var.client_id
  secret    = var.secret
  tenant_id = var.tenant_id
}

resource "powerplatform_environment" "Hack-Dev" {
  display_name      = "Hack-Dev"
  location          = "unitedstates"
  language_code     = "1033"
  currency_code     = "USD"
  environment_type  = "SubscriptionBasedTrial"
  domain            = "org4bfeba3b"
  security_group_id = "00000000-0000-0000-0000-000000000000"
}
