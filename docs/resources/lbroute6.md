---
subcategory: "Load Balancing"
---

# Resource: lbroute6

The lbroute6 resource is used to configure lbroute6.


## Example usage

```hcl
resource "citrixadc_nsip6" "tf1_nsip6" {
    ipv6address = "22::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_nsip6" "tf2_nsip6" {
    ipv6address = "33::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_nsip6" "tf3_nsip6" {
    ipv6address = "44::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_lbvserver" "llb6" {
    name = "llb6"
    servicetype = "ANY"
    persistencetype = "NONE"
    lbmethod = "ROUNDROBIN"
}

resource "citrixadc_service" "r4" {
    name = "r4"
    ip = "22::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf1_nsip6]
}

resource "citrixadc_service" "r5" {
    name = "r5"
    ip = "33::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf2_nsip6]

}

resource "citrixadc_service" "r6" {
    name = "r6"
    ip = "44::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf3_nsip6]

}

resource "citrixadc_lbvserver_service_binding" "tf_binding4" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r4.name
  weight = 10
}

resource "citrixadc_lbvserver_service_binding" "tf_binding5" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r5.name
  weight = 10
}

resource "citrixadc_lbvserver_service_binding" "tf_binding6" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r6.name
  weight = 10
}

resource "citrixadc_lbroute6" "demo_route6" {
    network = "66::/64"
    gatewayname = citrixadc_lbvserver.llb6.name
}
```


## Argument Reference

* `network` - (Required) The destination network.
* `gatewayname` - (Required) The name of the route.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbroute6. It has the same value as the `gatewayname` attribute.


## Import

A lbroute6 can be imported using its name, e.g.

```shell
terraform import citrixadc_lbroute6.tf_lbroute6 tf_lbroute6
```
