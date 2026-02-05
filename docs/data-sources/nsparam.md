---
subcategory: "NS"
---

# citrixadc_nsparam (Data Source)

Data source for querying Citrix ADC global system parameters. This data source retrieves information about the system-wide configuration settings on the ADC appliance.

## Example Usage

```hcl
data "citrixadc_nsparam" "example" {
}

# Output system parameters
output "timezone" {
  value = data.citrixadc_nsparam.example.timezone
}

output "ip_ttl" {
  value = data.citrixadc_nsparam.example.ipttl
}

output "use_proxy_port" {
  value = data.citrixadc_nsparam.example.useproxyport
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nsparam datasource (always "nsparam-config").
* `advancedanalyticsstats` - Enable or disable advanced analytics statistics collection.
* `aftpallowrandomsourceport` - Allow the FTP server to use a random source port for active FTP data connections.
* `cip` - Enable or disable the insertion of the actual client IP address into the HTTP header request.
* `cipheader` - Text that will be used as the client IP address header.
* `cookieversion` - Version of the cookie inserted by the system.
* `crportrange` - Port range for cache redirection services.
* `exclusivequotamaxclient` - Percentage of maxClient threshold to be divided equally among packet engines.
* `exclusivequotaspillover` - Percentage of spillover threshold to be divided equally among packet engines.
* `ftpportrange` - Minimum and maximum port (port range) that FTP services are allowed to use.
* `grantquotamaxclient` - Percentage of shared pool value granted to packet engine once it exhausts the local exclusive quota for maxClient.
* `grantquotaspillover` - Percentage of shared pool value granted to packet engine once it exhausts the local exclusive quota for spillover.
* `httpport` - HTTP ports on the web server for connection off-load.
* `icaports` - ICA ports on the web server for connection off-load.
* `internaluserlogin` - Enable or disable internal user login to the appliance.
* `ipttl` - IP Time to Live (TTL) and Hop Limit value for all outgoing packets from Citrix ADC.
* `maxconn` - Maximum number of connections from the appliance to attached servers.
* `maxreq` - Maximum number of requests on a connection between the appliance and a server.
* `mgmthttpport` - Management HTTP port configuration.
* `mgmthttpsport` - Management HTTPS port configuration.
* `pmtumin` - Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message.
* `pmtutimeout` - Interval, in minutes, for flushing the PMTU entries.
* `proxyprotocol` - Enable or disable proxy protocol header (v1 or v2) for client info insertion.
* `securecookie` - Enable or disable secure flag for persistence cookie.
* `secureicaports` - Secure ICA ports on the web server for connection off-load.
* `servicepathingressvlan` - VLAN on which the subscriber traffic arrives on the appliance.
* `tcpcip` - Enable or disable the insertion of the client TCP/IP header in TCP payload.
* `timezone` - Time zone for the Citrix ADC appliance.
* `useproxyport` - Enable or disable use_proxy_port setting.

## Notes

The nsparam resource is a singleton resource on the Citrix ADC appliance that contains global system parameters. These parameters affect the overall behavior and configuration of the ADC.
