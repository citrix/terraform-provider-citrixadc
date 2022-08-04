resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
  name      = "test_policy"
  action    = "test_action"
  rule      = "client.TCP.DSTPORT.EQ(22)"
}

# -------------------- ADC CLI ----------------------------
#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector

# ---------------- NOT YET IMPLEMENTED -------------------
# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name = "test_action"
#   collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name,
#                    citrixadc_appflowcollector.tf_appflowcollector2.name, ]
#   securityinsight = "ENABLED"
#   botinsight      = "ENABLED"
#   videoanalytics  = "ENABLED"
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
#   name      = "tf2_collector"
#   ipaddress = "192.168.2.3"
#   port      = 80
# }