---
subcategory: "RDP"
---

# Data Source: rdpserverprofile

The `citrixadc_rdpserverprofile` data source is used to retrieve information about a specific RDP server profile configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_rdpserverprofile" "example" {
  name = "my_rdpserverprofile"
}
```

## Argument Reference

* `name` - (Required) The name of the RDP server profile.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the RDP server profile.
* `psk` - Pre shared key value.
* `rdpip` - IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.
* `rdpport` - TCP port on which the RDP connection is established.
* `rdpredirection` - Enable/Disable RDP redirection support.

### Read-only rdpserverprofile metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_rdpserverprofile` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
