---
subcategory: "AppExpert"
---

# Resource: application_import

This resource is used to import an AppExpert application onto the Citrix ADC.


## Example usage

The template file (and, optionally, the deployment file) must already exist on the Citrix ADC appliance filesystem before this resource is applied. Upload or create those files out of band; this resource only triggers the import.

```hcl
resource "citrixadc_application_import" "tf_application_import" {
  appname             = "myapp"
  apptemplatefilename = "myapp_template.xml"
  deploymentfilename  = "myapp_deployment.xml"
}
```

Minimal configuration without a deployment file:

```hcl
resource "citrixadc_application_import" "tf_application_import" {
  appname             = "myapp"
  apptemplatefilename = "myapp_template.xml"
}
```


## Argument Reference

* `appname` - (Required) Name to assign to the application on the Citrix ADC. This value is also used as the resource ID and as the key for deleting the imported application. Changing this attribute forces a new resource to be created.
* `apptemplatefilename` - (Required) Name of the AppExpert application template file. This file must already exist on the Citrix ADC appliance filesystem; it is read by the import action to create the application. Changing this attribute forces a new resource to be created.
* `deploymentfilename` - (Optional) Name of the deployment file. When supplied, it must already exist on the Citrix ADC appliance filesystem. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the application resource. It has the same value as the `appname` attribute.
