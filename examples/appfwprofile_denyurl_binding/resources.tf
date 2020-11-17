resource citrixadc_appfwprofile demo_appfw {
    name = "demo_appfwprofile"
    bufferoverflowaction = ["none"]
    contenttypeaction = ["none"]
    cookieconsistencyaction = ["none"]
    creditcard = ["none"]
    creditcardaction = ["none"]
    crosssitescriptingaction = ["none"]
    csrftagaction = ["none"]
    denyurlaction = ["none"]
    dynamiclearning = ["none"]
    fieldconsistencyaction = ["none"]
    fieldformataction = ["none"]
    fileuploadtypesaction = ["none"]
    inspectcontenttypes = ["none"]
    jsondosaction = ["none"]
    jsonsqlinjectionaction = ["none"]
    jsonxssaction = ["none"]
    multipleheaderaction = ["none"]
    sqlinjectionaction = ["none"]
    starturlaction = ["none"]
    type = ["HTML"]
    xmlattachmentaction = ["none"]
    xmldosaction = ["none"]
    xmlformataction = ["none"]
    xmlsoapfaultaction = ["none"]
    xmlsqlinjectionaction = ["none"]
    xmlvalidationaction = ["none"]
    xmlwsiaction = ["none"]
    xmlxssaction = ["none"]
}

resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl1 {
    name = citrixadc_appfwprofile.demo_appfw.name
    denyurl = "debug[.][^/?]*(|[?].*)$"

    # Below attributes are to be provided to maintain idempotency
    alertonly      = "OFF"
    isautodeployed = "NOTAUTODEPLOYED"
    state          = "ENABLED"
}

resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl2 {
    name = citrixadc_appfwprofile.demo_appfw.name
    denyurl = "^[^?]*/default[.]ida[?]N+"

    # Below attributes are to be provided to maintain idempotency
    alertonly      = "OFF"
    isautodeployed = "NOTAUTODEPLOYED"
    state          = "ENABLED"
}
