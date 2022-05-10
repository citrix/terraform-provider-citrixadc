resource "citrixadc_systemfile" "tf_xmlerrorpage" {
  filename     = "appfwxmlerrorpage.xml"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwxmlerrorpage.xml")
}
resource "citrixadc_appfwxmlerrorpage" "tf_appfwxmlerrorpage" {
  name       = "tf_appfwxmlerrorpage"
  src        = "local://appfwxmlerrorpage.xml"
  depends_on = [citrixadc_systemfile.tf_xmlerrorpage]
  comment    = "TestingExample"
}