
resource "citrixadc_appflowaction" "tf_appflowaction" {
  name = "test_action"
  collectors = ["tf_collector", "tf2_collector" ]
  securityinsight = "ENABLED"
  botinsight      = "ENABLED"
  videoanalytics  = "ENABLED"
}

# -------------------- ADC CLI ----------------------------
#add appflow collector tf2_collector -IPAddress 192.168.2.3
#add appflow collector tf_collector -IPAddress 192.168.2.2

# ----------------- NOT YET IMPLEMENTED -----------------------
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