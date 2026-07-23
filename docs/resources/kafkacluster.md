---
subcategory: "Kafka"
---

# Resource: kafkacluster

This resource is used to manage a Kafka cluster on the Citrix ADC.

~> **Prerequisite:** The `kafka` feature must be licensed and enabled. This is a create-only resource; changing `name` forces re-creation.


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
