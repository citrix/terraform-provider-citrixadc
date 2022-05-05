resource "citrixadc_systemfile" "tf_jsonerrorpage" {
  filename     = "appfwjsonerrorpage.json"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwjsonerrorpage.json")
}
resource "citrixadc_appfwjsonerrorpage" "tf_appfwjsonerrorpage" {
  name       = "tf_appfwjsonerrorpage"
  src        = "local://appfwjsonerrorpage.json"
  depends_on = [citrixadc_systemfile.tf_jsonerrorpage]
  comment    = "TestingExample"
}