---
subcategory: "Utility"
---

# Resource: callhome

The callhome resource configures the Citrix ADC Call Home feature, which allows the appliance to automatically report failures and periodic heartbeats to Citrix technical support. Configure it to set the administrator contact email, the reporting mode, and an optional proxy server to use when the ADC reaches out over the internet.

Call Home is a singleton feature: there is exactly one Call Home configuration per Citrix ADC. Applying this resource sets the global Call Home configuration. There is no create or delete operation on the ADC; running `terraform destroy` only removes the resource from Terraform state and leaves the configuration in place on the appliance.


## Example usage

### Basic usage

```hcl
resource "citrixadc_callhome" "tf_callhome" {
  mode             = "Default"
  emailaddress     = "admin@example.com"
  hbcustominterval = 7
}
```

### Using a proxy server by IP address

```hcl
resource "citrixadc_callhome" "tf_callhome" {
  mode         = "Default"
  emailaddress = "admin@example.com"
  proxymode    = "YES"
  ipaddress    = "192.0.2.10"
  port         = 3128
}
```

### Using a proxy server by service name

```hcl
resource "citrixadc_callhome" "tf_callhome" {
  mode             = "Default"
  emailaddress     = "admin@example.com"
  proxymode        = "YES"
  proxyauthservice = "proxy_svc"
  port             = 3128
}
```


## Argument Reference

* `mode` - (Optional) CallHome mode of operation. Defaults to `"Default"`. Possible values: [ Default, CSP, Connector ]
* `emailaddress` - (Optional) Email address of the contact administrator.
* `proxymode` - (Optional) Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server. Defaults to `"NO"`. Possible values: [ YES, NO ]
* `ipaddress` - (Optional) IP address of the proxy server. Mutually exclusive with `proxyauthservice` - specify only one of the two to identify the proxy server.
* `proxyauthservice` - (Optional) Name of the service that represents the proxy server. Mutually exclusive with `ipaddress` - specify only one of the two to identify the proxy server.
* `port` - (Optional) HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.
* `hbcustominterval` - (Optional) Interval (in days) between CallHome heartbeats. Defaults to `7`. Minimum value = 1 Maximum value = 30


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the callhome resource. Because callhome is a singleton, this is always the static string `"callhome"`.
* `nodeid` - Unique number that identifies the cluster node. This is a read-only, GET-only cluster-node filter that is populated by the ADC and cannot be set.
