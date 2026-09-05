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
* `id` - The id of the subscriberradiusinterface. It is a system-generated identifier.

### Read-only subscriberradiusinterface metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_subscriberradiusinterface` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `svrstate` - The state of the radius service.
