resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
  name                       = "my_tmsession_action"
  sesstimeout                = 10
  defaultauthorizationaction = "ALLOW"
  sso                        = "OFF"
}
