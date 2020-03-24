resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tfUnit_sslprofile-hello"

  // `ecccurvebindings` is REQUIRED attribute.
  // The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
  // To unbind all the ecccurvebindings, an empty list `[]` is to be assinged to `ecccurvebindings` attribute
  ecccurvebindings = ["P_256"]

  // `cipherbindings` attribute block is OPTIONAL
  // If not given, all the default cipher bindings will be deleted and the ones given explicitly are retained/created.
  # cipherbindings {
  #   ciphername     = "HIGH"
  #   cipherpriority = 10
  # }
}

