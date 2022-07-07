resource "citrixadc_hafailover" "tf_failover" {
  ipaddress = "10.222.74.152"
  state     = "Primary"
  force     = true
}
