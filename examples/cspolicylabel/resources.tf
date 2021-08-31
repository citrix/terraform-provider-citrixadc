resource "citrixadc_cspolicylabel" "tf_policylabel" {
	cspolicylabeltype = "HTTP"
	labelname = "tf_policylabel"
}