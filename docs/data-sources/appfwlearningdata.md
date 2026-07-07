---
subcategory: "Application Firewall"
---

# Data Source: appfwlearningdata

The appfwlearningdata data source reads the Citrix ADC Application-Firewall **learned-data** table.

~> **BEST-EFFORT MODEL** NITRO exposes only a `get(all)` view for `appfwlearningdata` (there is no per-object GET). This data source reads the table — scoped by `profilename` (and optionally `securitycheck`) when supplied — and surfaces the **first** matching learned-data entry. NITRO can return many rows; this data source reports a single representative row, so verify against the live table if you need the full set. The learned-data table is only populated after the App-Firewall has accumulated learning for the given profile / security check.


## Example usage

```terraform
data "citrixadc_appfwlearningdata" "example" {
  profilename   = "my_appfwprofile"
  securitycheck = "startURL"
}

output "learned_url" {
  value = data.citrixadc_appfwlearningdata.example.url
}
```


## Argument Reference

* `profilename` - (Optional) Name of the profile to look up learned data for. When set, it scopes the NITRO lookup; otherwise the whole table is read.
* `securitycheck` - (Optional) Name of the security check to look up learned data for.


## Attribute Reference

In addition to the arguments, the following attributes are available (from the first matching learned-data entry):

* `id` - Synthetic id of the data source read (`appfwlearningdata-config`).
* `url` - Learnt URL.
* `name` - Learnt field name.
* `fieldtype` - Learnt field type.
* `fieldformatminlength` - The minimum allowed length for data in this form field.
* `fieldformatmaxlength` - The maximum allowed length for data in this form field.
* `fieldformatcharmappcre` - Form field value allowed character map.
* `value_type` - Learnt field value type.
* `value` - Learnt field value.
* `hits` - Learnt entity hit count.
* `data` - Learned data.
