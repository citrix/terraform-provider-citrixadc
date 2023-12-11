
//configure HTTP vServer
resource "citrixadc_lbvserver" "tf_targetlb2" {
  name            = var.targetlb2_config["name"]
  ipv46           = var.targetlb2_config["vip"]
  port            = var.targetlb2_config["port"]
  servicetype     = var.targetlb2_config["servicetype"]
  lbmethod        = var.targetlb2_config["lbmethod"]
  persistencetype = var.targetlb2_config["persistencetype"]
  timeout = var.targetlb2_config["timeout"]
  redirurl        = var.targetlb2_config["redirurl"]
  //httpprofilename = var.targetlb2_config["httpprofile"]
}
//output LBvservers "targetlb2"
output "targetlb2name" {
  value = citrixadc_lbvserver.tf_targetlb2.name
}
/*
//call module "lbmonitor"
module "lbmonitor" {
  source         = "../../../commonmodules/lbmonitor"
  monitor_config = var.monitor_config
}
output "showlbmonitor" {
  value = module.lbmonitor.lbcustommonitor
}*/

//call module "lbservice" 
module "loadbalancingservices" {
  source         = "../../../commonmodules/lbservice"
  servers        = var.servers
  service_config = var.service_config
}
//output module "lbservices"
output "lbservices" {
  value = module.loadbalancingservices.lbservices
}
data "terraform_remote_state" "cipher_svc" {
  backend = "local"
  config = {
    path = "../../../commonmodules/cipher_svc/terraform.tfstate"
  }
}
//bind cipher group to SSL services
resource "citrixadc_sslservice_sslciphersuite_binding" "tf_sslservice_sslcipher_binding" {
  for_each    = var.servers
  ciphername  = data.terraform_remote_state.cipher_svc.outputs.svccipher_name
  servicename = "${each.key}-${var.service_config["service_port"]}"
  depends_on  = [module.loadbalancingservices]
}
output "sslservice_cipherbinding" {
  value = [for ciphername in citrixadc_sslservice_sslciphersuite_binding.tf_sslservice_sslcipher_binding : ciphername.id]
}
//LB HTTP vServer-services binding
resource "citrixadc_lbvserver_service_binding" "tf_bindinghttp" {
  for_each    = var.servers
  name        = citrixadc_lbvserver.tf_targetlb2.name
  servicename = "${each.key}-${var.service_config["service_port"]}"
  depends_on  = [module.loadbalancingservices]
}

