resource "citrixadc_sslaction" "certinsertact" {
name       = "certinsertact"
clientcert = "ENABLED"
certheader = "CERT"
}

resource "citrixadc_sslpolicy" "certinsert_pol" {
name   = "certinsert_pol"
rule   = "false"
action = citrixadc_sslaction.certinsertact.name
}

resource "citrixadc_sslpolicylabel" "ssl_pol_label" {
	labelname = "ssl_pol_label"
	type = "DATA"	
}

resource "citrixadc_sslpolicylabel_sslpolicy_binding" "demo_sslpolicylabel_sslpolicy_binding" {
	gotopriorityexpression = "END"
	invoke = true
	labelname = citrixadc_sslpolicylabel.ssl_pol_label.labelname
	labeltype = "policylabel"
	policyname = citrixadc_sslpolicy.certinsert_pol.name
	priority = 56       
	invokelabelname = "ssl_pol_label"
}
