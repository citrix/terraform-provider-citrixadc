---
subcategory: "Audit"
---

# Resource: auditsyslogaction

This resource is used to create audit syslog actions.


## Example usage

```hcl
resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}
```


## Argument Reference

* `name` - (Required) Name of the audit syslog action.
* `serverip` - (Optional) IP address of the syslog server.
* `serverdomainname` - (Optional) SYSLOG server name as a FQDN. Mutually exclusive with serverIP/lbVserverName.
* `domainresolveretry` - (Optional) Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.
* `lbvservername` - (Optional) Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName.
* `serverport` - (Optional) Port on which the syslog server accepts connections.
* `loglevel` - (Optional) Audit log level, which specifies the types of events to log. Possible values: [ ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE ]
* `dateformat` - (Optional) Format of dates in the logs. Possible values: [ MMDDYYYY, DDMMYYYY, DDMMYYYY ]
* `logfacility` - (Optional) Facility value, as defined in RFC 3164, assigned to the log message. Possible values: [ LOCAL0, LOCAL1, LOCAL2, LOCAL3, LOCAL4, LOCAL5, LOCAL6, LOCAL7 ]
* `tcp` - (Optional) Log TCP messages. Possible values: [ NONE, ALL ]
* `acl` - (Optional) Log access control list (ACL) messages. Possible values: [ ENABLED, DISABLED ]
* `timezone` - (Optional) Time zone used for date and timestamps in the logs. Possible values: [ GMT\_TIME, LOCAL\_TIME ]
* `userdefinedauditlog` - (Optional) Log user-configurable log messages to syslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria. Possible values: [ YES, NO ]
* `appflowexport` - (Optional) Export log messages to AppFlow collectors. Appflow collectors are entities to which log messages can be sent so that some action can be performed on them. Possible values: [ ENABLED, DISABLED ]
* `lsn` - (Optional) Log lsn info. Possible values: [ ENABLED, DISABLED ]
* `alg` - (Optional) Log alg info. Possible values: [ ENABLED, DISABLED ]
* `subscriberlog` - (Optional) Log subscriber session event information. Possible values: [ ENABLED, DISABLED ]
* `transport` - (Optional) Transport type used to send auditlogs to syslog server. Possible values: [ TCP, UDP ]
* `tcpprofilename` - (Optional) Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.
* `maxlogdatasizetohold` - (Optional) Max size of log data that can be held in NSB chain of server info.
* `dns` - (Optional) Log DNS related syslog messages. Possible values: [ ENABLED, DISABLED ]
* `contentinspectionlog` - (Optional) Log Content Inspection event information. Possible values: [ ENABLED, DISABLED ]
* `netprofile` - (Optional) Name of the network profile. The SNIP configured in the network profile will be used as source IP while sending log messages.
* `sslinterception` - (Optional) Log SSL Interception event information. Possible values: [ ENABLED, DISABLED ]
* `urlfiltering` - (Optional) Log URL filtering event information. Possible values: [ ENABLED, DISABLED ]
* `domainresolvenow` - (Optional) Immediately send a DNS query to resolve the server's domain name.
* `managementlog` - (Optional) Management log specifies the categories of log files to be exported. It use destination and transport from PE params. Possible values: [ ALL, SHELL, ACCESS, NSMGMT, NONE ]
* `mgmtloglevel` - (Optional) Management log level, which specifies the types of events to log. Possible values: [ ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE ]
* `syslogcompliance` - (Optional) Setting this parameter ensures that all the Audit Logs generated for this Syslog Action comply with an RFC.
* `httpauthtoken` - (Optional) Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter.
* `httpendpointurl` - (Optional) The URL at which to upload the logs messages on the endpoint.
* `streamanalytics` - (Optional) Export log stream analytics statistics to syslog server.

## Attributes

In addition to the arguments, the following attributes are available:

* `id` - The id of the audit syslog action. It has the same value as the `name` attribute.


## Import

An instance of the resource can be imported using its name, e.g.

```shell
terraform import citrixadc_auditsyslogaction.tf_syslogaction tf_syslogaction
```
