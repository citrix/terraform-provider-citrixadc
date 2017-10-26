# terraform-provider-netscaler

[Terraform](https://www.terraform.io) Custom Provider for [Citrix NetScaler](https://www.citrix.com/products/netscaler-adc/)

## Description

This project is a terraform custom provider for Citrix NetScaler. It uses the [Nitro API] (https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) to create/configure LB configurations. 

## Requirement

* [hashicorp/terraform](https://github.com/hashicorp/terraform)


## Usage

### Running
1. Copy the binary (either from the [build](#building) or from the 
[releases](https://github.com/citrix/terraform-provider-netscaler/releases) page) `terraform-provider-netscaler` to the 
[third-pary plugin](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) directory.

2. Run `terraform init`.

3. Run `terraform` as usual 

```
terraform plan
terraform apply
```
4. The provider will not commit the config changes to NetScaler's persistent store. To do this, run the shell script `ns_commit.sh`:

```
export NS_URL=http://<host>:<port>/
export NS_USER=nsroot
export NS_PASSWORD=nsroot
./ns_commit.sh
```

To ensure that the config is saved on every run, we can use something like `terraform apply && ns_commit.sh`

### Provider Configuration

```
provider "netscaler" {
    username = "${var.ns_user}"
    password = "${var.ns_password}"
    endpoint = "http://10.71.136.250/"
}
```

Optionally, you can disable SSL verification if your Netscaler is using a self-signed certificate.

```
provider "netscaler" {
    username = "${var.ns_user}"
    password = "${var.ns_password}"
    endpoint = "http://10.71.136.250/"
    sslverify = false
}
```

##### Argument Reference

The following arguments are supported.

* `username` - This is the user name to access to NetScaler. Defaults to `nsroot` unless environment variable `NS_LOGIN` has been set
* `password` - This is the password to access to NetScaler. Defaults to `nsroot` unless environment variable `NS_PASSWORD` has been set
* `endpoint` - (Required) Nitro API endpoint in the form `http://<NS_IP>/` or `http://<NS_IP>:<PORT>/`. Can be specified in environment variable `NS_URL`
* `sslverify` - This controls SSL Verification. Set to `false` to disable SSL verification. Defaults to `true`. Can be specified in environment variable `NS_SSLVerify`. 

The username, password and endpoint can be provided in environment variables `NS_LOGIN`, `NS_PASSWORD` and `NS_URL`. 

### Resource Configuration

#### `netscaler_lbvserver`

```
resource "netscaler_lbvserver" "foo" {
  name = "sample_lb"
  ipv46 = "10.71.136.150"
  port = 443
  servicetype = "SSL"
  lbmethod = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  sslcertkey = "${netscaler_sslcertkey.foo.certkey}"
}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/load-balancing/lbvserver.html> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the SSL `certkey` to be bound to this `lbvserver` using the `sslcertkey` parameter

#### `netscaler_service`

```
resource "netscaler_service" "backend_1" {
  ip = "10.33.44.55"
  port = 80
  servicetype = "HTTP"
  lbvserver = "${netscaler_lbvserver.foo.name}"
  lbmonitor = "${netscaler_lbmonitor.foo.name}"
}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/basic/service.html> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the LB vserver  to be bound to this service  using the `lbvserver` parameter, and the `lbmonitor` parameter specifies the LB monitor to be bound.

#### `netscaler_servicegroup`

```
resource "netscaler_servicegroup" "backend_1" {
  servicegroupname = "backend_group_1"
  servicetype = "HTTP"
  lbvservers = ["${netscaler_lbvserver.foo.name}]"
  lbmonitor = "${netscaler_lbmonitor.foo.name}"
  servicegroupmembers = ["172.20.0.20:200:50","172.20.0.101:80:10",  "172.20.0.10:80:40"]

}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/basic/servicegroup.html> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the LB vservers  to be bound to this service using the `lbvservers` parameter. The `lbmonitor` parameter specifies the LB monitor to be bound.

#### `netscaler_csvserver`

```
resource "netscaler_csvserver" "foo" {
  name = "sample_cs"
  ipv46 = "10.71.139.151"
  servicetype = "SSL"
  port = 443
}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/content-switching/csvserver.html> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the SSL cert to be bound using the `sslcertkey` parameter


#### `netscaler_sslcertkey`

```
resource "netscaler_sslcertkey" "foo" {
  certkey = "sample_ssl_cert"
  cert = "/var/certs/server.crt"
  key = "/var/certs/server.key"
  expirymonitor = "ENABLED"
  notificationperiod = 90
}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/ssl/sslcertkey.html> for possible values for these arguments and for an exhaustive list of arguments. 


#### `netscaler_cspolicy`

```
resource "netscaler_cspolicy" "foo" {
  policyname = "sample_cspolicy"
  url = "/cart/*"
  csvserver = "${netscaler_csvserver.foo.name}"
  targetlbvserver = "${netscaler_lbvserver.foo.name}"
}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/content-switching/cspolicy.html> for possible values for these arguments and for an exhaustive list of arguments. 


#### `netscaler_lbmonitor`

```
resource "netscaler_lbmonitor" "foo" {
  monitorname = "sample_lb_monitor"
  type = "HTTP"
  interval = 350
  resptimeout = 250
}
```

##### Argument Reference
See <https://docs.citrix.com/en-us/netscaler/11-1/nitro-api/nitro-rest/api-reference/configuration/load-balancing/lbmonitor.html> for possible values for these arguments and for an exhaustive list of arguments. 

## Building
### Assumption
* You have (some) experience with Terraform, the different provisioners and providers that come out of the box,
its configuration files, tfstate files, etc.
* You are comfortable with the Go language and its code organization.

1. Install `terraform` from <https://www.terraform.io/downloads.html>
2. `go get -u github.com/hashicorp/terraform`
3. Install `godep` (<https://github.com/tools/godep>)
4. Check out this code: `git clone https://<>`
5. Build this code using `make build`



## Samples
See the `examples` directory for various LB topologies that can be driven from this terraform provider.

