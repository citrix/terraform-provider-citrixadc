resource "citrixadc_nsratecontrol" "tf_nsratecontrol" {
  tcpthreshold    = 10
  udpthreshold    = 10
  icmpthreshold   = 100
  tcprstthreshold = 100
}