// Gratuituous comment
resource "citrixadc_cluster" "tf_cluster" {
    clid = 1
    clip = "192.168.5.55"
    hellointerval = 200
    inc = "ENABLED"

    clusternodegroup {
        name = "ng0"
        strict = "YES"
    }


    clusternode { 
        nodeid = 0
        delay = 0
        priority = 30
        endpoint = "http://192.168.5.127"
        ipaddress = "192.168.5.127"
        tunnelmode = "GRE"
        nodegroup = "ng0"

        state = "ACTIVE"
    }

    clusternodegroup {
        name = "ng1"
        strict = "YES"
    }
    clusternode { 
        nodeid = 1
        delay = 0
        priority = 31
        endpoint = "http://192.168.7.146"
        ipaddress = "192.168.7.146"
        tunnelmode = "GRE"
        nodegroup = "ng1"

        state = "ACTIVE"
    }
    clusternodegroup {
        name = "ng2"
        strict = "YES"
    }

    clusternode { 
        nodeid = 2
        delay = 0
        priority = 31
        endpoint = "http://192.168.6.9"
        ipaddress = "192.168.6.9"
        tunnelmode = "GRE"
        nodegroup = "ng2"

        state = "ACTIVE"
    }
}
