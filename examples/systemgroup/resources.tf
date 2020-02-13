resource "citrixadc_systemgroup" "systemgroup" {
    groupname = "testgroupname"
    timeout = 999
    promptstring = "bye>"

    cmdpolicybinding { 
        policyname = "superuser"
        priority = 100
    }

    systemusers = [
        "nestor",
        "george",
    ]
}
