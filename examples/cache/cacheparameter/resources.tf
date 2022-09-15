resource "citrixadc_cacheparameter" "tf_cacheparameter" {
  memlimit    = "3500"
  maxpostlen  = "6000"
  verifyusing = "HOSTNAME"
}
