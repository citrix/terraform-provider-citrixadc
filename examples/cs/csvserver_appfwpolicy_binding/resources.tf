resource citrixadc_csvserver_appfwpolicy_binding demo_csvserver_appfwpolicy_binding {
  name                   = citrixadc_csvserver.demo_cs.name
  priority               = 100
  policyname             = citrixadc_appfwpolicy.demo_appfwpolicy1.name
  gotopriorityexpression = "END"
}
resource "citrixadc_csvserver" "demo_cs" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource citrixadc_appfwprofile demo_appfwprofile {
  name                     = "demo_appfwprofile"
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

resource citrixadc_appfwpolicy demo_appfwpolicy1 {
  name        = "demo_appfwpolicy1"
  profilename = citrixadc_appfwprofile.demo_appfwprofile.name
  rule        = "true"
}
