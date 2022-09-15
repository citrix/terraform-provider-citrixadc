resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
  name           = "my_rdpserverprofile"
  psk            = "key"
  rdpredirection = "ENABLE"
  rdpport        = 4000
}
