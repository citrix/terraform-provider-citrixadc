# Lb_Vserver1
lbvserver1_name        = "Vserver-LB-HTML"
lbvserver1_ip          = "10.1.1.2"
lbvserver1_port        = 80
lbvserver1_servicetype = "HTTP"

# Lb_Vserver2
lbvserver2_name        = "Vserver-LB-Image"
lbvserver2_ip          = "10.1.1.3"
lbvserver2_port        = 80
lbvserver2_servicetype = "HTTP"

# Service 1
service1_name        = "s1"
service1_ip          = "10.1.1.4"
service1_port        = 80
service1_servicetype = "HTTP"

# Service 2
service2_name        = "s2"
service2_ip          = "10.1.1.5"
service2_port        = 80
service2_servicetype = "HTTP"

# Service 3
service3_name        = "s3"
service3_ip          = "10.1.1.6"
service3_port        = 80
service3_servicetype = "HTTP"

# Service 4
service4_name        = "s4"
service4_ip          = "10.1.1.7"
service4_port        = 80
service4_servicetype = "HTTP"

# CS Vserver
csvserver_name        = "Vserver-CS-SSL"
csvserver_ipv46       = "10.1.1.1"
csvserver_port        = 443
csvserver_servicetype = "SSL"

# CS Policy 1
cspolicy1_name = "pol1"
cspolicy1_rule  = "HTTP.REQ.URL.SUFFIX.EQ(\"cgi\")"

# CS Policy 2
cspolicy2_name = "pol2"
cspolicy2_rule  = "HTTP.REQ.URL.SUFFIX.EQ(\"asp\")"

# CS Policy 3
cspolicy3_name = "pol3"
cspolicy3_rule  = "HTTP.REQ.URL.SUFFIX.EQ(\"gif\")"

# CS Policy 4
cspolicy4_name = "pol4"
cspolicy4_rule  = "HTTP.REQ.URL.SUFFIX.EQ(\"jpeg\")"

# SSL CertKey
sslcertkey_name = "mykey"
sslcertkey_cert = "/nsconfig/ssl/ns-root.cert"
sslcertkey_key  = "/nsconfig/ssl/ns-root.key"

