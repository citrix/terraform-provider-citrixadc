resource "citrixadc_channel" "tf_channel" {
  channel_id = "LA/3"
  tagall     = "ON"
  speed      = "1000"
}
