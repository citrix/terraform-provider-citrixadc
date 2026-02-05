---
subcategory: "Network"
---

# Data Source: citrixadc_forwardingsession

Use this data source to retrieve information about an existing Forwarding Session.

The `citrixadc_forwardingsession` data source allows you to retrieve details of a forwarding session by its name. This is useful for referencing existing forwarding session rules in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing forwarding session
data "citrixadc_forwardingsession" "example" {
  name = "my_forwarding_session"
}

# Reference forwarding session attributes
output "forwarding_network" {
  value = data.citrixadc_forwardingsession.example.network
}

output "forwarding_netmask" {
  value = data.citrixadc_forwardingsession.example.netmask
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the forwarding session rule. Can begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the forwarding session (same as name).

* `acl6name` - Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as a forwarding session rule.

* `aclname` - Name of any configured ACL whose action is ALLOW. The rule of the ACL is used as a forwarding session rule.

* `connfailover` - Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the forwarding session. Possible values: `ENABLED`, `DISABLED`.

* `netmask` - Subnet mask associated with the network.

* `network` - An IPv4 network address or IPv6 prefix of a network from which the forwarded traffic originates or to which it is destined.

* `processlocal` - Enabling this option on forwarding session will not steer the packet to flow processor. Instead, packet will be routed. Possible values: `ENABLED`, `DISABLED`.

* `sourceroutecache` - Cache the source ip address and mac address of the DA servers. Possible values: `ENABLED`, `DISABLED`.

* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Common Use Cases

### Retrieve Forwarding Session for Network Configuration

```hcl
data "citrixadc_forwardingsession" "app_forwarding" {
  name = "production_forwarding_session"
}

# Use the retrieved forwarding session details
output "forwarding_session_network" {
  value = data.citrixadc_forwardingsession.app_forwarding.network
}

output "ha_sync_enabled" {
  value = data.citrixadc_forwardingsession.app_forwarding.connfailover
}
```

### Reference Forwarding Session for Validation

```hcl
data "citrixadc_forwardingsession" "existing_session" {
  name = "existing_forwarding_session"
}

# Verify session configuration
locals {
  is_ha_enabled = data.citrixadc_forwardingsession.existing_session.connfailover == "ENABLED"
  uses_source_route_cache = data.citrixadc_forwardingsession.existing_session.sourceroutecache == "ENABLED"
}

output "forwarding_session_info" {
  value = {
    name = data.citrixadc_forwardingsession.existing_session.name
    network = data.citrixadc_forwardingsession.existing_session.network
    ha_enabled = local.is_ha_enabled
  }
}
```
