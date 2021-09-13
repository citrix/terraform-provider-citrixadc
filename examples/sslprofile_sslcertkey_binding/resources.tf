resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tfUnit_sslprofile-hello"

  // `ecccurvebindings` is REQUIRED attribute.
  // The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
  // To unbind all the ecccurvebindings, an empty list `[]` is to be assinged to `ecccurvebindings` attribute
  ecccurvebindings = ["P_256"]
  sslinterception = "ENABLED"

}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
}
  
resource "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	name = citrixadc_sslprofile.tf_sslprofile.name
	sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey 
}
