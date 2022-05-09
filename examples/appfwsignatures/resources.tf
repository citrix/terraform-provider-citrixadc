resource "citrixadc_systemfile" "tf_signature" {
  filename     = "appfw_signatures.xml"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfw_signatures.xml")
}
resource "citrixadc_appfwsignatures" "tf_appfwsignatures" {
  name       = "tf_appfwsignatures"
  src        = "local://appfw_signatures.xml"
  depends_on = [citrixadc_systemfile.tf_signature]
  comment    = "TestingExample"
}