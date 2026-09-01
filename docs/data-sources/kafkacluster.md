---
subcategory: "Kafka"
---

# Data Source: kafkacluster

The kafkacluster data source allows you to retrieve information about a Kafka cluster on the Citrix ADC.


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

### Read-only kafkacluster metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_kafkacluster` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `activesvc` - Total active services bound to servicegroup.
* `totalsvc` - Total services bound to servicegroup.
* `topicname` - Topic of the servicegroup.
* `numtopics` - Total number of topic servicegroups bound.
