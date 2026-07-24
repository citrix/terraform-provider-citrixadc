---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_restvalidation_binding

This resource is used to manage REST validation bindings of an application firewall profile.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}

resource "citrixadc_appfwprofile_restvalidation_binding" "tf_binding" {
  name                   = citrixadc_appfwprofile.tf_appfwprofile.name
  restvalidation         = "GET:/v1/bookstore/viewbooks"
  rest_validation_action = "log"
  state                  = "ENABLED"
  comment                = "Exempt the public catalog endpoint from API schema validation"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `restvalidation` - (Required) Exempt REST endpoints or any other URLs matching the given pattern from the API schema validation check. Example: `GET:/v1/bookstore/viewbooks`.
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `rest_validation_action` - (Optional) Action to be taken for traffic matching the configured relaxation rule. Defaults to `"none"`. Possible values: [ log, none ]
* `state` - (Optional) Enables or disables the binding. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_restvalidation_binding. It is the concatenation of the `name`, `rest_validation_action` and `restvalidation` attributes separated by a comma.
* `alertonly` - (Read-only) Send SNMP alert? This value is derived by the server.
* `isautodeployed` - (Read-only) Indicates whether the rule was auto deployed by a dynamic profile. This value is derived by the server. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `resourceid` - (Read-only) A server-assigned identifier that identifies the rule. It is populated from the NITRO GET response and is useful when importing an existing binding into Terraform state.


## Import

A appfwprofile_restvalidation_binding can be imported using its id. The id is a comma-separated list of `key:value` pairs in the order `name`, `rest_validation_action`, `restvalidation`, e.g.

```shell
terraform import citrixadc_appfwprofile_restvalidation_binding.tf_binding name:tf_appfwprofile,rest_validation_action:log,restvalidation:GET:/v1/bookstore/viewbooks
```
