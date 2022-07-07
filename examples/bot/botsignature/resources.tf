resource "citrixadc_systemfile" "tf_signature" {
  filename     = "bot_signature.json"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/bot_signatures.json")
}
resource "citrixadc_botsignature" "tf_botsignature" {
  name       = "tf_botsignature"
  src        = "local://bot_signature.json"
  depends_on = [citrixadc_systemfile.tf_signature]
  comment    = "TestingExample"
}