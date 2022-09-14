resource "citrixadc_locationfile" "tf_locationfile" {
  locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"
  format       = "netscaler"
}

resource "citrixadc_locationfile_import" "tf_locationfile_import" {
  locationfile = "my_file"
  src          = "local://my_location_file"
}
