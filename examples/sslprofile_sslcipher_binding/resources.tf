resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"

  ecccurvebindings = []
  // Don't define any bindings here
  // since we are using the explicit bindings resource

}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "HIGH"
    cipherpriority = 10
}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding2" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "LOW"
    cipherpriority = 20
}
