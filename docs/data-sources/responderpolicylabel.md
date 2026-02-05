---
subcategory: "Responder"
---

# Data Source: citrixadc_responderpolicylabel

The `citrixadc_responderpolicylabel` data source is used to retrieve information about an existing responder policy label configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_responderpolicylabel" "example" {
  labelname = "my_responderpolicylabel"
}
```

## Argument Reference

The following arguments are supported:

* `labelname` - (Required) Name of the responder policy label to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the responder policy label (same as labelname).
* `comment` - Any comments to preserve information about this responder policy label.
* `newname` - New name for the responder policy label (if renamed).
* `policylabeltype` - Type of responses sent by the policies bound to this policy label. Possible values:
  * `HTTP` - HTTP responses.
  * `OTHERTCP` - NON-HTTP TCP responses.
  * `SIP_UDP` - SIP responses.
  * `SIP_TCP` - SIP responses.
  * `MYSQL` - MYSQL responses.
  * `MSSQL` - MSSQL responses.
  * `NAT` - NAT responses.
  * `DIAMETER` - DIAMETER responses.
  * `RADIUS` - RADIUS responses.
  * `DNS` - DNS responses.
