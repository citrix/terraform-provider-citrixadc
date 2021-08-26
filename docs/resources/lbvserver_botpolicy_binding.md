---
subcategory: "Load Balancing"
---

# Resource:  lbvserver\_botpolicy_\_binding

The lbvserver_botpolicy_binding resource is used to bind load balancing virtual servers with bot policies.


## Example usage

```hcl
resource citrixadc_lbvserver_botpolicy_binding demo_lbvserver_botpolicy_binding {
  name                   = citrixadc_lbvserver.demo_lb.name
  policyname             = citrixadc_botpolicy.demo_botpolicy.name
  labeltype              = "reqvserver" # Possible values = reqvserver, resvserver, policylabel
  labelname              = citrixadc_lbvserver.demo_lb.name
  priority               = 100
  bindpoint              = "REQUEST" # Possible values = REQUEST, RESPONSE
  gotopriorityexpression = "END"
  invoke                 = true         # boolean
}

resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  servicetype = "HTTP"
}

resource "citrixadc_botpolicy" "demo_botpolicy" {
  name        = "demo_botpolicy"
  profilename = citrixadc_botprofile.tf_botprofile.name
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}

resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "tf_botprofile comment"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	devicefingerprintaction = ["LOG", "RESET"]
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	trapaction = ["LOG", "RESET"]
	bot_enable_tps = "ON"
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the citrixadc_lbvserver_botpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A citrixadc_lbvserver_botpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_citrixadc_lbvserver_botpolicy_binding.tf_citrixadc_lbvserver_botpolicy_binding tf_citrixadc_lbvserver_botpolicy_binding
```
