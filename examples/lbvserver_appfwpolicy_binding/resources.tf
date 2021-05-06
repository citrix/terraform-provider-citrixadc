resource citrixadc_lbvserver_appfwpolicy_binding demo_lbvserver_appfwpolicy_binding {
  name                   = citrixadc_lbvserver.demo_lb.name
  policyname             = citrixadc_appfwpolicy.demo_appfwpolicy.name
  labeltype              = "reqvserver" # Possible values = reqvserver, resvserver, policylabel
  labelname              = citrixadc_lbvserver.demo_lb.name
  priority               = 100
  bindpoint              = "REQUEST" # Possible values = REQUEST, RESPONSE
  gotopriorityexpression = "END"
  invoke                 = true         # boolean
}

resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  ipv46       = "1.1.1.1"
  port        = "80"
  servicetype = "HTTP"
}

resource citrixadc_appfwpolicy demo_appfwpolicy {
  name        = "demo_appfwpolicy"
  profilename = citrixadc_appfwprofile.demo_appfwprofile.name
  rule        = "true"
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
