---
subcategory: "AppExpert"
---

# Resource: application_export

Exports a configured AppExpert application from the Citrix ADC into an application template file (and, optionally, a deployment file) on the appliance filesystem. This is the inverse of the import action: use it to package a deployed application so it can be backed up, version-controlled, or moved to another appliance.

This is an action-only resource: it maps to the NITRO `application` object's `export` action. Each apply invokes the export and writes the template file on the appliance. The Citrix ADC provides no GET endpoint for the export state, so the provider cannot read anything back, and there is no data source. Read, update, and delete are no-ops (delete simply drops the resource from Terraform state — there is no inverse "un-export" API). Changing any configured attribute re-triggers the export.


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

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `application_export`. It does not correspond to any object on the Citrix ADC.
