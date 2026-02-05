---
subcategory: "Network"
---

# Data Source `ptp`

The ptp data source allows you to retrieve information about the Precision Time Protocol (PTP) configuration on the appliance.

## Example usage

```terraform
data "citrixadc_ptp" "tf_ptp" {
}

output "ptp_state" {
  value = data.citrixadc_ptp.tf_ptp.state
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

The following attributes are available:

* `id` - The id of the ptp resource.
* `state` - Enables or disables Precision Time Protocol (PTP) on the appliance. If you disable PTP, make sure you enable Network Time Protocol (NTP) on the cluster.
