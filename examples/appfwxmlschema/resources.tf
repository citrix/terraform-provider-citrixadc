resource "citrixadc_systemfile" "tf_xmlschema" {
  filename     = "appfwxmlschema.xml"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwxmlschema.xml")
}
resource "citrixadc_appfwxmlschema" "tf_appfwxmlschema" {
  name       = "tf_appfwxmlschema"
  src        = "local://appfwxmlschema.xml"
  depends_on = [citrixadc_systemfile.tf_xmlschema]
  comment    = "TestingExample"
}