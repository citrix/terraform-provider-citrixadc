---
subcategory: "IPSECALG"
---

# Data Source: ipsecalgsession

The ipsecalgsession data source retrieves information about active IPSec ALG sessions on the Citrix ADC. It performs a `show ipsecalgsession` (get-all) and returns the first session matching the optional source, NAT, or destination IP filters you supply.

-> **Note:** The IPSec ALG feature must be in use for sessions to exist. An empty session table (no active IPSec ALG sessions, or no session matching the supplied filters) is a valid result and does not produce an error.

## Example usage

Look up a session by original source IP address:

```terraform
data "citrixadc_ipsecalgsession" "example" {
  sourceip = "192.168.10.5"
}

output "ipsecalg_natip" {
  value = data.citrixadc_ipsecalgsession.example.natip
}
```

## Argument Reference

All arguments are optional filters. The data source returns the first session that matches every supplied filter.

* `sourceip` - (Optional) Original source IP address to filter on.
* `natip` - (Optional) Natted source IP address to filter on.
* `destip` - (Optional) Destination IP address to filter on.
* `sourceip_alg` - (Optional) Original source IP address (CLI-name twin) to filter on.
* `natip_alg` - (Optional) Natted source IP address (CLI-name twin) to filter on.
* `destip_alg` - (Optional) Destination IP address (CLI-name twin) to filter on.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the data source read (`ipsecalgsession`).
* `sourceip` - Original source IP address of the matched session.
* `natip` - Natted source IP address of the matched session.
* `destip` - Destination IP address of the matched session.
* `sourceip_alg` - Original source IP address of the matched session (CLI-name twin).
* `natip_alg` - Natted source IP address of the matched session (CLI-name twin).
* `destip_alg` - Destination IP address of the matched session (CLI-name twin).

Additional read-only fields returned by the Citrix ADC when present in the session, such as `spiin` and `spiout` (the inbound and outbound Security Parameter Index values), are included in the underlying NITRO response.
