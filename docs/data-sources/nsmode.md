---
subcategory: "NS"
---

# citrixadc_nsmode (Data Source)

Data source for querying Citrix ADC mode settings. This data source retrieves information about the currently configured modes on the ADC appliance.

## Example Usage

```hcl
data "citrixadc_nsmode" "example" {
}

# Output mode settings
output "usip_enabled" {
  value = data.citrixadc_nsmode.example.usip
}

output "l3_mode" {
  value = data.citrixadc_nsmode.example.l3
}

output "edge_mode" {
  value = data.citrixadc_nsmode.example.edge
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nsmode datasource (always "nsmode").
* `fr` - Fast Ramp mode. When enabled, speeds up the initial data transfer.
* `l2` - Layer 2 mode. When enabled, the ADC acts as a bridge.
* `usip` - Use Source IP mode. When enabled, the ADC uses the client's source IP address for backend connections.
* `cka` - Client Keep-Alive mode. When enabled, maintains client connections even when backend connections are unavailable.
* `tcpb` - TCP Buffering mode. When enabled, buffers TCP packets to optimize performance.
* `mbf` - MAC-based forwarding mode. When enabled, forwards packets based on MAC addresses.
* `edge` - Edge configuration mode. When enabled, the ADC operates in edge mode.
* `usnip` - Use Subnet IP mode. When enabled, the ADC uses subnet IP addresses.
* `l3` - Layer 3 mode. When enabled, the ADC operates at Layer 3 of the OSI model.
* `pmtud` - Path MTU Discovery mode. When enabled, discovers the maximum transmission unit of network paths.
* `mediaclassification` - Media classification mode. When enabled, classifies media traffic for optimization.
* `sradv` - Static Route Advertisement mode. When enabled, advertises static routes.
* `dradv` - Dynamic Route Advertisement mode. When enabled, advertises dynamic routes.
* `iradv` - Intranet Route Advertisement mode. When enabled, advertises intranet routes.
* `sradv6` - Static Route Advertisement mode for IPv6. When enabled, advertises static IPv6 routes.
* `dradv6` - Dynamic Route Advertisement mode for IPv6. When enabled, advertises dynamic IPv6 routes.
* `bridgebpdus` - Bridge BPDUs mode. When enabled, bridges Bridge Protocol Data Units.
* `ulfd` - Use Link Failure Detection mode. When enabled, detects link failures.

## Notes

The nsmode resource is a singleton resource on the Citrix ADC appliance. All modes are boolean values indicating whether the particular mode is enabled or disabled on the appliance.
