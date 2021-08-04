---
subcategory: "Load Balancing"
---

# Resource: lbroute

The lbroute resource is used to bind a route VIP to the route structure.


## Example usage

```hcl
resource "citrixadc_lbroute" "tf_lbroute" {
	network = "55.0.0.0"
	netmask = "255.0.0.0"
	gatewayname = citrixadc_lbvserver.tf_lbvserver.name

	depends_on = [citrixadc_lbvserver_service_binding.tf_lbvserver_service_binding, citrixadc_nsip.nsip]
}

resource "citrixadc_nsip" "nsip" {
	ipaddress = "22.2.2.1"
	netmask   = "255.255.255.0"
}

resource "citrixadc_lbvserver_service_binding" "tf_lbvserver_service_binding" {
	name = citrixadc_lbvserver.tf_lbvserver.name
	servicename = citrixadc_service.tf_service.name
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	port = 65535
	ip = "22.2.2.2"
	servicetype = "ANY"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name = "tf_lbvserver"
	ipv46 = "0.0.0.0"
	servicetype = "ANY"
	lbmethod = "ROUNDROBIN"
	persistencetype = "NONE"
	clttimeout = 120
	port = 0
}
```


## Argument Reference

* `network` - (Required) The IP address of the network to which the route belongs.
* `netmask` - (Required) The netmask to which the route belongs.
* `gatewayname` - (Required) The name of the route.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbroute. It is the concatenation of the `network`, `netmask` and `gatewayname` attributes separated by commas.


## Import

An lbroute can be imported using its name, e.g.

```shell
terraform import citrixadc_lbroute.tf_lbroute 55.0.0.0,255.0.0.0,tf_lbvserver
```
