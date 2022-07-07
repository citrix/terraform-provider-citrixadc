resource "citrixadc_vpnparameter" "tf_vpnparameter" {
  splitdns              = "LOCAL"
  sesstimeout           = 30
  clientsecuritylog     = "OFF"
  smartgroup            = 10
  splittunnel           = "OFF"
  locallanaccess        = "OFF"
  winsip                = "4.45.5.4"
  samesite              = "None"
  backendcertvalidation = "DISABLED"
  backendserversni      = "DISABLED"
  icasessiontimeout     = "OFF"
  iconwithreceiver      = "OFF"
  linuxpluginupgrade    = "Always"
  uitheme               = "DEFAULT"
  httpport              = [80]
}