package netscaler

import (
        "github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)


func resourceNetScalerLbvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
                Create:        createLbvserverFunc,
                Read:          readLbvserverFunc,
                Update:        updateLbvserverFunc,
                Delete:        deleteLbvserverFunc,
		Schema: map[string]*schema.Schema{
                       "appflowlog": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "authentication": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "authenticationhost": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "authn401": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "authnprofile": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "authnvsname": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "backuppersistencetimeout": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "backupvserver": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "bypassaaaa": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "cacheable": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "clttimeout": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "comment": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "connfailover": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "cookiename": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "datalength": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "dataoffset": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "dbprofilename": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "dbslb": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "disableprimaryondown": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "dns64": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "downstateflush": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "hashlength": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "healththreshold": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "httpprofilename": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "icmpvsrresponse": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "insertvserveripport": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "ipmask": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "ippattern": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "ipv46": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "l2conn": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "lbmethod": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "listenpolicy": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "listenpriority": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "m": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "macmoderetainvlan": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "maxautoscalemembers": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "minautoscalemembers": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "mssqlserverversion": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "mysqlcharacterset": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "mysqlprotocolversion": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "mysqlservercapabilities": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "mysqlserverversion": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "name": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "netmask": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "netprofile": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "newname": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "newservicerequest": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "newservicerequestincrementinterval": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "newservicerequestunit": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "persistencebackup": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "persistencetype": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "persistmask": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "port": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "pq": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "push": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "pushlabel": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "pushmulticlients": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "pushvserver": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "range": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "recursionavailable": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "redirectportrewrite": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "redirurl": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "redirurlflags": &schema.Schema{
                               Type:     schema.TypeBool,
			       Optional: true,
                               Computed: true,
			},
                        "resrule": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "rtspnat": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "rule": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "sc": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "servicename": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "servicetype": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "sessionless": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "skippersistency": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "sobackupaction": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "somethod": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "sopersistence": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "sopersistencetimeout": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "sothreshold": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "state": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "tcpprofilename": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "td": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "timeout": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "tosid": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "v6netmasklen": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "v6persistmasklen": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        "vipheader": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "weight": &schema.Schema{
                               Type:     schema.TypeInt,
			       Optional: true,
                               Computed: true,
			},
                        
		},
	}
}

func createLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
        var lbvserverName string
	if v, ok := d.GetOk("name"); ok {
             lbvserverName = v.(string)
	} else {
             lbvserverName= resource.PrefixedUniqueId("tf-lbvserver-")
             d.Set("name", lbvserverName)
	}
        lbvserver := lb.Lbvserver{
                Name:            lbvserverName,
                Appflowlog:           d.Get("appflowlog").(string),
                Authentication:           d.Get("authentication").(string),
                Authenticationhost:           d.Get("authenticationhost").(string),
                Authn401:           d.Get("authn401").(string),
                Authnprofile:           d.Get("authnprofile").(string),
                Authnvsname:           d.Get("authnvsname").(string),
                Backuppersistencetimeout:           d.Get("backuppersistencetimeout").(int),
                Backupvserver:           d.Get("backupvserver").(string),
                Bypassaaaa:           d.Get("bypassaaaa").(string),
                Cacheable:           d.Get("cacheable").(string),
                Clttimeout:           d.Get("clttimeout").(int),
                Comment:           d.Get("comment").(string),
                Connfailover:           d.Get("connfailover").(string),
                Cookiename:           d.Get("cookiename").(string),
                Datalength:           d.Get("datalength").(int),
                Dataoffset:           d.Get("dataoffset").(int),
                Dbprofilename:           d.Get("dbprofilename").(string),
                Dbslb:           d.Get("dbslb").(string),
                Disableprimaryondown:           d.Get("disableprimaryondown").(string),
                Dns64:           d.Get("dns64").(string),
                Downstateflush:           d.Get("downstateflush").(string),
                Hashlength:           d.Get("hashlength").(int),
                Healththreshold:           d.Get("healththreshold").(int),
                Httpprofilename:           d.Get("httpprofilename").(string),
                Icmpvsrresponse:           d.Get("icmpvsrresponse").(string),
                Insertvserveripport:           d.Get("insertvserveripport").(string),
                Ipmask:           d.Get("ipmask").(string),
                Ippattern:           d.Get("ippattern").(string),
                Ipv46:           d.Get("ipv46").(string),
                L2conn:           d.Get("l2conn").(string),
                Lbmethod:           d.Get("lbmethod").(string),
                Listenpolicy:           d.Get("listenpolicy").(string),
                Listenpriority:           d.Get("listenpriority").(int),
                M:           d.Get("m").(string),
                Macmoderetainvlan:           d.Get("macmoderetainvlan").(string),
                Maxautoscalemembers:           d.Get("maxautoscalemembers").(int),
                Minautoscalemembers:           d.Get("minautoscalemembers").(int),
                Mssqlserverversion:           d.Get("mssqlserverversion").(string),
                Mysqlcharacterset:           d.Get("mysqlcharacterset").(int),
                Mysqlprotocolversion:           d.Get("mysqlprotocolversion").(int),
                Mysqlservercapabilities:           d.Get("mysqlservercapabilities").(int),
                Mysqlserverversion:           d.Get("mysqlserverversion").(string),
                Netmask:           d.Get("netmask").(string),
                Netprofile:           d.Get("netprofile").(string),
                Newname:           d.Get("newname").(string),
                Newservicerequest:           d.Get("newservicerequest").(int),
                Newservicerequestincrementinterval:           d.Get("newservicerequestincrementinterval").(int),
                Newservicerequestunit:           d.Get("newservicerequestunit").(string),
                Persistencebackup:           d.Get("persistencebackup").(string),
                Persistencetype:           d.Get("persistencetype").(string),
                Persistmask:           d.Get("persistmask").(string),
                Port:           d.Get("port").(int),
                Pq:           d.Get("pq").(string),
                Push:           d.Get("push").(string),
                Pushlabel:           d.Get("pushlabel").(string),
                Pushmulticlients:           d.Get("pushmulticlients").(string),
                Pushvserver:           d.Get("pushvserver").(string),
                Range:           d.Get("range").(int),
                Recursionavailable:           d.Get("recursionavailable").(string),
                Redirectportrewrite:           d.Get("redirectportrewrite").(string),
                Redirurl:           d.Get("redirurl").(string),
                Redirurlflags:           d.Get("redirurlflags").(bool),
                Resrule:           d.Get("resrule").(string),
                Rtspnat:           d.Get("rtspnat").(string),
                Rule:           d.Get("rule").(string),
                Sc:           d.Get("sc").(string),
                Servicename:           d.Get("servicename").(string),
                Servicetype:           d.Get("servicetype").(string),
                Sessionless:           d.Get("sessionless").(string),
                Skippersistency:           d.Get("skippersistency").(string),
                Sobackupaction:           d.Get("sobackupaction").(string),
                Somethod:           d.Get("somethod").(string),
                Sopersistence:           d.Get("sopersistence").(string),
                Sopersistencetimeout:           d.Get("sopersistencetimeout").(int),
                Sothreshold:           d.Get("sothreshold").(int),
                State:           d.Get("state").(string),
                Tcpprofilename:           d.Get("tcpprofilename").(string),
                Td:           d.Get("td").(int),
                Timeout:           d.Get("timeout").(int),
                Tosid:           d.Get("tosid").(int),
                V6netmasklen:           d.Get("v6netmasklen").(int),
                V6persistmasklen:           d.Get("v6persistmasklen").(int),
                Vipheader:           d.Get("vipheader").(string),
                Weight:           d.Get("weight").(int),
                
	}

        _, err := client.AddResource(netscaler.Lbvserver.Type(), lbvserverName, &lbvserver)
	if err != nil {
		return err
	}

        d.SetId(lbvserverName)
        err = readLbvserverFunc(d, meta)
	if err != nil {
                log.Printf("?? we just created this lbvserver but we can't read it ?? %s", lbvserverName)
		return nil
	}
	return nil
}

func readLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
        lbvserverName:= d.Id()
        log.Printf("Reading lbvserver state %s", lbvserverName)
        data, err := client.FindResource(netscaler.Lbvserver.Type(), lbvserverName)
	if err != nil {
        log.Printf("Clearing lbvserver state %s", lbvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
        d.Set("appflowlog", data["appflowlog"])
        d.Set("authentication", data["authentication"])
        d.Set("authenticationhost", data["authenticationhost"])
        d.Set("authn401", data["authn401"])
        d.Set("authnprofile", data["authnprofile"])
        d.Set("authnvsname", data["authnvsname"])
        d.Set("backuppersistencetimeout", data["backuppersistencetimeout"])
        d.Set("backupvserver", data["backupvserver"])
        d.Set("bypassaaaa", data["bypassaaaa"])
        d.Set("cacheable", data["cacheable"])
        d.Set("clttimeout", data["clttimeout"])
        d.Set("comment", data["comment"])
        d.Set("connfailover", data["connfailover"])
        d.Set("cookiename", data["cookiename"])
        d.Set("datalength", data["datalength"])
        d.Set("dataoffset", data["dataoffset"])
        d.Set("dbprofilename", data["dbprofilename"])
        d.Set("dbslb", data["dbslb"])
        d.Set("disableprimaryondown", data["disableprimaryondown"])
        d.Set("dns64", data["dns64"])
        d.Set("downstateflush", data["downstateflush"])
        d.Set("hashlength", data["hashlength"])
        d.Set("healththreshold", data["healththreshold"])
        d.Set("httpprofilename", data["httpprofilename"])
        d.Set("icmpvsrresponse", data["icmpvsrresponse"])
        d.Set("insertvserveripport", data["insertvserveripport"])
        d.Set("ipmask", data["ipmask"])
        d.Set("ippattern", data["ippattern"])
        d.Set("ipv46", data["ipv46"])
        d.Set("l2conn", data["l2conn"])
        d.Set("lbmethod", data["lbmethod"])
        d.Set("listenpolicy", data["listenpolicy"])
        d.Set("listenpriority", data["listenpriority"])
        d.Set("m", data["m"])
        d.Set("macmoderetainvlan", data["macmoderetainvlan"])
        d.Set("maxautoscalemembers", data["maxautoscalemembers"])
        d.Set("minautoscalemembers", data["minautoscalemembers"])
        d.Set("mssqlserverversion", data["mssqlserverversion"])
        d.Set("mysqlcharacterset", data["mysqlcharacterset"])
        d.Set("mysqlprotocolversion", data["mysqlprotocolversion"])
        d.Set("mysqlservercapabilities", data["mysqlservercapabilities"])
        d.Set("mysqlserverversion", data["mysqlserverversion"])
        d.Set("name", data["name"])
        d.Set("netmask", data["netmask"])
        d.Set("netprofile", data["netprofile"])
        d.Set("newname", data["newname"])
        d.Set("newservicerequest", data["newservicerequest"])
        d.Set("newservicerequestincrementinterval", data["newservicerequestincrementinterval"])
        d.Set("newservicerequestunit", data["newservicerequestunit"])
        d.Set("persistencebackup", data["persistencebackup"])
        d.Set("persistencetype", data["persistencetype"])
        d.Set("persistmask", data["persistmask"])
        d.Set("port", data["port"])
        d.Set("pq", data["pq"])
        d.Set("push", data["push"])
        d.Set("pushlabel", data["pushlabel"])
        d.Set("pushmulticlients", data["pushmulticlients"])
        d.Set("pushvserver", data["pushvserver"])
        d.Set("range", data["range"])
        d.Set("recursionavailable", data["recursionavailable"])
        d.Set("redirectportrewrite", data["redirectportrewrite"])
        d.Set("redirurl", data["redirurl"])
        d.Set("redirurlflags", data["redirurlflags"])
        d.Set("resrule", data["resrule"])
        d.Set("rtspnat", data["rtspnat"])
        d.Set("rule", data["rule"])
        d.Set("sc", data["sc"])
        d.Set("servicename", data["servicename"])
        d.Set("servicetype", data["servicetype"])
        d.Set("sessionless", data["sessionless"])
        d.Set("skippersistency", data["skippersistency"])
        d.Set("sobackupaction", data["sobackupaction"])
        d.Set("somethod", data["somethod"])
        d.Set("sopersistence", data["sopersistence"])
        d.Set("sopersistencetimeout", data["sopersistencetimeout"])
        d.Set("sothreshold", data["sothreshold"])
        d.Set("state", data["state"])
        d.Set("tcpprofilename", data["tcpprofilename"])
        d.Set("td", data["td"])
        d.Set("timeout", data["timeout"])
        d.Set("tosid", data["tosid"])
        d.Set("v6netmasklen", data["v6netmasklen"])
        d.Set("v6persistmasklen", data["v6persistmasklen"])
        d.Set("vipheader", data["vipheader"])
        d.Set("weight", data["weight"])
        

	return nil

}

func updateLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] In update func")
	client := meta.(*NetScalerNitroClient).client
        lbvserverName := d.Get("name").(string)

        lbvserver := lb.Lbvserver{
		Name: d.Get("name").(string),
	}
        if d.HasChange("appflowlog") {
                log.Printf("[DEBUG] Appflowlog has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Appflowlog = d.Get("appflowlog").(string)
	}
        if d.HasChange("authentication") {
                log.Printf("[DEBUG] Authentication has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Authentication = d.Get("authentication").(string)
	}
        if d.HasChange("authenticationhost") {
                log.Printf("[DEBUG] Authenticationhost has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Authenticationhost = d.Get("authenticationhost").(string)
	}
        if d.HasChange("authn401") {
                log.Printf("[DEBUG] Authn401 has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Authn401 = d.Get("authn401").(string)
	}
        if d.HasChange("authnprofile") {
                log.Printf("[DEBUG] Authnprofile has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Authnprofile = d.Get("authnprofile").(string)
	}
        if d.HasChange("authnvsname") {
                log.Printf("[DEBUG] Authnvsname has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Authnvsname = d.Get("authnvsname").(string)
	}
        if d.HasChange("backuppersistencetimeout") {
                log.Printf("[DEBUG] Backuppersistencetimeout has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Backuppersistencetimeout = d.Get("backuppersistencetimeout").(int)
	}
        if d.HasChange("backupvserver") {
                log.Printf("[DEBUG] Backupvserver has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Backupvserver = d.Get("backupvserver").(string)
	}
        if d.HasChange("bypassaaaa") {
                log.Printf("[DEBUG] Bypassaaaa has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Bypassaaaa = d.Get("bypassaaaa").(string)
	}
        if d.HasChange("cacheable") {
                log.Printf("[DEBUG] Cacheable has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Cacheable = d.Get("cacheable").(string)
	}
        if d.HasChange("clttimeout") {
                log.Printf("[DEBUG] Clttimeout has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Clttimeout = d.Get("clttimeout").(int)
	}
        if d.HasChange("comment") {
                log.Printf("[DEBUG] Comment has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Comment = d.Get("comment").(string)
	}
        if d.HasChange("connfailover") {
                log.Printf("[DEBUG] Connfailover has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Connfailover = d.Get("connfailover").(string)
	}
        if d.HasChange("cookiename") {
                log.Printf("[DEBUG] Cookiename has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Cookiename = d.Get("cookiename").(string)
	}
        if d.HasChange("datalength") {
                log.Printf("[DEBUG] Datalength has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Datalength = d.Get("datalength").(int)
	}
        if d.HasChange("dataoffset") {
                log.Printf("[DEBUG] Dataoffset has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Dataoffset = d.Get("dataoffset").(int)
	}
        if d.HasChange("dbprofilename") {
                log.Printf("[DEBUG] Dbprofilename has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Dbprofilename = d.Get("dbprofilename").(string)
	}
        if d.HasChange("dbslb") {
                log.Printf("[DEBUG] Dbslb has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Dbslb = d.Get("dbslb").(string)
	}
        if d.HasChange("disableprimaryondown") {
                log.Printf("[DEBUG] Disableprimaryondown has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
	}
        if d.HasChange("dns64") {
                log.Printf("[DEBUG] Dns64 has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Dns64 = d.Get("dns64").(string)
	}
        if d.HasChange("downstateflush") {
                log.Printf("[DEBUG] Downstateflush has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Downstateflush = d.Get("downstateflush").(string)
	}
        if d.HasChange("hashlength") {
                log.Printf("[DEBUG] Hashlength has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Hashlength = d.Get("hashlength").(int)
	}
        if d.HasChange("healththreshold") {
                log.Printf("[DEBUG] Healththreshold has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Healththreshold = d.Get("healththreshold").(int)
	}
        if d.HasChange("httpprofilename") {
                log.Printf("[DEBUG] Httpprofilename has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Httpprofilename = d.Get("httpprofilename").(string)
	}
        if d.HasChange("icmpvsrresponse") {
                log.Printf("[DEBUG] Icmpvsrresponse has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
	}
        if d.HasChange("insertvserveripport") {
                log.Printf("[DEBUG] Insertvserveripport has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Insertvserveripport = d.Get("insertvserveripport").(string)
	}
        if d.HasChange("ipmask") {
                log.Printf("[DEBUG] Ipmask has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Ipmask = d.Get("ipmask").(string)
	}
        if d.HasChange("ippattern") {
                log.Printf("[DEBUG] Ippattern has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Ippattern = d.Get("ippattern").(string)
	}
        if d.HasChange("ipv46") {
                log.Printf("[DEBUG] Ipv46 has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Ipv46 = d.Get("ipv46").(string)
	}
        if d.HasChange("l2conn") {
                log.Printf("[DEBUG] L2conn has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.L2conn = d.Get("l2conn").(string)
	}
        if d.HasChange("lbmethod") {
                log.Printf("[DEBUG] Lbmethod has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Lbmethod = d.Get("lbmethod").(string)
	}
        if d.HasChange("listenpolicy") {
                log.Printf("[DEBUG] Listenpolicy has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Listenpolicy = d.Get("listenpolicy").(string)
	}
        if d.HasChange("listenpriority") {
                log.Printf("[DEBUG] Listenpriority has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Listenpriority = d.Get("listenpriority").(int)
	}
        if d.HasChange("m") {
                log.Printf("[DEBUG] M has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.M = d.Get("m").(string)
	}
        if d.HasChange("macmoderetainvlan") {
                log.Printf("[DEBUG] Macmoderetainvlan has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Macmoderetainvlan = d.Get("macmoderetainvlan").(string)
	}
        if d.HasChange("maxautoscalemembers") {
                log.Printf("[DEBUG] Maxautoscalemembers has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Maxautoscalemembers = d.Get("maxautoscalemembers").(int)
	}
        if d.HasChange("minautoscalemembers") {
                log.Printf("[DEBUG] Minautoscalemembers has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Minautoscalemembers = d.Get("minautoscalemembers").(int)
	}
        if d.HasChange("mssqlserverversion") {
                log.Printf("[DEBUG] Mssqlserverversion has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Mssqlserverversion = d.Get("mssqlserverversion").(string)
	}
        if d.HasChange("mysqlcharacterset") {
                log.Printf("[DEBUG] Mysqlcharacterset has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Mysqlcharacterset = d.Get("mysqlcharacterset").(int)
	}
        if d.HasChange("mysqlprotocolversion") {
                log.Printf("[DEBUG] Mysqlprotocolversion has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Mysqlprotocolversion = d.Get("mysqlprotocolversion").(int)
	}
        if d.HasChange("mysqlservercapabilities") {
                log.Printf("[DEBUG] Mysqlservercapabilities has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Mysqlservercapabilities = d.Get("mysqlservercapabilities").(int)
	}
        if d.HasChange("mysqlserverversion") {
                log.Printf("[DEBUG] Mysqlserverversion has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Mysqlserverversion = d.Get("mysqlserverversion").(string)
	}
        if d.HasChange("name") {
                log.Printf("[DEBUG] Name has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Name = d.Get("name").(string)
	}
        if d.HasChange("netmask") {
                log.Printf("[DEBUG] Netmask has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Netmask = d.Get("netmask").(string)
	}
        if d.HasChange("netprofile") {
                log.Printf("[DEBUG] Netprofile has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Netprofile = d.Get("netprofile").(string)
	}
        if d.HasChange("newname") {
                log.Printf("[DEBUG] Newname has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Newname = d.Get("newname").(string)
	}
        if d.HasChange("newservicerequest") {
                log.Printf("[DEBUG] Newservicerequest has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Newservicerequest = d.Get("newservicerequest").(int)
	}
        if d.HasChange("newservicerequestincrementinterval") {
                log.Printf("[DEBUG] Newservicerequestincrementinterval has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Newservicerequestincrementinterval = d.Get("newservicerequestincrementinterval").(int)
	}
        if d.HasChange("newservicerequestunit") {
                log.Printf("[DEBUG] Newservicerequestunit has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Newservicerequestunit = d.Get("newservicerequestunit").(string)
	}
        if d.HasChange("persistencebackup") {
                log.Printf("[DEBUG] Persistencebackup has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Persistencebackup = d.Get("persistencebackup").(string)
	}
        if d.HasChange("persistencetype") {
                log.Printf("[DEBUG] Persistencetype has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Persistencetype = d.Get("persistencetype").(string)
	}
        if d.HasChange("persistmask") {
                log.Printf("[DEBUG] Persistmask has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Persistmask = d.Get("persistmask").(string)
	}
        if d.HasChange("port") {
                log.Printf("[DEBUG] Port has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Port = d.Get("port").(int)
	}
        if d.HasChange("pq") {
                log.Printf("[DEBUG] Pq has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Pq = d.Get("pq").(string)
	}
        if d.HasChange("push") {
                log.Printf("[DEBUG] Push has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Push = d.Get("push").(string)
	}
        if d.HasChange("pushlabel") {
                log.Printf("[DEBUG] Pushlabel has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Pushlabel = d.Get("pushlabel").(string)
	}
        if d.HasChange("pushmulticlients") {
                log.Printf("[DEBUG] Pushmulticlients has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Pushmulticlients = d.Get("pushmulticlients").(string)
	}
        if d.HasChange("pushvserver") {
                log.Printf("[DEBUG] Pushvserver has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Pushvserver = d.Get("pushvserver").(string)
	}
        if d.HasChange("range") {
                log.Printf("[DEBUG] Range has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Range = d.Get("range").(int)
	}
        if d.HasChange("recursionavailable") {
                log.Printf("[DEBUG] Recursionavailable has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Recursionavailable = d.Get("recursionavailable").(string)
	}
        if d.HasChange("redirectportrewrite") {
                log.Printf("[DEBUG] Redirectportrewrite has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
	}
        if d.HasChange("redirurl") {
                log.Printf("[DEBUG] Redirurl has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Redirurl = d.Get("redirurl").(string)
	}
        if d.HasChange("redirurlflags") {
                log.Printf("[DEBUG] Redirurlflags has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Redirurlflags = d.Get("redirurlflags").(bool)
	}
        if d.HasChange("resrule") {
                log.Printf("[DEBUG] Resrule has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Resrule = d.Get("resrule").(string)
	}
        if d.HasChange("rtspnat") {
                log.Printf("[DEBUG] Rtspnat has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Rtspnat = d.Get("rtspnat").(string)
	}
        if d.HasChange("rule") {
                log.Printf("[DEBUG] Rule has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Rule = d.Get("rule").(string)
	}
        if d.HasChange("sc") {
                log.Printf("[DEBUG] Sc has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Sc = d.Get("sc").(string)
	}
        if d.HasChange("servicename") {
                log.Printf("[DEBUG] Servicename has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Servicename = d.Get("servicename").(string)
	}
        if d.HasChange("servicetype") {
                log.Printf("[DEBUG] Servicetype has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Servicetype = d.Get("servicetype").(string)
	}
        if d.HasChange("sessionless") {
                log.Printf("[DEBUG] Sessionless has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Sessionless = d.Get("sessionless").(string)
	}
        if d.HasChange("skippersistency") {
                log.Printf("[DEBUG] Skippersistency has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Skippersistency = d.Get("skippersistency").(string)
	}
        if d.HasChange("sobackupaction") {
                log.Printf("[DEBUG] Sobackupaction has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Sobackupaction = d.Get("sobackupaction").(string)
	}
        if d.HasChange("somethod") {
                log.Printf("[DEBUG] Somethod has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Somethod = d.Get("somethod").(string)
	}
        if d.HasChange("sopersistence") {
                log.Printf("[DEBUG] Sopersistence has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Sopersistence = d.Get("sopersistence").(string)
	}
        if d.HasChange("sopersistencetimeout") {
                log.Printf("[DEBUG] Sopersistencetimeout has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
	}
        if d.HasChange("sothreshold") {
                log.Printf("[DEBUG] Sothreshold has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Sothreshold = d.Get("sothreshold").(int)
	}
        if d.HasChange("state") {
                log.Printf("[DEBUG] State has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.State = d.Get("state").(string)
	}
        if d.HasChange("tcpprofilename") {
                log.Printf("[DEBUG] Tcpprofilename has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
	}
        if d.HasChange("td") {
                log.Printf("[DEBUG] Td has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Td = d.Get("td").(int)
	}
        if d.HasChange("timeout") {
                log.Printf("[DEBUG] Timeout has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Timeout = d.Get("timeout").(int)
	}
        if d.HasChange("tosid") {
                log.Printf("[DEBUG] Tosid has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Tosid = d.Get("tosid").(int)
	}
        if d.HasChange("v6netmasklen") {
                log.Printf("[DEBUG] V6netmasklen has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.V6netmasklen = d.Get("v6netmasklen").(int)
	}
        if d.HasChange("v6persistmasklen") {
                log.Printf("[DEBUG] V6persistmasklen has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.V6persistmasklen = d.Get("v6persistmasklen").(int)
	}
        if d.HasChange("vipheader") {
                log.Printf("[DEBUG] Vipheader has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Vipheader = d.Get("vipheader").(string)
	}
        if d.HasChange("weight") {
                log.Printf("[DEBUG] Weight has changed for lbvserver %s, starting update", lbvserverName)
                lbvserver.Weight = d.Get("weight").(int)
	}
        

        _, err := client.UpdateResource(netscaler.Lbvserver.Type(), lbvserverName, &lbvserver)
	if err != nil {
        return fmt.Errorf("Error updating lbvserver %s", lbvserverName)
	}
        return readLbvserverFunc(d, meta)
}

func deleteLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
        lbvserverName := d.Id()
        err := client.DeleteResource(netscaler.Lbvserver.Type(), lbvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
