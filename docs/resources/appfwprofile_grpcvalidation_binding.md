---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_grpcvalidation_binding

This resource is used to bind a gRPC validation relaxation rule to an application firewall profile.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
}

resource "citrixadc_appfwprofile_grpcvalidation_binding" "tf_binding" {
  name                         = citrixadc_appfwprofile.tf_appfwprofile.name
  grpcvalidation               = "bookstore.api.doc.AddBook"
  grpc_relax_validation_action = "log"
  state                        = "ENABLED"
  comment                      = "Relax schema validation for the AddBook gRPC method"
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the gRPC validation rule is bound. Changing this value forces a new resource to be created.
* `grpcvalidation` - (Required) Exempt any gRPC method matching the given pattern from the API schema validation check. Example: `bookstore.api.doc.AddBook`. Changing this value forces a new resource to be created.
* `grpc_relax_validation_action` - (Optional) Action to be taken for traffic matching the configured relaxation rule. Defaults to `none`. Changing this value forces a new resource to be created. Possible values: [ log, none ]
* `comment` - (Optional) Any comments about the purpose of the profile, or other useful information about the profile. Changing this value forces a new resource to be created.
* `state` - (Optional) Enable or disable the gRPC validation rule. Changing this value forces a new resource to be created. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - (Optional) Indicates whether the rule was auto-deployed by a dynamic profile. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_grpcvalidation_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `grpc_relax_validation_action:<grpc_relax_validation_action>,grpcvalidation:<grpcvalidation>,name:<name>`.
* `resourceid` - (Read-only) A system-generated identifier that identifies the rule.
* `alertonly` - (Read-only) Indicates whether an SNMP alert is sent.


## Import

A appfwprofile_grpcvalidation_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_appfwprofile_grpcvalidation_binding.tf_binding "grpc_relax_validation_action:log,grpcvalidation:bookstore.api.doc.AddBook,name:tf_appfwprofile"
```
