resource "citrixadc_appfwpolicylabel" "tf_appfwpolicylabel" {
  labelname       = "tf_appfwpolicylabel"
  policylabeltype = "http_req"
}
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}

resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
  name        = "tf_appfwpolicy"
  profilename = citrixadc_appfwprofile.tf_appfwprofile.name
  rule        = "true"
}
resource "citrixadc_appfwpolicylabel_appfwpolicy_binding" "tf_binding" {
  labelname  = citrixadc_appfwpolicylabel.tf_appfwpolicylabel.labelname
  policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
  priority   = 90
}