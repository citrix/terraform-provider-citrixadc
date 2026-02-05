---
subcategory: "Policy"
---

# Data Source `policymap`

The policymap data source allows you to retrieve information about an existing policy map configuration.

## Example usage

```terraform
data "citrixadc_policymap" "tf_policymap" {
  mappolicyname = "my_policymap"
}

output "mappolicyname" {
  value = data.citrixadc_policymap.tf_policymap.mappolicyname
}

output "source_domain" {
  value = data.citrixadc_policymap.tf_policymap.sd
}
```

## Argument Reference

* `mappolicyname` - (Required) Name for the map policy. Must begin with a letter, number, or the underscore (_) character and must consist only of letters, numbers, and the hash (#), period (.), colon (:), space ( ), at (@), equals (=), hyphen (-), and underscore (_) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policymap resource. It has the same value as the `mappolicyname` attribute.
* `sd` - Publicly known source domain name. This is the domain name with which a client request arrives at a reverse proxy virtual server for cache redirection. If you specify a source domain, you must specify a target domain.
* `su` - Source URL. Specify all or part of the source URL, in the following format: /[[prefix] [*]] [.suffix].
* `td` - Target domain name sent to the server. The source domain name is replaced with this domain name.
* `tu` - Target URL. Specify the target URL in the following format: /[[prefix] [*]][.suffix].
