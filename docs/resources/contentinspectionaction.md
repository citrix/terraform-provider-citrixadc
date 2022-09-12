---
subcategory: "CI"
---

# Resource: contentinspectionaction

The contentinspectionaction resource is used to create contentinspectionaction.


## Example usage

```hcl
resource "citrixadc_contentinspectionaction" "tf_contentinspectionaction" {
  name            = "my_ci_action"
  type            = "ICAP"
  icapprofilename = "reqmod-profile"
  servername      = "vicap"
  ifserverdown    = "DROP"
}
```


## Argument Reference

* `name` - (Required) Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `type` - (Required) Type of operation this action is going to perform. following actions are available to configure: * ICAP - forward the incoming request or response to an ICAP server for modification. * INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention. * MIRROR - Forwards cloned packets for Intrusion Detection. * NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device. * NSTRACE - capture current and further incoming packets on this transaction.
* `icapprofilename` - (Optional) Name of the ICAP profile to be attached to the contentInspection action.
* `ifserverdown` - (Optional) Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are: * RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired. * DROP - Drop the request without sending a response to the user. * CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.
* `serverip` - (Optional) IP address of remoteService
* `servername` - (Optional) Name of the LB vserver or service
* `serverport` - (Optional) Port of remoteService


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionaction. It has the same value as the `name` attribute.


## Import

A contentinspectionaction can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionaction.tf_contentinspectionaction my_ci_action
```
