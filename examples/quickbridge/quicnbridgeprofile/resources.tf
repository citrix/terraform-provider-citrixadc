resource citrixadc_quicbridgeprofile demo_quicbridge {
  name             = "demo_quicbridge"
  routingalgorithm = "PLAINTEXT" # OPTIONAL
  serveridlength   = 4           # OPTIONAL
}
resource "citrixadc_lbvserver" "demo_quicbridge_lbvserver" {
  name                  = "demo_quicbridge_vserver"
  ipv46                 = "10.202.11.11"
  lbmethod              = "TOKEN"
  persistencetype       = "CUSTOMSERVERID"
  port                  = 8080
  servicetype           = "QUIC_BRIDGE"
  quicbridgeprofilename = citrixadc_quicbridgeprofile.demo_quicbridge.name
}
