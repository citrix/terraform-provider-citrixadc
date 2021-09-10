# Need the following cli commands since no resource yet exists
# add lb metricTable tab1
# bind metrictable tab1 metric1 1.3.6.1.4.1.5951.4.1.1.8.0
# add lb monitor mload1 LOAD

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "LOAD"
  metrictable = "tab1"
}

resource citrixadc_lbmonitor_metric_binding demo_binding1 {
 monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
 metric = "metric1"
 metricthreshold = 100
}


