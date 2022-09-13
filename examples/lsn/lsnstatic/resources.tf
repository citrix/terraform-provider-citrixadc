resource "citrixadc_lsnstatic" "tf_lsnstatic" {
  name              = "my_lsn_static"
  transportprotocol = "TCP"
  subscrip          = "10.222.74.128"
  subscrport        = 3000
}
