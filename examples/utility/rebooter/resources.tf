resource "citrixadc_rebooter" "tf_rebooter" {
  timestamp            = timestamp()
  warm                 = false
  wait_until_reachable = true


  # Wait for 10m in total
  reachable_timeout = "10m"

  # First poll after 60s
  reachable_poll_delay = "60s"

  # Subsequent polls each 60s
  reachable_poll_interval = "60s"

  # Timeout each single HTTP request after 20s
  reachable_poll_timeout = "20s"
}
