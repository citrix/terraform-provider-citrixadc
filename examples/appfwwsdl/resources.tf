resource "citrixadc_systemfile" "tf_wsdl" {
  filename     = "sample.wsdl"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/sample.wsdl")
}
resource "citrixadc_appfwwsdl" "tf_appfwwsdl" {
  name       = "tf_appfwwsdl"
  src        = "local://sample.wsdl"
  depends_on = [citrixadc_systemfile.tf_wsdl]
  comment    = "TestingExample"
}