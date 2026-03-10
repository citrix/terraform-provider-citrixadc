terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

# provider "citrixadc" {
#   endpoint   = "http://10.102.201.246/"
#   username   = "nsroot"
#   password   = "ConfigADC#123dsd"
#   proxied_ns = "10.102.201.73"
# }

# NetScaler Console Cloud Cloud Proxied API Configuration
# When using NetScaler Console Cloud:
# - endpoint: Your NetScaler Console Cloud URL (e.g., https://alps.adm.cloudburrito.com/)
# - username: API Client ID 
# - password: API Client Secret 
# - proxied_ns: IP of the target NetScaler instance (must be managed by NetScaler Console)
# - is_cloud: Set to true for NetScaler Console Cloud
# - do_login: Set to true to establish session
provider "citrixadc" {
  endpoint   = "https://alps.adm.cloudburrito.com/"
  username   = "4tytuyuutf-4713-46f6-b54e-3d986493a3d8"
  password   = "nhwopMqXlXdhFfAPbnTPxw==="
  proxied_ns = "10.146.88.126"
  is_cloud   = true
  do_login   = true
}

resource "citrixadc_lbvserver" "tf_test_backup_lb" {
	name = "bkp1"
	ipv46 = "10.202.11.20"
	port = 80
	servicetype = "HTTP"
}
 