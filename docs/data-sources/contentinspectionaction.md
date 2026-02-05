---
subcategory: "Content Inspection"
---

# Data Source `contentinspectionaction`

The contentinspectionaction data source allows you to retrieve information about a content inspection action.


## Example usage

```terraform
data "citrixadc_contentinspectionaction" "tf_contentinspectionaction" {
  name = "my_ci_action"
}

output "type" {
  value = data.citrixadc_contentinspectionaction.tf_contentinspectionaction.type
}

output "serverip" {
  value = data.citrixadc_contentinspectionaction.tf_contentinspectionaction.serverip
}
```


## Argument Reference

* `name` - (Required) Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `icapprofilename` - Name of the ICAP profile to be attached to the contentInspection action.
* `ifserverdown` - Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are:
  * RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
  * DROP - Drop the request without sending a response to the user.
  * CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.
* `serverip` - IP address of remoteService
* `servername` - Name of the LB vserver or service
* `serverport` - Port of remoteService
* `type` - Type of operation this action is going to perform. following actions are available to configure:
  * ICAP - forward the incoming request or response to an ICAP server for modification.
  * INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention.
  * MIRROR - Forwards cloned packets for Intrusion Detection.
  * NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device.
  * NSTRACE - capture current and further incoming packets on this transaction.

## Attribute Reference

* `id` - The id of the contentinspectionaction. It has the same value as the `name` attribute.


## Import

A contentinspectionaction can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionaction.tf_contentinspectionaction my_ci_action
```
