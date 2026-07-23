---
subcategory: "Kafka"
---

# Data Source: kafkacluster_servicegroup_binding

The kafkacluster_servicegroup_binding data source allows you to retrieve information about a service group bound to a Kafka cluster.


## Example usage

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
