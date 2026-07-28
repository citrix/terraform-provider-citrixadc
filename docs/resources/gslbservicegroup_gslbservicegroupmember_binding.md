---
subcategory: "GSLB"
---

# Resource: gslbservicegroup_gslbservicegroupmember_binding

This resource is used to bind a member to a GSLB service group.


## Example usage

```hcl
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "gslbsvcgrp1"
  servicetype      = "HTTP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}

# Bind a member by IP address and port (the IP path).
resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "member_by_ip" {
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  ip               = "192.0.2.10"
  port             = 80
  weight           = 20
}
```

A member can be identified in one of two mutually-exclusive ways. Set exactly one of `ip` or `servername` (not both):

```hcl
# Bind a member by referencing a pre-configured server (the server-name path).
resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "member_by_server" {
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  servername       = "web-server-1"
  port             = 80
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the GSLB service group.
* `port` - (Required) Server port number.
* `ip` - (Optional) IP Address of the member. Set exactly one of `ip` or `servername`; the two are mutually exclusive and one of them is required.
* `servername` - (Optional) Name of a pre-configured server to which to bind the service group. Set exactly one of `servername` or `ip`; the two are mutually exclusive and one of them is required.
* `weight` - (Optional, Computed) Weight to assign to the member. Specifies the capacity of the server relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service. If not set, the ADC defaults this to `1`.
* `state` - (Optional, Computed) Initial state of the member. If not set, the ADC defaults this to `ENABLED`. Possible values: [ ENABLED, DISABLED ]
* `hashid` - (Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `publicip` - (Optional) The public IP address that a NAT device translates to the GSLB service's private IP address.
* `publicport` - (Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service.
* `siteprefix` - (Optional) The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string `NONE` is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `order` - (Optional) Order number to be assigned to the GSLB service group member.

~> **NOTE:** Every argument forces replacement when changed. There is no in-place update for this binding; the NITRO API exposes only bind (add) and unbind (delete) operations.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding. It is a composite, comma-separated string of `key:value` pairs (values URL-encoded) in the form `ip:<ip>,port:<port>,servername:<servername>,servicegroupname:<servicegroupname>`. The unused member field (whichever of `ip` or `servername` you did not set) is left empty.


## Import

A gslbservicegroup_gslbservicegroupmember_binding can be imported using its composite ID. The ID is the same comma-separated `key:value` form described above.

```shell
terraform import citrixadc_gslbservicegroup_gslbservicegroupmember_binding.member_by_ip "ip:192.0.2.10,port:80,servername:,servicegroupname:gslbsvcgrp1"
```

For a member bound by server name, leave the `ip` value empty instead:

```shell
terraform import citrixadc_gslbservicegroup_gslbservicegroupmember_binding.member_by_server "ip:,port:80,servername:web-server-1,servicegroupname:gslbsvcgrp1"
```
