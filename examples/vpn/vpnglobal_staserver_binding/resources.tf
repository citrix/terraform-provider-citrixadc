resource "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
  staserver      = "http://www.example.com/"
  staaddresstype = "IPV4"
}