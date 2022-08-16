resource "citrixadc_aaakcdaccount" "tf_aaakcdaccount" {
  kcdaccount    = "my_kcdaccount"
  delegateduser = "john"
  kcdpassword   = "my_password"
  realmstr      = "my_realm"
}
