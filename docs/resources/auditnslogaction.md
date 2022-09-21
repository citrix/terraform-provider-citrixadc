---
subcategory: "Audit"
---

# Resource: auditnslogaction

The auditnslogaction resource is used to create auditnslogaction.


## Example usage

```hcl
resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
  name     = "my_auditnslogaction"
  serverip = "10.222.74.180"
  loglevel = ["ALERT", "CRITICAL"]
  tcp      = "ALL"
  acl      = "ENABLED"
}
```


## Argument Reference

* `loglevel` - (Required) Audit log level, which specifies the types of events to log.  Available settings function as follows:  * ALL - All events. * EMERGENCY - Events that indicate an immediate crisis on the server. * ALERT - Events that might require action. * CRITICAL - Events that indicate an imminent server crisis. * ERROR - Events that indicate some type of error. * WARNING - Events that require action in the near future. * NOTICE - Events that the administrator should know about. * INFORMATIONAL - All but low-level events. * DEBUG - All events, in extreme detail. * NONE - No events.
* `name` - (Required) Name of the nslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog action is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my nslog action" or 'my nslog action').
* `acl` - (Optional) Log access control list (ACL) messages.
* `alg` - (Optional) Log the ALG messages
* `appflowexport` - (Optional) Export log messages to AppFlow collectors. Appflow collectors are entities to which log messages can be sent so that some action can be performed on them.
* `contentinspectionlog` - (Optional) Log Content Inspection event information
* `dateformat` - (Optional) Format of dates in the logs. Supported formats are:  * MMDDYYYY - U.S. style month/date/year format. * DDMMYYYY - European style date/month/year format. * YYYYMMDD - ISO style year/month/date format.
* `domainresolvenow` - (Optional) Immediately send a DNS query to resolve the server's domain name.
* `domainresolveretry` - (Optional) Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the audit server if the last query failed.
* `logfacility` - (Optional) Facility value, as defined in RFC 3164, assigned to the log message.  Log facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.
* `lsn` - (Optional) Log the LSN messages
* `serverdomainname` - (Optional) Auditserver name as a FQDN. Mutually exclusive with serverIP
* `serverip` - (Optional) IP address of the nslog server.
* `serverport` - (Optional) Port on which the nslog server accepts connections.
* `sslinterception` - (Optional) Log SSL Interception event information
* `subscriberlog` - (Optional) Log subscriber session event information
* `tcp` - (Optional) Log TCP messages.
* `timezone` - (Optional) Time zone used for date and timestamps in the logs.  Available settings function as follows:  * GMT_TIME. Coordinated Universal Time. * LOCAL_TIME. The server's timezone setting.
* `urlfiltering` - (Optional) Log URL filtering event information
* `userdefinedauditlog` - (Optional) Log user-configurable log messages to nslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditnslogaction. It has the same value as the `name` attribute.


## Import

A auditnslogaction can be imported using its name, e.g.

```shell
terraform import citrixadc_auditnslogaction.tf_auditnslogaction my_auditnslogaction
```
