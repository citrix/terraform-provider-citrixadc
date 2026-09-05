---
subcategory: "Content Switching"
---

# Data Source: cspolicylabel

The cspolicylabel data source allows you to retrieve information about Content Switching policy labels.


## Example usage

```terraform
data "citrixadc_cspolicylabel" "tf_cspolicylabel" {
  labelname = "tf_policylabel"
}

output "cspolicylabeltype" {
  value = data.citrixadc_cspolicylabel.tf_cspolicylabel.cspolicylabeltype
}

output "newname" {
  value = data.citrixadc_cspolicylabel.tf_cspolicylabel.newname
}
```


## Argument Reference

* `labelname` - (Required) Name for the policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. The label name must be unique within the list of policy labels for content switching.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cspolicylabeltype` - Protocol supported by the policy label. All policies bound to the policy label must either match the specified protocol or be a subtype of that protocol. Available settings function as follows:
  * HTTP - Supports policies that process HTTP traffic. Used to access unencrypted Web sites. (The default.)
  * SSL - Supports policies that process HTTPS/SSL encrypted traffic. Used to access encrypted Web sites.
  * TCP - Supports policies that process any type of TCP traffic, including HTTP.
  * SSL_TCP - Supports policies that process SSL-encrypted TCP traffic, including SSL.
  * UDP - Supports policies that process any type of UDP-based traffic, including DNS.
  * DNS - Supports policies that process DNS traffic.
  * ANY - Supports all types of policies except HTTP, SSL, and TCP.
  * SIP_UDP - Supports policies that process UDP based Session Initiation Protocol (SIP) traffic. SIP initiates, manages, and terminates multimedia communications sessions, and has emerged as the standard for Internet telephony (VoIP).
  * RTSP - Supports policies that process Real Time Streaming Protocol (RTSP) traffic. RTSP provides delivery of multimedia and other streaming data, such as audio, video, and other types of streamed media.
  * RADIUS - Supports policies that process Remote Authentication Dial In User Service (RADIUS) traffic. RADIUS supports combined authentication, authorization, and auditing services for network management.
  * MYSQL - Supports policies that process MYSQL traffic.
  * MSSQL - Supports policies that process Microsoft SQL traffic.
* `newname` - The new name of the content switching policylabel.
* `id` - The id of the cspolicylabel. It has the same value as the `labelname` attribute.

### Read-only cspolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_cspolicylabel` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `numpol` - Number of policies bound to the label.
* `hits` - Number of times the policy label was invoked.
* `policyname` - Name of the content switching policy.
* `priority` - Specifies the priority of the policy.
* `targetvserver` - Name of the virtual server to which to forward requests that match the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation (for example `policylabel`).
* `invoke_labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
