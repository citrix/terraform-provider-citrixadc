resource "citrixadc_nsconfig_clear" "tfnsaction" {
  force = false
  level = "full"
  #  rbaconfig = "YES"
  timestamp = "2020-04-07T12:44:44Z" #timestamp()
}

resource "citrixadc_nsconfig_save" "tfnssave" {
  all       = true
  timestamp = "2020-04-07T12:44:44Z" #timestamp()
}

resource "citrixadc_nsconfig_update" "tfnsupdate" {
  ipaddress = "10.0.1.165"
  netmask   = "255.255.255.0"
  nsvlan    = 100
  ifnum     = ["1/1", "1/2"]
  tagged    = "YES"
}
