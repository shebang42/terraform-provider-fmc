---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_icmpv4_objects Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This plural resource manages a bulk of ICMPv4. The FMC API supports quick bulk creation of this resource. Deletion of this resource is done one-by-one or in bulk, depending of FMC version. Modification is always done one-by-one. Updating/deleting fmc_icmpv4_objects can thus take much more time than creating it
---

# fmc_icmpv4_objects (Resource)

This plural resource manages a bulk of ICMPv4. The FMC API supports quick bulk creation of this resource. Deletion of this resource is done one-by-one or in bulk, depending of FMC version. Modification is always done one-by-one. Updating/deleting `fmc_icmpv4_objects` can thus take much more time than creating it

## Example Usage

```terraform
resource "fmc_icmpv4_objects" "example" {
  items = {
    icmpv4_1 = {
      description = "ICMPv4 network unreachable response, type 3, code 0"
      icmp_type   = 3
      code        = 0
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `items` (Attributes Map) Map of icmpv4s. The key of the map is the name of the individual ICMPv4 Object. Renaming ICMPv4 objects in bulk is not yet implemented. (see [below for nested schema](#nestedatt--items))

### Optional

- `domain` (String) The name of the FMC domain

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Optional:

- `code` (Number) ICMPv4 [code number](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml) subordinate to the given `icmp_type`.
  - Range: `0`-`255`
- `description` (String) Optional description of the resource.
- `icmp_type` (Number) ICMPv4 [type number](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml).
  - Range: `0`-`255`
- `overridable` (Boolean) Indicates whether object values can be overridden.

Read-Only:

- `id` (String) UUID of the managed ICMPv4 object.

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_icmpv4_objects.example "<domain>,[<icmpv4_objects_name>]"
```