---
subcategory: "Audit"
---

# Resource: auditnslogparams

The auditnslogparams resource is used to update auditnslogparams.


## Example usage

```hcl
resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
  dateformat = "DDMMYYYY"
  loglevel   = ["EMERGENCY"]
  tcp        = "ALL"
}
```


## Argument Reference

* `serverip` - (Optional) IP address of the nslog server. Minimum length =  1
* `serverport` - (Optional) Port on which the nslog server accepts connections. Minimum value =  1
* `dateformat` - (Optional) Format of dates in the logs. Supported formats are: * MMDDYYYY - U.S. style month/date/year format. * DDMMYYYY - European style date/month/year format. * YYYYMMDD - ISO style year/month/date format. Possible values: [ MMDDYYYY, DDMMYYYY, YYYYMMDD ]
* `loglevel` - (Optional) Types of information to be logged. Available settings function as follows: * ALL - All events. * EMERGENCY - Events that indicate an immediate crisis on the server. * ALERT - Events that might require action. * CRITICAL - Events that indicate an imminent server crisis. * ERROR - Events that indicate some type of error. * WARNING - Events that require action in the near future. * NOTICE - Events that the administrator should know about. * INFORMATIONAL - All but low-level events. * DEBUG - All events, in extreme detail. * NONE - No events. Possible values: [ ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE ]
* `logfacility` - (Optional) Facility value, as defined in RFC 3164, assigned to the log message. Log facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external. Possible values: [ LOCAL0, LOCAL1, LOCAL2, LOCAL3, LOCAL4, LOCAL5, LOCAL6, LOCAL7 ]
* `tcp` - (Optional) Configure auditing to log TCP messages. Possible values: [ NONE, ALL ]
* `acl` - (Optional) Configure auditing to log access control list (ACL) messages. Possible values: [ ENABLED, DISABLED ]
* `timezone` - (Optional) Time zone used for date and timestamps in the logs. Supported settings are: * GMT_TIME - Coordinated Universal Time. * LOCAL_TIME - Use the server's timezone setting. Possible values: [ GMT_TIME, LOCAL_TIME ]
* `userdefinedauditlog` - (Optional) Log user-configurable log messages to nslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria. Possible values: [ YES, NO ]
* `appflowexport` - (Optional) Export log messages to AppFlow collectors. Appflow collectors are entities to which log messages can be sent so that some action can be performed on them. Possible values: [ ENABLED, DISABLED ]
* `lsn` - (Optional) Log the LSN messages. Possible values: [ ENABLED, DISABLED ]
* `alg` - (Optional) Log the ALG messages. Possible values: [ ENABLED, DISABLED ]
* `subscriberlog` - (Optional) Log subscriber session event information. Possible values: [ ENABLED, DISABLED ]
* `sslinterception` - (Optional) Log SSL Interception event information. Possible values: [ ENABLED, DISABLED ]
* `urlfiltering` - (Optional) Log URL filtering event information. Possible values: [ ENABLED, DISABLED ]
* `contentinspectionlog` - (Optional) Log Content Inspection event information. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditnslogparams. It is a unique string prefixed with  `tf-auditnslogparams-` attribute.
