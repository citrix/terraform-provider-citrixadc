resource "citrixadc_systemfile" "tf_errorpage" {
  filename     = "appfwhtmlerrorpage.html"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwhtmlerrorpage.html")
}
resource "citrixadc_appfwhtmlerrorpage" "tf_appfwhtmlerrorpage" {
  name       = "tf_appfwhtmlerrorpage"
  src        = "local://appfwhtmlerrorpage.html"
  depends_on = [citrixadc_systemfile.tf_errorpage]
  comment    = "TestingExample"
}