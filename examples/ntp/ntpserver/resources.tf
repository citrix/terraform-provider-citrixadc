#example with servername
resource "citrixadc_ntpserver" "tf_ntpserver" {
  servername          = "www.example.com"
  minpoll            = 6
  maxpoll            = 10
  preferredntpserver = "YES"

}

#example with serverip
resource "citrixadc_ntpserver" "tf_ntpserver" {
	serverip          = "10.222.74.200"
	minpoll            = 5
	maxpoll            = 9
	preferredntpserver = "NO"
  
  }