resource "citrixadc_hanode" "local_node" {
  hanode_id     = 0       //the id of local_node is always 0
  hellointerval = 400
  deadinterval = 30
}

resource "citrixadc_hanode" "remote_node" {
  //we can only add remote add and not configure anything
  //if we want to configure this, we have to connect to the remote_node NSIP
  hanode_id = 3          
  ipaddress = "10.222.74.145"
}

