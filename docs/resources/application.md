---
subcategory: "AppExpert"
---

# Resource: application

Imports an AppExpert application onto the Citrix ADC from an application template file that already resides on the appliance filesystem. Use this resource to deploy a packaged AppExpert application (its policies, expressions, and configuration) so that it becomes an active application on the appliance without recreating each component by hand.

This is an action-only resource: it maps to the NITRO `application` object's `Import` action. The Citrix ADC provides no GET endpoint for an imported application, so the provider cannot read the object back to detect drift, and there is no update operation. Any change to a configured attribute forces the application to be deleted and re-imported.


## Example usage

The template file (and, optionally, the deployment file) must already exist on the Citrix ADC appliance filesystem before this resource is applied. Upload or create those files out of band; this resource only triggers the import.

```hcl
resource "citrixadc_application" "tf_application" {
  appname             = "myapp"
  apptemplatefilename = "myapp_template.xml"
  deploymentfilename  = "myapp_deployment.xml"
}
```

Minimal configuration without a deployment file:

```hcl
resource "citrixadc_application" "tf_application" {
  appname             = "myapp"
  apptemplatefilename = "myapp_template.xml"
}
```


## Argument Reference

* `appname` - (Required) Name to assign to the application on the Citrix ADC. Also used as the resource ID and as the key for deleting the imported application. Changing this attribute forces a new resource to be created.
* `apptemplatefilename` - (Required) Name of the AppExpert application template file. This file must already exist on the Citrix ADC appliance filesystem; it is read by the import action to create the application. Changing this attribute forces a new resource to be created.
* `deploymentfilename` - (Optional) Name of the deployment file. When supplied, it must already exist on the Citrix ADC appliance filesystem. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the application. It has the same value as the `appname` attribute.
