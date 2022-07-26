resource "citrixadc_appflowpolicylabel_appflowpolicy_binding" "tf_appflowpolicylabel_appflowpolicy_binding" {
  labelname  = "tf_policylabel"
  policyname = "test_policy"
  priority   = 30
}
# -------------------- ADC CLI ----------------------------

#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector
#add appflowpolicy test_policy client.TCP.DSTPORT.EQ(22) test_action
#add appflowpolicylabel tf_policylabel -policylabeltype OTHERTCP


# ---------------- NOT YET IMPLEMENTED -------------------
# resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
#   labelname       = "tf_policylabel"
#   policylabeltype = "OTHERTCP"
# }

# resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
#   name      = "test_policy"
#   action    = citrixadc_appflowaction.tf_appflowaction.name
#   rule      = "client.TCP.DSTPORT.EQ(22)"
# }
# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name = "test_action"
#   collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
#   securityinsight = "ENABLED"
#   botinsight      = "ENABLED"
#   videoanalytics  = "ENABLED"
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }