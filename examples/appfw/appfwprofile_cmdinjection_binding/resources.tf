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
resource "citrixadc_appfwprofile_cmdinjection_binding" "tf_binding" {
  name                 = citrixadc_appfwprofile.tf_appfwprofile.name
  cmdinjection         = "tf_cmdinjection"
  formactionurl_cmd    = "http://10.10.10.10/"
  as_scan_location_cmd = "HEADER"
  as_value_type_cmd    = "Keyword"
  as_value_expr_cmd    = "[a-z]+grep"
  alertonly            = "OFF"
  isvalueregex_cmd     = "REGEX"
  isautodeployed       = "NOTAUTODEPLOYED"
  comment              = "Testing"
}