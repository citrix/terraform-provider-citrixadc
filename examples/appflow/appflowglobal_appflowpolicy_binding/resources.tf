resource "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
  policyname     = "test_policy"
  globalbindtype = "SYSTEM_GLOBAL"
  type           = "REQ_OVERRIDE"
  priority       = 55
}

# -------------------- ADC CLI ----------------------------
#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector
#add appflowpolicy test_policy client.TCP.DSTPORT.EQ(22) test_action


# ---------------- NOT YET IMPLEMENTED -------------------
# resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
#   name   = "test_policy"
#   action = citrixadc_appflowaction.tf_appflowaction.name
#   rule   = "client.TCP.DSTPORT.EQ(22)"
# }
# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name            = "test_action"
#   collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
#   securityinsight = "ENABLED"
#   botinsight      = "ENABLED"
#   videoanalytics  = "ENABLED"
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }
