---
subcategory: "NS"
---

# Data Source: nsaptlicense

The nsaptlicense data source allows you to retrieve information about an allocated Citrix ADC APT/CADS pooled license record on the appliance, looked up by its hardware serial number / license activation code.


## Example usage

```terraform
data "citrixadc_nsaptlicense" "example" {
  serialno = "ABC123XYZ789"
}

output "nsaptlicense_countavailable" {
  value = data.citrixadc_nsaptlicense.example.countavailable
}
```


## Argument Reference

* `serialno` - (Required) Hardware Serial Number/License Activation Code (LAC) used to look up the allocated license record.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - License ID.
* `bindtype` - Bind type.
* `countavailable` - The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses.
* `licensedir` - License Directory.
* `sessionid` - Session ID.
* `useproxy` - Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option.
