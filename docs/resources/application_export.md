---
subcategory: "AppExpert"
---

# Resource: application_export

This resource is used to export an AppExpert application to a template file on the Citrix ADC.


## Example usage

The application named by `appname` must already exist on the Citrix ADC before this resource is applied. The export writes the template file to the appliance filesystem.

```hcl
resource "citrixadc_application_export" "tf_application_export" {
  appname             = "myapp"
  apptemplatefilename = "myapp_template.xml"
  deploymentfilename  = "myapp_deployment.xml"
}
```

Minimal configuration exporting to a template file only:

```hcl
resource "citrixadc_application_export" "tf_application_export" {
  appname             = "myapp"
  apptemplatefilename = "myapp_template.xml"
}
```


## Argument Reference

Changing any of these arguments forces a new export action to be performed.

* `appname` - (Required) Name of the application on the Citrix ADC to export. Changing this attribute re-triggers the export.
* `apptemplatefilename` - (Optional) Name of the AppExpert application template file to write on the appliance filesystem. Changing this attribute re-triggers the export.
* `deploymentfilename` - (Optional) Name of the deployment file to write on the appliance filesystem. Changing this attribute re-triggers the export.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the application_export resource. It is set to `application_export`.
