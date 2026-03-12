---
subcategory: "SMPP"
---

# citrixadc_smppuser (Data Source)

Data source for querying Citrix ADC SMPP users. This data source retrieves information about SMPP users configured on the ADC appliance for SMPP protocol operations.

## Example Usage

```hcl
data "citrixadc_smppuser" "example" {
  username = "smpp_user1"
}

# Output user attributes
output "smpp_password" {
  value = data.citrixadc_smppuser.example.password
  sensitive = true
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) Name of the SMPP user. Must be the same as the user name specified in the SMPP server.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the smppuser datasource.
* `password` - Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server.

## Notes

SMPP (Short Message Peer-to-Peer) users are configured for authentication when the ADC communicates with SMPP servers for SMS message handling.
