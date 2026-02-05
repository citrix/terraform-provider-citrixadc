---
subcategory: "Subscriber"
---

# Data Source: subscriberradiusinterface

The subscriberradiusinterface data source allows you to retrieve information about the subscriber RADIUS interface configuration.

## Example usage

```terraform
data "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
}

output "listeningservice" {
  value = data.citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface.listeningservice
}

output "radiusinterimasstart" {
  value = data.citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface.radiusinterimasstart
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `listeningservice` - Name of RADIUS LISTENING service that will process RADIUS accounting requests.
* `radiusinterimasstart` - Treat radius interim message as start radius messages.

## Attribute Reference

* `id` - The id of the subscriberradiusinterface. It is a system-generated identifier.
