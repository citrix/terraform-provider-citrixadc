resource "citrixadc_userprotocol" "tf_userprotocol" {
  name      = "my_userprotocol"
  transport = "TCP"
  extension = "my_extension"
  comment   = "my_comment"
}
