---
subcategory: "NS"
---

# Resource: nsconfig_unset

The nsconfig_unset resource is used to apply the unset operation for ns config. It
resets the named settable `ns config` parameters to their appliance defaults
(NITRO `unset ns config`).


## Example usage

```hcl
resource "citrixadc_nsconfig_unset" "tf_nsunset" {
  attributes = [
    "maxconn",
    "maxreq",
    "cip",
    "timezone",
    "grantquotamaxclient",
  ]
  timestamp = "2024-06-01T12:00:00"
}
```

~> **WARNING** Unsetting `ipaddress`, `netmask`, `ifnum`, `nsvlan` or `tagged`
resets the appliance management addressing and can disconnect the box. Only unset
those parameters if you understand the impact.


## Argument Reference

* `attributes` - (Required) Set of nsconfig parameter names to reset to their appliance defaults. Allowed values: `nsvlan`, `securemanagementtraffic`, `ftpportrange`, `crportrange`, `timezone`, `ipaddress`, `netmask`, `ifnum`, `tagged`, `httpport`, `maxconn`, `maxreq`, `cip`, `cipheader`, `cookieversion`, `securecookie`, `pmtumin`, `pmtutimeout`, `grantquotamaxclient`, `exclusivequotamaxclient`, `grantquotaspillover`, `exclusivequotaspillover`.
* `timestamp` - (Required) Timestamp marker used as the resource ID. Change it to re-run the `unset ns config` action. All attributes force replacement, so bumping this value is how the action is re-applied.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the nsconfig_unset resource (equals the configured timestamp).
