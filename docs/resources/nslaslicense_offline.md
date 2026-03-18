---
subcategory: "NS"
---

# Resource: nslaslicense_offline

The nslaslicense_offline resource is used to generate and apply offline LAS licenses for NetScaler VPX and MPX appliances. This resource orchestrates the complete offline licensing workflow including request generation, LAS server interaction, and license application to the device.


## Example usage

```hcl

resource "citrixadc_nslaslicense_offline" "license_vpx" {
  request_pem      = "CNS_V10000_SERVER"
  request_edition  = "Premium"
  is_fips          = false
  las_secrets_json = "${path.module}/las_secrets.json"
}

output "license_status" {
  value = {
    status       = citrixadc_nslaslicense_offline.license_vpx.status
    build       = citrixadc_nslaslicense_offline.license_vpx.build
    version      = citrixadc_nslaslicense_offline.license_vpx.version
    blob_path    = citrixadc_nslaslicense_offline.license_vpx.license_blob_path
  }
}
```

## LAS Secrets File

The `las_secrets_json` file must contain the following JSON structure with your actual credentials:

```json
{
  "ccid": "<your_citrix_customer_id>",
  "client": "<your_client_id>",
  "password": "<your_client_secret>",
  "las_endpoint": "https://las.cloud.com",
  "cc_endpoint": "https://trust.citrixworkspacesapi.net/root/tokens/clients"
}
```


## Argument Reference

* `request_pem` - (Required) Platform Entitlement Model (PEM) code. Examples: CNS_V1000_SERVER, CNS_V3000_SERVER, CNS_V10000_SERVER. Required for NetScaler VPX and MPX.
* `request_edition` - (Required) License edition. Valid values: Advanced, Standard, Premium.
* `las_secrets_json` - (Required) File path containing LAS authentication secrets and endpoints (ccid, client, password, las_endpoint, cc_endpoint).
* `is_fips` - (Optional) Whether this is a FIPS-enabled device. Default: false.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The system generated id of the nslaslicense_offline resource.
* `lsguid` - License Server GUID extracted from the device.
* `version` - NetScaler software version detected on the device.
* `build` - NetScaler build number detected on the device.
* `license_blob_path` - Local file path where the license blob (.tgz) is saved.
* `status` - License application status (e.g., "applied").
* `last_updated` - Timestamp of when the license was last applied.


## Notes

* This resource requires SSH/SFTP access to the NetScaler device for license application.
* The provider's `username` must be "nsroot" for offline licensing operations.
* License blobs are saved locally in `/tmp/offline_token_<device_ip>_ns_activation.blob.tgz`.
* The resource performs a complete offline licensing workflow: version check, request generation, LAS server interaction, and license application.
* On resource deletion, the license remains active on the device; only the Terraform state is removed.
