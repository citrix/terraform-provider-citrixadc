# Citrix ADC Provider

The Citrix ADC provider is used to configure target ADC instances
using the NITRO API.

## Example Usage

```hcl
# Simplest and least secure configuration
# Use http and default values for username and password
provider "citrixadc" {
  endpoint = "http://10.0.0.1"
}


# Use https and non default password
provider "citrixadc" {
  endpoint = "https://10.0.0.1"
  username = "nsroot"
  password = "secret"

  # Do not error due to non signed ADC TLS certificate
  # Can skip this if ADC TLS certificate is trusted
  insecure_skip_verify = true
}


# Proxy calls through NetScaler Console
# Login credentials refer to NetScaler Console
# Target ADC is referred by its ip address
provider "citrixadc" {
  endpoint   = "https://10.22.0.1"
  username   = "nsroot"
  password   = "admpassword"
  proxied_ns = "10.0.0.1"
}

# Target non default partition
provider "citrixadc" {
  endpoint   = "https://10.22.0.1"
  username   = "nsroot"
  password   = "admpassword"
  do_login   = true
  partition  = "par1"
}

# Proxy calls through NetScaler Console Service
# When using NetScaler Console Cloud:
# - endpoint: Your NetScaler Console Cloud URL (e.g., https://alps.adm.cloudburrito.com/)
# - username: API Client ID 
# - password: API Client Secret 
# - proxied_ns: IP of the target NetScaler instance (must be managed by NetScaler Console)
# - is_cloud: Set to true for NetScaler Console Cloud
# - do_login: Set to true to establish session
provider "citrixadc" {
  endpoint   = "https://<<ADM_CLOUD_URL>>/"
  username   = "<<USERNAME>>"
  password   = "<<PASSWORD>>"
  proxied_ns = "<<NS_IP>>"
  is_cloud   = true
  do_login   = true
}

```

## Argument Reference

The following arguments are supported.

* `endpoint` - (Required) Defines the NITRO API endpoint prefix. Can use either `http` or `https` protocol.
* `username` - (Required) Defines the username that will be used by the NITRO API for authentication. Can be sourced from the `NS_LOGIN` environment variable. Defaults to `nsroot`.
* `password` - (Required) Defines the password that will be used by the NITRO API for authentication. Can be sourced from the `NS_PASSWORD` environment variable. Defaults to `nsroot`.
* `insecure_skip_verify` - (Optional) Boolean variable that defines if an error should be thrown if the target ADC's TLS certificate is not trusted. When `true` the error will be ignored. When `false` such an error will cause the failure of any provider operation. Defaults to `false`.
* `proxied_ns` - (Optional) When defined use NetScaler Console as a proxy for the NITRO API calls. All credentials refer to the NetScaler Console. The value of this attribute is the target ADC's ip address. Can be sourced from the `_MPS_API_PROXY_MANAGED_INSTANCE_IP` environment variable.
* `do_login` - (Optional) When set to true the NITRO client will perform the login operation and acquire a session token which will be used for all subsequent operations. This is required when targeting a non default admin partition.
* `is_cloud` - (Optional) Boolean variable that defines whether NetScaler Console Service is used for proxied calls. When `true`, `username`, `password` and `endpoint` must refer to the Console Service. Defaults to `false`.
* `partition` - (Optional) Partition to target. All resources utilizing this provider instance will reside on the target admin partition.

!> Avoid hard coding credentials in terraform configuration files. It presents a security risk especially if they are committed and published in version control systems.
