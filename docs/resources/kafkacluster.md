---
subcategory: "Kafka"
---

# Resource: kafkacluster

The kafkacluster resource defines a named Kafka cluster on the Citrix ADC. The cluster acts as a logical grouping that Kafka broker service groups are later bound to, allowing the ADC to load balance and proxy traffic across the brokers that make up the cluster.

The `kafka` feature must be licensed and enabled on the Citrix ADC before this resource can be created.

Note: This is a create-only resource. The Kafka cluster has no updatable attributes; changing `name` forces the resource to be destroyed and re-created.


## Example usage

```hcl
resource "citrixadc_kafkacluster" "tf_kafkacluster" {
  name = "kafka-cluster-1"
}
```


## Argument Reference

* `name` - (Required) Name for the Kafka cluster. Maximum length = 127. Cannot be changed after the cluster is created; changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the kafkacluster. It has the same value as the `name` attribute.


## Import

A kafkacluster can be imported using its name, e.g.

```shell
terraform import citrixadc_kafkacluster.tf_kafkacluster kafka-cluster-1
```
