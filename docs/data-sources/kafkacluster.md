---
subcategory: "Kafka"
---

# Data Source: kafkacluster

The kafkacluster data source allows you to retrieve information about an existing Kafka cluster configured on the Citrix ADC by looking it up by name.

The `kafka` feature must be licensed and enabled on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_kafkacluster" "tf_kafkacluster" {
  name = "kafka-cluster-1"
}

output "kafkacluster_id" {
  value = data.citrixadc_kafkacluster.tf_kafkacluster.id
}
```


## Argument Reference

* `name` - (Required) Name of the Kafka cluster to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the kafkacluster. It has the same value as the `name` attribute.
* `name` - Name of the Kafka cluster.
