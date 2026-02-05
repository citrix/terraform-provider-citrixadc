---
subcategory: "Audit"
---

# Data Source: auditsyslogaction

The auditsyslogaction data source allows you to retrieve information about an existing audit syslog action.

## Example usage

```terraform
data "citrixadc_auditsyslogaction" "example" {
  name = "my_syslog_action"
}

output "server_ip" {
  value = data.citrixadc_auditsyslogaction.example.serverip
}

output "server_port" {
  value = data.citrixadc_auditsyslogaction.example.serverport
}

output "transport" {
  value = data.citrixadc_auditsyslogaction.example.transport
}
```

## Argument Reference

* `name` - (Required) Name of the syslog action to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditsyslogaction. It has the same value as the `name` attribute.
* `acl` - Log access control list (ACL) messages.
* `alg` - Log alg info.
* `appflowexport` - Export log messages to AppFlow collectors.
* `contentinspectionlog` - Log Content Inspection event information.
* `dateformat` - Format of dates in the logs. Supported formats are: MMDDYYYY, DDMMYYYY, YYYYMMDD.
* `dns` - Log DNS related syslog messages.
* `domainresolvenow` - Immediately send a DNS query to resolve the server's domain name.
* `domainresolveretry` - Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.
* `httpauthtoken` - Token for authenticating with the endpoint.
* `httpendpointurl` - The URL at which to upload the logs messages on the endpoint.
* `lbvservername` - Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName.
* `logfacility` - Facility value, as defined in RFC 3164, assigned to the log message.
* `loglevel` - Audit log level, which specifies the types of events to log. Available values: ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE.
* `lsn` - Log lsn info.
* `managementlog` - Management log specifies the categories of log files to be exported. Available values: ALL, SHELL, ACCESS, NSMGMT, NONE.
* `maxlogdatasizetohold` - Max size of log data that can be held in NSB chain of server info.
* `mgmtloglevel` - Management log level, which specifies the types of events to log. Available values: ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE.
* `netprofile` - Name of the network profile. The SNIP configured in the network profile will be used as source IP while sending log messages.
* `protocolviolations` - Log protocol violations.
* `serverdomainname` - SYSLOG server name as a FQDN. Mutually exclusive with serverIP/lbVserverName.
* `serverip` - IP address of the syslog server.
* `serverport` - Port on which the syslog server accepts connections.
* `sslinterception` - Log SSL Interception event information.
* `streamanalytics` - Export log stream analytics statistics to syslog server.
* `subscriberlog` - Log subscriber session event information.
* `syslogcompliance` - Setting this parameter ensures that all the Audit Logs generated for this Syslog Action comply with an RFC (e.g., RFC5424).
* `tcp` - Log TCP messages.
* `tcpprofilename` - Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.
* `timezone` - Time zone used for date and timestamps in the logs. Supported settings: GMT_TIME, LOCAL_TIME.
* `transport` - Transport type used to send auditlogs to syslog server. Default type is UDP.
* `urlfiltering` - Log URL filtering event information.
* `userdefinedauditlog` - Log user-configurable log messages to syslog.

## Import

A auditsyslogaction can be imported using its name, e.g.

```shell
terraform import citrixadc_auditsyslogaction.example my_syslog_action
```
