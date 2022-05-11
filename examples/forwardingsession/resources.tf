resource "citrixadc_forwardingsession" "tf_forwarding" {
  name             = "tf_forwarding"
  network          = "10.102.105.90"
  netmask          = "255.255.255.255"
  connfailover     = "ENABLED"
  sourceroutecache = "ENABLED"
  processlocal     = "DISABLED"
}