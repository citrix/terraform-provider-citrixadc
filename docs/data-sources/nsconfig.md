---
subcategory: "NS"
---

# Data Source: nsconfig

The nsconfig data source retrieves the system-wide configuration of the Citrix ADC, including the management IP (NSIP) address and netmask, the management VLAN (NSVLAN) and its bound interfaces, the secure management traffic settings, and a rich set of read-only status fields. Because nsconfig is a singleton, no lookup argument is required.


## Example usage

```hcl
data "citrixadc_nsconfig" "current" {
}

output "nsip_address" {
  value = data.citrixadc_nsconfig.current.ipaddress
}

output "nsip_netmask" {
  value = data.citrixadc_nsconfig.current.netmask
}

output "secure_management_traffic" {
  value = data.citrixadc_nsconfig.current.securemanagementtraffic
}
```


## Argument Reference

This data source takes no lookup arguments; it always reads the single nsconfig object present on the appliance.


## Attribute Reference

In addition to `id`, the following attributes are available:

* `id` - The ID of the nsconfig data source. Constant value `nsconfig-config`.
* `ipaddress` - IP address of the Citrix ADC (NSIP address).
* `netmask` - Netmask corresponding to the NSIP address.
* `nsvlan` - VLAN (NSVLAN) for the subnet on which the NSIP address resides.
* `ifnum` - List of interfaces bound to the NSVLAN.
* `tagged` - Whether the NSVLAN interfaces are added as 802.1q tagged interfaces.
* `securemanagementtraffic` - Whether secure management traffic handling is enabled.
* `securemanagementtd` - Management traffic domain identifier used for secure management traffic.
* `Async` - Whether the operation runs in async mode and returns a job ID.
* `all` - Whether saveconfig is performed for all partitions.
* `changedpassword` - Lists all passwords changed that would not work when downgraded to older releases.
* `cip` - Controls insertion of the actual client IP address into the HTTP header request passed to the attached servers.
* `cipheader` - Text used as the client IP header.
* `config` - Configuration file used to find weak passwords.
* `config1` - Location of the configurations.
* `config2` - Location of the configurations.
* `configfile` - Full path of the config file to be converted to NITRO.
* `cookieversion` - Version of the cookie inserted by the system.
* `crportrange` - Port range for cache redirection services.
* `exclusivequotamaxclient` - Percentage of maxClient given to PEs.
* `exclusivequotaspillover` - Percentage of spillover threshold given to PEs.
* `force` - Whether configurations are cleared without prompting for confirmation.
* `ftpportrange` - Port range configured for FTP services.
* `grantquotamaxclient` - Percentage of shared quota granted at a time for maxClient.
* `grantquotaspillover` - Percentage of shared quota granted at a time for spillover.
* `httpport` - List of HTTP ports on the web server used for connection off-load.
* `ignoredevicespecific` - Whether device-specific differences are suppressed.
* `level` - Types of configurations to be cleared (basic, extended, full).
* `maxconn` - Maximum number of connections made from the system to the attached web servers.
* `maxreq` - Maximum number of requests passed on a single connection to an attached server.
* `outtype` - Format used to display the difference in configurations.
* `pmtumin` - Minimum Path MTU.
* `pmtutimeout` - Path MTU timeout value in minutes.
* `rbaconfig` - Whether RBA configurations and TACACS policies bound to system global are preserved during a basic clear.
* `responsefile` - Full path of the file used to store the NITRO graph.
* `securecookie` - Whether the secure flag is set for the persistence cookie.
* `template` - File that contains the commands to be compared.
* `timezone` - Name of the timezone.
* `weakpassword` - Lists all weak passwords not adhering to strong password requirements.
