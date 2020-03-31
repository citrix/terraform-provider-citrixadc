resource "citrixadc_sslparameter" "default" {
  pushflag       = "2"
  denysslreneg   = "NONSECURE"
  defaultprofile = "ENABLED"
}
