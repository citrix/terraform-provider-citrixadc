// `network`, `netmask`, and `gateway` are MANDATORY attributes
resource "citrixadc_route" "route1" {
  depends_on = [citrixadc_nsip.nsip]
  network    = "100.0.100.0"
  netmask    = "255.255.255.0"
  gateway    = "100.0.1.1"
  advertise  = "ENABLED"
}

resource "citrixadc_nsip" "nsip" {
  ipaddress = "100.0.1.100"
  netmask   = "255.255.255.0"
}
