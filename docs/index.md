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


# Proxy calls through ADM
# Login credentials refer to ADM
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
```

## Argument Reference

The following arguments are supported.

* `endpoint` - (Required) Defines the NITRO API endpoint prefix. Can use either `http` or `https` protocol.
* `username` - (Required) Defines the username that will be used by the NITRO API for authentication. Can be sourced from the `NS_LOGIN` environment variable. Defaults to `nsroot`.
* `password` - (Required) Defines the password that will be used by the NITRO API for authentication. Can be sourced from the `NS_PASSWORD` environment variable. Defaults to `nsroot`.
* `insecure_skip_verify` - (Optional) Boolean variable that defines if an error should be thrown if the target ADC's TLS certificate is not trusted. When `true` the error will be ignored. When `false` such an error will cause the failure of any provider operation. Defaults to `false`.
* `proxied_ns` - (Optional) When defined use ADM as a proxy for the NITRO API calls. All credentials refer to the ADM. The value of this attribute is the target ADC's ip address. Can be sourced from the `_MPS_API_PROXY_MANAGED_INSTANCE_IP` environment variable.
* `do_login` - (Optional) When set to true the NITRO client will perform the login operation and acquire a session token which will be used for all subsequent operations. This is required when targeting a non default admin partition.
* `partition` - (Optional) Partition to target. All resources utilizing this provider instance will reside on the target admin partition.

!> Avoid hard coding credentials in terraform configuration files. It presents a security risk especially if they are committed and published in version control systems.
