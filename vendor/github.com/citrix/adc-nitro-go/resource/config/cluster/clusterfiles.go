/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package cluster

/**
* Configuration for files resource.
*/
type Clusterfiles struct {
	/**
	* The directories and files to be synchronized. The available settings function as follows:
		Mode    Paths
		all           /nsconfig/ssl/
		/var/netscaler/ssl/
		/var/vpn/bookmark/
		/nsconfig/dns/
		/nsconfig/htmlinjection/
		/netscaler/htmlinjection/ens/
		/nsconfig/monitors/
		/nsconfig/nstemplates/
		/nsconfig/ssh/
		/nsconfig/rc.netscaler
		/nsconfig/resolv.conf
		/nsconfig/inetd.conf
		/nsconfig/syslog.conf
		/nsconfig/ntp.conf
		/nsconfig/httpd.conf
		/nsconfig/sshd_config
		/nsconfig/hosts
		/nsconfig/enckey
		/var/nslw.bin/etc/krb5.conf
		/var/nslw.bin/etc/krb5.keytab
		/var/lib/likewise/db/
		/var/download/
		/var/wi/tomcat/webapps/
		/var/wi/tomcat/conf/Catalina/localhost/
		/var/wi/java_home/lib/security/cacerts
		/var/wi/java_home/jre/lib/security/cacerts
		/var/netscaler/locdb/
		ssl            /nsconfig/ssl/
		/var/netscaler/ssl/
		bookmarks     /var/vpn/bookmark/
		dns                  /nsconfig/dns/
		htmlinjection    /nsconfig/htmlinjection/
		imports          /var/download/
		misc               /nsconfig/license/
		/nsconfig/rc.conf
		all_plus_misc    Includes *all* files and /nsconfig/license/ and /nsconfig/rc.conf.
		Default value: all
	*/
	Mode []string `json:"mode,omitempty"`

}
