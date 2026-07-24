---
subcategory: "NS"
---

# Resource: nsextension

This resource is used to import NetScaler extensions.


## Example usage

```hcl
resource "citrixadc_nsextension" "tf_nsextension" {
  name           = "myextension"
  src            = "local://myextension.lua"
  comment        = "Custom protocol parser"
  trace          = "calls"
  tracefunctions = "handler1,handler2"
  tracevariables = "ctx,buf"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the extension object on the Citrix ADC. Changing this value forces a new resource to be created.
* `src` - (Required) Local path to and name of, or URL (protocol, host, path, and file name) for, the file from which the extension is imported (for example `local://myextension.lua` or `http://192.0.2.10/extensions/myextension.lua`). NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server. This value is a write-only import input: it is sent to the appliance but is not returned by a GET, so it is not recoverable on read or import. Changing this value forces a new resource to be created.
* `overwrite` - (Optional) Overwrites the existing file when importing. Changing this value forces a new resource to be created.
* `comment` - (Optional) Any comments to preserve information about the extension object. This attribute is updateable in place.
* `trace` - (Optional) Enables tracing to the NS log file of extension execution. `off` turns off tracing; `calls` traces extension function calls with arguments and function returns with the first return value; `lines` adds line numbers for executed extension lines; `all` adds local variables changed by executed extension lines. Note that the DEBUG log level must be enabled to see extension tracing (`set audit syslogParams -loglevel ALL` or `-loglevel DEBUG`). This attribute is updateable in place. Possible values: [ off, calls, lines, all ]
* `tracefunctions` - (Optional) Comma-separated list of extension functions to trace. By default, all extension functions are traced. This attribute is updateable in place.
* `tracevariables` - (Optional) Comma-separated list of variables (in traced extension functions) to trace. By default, all variables are traced. This attribute is updateable in place.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsextension. It has the same value as the `name` attribute.


## Import

A nsextension can be imported using its name, e.g.

```shell
terraform import citrixadc_nsextension.tf_nsextension myextension
```

~> **Note** Because `src` is not returned by the appliance, after import you must add `src` (and, if used, `overwrite`) to your configuration to match the original import source.
