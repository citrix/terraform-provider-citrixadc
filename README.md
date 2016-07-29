# terraform-provider-netscaler

[Terraform](http://www.terraform.io) Custom Provider for [Citrix NetScaler](https://www.citrix.com/products/netscaler-adc/)

## Description

This project is a terraform custom provider for Citrix NetScaler. It uses the [Nitro API] (http://docs.citrix.com/en-us/netscaler/11/nitro-api.html) to create/configure simple LB configurations (`lb vserver` which is bound to a list of `service`)

## Requirement

* [hashicorp/terraform](https://github.com/hashicorp/terraform)


## Usage

### Provider Configuration

#### `netscaler`

```
provider "netscaler" {
    username = "${var.ns_user}"
    password = "${var.ns_password}"
    endpoint = "http://10.71.136.250/"
}
```

##### Argument Reference

The following arguments are supported.

* `username` - (Required) This is the user name to access to NetScaler.
* `password` - (Required) This is the password to access to NetScaler.
* `endpoint` - (Required) Nitro API endpoint in the form `http://<NS_IP>/`"

### Resource Configuration

#### `netscaler_lb`

```
resource "netscaler_lb" "foo" {
  name = "sample_lb"
  vip = "10.71.136.150"
  port = 443
  service_type = "SSL"
  lb_method = "ROUNDROBIN"
  persistence_type = "COOKIEINSERT"
}
```

##### Argument Reference
See <http://docs.citrix.com/en-us/netscaler/11-1/load-balancing/load-balancing-setup.html> for possible values for these arguments.

The following arguments are supported.

* `name` - (Optional) name of the lb vserver in NetScaler
* `vip` - (Required) The VIP for the lb vserver
* `port` - (Required) Port e.g., 80.
* `service_type` - (Optional) Usually `HTTP` or `SSL`. NetScaler will default to `HTTP`
* `lb_method` - (Optional) Usually `LEASTCONNECTION` or `ROUNDROBIN`, `LEASTRESPONSETIME`. See NetScaler docs for more options
* `persistence_type` - (Optional) Usually `COOKIEINSERT`. See NetScaler docs for more options

#### `netscaler_svc`

```
resource "netscaler_svc" "backend_1" {
  lb = "${netscaler_lb.foo.name}"
  ip = "10.33.44.55"
  port = 80
}
```
##### Argument Reference
See <http://docs.citrix.com/en-us/netscaler/11-1/load-balancing/load-balancing-setup.html> for possible values for these arguments.

Each `netscaler_svc` models a NetScaler `service` object. The NetScaler docs have more values for service type etc.

* `lb` - (Required) The name of the `lb vserver` to bind to. Usually you refer to a previously declared `lb vserver` in the same config, i.e., `${netscaler_lb.foolb.name}`
* `ip` - (Required) IP address. 
* `port` - (Required) Port
* `service_type` - (Optional) This has to be compatible with the `service_type` declared in the `netscaler_lb`. Defaults to `HTTP`


##### For example

Example 1:

```
resource "netscaler_lb" foolb" {
  name = "sample_lb"
  vip = "10.71.136.151"
  port = 443
  service_type = "SSL"
}

resource "netscaler_svc" "backend_1" {
  lb = "${netscaler_lb.foolb.name}"
  ip = "10.33.44.55"
  port = 80
}

resource "netscaler_svc" "backend_2" {
  lb = "${netscaler_lb.foolb.name}"
  ip = "10.33.44.54"
  port = 80
}
```



## Running
### Assumption
* You have (some) experience with Terraform, the different provisioners and providers that come out of the box,
its configuration files, tfstate files, etc.
* You are comfortable with the Go language and its code organization.

1. Install `terraform` from <https://www.terraform.io/downloads.html>
2. `go get -u github.com/hashicorp/terraform`
3. Check out this code: `git clone http://<>`
4. Build this code using `make`
5. Copy the resulting binary `terraform-provider-netscaler` to the same directory as the `terraform` install (use `which terraform`)
6. Run `terraform` as usual 



## Author

[chiradeep](https://github.com/chiradeep)