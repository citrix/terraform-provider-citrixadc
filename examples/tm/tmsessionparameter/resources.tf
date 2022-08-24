resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
  sesstimeout                = 40
  defaultauthorizationaction = "ALLOW"
  sso                        = "OFF"
  ssodomain                  = 3
}
