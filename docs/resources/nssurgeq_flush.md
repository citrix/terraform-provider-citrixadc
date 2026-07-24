---
subcategory: "NS"
---

# Resource: nssurgeq_flush

This resource is used to flush the surge queue on the Citrix ADC.


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

* `id` - The id of the nssurgeq_flush resource. It is set to `nssurgeq_flush`.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it flushes the surge queue at the moment of apply. To flush the queue again, taint the resource or change any of the scoping arguments. Importing this resource is not meaningful.
