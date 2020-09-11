---
subcategory: "Audit"
---

# Resource: auditmessageaction

This resource is used to create audit message actions.


## Example usage

```hcl
resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"message to log\""
    logtonewnslog = "YES"
}
```


## Argument Reference

* `name` - (Required) Name of the audit message action.
* `loglevel` - (Optional) Audit log level, which specifies the severity level of the log message being generated..

    The following loglevels are valid:
    * EMERGENCY - Events that indicate an immediate crisis on the server.
    * ALERT - Events that might require action.
    * CRITICAL - Events that indicate an imminent server crisis.
    * ERROR - Events that indicate some type of error.
    * WARNING - Events that require action in the near future.
    * NOTICE - Events that the administrator should know about.
    * INFORMATIONAL - All but low-level events.
    * DEBUG - All events, in extreme detail.

* `stringbuilderexpr` - (Optional) Default-syntax expression that defines the format and content of the log message.
* `logtonewnslog` - (Optional) Send the message to the new nslog. Possible values: [YES, NO].
* `bypasssafetycheck` - (Optional) Bypass the safety check and allow unsafe expressions. Possible values: [YES, NO].

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the resource. It has the same value as the `name` attribute.


## Import

A auditmessageaction can be imported using its name, e.g.

```shell
terraform import citrixadc_auditmessageaction.tf_msgaction tf_msgaction
```
