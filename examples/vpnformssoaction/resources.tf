resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
	name = "tf_vpnformssoaction"
	actionurl = "/home"
	userfield = "username"
	passwdfield = "password"
	ssosuccessrule = "true"
	namevaluepair = "name1=value1&name2=value2"
	nvtype = "STATIC"
	responsesize = "150"
	submitmethod = "POST"
}