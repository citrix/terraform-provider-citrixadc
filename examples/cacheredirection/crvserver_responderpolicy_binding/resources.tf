resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_crvserver_responderpolicy_binding" "crvserver_responderpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_responderpolicy.tf_responderpolicy.name
    priority = 10
  
}