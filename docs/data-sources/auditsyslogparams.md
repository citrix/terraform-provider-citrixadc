---
subcategory: "Audit"
---

# Data Source: auditsyslogparams

The auditsyslogparams data source allows you to retrieve the global audit syslog parameters configuration.

## Example usage

```terraform
data "citrixadc_auditsyslogparams" "example" {
}

output "dateformat" {
  value = data.citrixadc_auditsyslogparams.example.dateformat
}

output "loglevel" {
  value = data.citrixadc_auditsyslogparams.example.loglevel
}

output "tcp" {
  value = data.citrixadc_auditsyslogparams.example.tcp
}
```

## Argument Reference

This datasource does not require any arguments as it retrieves the global audit syslog parameters configuration.

## Attribute Reference

The following attributes are available:

* `id` - The id of the auditsyslogparams configuration.
* `acl` - Log access control list (ACL) messages.
* `alg` - Log the ALG messages.
* `appflowexport` - Export log messages to AppFlow collectors.
* `contentinspectionlog` - Log Content Inspection event information.
* `dateformat` - Format of dates in the logs. Supported formats: MMDDYYYY, DDMMYYYY, YYYYMMDD.
* `dns` - Log DNS related syslog messages.
* `logfacility` - Facility value, as defined in RFC 3164, assigned to the log message.
* `loglevel` - Types of information to be logged. Available values: ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, NONE.
* `lsn` - Log the LSN messages.
* `protocolviolations` - Log protocol violations.
* `serverip` - IP address of the syslog server.
* `serverport` - Port on which the syslog server accepts connections.
* `sslinterception` - Log SSL Interception event information.
* `streamanalytics` - Export log stream analytics statistics to syslog server.
* `subscriberlog` - Log subscriber session event information.
* `tcp` - Log TCP messages.
* `timezone` - Time zone used for date and timestamps in the logs. Available settings: GMT_TIME, LOCAL_TIME.
* `urlfiltering` - Log URL filtering event information.
* `userdefinedauditlog` - Log user-configurable log messages to syslog.
