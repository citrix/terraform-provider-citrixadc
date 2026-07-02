---
subcategory: "Kafka"
---

# Resource: kafkacluster_servicegroup_binding

Binds a service group to a Kafka cluster so that the Citrix ADC can distribute Kafka broker traffic across the members of that service group. Use this binding to associate the back-end brokers (grouped in a service group) with the Kafka cluster definition that fronts them. The Kafka feature must be available and licensed on the appliance.

~> **Immutable binding.** This binding supports bind and unbind operations only; there is no NITRO set/update endpoint. Both `name` and `servicegroupname` force replacement (`RequiresReplace`), so changing either attribute unbinds the existing pairing and creates a new one.


## Example usage

```hcl
resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "kafka_brokers"
  servicetype      = "TCP"
}

resource "citrixadc_kafkacluster_servicegroup_binding" "tf_binding" {
  name             = "kafkacluster1"
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Kafka cluster to which the service group is bound. Changing this forces a new resource to be created.
* `servicegroupname` - (Required) Name of the service group to bind to the Kafka cluster. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `kafkacluster_servicegroup_binding` resource. It is a comma-separated, URL-encoded composite key of the form `name:<name>,servicegroupname:<servicegroupname>`, for example `name:kafkacluster1,servicegroupname:kafka_brokers`.


## Import

A `kafkacluster_servicegroup_binding` resource can be imported using its composite id (`name:<name>,servicegroupname:<servicegroupname>`), e.g.

```shell
terraform import citrixadc_kafkacluster_servicegroup_binding.tf_binding name:kafkacluster1,servicegroupname:kafka_brokers
```
