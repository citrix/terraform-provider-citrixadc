---
subcategory: "SSL"
---

# Data Source: citrixadc_sslcrl

The sslcrl data source allows you to retrieve information about a Certificate Revocation List (CRL).

## Example usage

```terraform
data "citrixadc_sslcrl" "tf_sslcrl" {
  crlname = "tf_sslcrl"
}

output "crlpath" {
  value = data.citrixadc_sslcrl.tf_sslcrl.crlpath
}

output "cacert" {
  value = data.citrixadc_sslcrl.tf_sslcrl.cacert
}

output "refresh" {
  value = data.citrixadc_sslcrl.tf_sslcrl.refresh
}
```

## Argument Reference

* `crlname` - (Required) Name for the Certificate Revocation List (CRL). Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the CRL is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `basedn` - Base distinguished name (DN), which is used in an LDAP search to search for a CRL. Citrix recommends searching for the Base DN instead of the Issuer Name from the CA certificate, because the Issuer Name field might not exactly match the LDAP directory structure's DN.
* `binary` - Set the LDAP-based CRL retrieval mode to binary.
* `binddn` - Bind distinguished name (DN) to be used to access the CRL object in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.
* `cacert` - CA certificate that has issued the CRL. Required if CRL Auto Refresh is selected. Install the CA certificate on the appliance before adding the CRL.
* `cacertfile` - Name of and, optionally, path to the CA certificate file. /nsconfig/ssl/ is the default path.
* `cakeyfile` - Name of and, optionally, path to the CA key file. /nsconfig/ssl/ is the default path.
* `crlpath` - Path to the CRL file. /var/netscaler/ssl/ is the default path.
* `day` - Day on which to refresh the CRL, or, if the Interval parameter is not set, the number of days after which to refresh the CRL. If Interval is set to MONTHLY, specify the date. If Interval is set to WEEKLY, specify the day of the week (for example, Sun=0 and Sat=6). This parameter is not applicable if the Interval is set to DAILY.
* `gencrl` - Name of and, optionally, path to the CRL file to be generated. The list of certificates that have been revoked is obtained from the index file. /nsconfig/ssl/ is the default path.
* `id` - The id of the sslcrl. It has the same value as the `crlname` attribute.
* `indexfile` - Name of and, optionally, path to the file containing the serial numbers of all the certificates that are revoked. Revoked certificates are appended to the file. /nsconfig/ssl/ is the default path.
* `inform` - Input format of the CRL file. The two formats supported on the appliance are: PEM - Privacy Enhanced Mail. DER - Distinguished Encoding Rule.
* `interval` - CRL refresh interval. Use the NONE setting to unset this parameter.
* `method` - Method for CRL refresh. If LDAP is selected, specify the method, CA certificate, base DN, port, and LDAP server name. If HTTP is selected, specify the CA certificate, method, URL, and port. Cannot be changed after a CRL is added.
* `password` - Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.
* `port` - Port for the LDAP server.
* `refresh` - Set CRL auto refresh.
* `revoke` - Name of and, optionally, path to the certificate to be revoked. /nsconfig/ssl/ is the default path.
* `scope` - Extent of the search operation on the LDAP server. Available settings function as follows: One - One level below Base DN. Base - Exactly the same level as Base DN.
* `server` - IP address of the LDAP server from which to fetch the CRLs.
* `time` - Time, in hours (1-24) and minutes (1-60), at which to refresh the CRL.
* `url` - URL of the CRL distribution point.

## Import

A sslcrl can be imported using its crlname, e.g.

```shell
terraform import citrixadc_sslcrl.tf_sslcrl tf_sslcrl
```
