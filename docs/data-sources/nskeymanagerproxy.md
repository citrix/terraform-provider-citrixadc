---
subcategory: "NS"
---

# Data Source: nskeymanagerproxy

The nskeymanagerproxy data source allows you to retrieve information about a Key Manager proxy server configured on the Citrix ADC, looking it up by its server IP address.


## Example usage

```terraform
data "citrixadc_nskeymanagerproxy" "example" {
  serverip = "192.168.20.30"
}

output "keymanagerproxy_port" {
  value = data.citrixadc_nskeymanagerproxy.example.port
}
```


## Argument Reference

* `serverip` - (Optional) IP address of the Key Manager proxy server.
* `servername` - (Optional) Fully qualified domain name of the Key Manager proxy server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nskeymanagerproxy. It is the `serverip` value when set; otherwise it falls back to the `servername` value.
* `port` - Key Manager proxy server port.
* `nodeid` - Unique number that identifies the cluster node.

### Read-only nskeymanagerproxy metadata

This attribute is returned by the appliance on a GET (it is not configurable on the `citrixadc_nskeymanagerproxy` resource). It is Computed/GET-only, and is `null` when the appliance does not return it.

* `status` - Status of Key Manager proxy server connectivity.
