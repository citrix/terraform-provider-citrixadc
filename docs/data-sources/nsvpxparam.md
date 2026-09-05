---
subcategory: "NS"
---

# citrixadc_nsvpxparam (Data Source)

Data source for querying Citrix ADC VPX parameters. This data source retrieves VPX-specific configuration parameters for virtual appliances.

## Example Usage

```hcl
data "citrixadc_nsvpxparam" "example" {
  ownernode = 0
}

# Output VPX parameters
output "cpuyield" {
  value = data.citrixadc_nsvpxparam.example.cpuyield
}

output "kvmvirtiomultiqueue" {
  value = data.citrixadc_nsvpxparam.example.kvmvirtiomultiqueue
}
```

## Argument Reference

The following arguments are supported:

* `ownernode` - (Required) ID of the cluster node for which you are querying the parameters. It can be configured only through the cluster IP address.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nsvpxparam datasource.
* `cpuyield` - CPU yield setting for the virtual appliance. Options: YES (allow vCPU yield), NO (no vCPU yield), DEFAULT (restore default behavior).
* `kvmvirtiomultiqueue` - Multi-queue setting for KVM VPX with virtio NICs. Options: YES (use multiple queues), NO (use single queue).
* `masterclockcpu1` - (Deprecated) Master clock CPU setting.

### Read-only nsvpxparam metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_nsvpxparam` resource). They are Computed / GET-only, and any attribute the appliance does not return is `null`.

* `vpxenvironment` - Shows the VPX running environment (VPX, AWSCLOUD, OPENSTACKCLOUD, AZURECLOUD, GOOGLECLOUD, ALICLOUD, IBMCLOUD).
* `memorystatus` - Provides the information about memory status (Sufficent, Insufficent).
* `cloudproductcode` - Cloud Product Code Description.
* `vpxoemcode` - OEM Distribution Code.
* `technicalsupportpin` - Technical Support PIN for cloud subscription VMs.

