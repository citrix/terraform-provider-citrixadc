---
subcategory: "NS"
---

# Resource: nssurgeq_flush

Flushes the surge queue on the Citrix ADC. The surge queue holds client connections that are waiting for the appliance to open new server-side connections when a server is momentarily unable to accept the offered load. Flushing it clears those queued requests, which is helpful when draining a server, recovering from a backlog, or testing surge-protection behavior. The flush can be applied system-wide, or scoped to a specific virtual server, service, or service group member.

This is an **action-only** resource. Applying it triggers the NITRO `flush` action on the `nssurgeq` endpoint. There is no corresponding ADC object to read back, so this resource performs a one-shot action on create. Read and update are no-ops, and importing it is not meaningful.


## Example usage

System-wide flush (no arguments):

```hcl
resource "citrixadc_nssurgeq_flush" "flush_all" {}
```

Flush the surge queue of a specific service group member:

```hcl
resource "citrixadc_nssurgeq_flush" "flush_member" {
  name       = "svcgrp1"
  servername = "websrv1"
  port       = 80
}
```


## Argument Reference

* `name` - (Optional) Name of a virtual server, service, or service group for which the surge queue must be flushed. If omitted, the flush applies system-wide. Changing this value forces the action to be re-applied (replacement).
* `servername` - (Optional) Name of a service group member. This argument is needed when you want to flush the surge queue of a service group member; it requires `name` to be set. Changing this value forces the action to be re-applied (replacement).
* `port` - (Optional) Port on which the server is bound to the entity (service group). This argument requires `servername` to be set. Changing this value forces the action to be re-applied (replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the fixed value `nssurgeq_flush`. Because `nssurgeq_flush` is an action-only resource with no NITRO GET endpoint, the ID does not correspond to a readable ADC object.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it flushes the surge queue at the moment of apply; the result cannot be read back from the ADC, so subsequent plans will not detect drift. To flush the queue again, taint the resource or change any of the scoping arguments. Destroying the resource only removes it from Terraform state and has no effect on the appliance. Importing this resource is not meaningful.
