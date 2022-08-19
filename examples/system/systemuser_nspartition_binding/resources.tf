resource "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
  username      = citrixadc_systemuser.user.username
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
}

resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}


resource "citrixadc_systemuser" "user" {
  username = "george"
  password = "12345"
  timeout  = 900
}
