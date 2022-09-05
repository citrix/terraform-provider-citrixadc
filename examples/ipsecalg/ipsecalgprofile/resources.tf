resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
  name              = "my_ipsecalgprofile"
  ikesessiontimeout = 50
  espsessiontimeout = 20
  connfailover      = "DISABLED"
}
