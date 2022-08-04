# Since the analyticsprofile resource is not yet available on Terraform,
# the analyticsprofile policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add analyticsprofile <name> -type <type>

resource "citrixadc_appflowaction_analyticsprofile_binding" "tf_appflowaction_analyticsprofile_binding" {
  name      = "test_action"
  analyticsprofile = "ns_analytics_global_profile"
}


# -------------------- ADC CLI ----------------------------
#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector

#---------------------NOT YET IMPLEMENTED ------------------------

# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name = "test_action"
#   collectors = [citrixadc_appflowcollector.tf_appflowcollector.name,
#                 citrixadc_appflowcollector.tf_appflowcollector2.name,]
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