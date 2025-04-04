---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "grafana_fleet_management_collector Resource - terraform-provider-grafana"
subcategory: "Fleet Management"
description: |-
  Manages Grafana Fleet Management collectors.
  Official documentation https://grafana.com/docs/grafana-cloud/send-data/fleet-management/API documentation https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/collector-api/
  Required access policy scopes:
  fleet-management:readfleet-management:write
---

# grafana_fleet_management_collector (Resource)

Manages Grafana Fleet Management collectors.

* [Official documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/)
* [API documentation](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/collector-api/)

Required access policy scopes:

* fleet-management:read
* fleet-management:write

## Example Usage

```terraform
resource "grafana_fleet_management_collector" "test" {
  id = "my_collector"
  remote_attributes = {
    "env"   = "PROD",
    "owner" = "TEAM-A"
  }
  enabled = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) ID of the collector

### Optional

- `enabled` (Boolean) Whether the collector is enabled or not
- `remote_attributes` (Map of String) Remote attributes for the collector

## Import

Import is supported using the following syntax:

```shell
terraform import grafana_fleet_management_collector.name "{{ id }}"
```
