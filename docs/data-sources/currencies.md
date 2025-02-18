---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "powerplatform_currencies Data Source - powerplatform"
subcategory: ""
description: |-
  Fetches the list of available Dynamics 365 currencies. For more information see Power Platform Currencies https://learn.microsoft.com/en-us/power-platform/admin/manage-transactions-with-multiple-currencies
---

# powerplatform_currencies (Data Source)

Fetches the list of available Dynamics 365 currencies. For more information see [Power Platform Currencies](https://learn.microsoft.com/en-us/power-platform/admin/manage-transactions-with-multiple-currencies)

## Example Usage

```terraform
terraform {
  required_providers {
    powerplatform = {
      source  = "microsoft/power-platform"
    }
  }
}

provider "powerplatform" {
  use_client = true
}

data "powerplatform_locations" "all_locations" {}

data "powerplatform_currencies" "all_currencies_by_location" {
  location = data.powerplatform_locations.all_locations.locations[0].name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `location` (String) Location of the currencies

### Optional

- `id` (Number) Id of the read operation

### Read-Only

- `currencies` (Attributes List) List of available currencies (see [below for nested schema](#nestedatt--currencies))

<a id="nestedatt--currencies"></a>
### Nested Schema for `currencies`

Read-Only:

- `code` (String) Code of the location
- `id` (String) Unique identifier of the currency
- `is_tenant_default` (Boolean) Is the currency the default for the tenant
- `name` (String) Name of the currency
- `symbol` (String) Symbol of the currency
- `type` (String) Type of the currency
