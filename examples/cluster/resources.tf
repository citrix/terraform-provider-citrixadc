resource "citrixadc_cluster" "tf_cluster" {
    clid = 1
    clip = "10.78.60.15"
    hellointerval = 200

    clusternode { 
        nodeid = 0
        delay = 0
        priority = 30
        endpoint = "http://10.78.60.10"
        backplane = "0/1/1"
        ipaddress = "10.78.60.10"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }

    clusternode { 
        nodeid = 1
        delay = 0
        priority = 31
        endpoint = "http://10.78.60.11"
        ipaddress = "10.78.60.11"
        backplane = "1/1/1"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }

    clusternode { 
        nodeid = 2
        delay = 0
        priority = 31
        endpoint = "http://10.78.60.12"
        ipaddress = "10.78.60.12"
        backplane = "2/1/1"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }

}
