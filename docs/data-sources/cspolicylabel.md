---
subcategory: "Content Switching"
---

# Data Source `cspolicylabel`

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

## Attribute Reference

* `id` - The id of the cspolicylabel. It has the same value as the `labelname` attribute.


## Import

A cspolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_cspolicylabel.tf_cspolicylabel tf_policylabel
```
