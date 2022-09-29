//call module "lbvs-ssl" 
module "lbvserverssl" {
  source         = "../modules/lbvs-ssl"
  sslvip_config = var.sslvip_config
  sslvsparam = var.sslvsparam
}
//configure HTTP LB vServer
resource "citrixadc_lbvserver" "tf_lbvserverhttp" {
  name            = var.httpvip_config["name"]
  ipv46           = var.httpvip_config["vip"]
  port            = var.httpvip_config["port"]
  servicetype     = var.httpvip_config["servicetype"]
  lbmethod        = var.httpvip_config["lbmethod"]
  persistencetype = var.httpvip_config["persistencetype"]
  redirurl        = var.httpvip_config["redirurl"]
}
//call module "lbservice" 
module "loadbalancingservices" {
  source = "../modules/lbservice"
  servers = var.servers
  service_config = var.service_config
}
//output module "lbservices"
output "lbservices" {
  value = module.loadbalancingservices.lbservices
}
//LB vServer-services binding
resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  for_each = var.servers
  name = module.lbvserverssl.lbvserverssl
  servicename = "${each.key}-${var.service_config["service_port"]}"
  depends_on = [module.loadbalancingservices]
}
//call module "cipher_svc"
module "ciphergroupname" {
  source = "../modules/cipher_svc"
}
output "showsvcciphergroup" {
  value = module.ciphergroupname.cipher_name
}
//bind cipher group to SSL services
resource "citrixadc_sslservice_sslciphersuite_binding" "tf_sslservice_sslcipher_binding" {
  for_each    = var.servers
  ciphername  = module.ciphergroupname.cipher_name
  servicename = "${each.key}-${var.service_config["service_port"]}"
  depends_on  = [module.loadbalancingservices]
}
//call module "cipher_vs" to be bound to SSL vServer
module "vsciphergroupname" {
  source = "../modules/cipher_vs"
}
output "showvsciphergroup" {
  value = module.vsciphergroupname.vscipher_name
}
//bind cipher group to SSL vServer
resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
  ciphername  = module.vsciphergroupname.vscipher_name
  vservername = module.lbvserverssl.lbvserverssl
  depends_on  = [module.lbvserverssl]
}
//call module "sslcertkeypair" which install ssl certificate
module "sslcertkeyname" {
  source               = "../modules/sslcertkeypair"
  ssl_certificate_path = var.ssl_certificate_path
  ssl_certkey_name     = var.ssl_certkey_name
  ssl_key_path         = var.ssl_key_path
  filename             = var.filename
}
//bind ssl certificate to SSL LB vServer
resource "citrixadc_sslvserver_sslcertkey_binding" "cert_binding" {
  certkeyname = module.sslcertkeyname.sslcertname
  vservername = module.lbvserverssl.lbvserverssl
  depends_on  = [module.lbvserverssl]
}
