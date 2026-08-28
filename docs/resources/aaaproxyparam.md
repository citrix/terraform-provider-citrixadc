---
subcategory: "AAA"
---

# Resource: aaaproxyparam

The aaaproxyparam resource is used to configure the AAA proxy parameters. This is a
singleton (global) configuration resource.


## Example usage

```hcl
resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
  proxy              = "10.1.1.1:8080"
  proxyauthorization = "basic"
  proxyusername      = "proxyuser"
  proxypassword      = "proxypass123"
}
```

### Write-only (ephemeral) secret

Use `proxypassword_wo` to keep the secret out of Terraform state. Bump
`proxypassword_wo_version` whenever the secret value changes so the provider
re-applies it:

```hcl
variable "proxy_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
  proxy                    = "10.1.1.1:8080"
  proxyauthorization       = "basic"
  proxyusername            = "proxyuser"
  proxypassword_wo         = var.proxy_password
  proxypassword_wo_version = 1
}
```


## Argument Reference

* `proxy` - (Optional) IP address and Port of the proxy server to be used for HTTP access for this request. This can be configured in `ipaddress:port` format (for example `a.b.c.d:e`) or as a URL (for example `http://a.b.c.d` without a port or `http://a.b.c.d:8080` with a port).
* `proxyauthorization` - (Optional) This indicates whether Proxy-Authorization header will be sent or not. Possible values: [ disabled, basic ]
* `proxyusername` - (Optional) Username that will be sent as part of Basic Proxy-Authorization header. Maximum length =  256
* `proxypassword` - (Optional) Password that will be sent as part of Basic Proxy-Authorization header. Maximum length =  256. This is stored in Terraform state; prefer `proxypassword_wo` for ephemeral (write-only) handling.
* `proxypassword_wo` - (Optional, Write-only) Write-only (ephemeral) equivalent of `proxypassword`. The value is used only during the apply and is never persisted to Terraform state. Pair it with `proxypassword_wo_version` to trigger updates when the secret rotates. (Requires Terraform 1.11+.)
* `proxypassword_wo_version` - (Optional) Version tracker for `proxypassword_wo`. Increment this value to signal that the write-only secret changed so the provider pushes the new value to the appliance. Default value: 1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaproxyparam. Because this is a singleton resource, it has a fixed identifier.
