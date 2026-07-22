---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_grpcvalidation_binding

The appfwprofile_grpcvalidation_binding data source allows you to retrieve information about a gRPC validation relaxation binding on an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_grpcvalidation_binding" "tf_binding" {
  name                         = "tf_appfwprofile"
  grpcvalidation               = "bookstore.api.doc.AddBook"
  grpc_relax_validation_action = "log"
}

output "state" {
  value = data.citrixadc_appfwprofile_grpcvalidation_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_grpcvalidation_binding.tf_binding.resourceid
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the gRPC validation rule is bound.
* `grpcvalidation` - (Required) Exempt any gRPC method matching the given pattern from the API schema validation check. Example: `bookstore.api.doc.AddBook`.
* `grpc_relax_validation_action` - (Required) Action to be taken for traffic matching the configured relaxation rule. Possible values: [ log, none ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_grpcvalidation_binding. It is a composite identifier composed of comma-separated `key:value` pairs, in the format `grpc_relax_validation_action:<grpc_relax_validation_action>,grpcvalidation:<grpcvalidation>,name:<name>`.
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the gRPC validation rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `resourceid` - A system-generated identifier that identifies the rule.
* `alertonly` - Indicates whether an SNMP alert is sent.
```
