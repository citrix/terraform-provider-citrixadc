---
subcategory: "Kafka"
---

# Data Source: kafkacluster_servicegroup_binding

Retrieves information about a service group bound to a Kafka cluster. Look the binding up by the Kafka cluster name and the bound service group name. The Kafka feature must be available and licensed on the appliance.


## Example Usage

```terraform
data "citrixadc_kafkacluster_servicegroup_binding" "tf_binding" {
  name             = "kafkacluster1"
  servicegroupname = "kafka_brokers"
}

output "bound_servicegroup" {
  value = data.citrixadc_kafkacluster_servicegroup_binding.tf_binding.servicegroupname
}
```


## Argument Reference

* `name` - (Required) Name of the Kafka cluster.
* `servicegroupname` - (Required) Name of the bound service group used to identify the binding within the cluster.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The composite ID of the binding, in the form `name:<name>,servicegroupname:<servicegroupname>` (values URL-encoded).
