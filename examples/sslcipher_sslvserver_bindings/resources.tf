resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
	ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
	vservername = citrixadc_lbvserver.tf_sslvserver.name
}

/*
resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding2" {
	ciphername = citrixadc_sslcipher.tfsslcipher.ciphergroupname
	vservername = citrixadc_lbvserver.tf_sslvserver.name
}
*/


resource "citrixadc_lbvserver" "tf_sslvserver" {
	name = "tf_sslvserver"
	servicetype = "SSL"
	ipv46 = "5.5.5.5"
	port = 443
}

resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tfsslcipher"

  # ciphersuitebinding is MANDATORY attribute
  # Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
    cipherpriority = 1
  }
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
    cipherpriority = 2
  }
  ciphersuitebinding {
    ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
    cipherpriority = 3
  }
}

data "citrixadc_sslcipher_sslvserver_bindings" "sslbindings" {
    ciphername = "tfsslcipher"
}



