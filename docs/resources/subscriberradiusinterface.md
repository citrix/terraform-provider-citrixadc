---
subcategory: "Subscriber"
---

# Resource: subscriberradiusinterface

The subscriberradiusinterface resource is used to create subscriberradiusinterface.


## Example usage

```hcl
resource "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
  listeningservice     = citrixadc_service.tf_service.name
  radiusinterimasstart = "ENABLED"
}

resource "citrixadc_service" "tf_service" {
  name        = "srad1"
  port        = 1813
  ip          = "192.0.0.206"
  servicetype = "RADIUSListener"
}
```


## Argument Reference

* `listeningservice` - (Required) Name of RADIUS LISTENING service that will process RADIUS accounting requests.
* `radiusinterimasstart` - (Optional) Treat radius interim message as start radius messages.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the subscriberradiusinterface. It is a unique string prefixed with `tf-subscriberradiusinterface-`.