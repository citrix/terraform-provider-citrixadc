---
subcategory: "Policy"
---

# Resource: policymap

The policymap resource is used to create a policy to map a publicly known domain name to a target domain name.


## Example usage

```hcl
resource "citrixadc_policymap" "tf_policymap" {
	mappolicyname = "tf_policymap"
	sd = "www.citrix.com"
	td = "www.google.com"
	su = "/www.citrix.com"
	tu = "/www.google.com"
}
```


## Argument Reference

* `mappolicyname` - (Required) Name for the map policy. Must begin with a letter, number, or the underscore (_) character and must consist only of letters, numbers, and the hash (#), period (.), colon (:), space ( ), at (@), equals (=), hyphen (-), and underscore (_) characters. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my map" or 'my map').
* `sd` - (Required) Publicly known source domain name. This is the domain name with which a client request arrives at a reverse proxy virtual server for cache redirection. If you specify a source domain, you must specify a target domain.
* `su` - (Optional) Source URL. Specify all or part of the source URL, in the following format: /[[prefix] [*]] [.suffix].
* `td` - (Required) Target domain name sent to the server. The source domain name is replaced with this domain name.
* `tu` - (Optional) Target URL. Specify the target URL in the following format: /[[prefix] [*]][.suffix].


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policymap. It has the same value as the `mappolicyname` attribute.


## Import

A policymap can be imported using its name, e.g.

```shell
terraform import citrixadc_policymap.tf_policymap tf_policymap
```
