
resource "netscaler_inat" "foo" {
  
  name = "ip4ip4"
  privateip = "192.168.2.5"
  publicip = "172.17.1.2"
}
