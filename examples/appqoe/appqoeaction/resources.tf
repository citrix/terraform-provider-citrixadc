resource "citrixadc_appqoeaction" "tf_appqoeaction" {
  name        = "my_appqoeaction"
  priority    = "LOW"
  respondwith = "NS"
  delay       = 40
}
