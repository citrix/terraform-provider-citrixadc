resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
  sipalgprofilename      = "my_lsn_sipalgprofile"
  datasessionidletimeout = 150
  sipsessiontimeout      = 150
  registrationtimeout    = 150
  sipsrcportrange        = "4200"
  siptransportprotocol   = "TCP"
}
