#create content switching SSL vServer
resource "citrixadc_csvserver" "tf_csvserverssl" {
  name            = var.sslcsvip_config["name"]
  ipv46           = var.sslcsvip_config["vip"]
  port            = var.sslcsvip_config["port"]
  servicetype     = var.sslcsvip_config["servicetype"]
  httpprofilename = var.sslcsvip_config["httpprofile"]
}
//configure SSL vServer settings
resource "citrixadc_sslvserver" "tf_cssslvserver" {
  ersa        = var.sslcsvip_config["ersa"]
  ssl3        = var.sslcsvip_config["ssl3"]
  tls1        = var.sslcsvip_config["tls1"]
  tls11       = var.sslcsvip_config["tls11"]
  vservername = citrixadc_csvserver.tf_csvserverssl.name
}
//output csvservers "csvserverssl"
output "csvserverssl" {
  value = citrixadc_csvserver.tf_csvserverssl.name
}

#create content switching HTTP vServer
resource "citrixadc_csvserver" "tf_csvserverhttp" {
  name            = var.httpcsvip_config["name"]
  ipv46           = var.httpcsvip_config["vip"]
  port            = var.httpcsvip_config["port"]
  servicetype     = var.httpcsvip_config["servicetype"]
  redirecturl     = var.httpcsvip_config["redirecturl"]
  httpprofilename = var.httpcsvip_config["httpprofile"]
}
//output csvservers "csvserverhttp"
output "csvserverhttp" {
  value = citrixadc_csvserver.tf_csvserverhttp.name
}
#create target HTTP LB vServer
resource "citrixadc_lbvserver" "tf_lbvserverhttp" {
  name            = var.httpvip_config["name"]
  ipv46           = var.httpvip_config["vip"]
  port            = var.httpvip_config["port"]
  servicetype     = var.httpvip_config["servicetype"]
  lbmethod        = var.httpvip_config["lbmethod"]
  persistencetype = var.httpvip_config["persistencetype"]
  timeout         = var.httpvip_config["timeout"]
  httpprofilename = var.httpvip_config["httpprofile"]
}
//output LBvservers "lbvserverhttp"
output "lbvserverhttp" {
  value = citrixadc_lbvserver.tf_lbvserverhttp.name
}

//call target LB1 
data "terraform_remote_state" "TargetLB1" {
  backend = "local"
  config = {
    path = "./TargetLB1/terraform.tfstate"
  }
}
data "terraform_remote_state" "TargetLB2" {
  backend = "local"
  config = {
    path = "./TargetLB2/terraform.tfstate"
  }
}
data "terraform_remote_state" "TargetLB3" {
  backend = "local"
  config = {
    path = "./TargetLB3/terraform.tfstate"
  }
}

//call module "lbservice" 
module "loadbalancingservices" {
  source         = "../../commonmodules/lbservice"
  servers        = var.servers
  service_config = var.service_config
  
}
//output module "lbservices"
output "lbservices" {
  value = module.loadbalancingservices.lbservices
}
//passing vs cipher group data value
data "terraform_remote_state" "cipher_vs" {
  backend = "local"
  config = {
    path = "../../commonmodules/cipher_vs/terraform.tfstate"
  }
}
//bind cipher group to SSL CS vServer
resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_cssslvserver_sslciphersuite_binding" {
  ciphername  = data.terraform_remote_state.cipher_vs.outputs.vscipher_name
  vservername = citrixadc_csvserver.tf_csvserverssl.name
  depends_on  = [citrixadc_csvserver.tf_csvserverssl]
}
output "cssslvserver_cipherbinding" {
  value = citrixadc_sslvserver_sslciphersuite_binding.tf_cssslvserver_sslciphersuite_binding.id
}
//LB SSL vServer-services binding
resource "citrixadc_lbvserver_service_binding" "tf_bindinghttp" {
  for_each    = var.servers
  name        = citrixadc_lbvserver.tf_lbvserverhttp.name
  servicename = "${each.key}-${var.service_config["service_port"]}"
  depends_on  = [module.loadbalancingservices]
}

//passing common ssl certificate keypair value
data "terraform_remote_state" "sslcertkeypair" {
  backend = "local"
  config = {
    path = "../../commonmodules/sslcertkeypair/terraform.tfstate"
  }
}
//bind ssl certkey pair to CS ssl vServer
resource "citrixadc_sslvserver_sslcertkey_binding" "cert_bindingcs" {
  vservername = citrixadc_csvserver.tf_csvserverssl.name
  certkeyname = data.terraform_remote_state.sslcertkeypair.outputs.sslcertname
  //depends_on  = [citrixadc_sslvserver.tf_csvserverssl]
}
output "cssslvserver_certbinding" {
  value = citrixadc_sslvserver_sslcertkey_binding.cert_bindingcs.id
}

resource "citrixadc_csvserver" "tf_defaulttargetbinding" {
  name             = citrixadc_csvserver.tf_csvserverssl.name
  lbvserverbinding = citrixadc_lbvserver.tf_lbvserverhttp.name
  servicetype      = var.sslcsvip_config["servicetype"]
}
#create content switching action
resource "citrixadc_csaction" "tf_csaction1" {
  name            = var.cs_action1["name"]
  targetlbvserver = data.terraform_remote_state.TargetLB1.outputs.targetlb1name
}
#create content switching policy 1 and bind to SSL CS vServer created, bind target LB vServer to CS action
resource "citrixadc_cspolicy" "tf_cspolicy1" {
  action     = citrixadc_csaction.tf_csaction1.name
  policyname = var.cs_pol1["name"]
  rule       = var.cs_pol1["rule"]
}
output "csaction1" {
  value = citrixadc_csaction.tf_csaction1.name
}
output "cspolicy1" {
  value = citrixadc_cspolicy.tf_cspolicy1.policyname
}
resource "citrixadc_csvserver_cspolicy_binding" "tf_cspolbind1" {
  name       = citrixadc_csvserver.tf_csvserverssl.name
  policyname = citrixadc_cspolicy.tf_cspolicy1.policyname
  priority   = var.cs_pol1["priority"]
}
output "cspolicy1binding" {
  value = citrixadc_csvserver_cspolicy_binding.tf_cspolbind1.id
}
#create content switching action
resource "citrixadc_csaction" "tf_csaction2" {
  name            = var.cs_action2["name"]
  targetlbvserver = data.terraform_remote_state.TargetLB2.outputs.targetlb2name
}
#create content switching policy 2 and bind to SSL CS vServer created, bind target LB vServer to CS action
resource "citrixadc_cspolicy" "tf_cspolicy2" {
  action     = citrixadc_csaction.tf_csaction2.name
  policyname = var.cs_pol2["name"]
  rule       = var.cs_pol2["rule"]
}
output "csaction2" {
  value = citrixadc_csaction.tf_csaction2.name
}
output "cspolicy2" {
  value = citrixadc_cspolicy.tf_cspolicy2.policyname
}
resource "citrixadc_csvserver_cspolicy_binding" "tf_cspolbind2" {
  name       = citrixadc_csvserver.tf_csvserverssl.name
  policyname = citrixadc_cspolicy.tf_cspolicy2.policyname
  priority   = var.cs_pol2["priority"]
}
output "cspolicy2binding" {
  value = citrixadc_csvserver_cspolicy_binding.tf_cspolbind2.id
}
#create content switching action
resource "citrixadc_csaction" "tf_csaction3" {
  name            = var.cs_action3["name"]
  targetlbvserver = data.terraform_remote_state.TargetLB3.outputs.targetlb3name
}
#create content switching policy 3 and bind to SSL CS vServer created, bind target LB vServer to CS action
resource "citrixadc_cspolicy" "tf_cspolicy3" {
  action     = citrixadc_csaction.tf_csaction3.name
  policyname = var.cs_pol3["name"]
  rule       = var.cs_pol3["rule"]
}
output "csaction3" {
  value = citrixadc_csaction.tf_csaction3.name
}
output "cspolicy3" {
  value = citrixadc_cspolicy.tf_cspolicy3.policyname
}
resource "citrixadc_csvserver_cspolicy_binding" "tf_cspolbind3" {
  name       = citrixadc_csvserver.tf_csvserverssl.name
  policyname = citrixadc_cspolicy.tf_cspolicy3.policyname
  priority   = var.cs_pol3["priority"]
}
output "cspolicy3binding" {
  value = citrixadc_csvserver_cspolicy_binding.tf_cspolbind3.id
}

//passing responder redirection policy 
data "terraform_remote_state" "responder_httpTohttps" {
  backend = "local"
  config = {
    path = "../../commonmodules/responder_httpTohttps/terraform.tfstate"
  }
}
#bind responder policy to HTTP CS vServer
resource "citrixadc_csvserver_responderpolicy_binding" "tf_bind" {
  name       = citrixadc_csvserver.tf_csvserverhttp.name
  policyname = data.terraform_remote_state.responder_httpTohttps.outputs.responder_httpTohttps
  priority   = 100
  bindpoint  = "REQUEST"
}