---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_port_groups Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source can read the Port Groups.
---

# fmc_port_groups (Data Source)

This data source can read the Port Groups.

## Example Usage

```terraform
data "fmc_port_groups" "example" {
  items = {
    "port_group_1" = {
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) The name of the FMC domain
- `items` (Attributes Map) Map of port groups. The key of the map is the name of the individual Port Group. (see [below for nested schema](#nestedatt--items))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `description` (String) Optional user-created description.
- `id` (String) UUID of the managed Port Groups.
- `objects` (Attributes Set) (see [below for nested schema](#nestedatt--items--objects))
- `overridable` (Boolean) Indicates whether object values can be overridden.
- `type` (String) Type of the object; this value is always 'PortObjectGroup'.

<a id="nestedatt--items--objects"></a>
### Nested Schema for `items.objects`

Read-Only:

- `id` (String) UUID of the port (such as fmc_port.test.id, etc.).
- `type` (String)
