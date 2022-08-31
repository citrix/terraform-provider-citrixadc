resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
  name                        = "my_tmsamlssoprofile"
  assertionconsumerserviceurl = "https://service.example.com"
  sendpassword                = "OFF"
  relaystaterule              = "true"
}
