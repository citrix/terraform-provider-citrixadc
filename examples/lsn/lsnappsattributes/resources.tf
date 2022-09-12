resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
  name              = "my_lsn_appattributes"
  transportprotocol = "TCP"
  port              = 90
  sessiontimeout    = 40
}
