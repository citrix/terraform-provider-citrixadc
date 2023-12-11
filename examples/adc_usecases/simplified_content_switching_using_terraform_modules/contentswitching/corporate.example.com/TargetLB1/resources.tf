
//configure HTTP vServer
resource "citrixadc_lbvserver" "tf_targetlb1" {
  name            = var.targetlb1_config["name"]
  ipv46           = var.targetlb1_config["vip"]
  port            = var.targetlb1_config["port"]
  servicetype     = var.targetlb1_config["servicetype"]
  lbmethod        = var.targetlb1_config["lbmethod"]
  persistencetype = var.targetlb1_config["persistencetype"]
  timeout         = var.targetlb1_config["timeout"]
  redirurl        = var.targetlb1_config["redirurl"]
}
//output LBvservers "targetlb1"
output "targetlb1name" {
  value = citrixadc_lbvserver.tf_targetlb1.name
}

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
  name        = citrixadc_lbvserver.tf_targetlb1.name
  servicename = "${each.key}-${var.service_config["service_port"]}"
  depends_on  = [module.loadbalancingservices]
}

