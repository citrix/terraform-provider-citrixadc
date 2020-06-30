resource "citrixadc_netprofile" "tf_netprofile" {
    name = "tf_netprofile"
    proxyprotocol = "ENABLED"
    proxyprotocoltxversion = "V1"
}
