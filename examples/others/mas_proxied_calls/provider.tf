terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
      #version = "0.2.0" 
    }
  }
}

# provider "citrixadc" {
#   endpoint   = "http://10.102.201.246/"
#   username   = "nsroot"
#   password   = "ConfigADC#123"
#   proxied_ns = "10.102.201.73"
# }

# ADM Cloud Proxied API Configuration
# When using ADM Cloud:
# - endpoint: Your ADM Cloud URL (e.g., https://alps.adm.cloudburrito.com/)
# - username: API Client ID 
# - password: API Client Secret 
# - proxied_ns: IP of the target NetScaler instance (must be managed by ADM)
# - is_cloud: Set to true for ADM Cloud
# - do_login: Set to true to establish session
provider "citrixadc" {
  endpoint   = "https://alps.adm.cloudburrito.com/"
  username   = "4857042f-4913-44f6-b54e-3d726493a3d8"
  password   = "nhwopMqXlXdhFfAPbnTPxw=="
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
 