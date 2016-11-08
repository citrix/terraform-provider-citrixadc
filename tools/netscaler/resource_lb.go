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
			"activeservices": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authenticationhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authn401": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authnprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authnvsname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backuppersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"backupvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bindpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bypassaaaa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cachevserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"connfailover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consolidatedlconn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consolidatedlconngbl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cookiedomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cookiename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"curstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datalength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dataoffset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dbslb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disableprimaryondown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dnsvservername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"effectivestate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groupname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gt2gb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hashlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"health": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"healththreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"homepage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icmpvsrresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"insertvserveripport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipmapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ippattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"isgslb": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"l2conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lbmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lbrrreason": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"listenpolicy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"listenpriority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"m": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"macmoderetainvlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"maxautoscalemembers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minautoscalemembers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mssqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mysqlcharacterset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mysqlprotocolversion": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mysqlservercapabilities": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mysqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"newservicerequest": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"newservicerequestincrementinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"newservicerequestunit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ngname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistencebackup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistencetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"precedence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pushlabel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pushmulticlients": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pushvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"recursionavailable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirurlflags": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rtspnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ruletype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sessionless": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skippersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sobackupaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"somethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sopersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sopersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sothreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"statechangetimemsec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"statechangetimesec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"statechangetimeseconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"thresholdvalue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tickssincelaststatechange": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"totalservices": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"v6netmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"v6persistmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
		lbvserverName = resource.PrefixedUniqueId("tf-lbvserver-")
		d.Set("name", lbvserverName)
	}
	lbvserver := lb.Lbvserver{
		Name:                               lbvserverName,
		Activeservices:                     d.Get("activeservices").(int),
		Appflowlog:                         d.Get("appflowlog").(string),
		Authentication:                     d.Get("authentication").(string),
		Authenticationhost:                 d.Get("authenticationhost").(string),
		Authn401:                           d.Get("authn401").(string),
		Authnprofile:                       d.Get("authnprofile").(string),
		Authnvsname:                        d.Get("authnvsname").(string),
		Backuppersistencetimeout:           d.Get("backuppersistencetimeout").(int),
		Backupvserver:                      d.Get("backupvserver").(string),
		Bindpoint:                          d.Get("bindpoint").(string),
		Bypassaaaa:                         d.Get("bypassaaaa").(string),
		Cacheable:                          d.Get("cacheable").(string),
		Cachevserver:                       d.Get("cachevserver").(string),
		Clttimeout:                         d.Get("clttimeout").(int),
		Comment:                            d.Get("comment").(string),
		Connfailover:                       d.Get("connfailover").(string),
		Consolidatedlconn:                  d.Get("consolidatedlconn").(string),
		Consolidatedlconngbl:               d.Get("consolidatedlconngbl").(string),
		Cookiedomain:                       d.Get("cookiedomain").(string),
		Cookiename:                         d.Get("cookiename").(string),
		Curstate:                           d.Get("curstate").(string),
		Datalength:                         d.Get("datalength").(int),
		Dataoffset:                         d.Get("dataoffset").(int),
		Dbprofilename:                      d.Get("dbprofilename").(string),
		Dbslb:                              d.Get("dbslb").(string),
		Disableprimaryondown:               d.Get("disableprimaryondown").(string),
		Dns64:                              d.Get("dns64").(string),
		Dnsvservername:                     d.Get("dnsvservername").(string),
		Domain:                             d.Get("domain").(string),
		Downstateflush:                     d.Get("downstateflush").(string),
		Effectivestate:                     d.Get("effectivestate").(string),
		Gotopriorityexpression:             d.Get("gotopriorityexpression").(string),
		Groupname:                          d.Get("groupname").(string),
		Gt2gb:                              d.Get("gt2gb").(string),
		Hashlength:                         d.Get("hashlength").(int),
		Health:                             d.Get("health").(int),
		Healththreshold:                    d.Get("healththreshold").(int),
		Homepage:                           d.Get("homepage").(string),
		Httpprofilename:                    d.Get("httpprofilename").(string),
		Icmpvsrresponse:                    d.Get("icmpvsrresponse").(string),
		Insertvserveripport:                d.Get("insertvserveripport").(string),
		Invoke:                             d.Get("invoke").(bool),
		Ipmapping:                          d.Get("ipmapping").(string),
		Ipmask:                             d.Get("ipmask").(string),
		Ippattern:                          d.Get("ippattern").(string),
		Ipv46:                              d.Get("ipv46").(string),
		Isgslb:                             d.Get("isgslb").(bool),
		L2conn:                             d.Get("l2conn").(string),
		Labelname:                          d.Get("labelname").(string),
		Labeltype:                          d.Get("labeltype").(string),
		Lbmethod:                           d.Get("lbmethod").(string),
		Lbrrreason:                         d.Get("lbrrreason").(int),
		Listenpolicy:                       d.Get("listenpolicy").(string),
		Listenpriority:                     d.Get("listenpriority").(int),
		M:                                  d.Get("m").(string),
		Macmoderetainvlan:                  d.Get("macmoderetainvlan").(string),
		Map:                                d.Get("map").(string),
		Maxautoscalemembers:                d.Get("maxautoscalemembers").(int),
		Minautoscalemembers:                d.Get("minautoscalemembers").(int),
		Mssqlserverversion:                 d.Get("mssqlserverversion").(string),
		Mysqlcharacterset:                  d.Get("mysqlcharacterset").(int),
		Mysqlprotocolversion:               d.Get("mysqlprotocolversion").(int),
		Mysqlservercapabilities:            d.Get("mysqlservercapabilities").(int),
		Mysqlserverversion:                 d.Get("mysqlserverversion").(string),
		Netmask:                            d.Get("netmask").(string),
		Netprofile:                         d.Get("netprofile").(string),
		Newname:                            d.Get("newname").(string),
		Newservicerequest:                  d.Get("newservicerequest").(int),
		Newservicerequestincrementinterval: d.Get("newservicerequestincrementinterval").(int),
		Newservicerequestunit:              d.Get("newservicerequestunit").(string),
		Ngname:                             d.Get("ngname").(string),
		Persistencebackup:                  d.Get("persistencebackup").(string),
		Persistencetype:                    d.Get("persistencetype").(string),
		Persistmask:                        d.Get("persistmask").(string),
		Policyname:                         d.Get("policyname").(string),
		Port:                               d.Get("port").(int),
		Pq:                                 d.Get("pq").(string),
		Precedence:                         d.Get("precedence").(string),
		Push:                               d.Get("push").(string),
		Pushlabel:                          d.Get("pushlabel").(string),
		Pushmulticlients:                   d.Get("pushmulticlients").(string),
		Pushvserver:                        d.Get("pushvserver").(string),
		Range:                              d.Get("range").(int),
		Recursionavailable:                 d.Get("recursionavailable").(string),
		Redirect:                           d.Get("redirect").(string),
		Redirectportrewrite:                d.Get("redirectportrewrite").(string),
		Redirurl:                           d.Get("redirurl").(string),
		Redirurlflags:                      d.Get("redirurlflags").(bool),
		Resrule:                            d.Get("resrule").(string),
		Rtspnat:                            d.Get("rtspnat").(string),
		Rule:                               d.Get("rule").(string),
		Ruletype:                           d.Get("ruletype").(int),
		Sc:                                 d.Get("sc").(string),
		Servicename:                        d.Get("servicename").(string),
		Servicetype:                        d.Get("servicetype").(string),
		Sessionless:                        d.Get("sessionless").(string),
		Skippersistency:                    d.Get("skippersistency").(string),
		Sobackupaction:                     d.Get("sobackupaction").(string),
		Somethod:                           d.Get("somethod").(string),
		Sopersistence:                      d.Get("sopersistence").(string),
		Sopersistencetimeout:               d.Get("sopersistencetimeout").(int),
		Sothreshold:                        d.Get("sothreshold").(int),
		State:                              d.Get("state").(string),
		Statechangetimemsec:                d.Get("statechangetimemsec").(int),
		Statechangetimesec:                 d.Get("statechangetimesec").(string),
		Statechangetimeseconds:             d.Get("statechangetimeseconds").(int),
		Status:                             d.Get("status").(int),
		Tcpprofilename:                     d.Get("tcpprofilename").(string),
		Td:                                 d.Get("td").(int),
		Thresholdvalue:                     d.Get("thresholdvalue").(int),
		Tickssincelaststatechange:          d.Get("tickssincelaststatechange").(int),
		Timeout:                            d.Get("timeout").(int),
		Tosid:                              d.Get("tosid").(int),
		Totalservices:                      d.Get("totalservices").(int),
		Type:                               d.Get("type").(string),
		V6netmasklen:                       d.Get("v6netmasklen").(int),
		V6persistmasklen:                   d.Get("v6persistmasklen").(int),
		Value:                              d.Get("value").(string),
		Version:                            d.Get("version").(int),
		Vipheader:                          d.Get("vipheader").(string),
		Weight:                             d.Get("weight").(int),
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
	lbvserverName := d.Id()
	log.Printf("Reading lbvserver state %s", lbvserverName)
	data, err := client.FindResource(netscaler.Lbvserver.Type(), lbvserverName)
	if err != nil {
		log.Printf("Clearing lbvserver state %s", lbvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("activeservices", data["activeservices"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("authentication", data["authentication"])
	d.Set("authenticationhost", data["authenticationhost"])
	d.Set("authn401", data["authn401"])
	d.Set("authnprofile", data["authnprofile"])
	d.Set("authnvsname", data["authnvsname"])
	d.Set("backuppersistencetimeout", data["backuppersistencetimeout"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("bindpoint", data["bindpoint"])
	d.Set("bypassaaaa", data["bypassaaaa"])
	d.Set("cacheable", data["cacheable"])
	d.Set("cachevserver", data["cachevserver"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("comment", data["comment"])
	d.Set("connfailover", data["connfailover"])
	d.Set("consolidatedlconn", data["consolidatedlconn"])
	d.Set("consolidatedlconngbl", data["consolidatedlconngbl"])
	d.Set("cookiedomain", data["cookiedomain"])
	d.Set("cookiename", data["cookiename"])
	d.Set("curstate", data["curstate"])
	d.Set("datalength", data["datalength"])
	d.Set("dataoffset", data["dataoffset"])
	d.Set("dbprofilename", data["dbprofilename"])
	d.Set("dbslb", data["dbslb"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("dns64", data["dns64"])
	d.Set("dnsvservername", data["dnsvservername"])
	d.Set("domain", data["domain"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("effectivestate", data["effectivestate"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupname", data["groupname"])
	d.Set("gt2gb", data["gt2gb"])
	d.Set("hashlength", data["hashlength"])
	d.Set("health", data["health"])
	d.Set("healththreshold", data["healththreshold"])
	d.Set("homepage", data["homepage"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("insertvserveripport", data["insertvserveripport"])
	d.Set("invoke", data["invoke"])
	d.Set("ipmapping", data["ipmapping"])
	d.Set("ipmask", data["ipmask"])
	d.Set("ippattern", data["ippattern"])
	d.Set("ipv46", data["ipv46"])
	d.Set("isgslb", data["isgslb"])
	d.Set("l2conn", data["l2conn"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("lbmethod", data["lbmethod"])
	d.Set("lbrrreason", data["lbrrreason"])
	d.Set("listenpolicy", data["listenpolicy"])
	d.Set("listenpriority", data["listenpriority"])
	d.Set("m", data["m"])
	d.Set("macmoderetainvlan", data["macmoderetainvlan"])
	d.Set("map", data["map"])
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
	d.Set("ngname", data["ngname"])
	d.Set("persistencebackup", data["persistencebackup"])
	d.Set("persistencetype", data["persistencetype"])
	d.Set("persistmask", data["persistmask"])
	d.Set("policyname", data["policyname"])
	d.Set("port", data["port"])
	d.Set("pq", data["pq"])
	d.Set("precedence", data["precedence"])
	d.Set("push", data["push"])
	d.Set("pushlabel", data["pushlabel"])
	d.Set("pushmulticlients", data["pushmulticlients"])
	d.Set("pushvserver", data["pushvserver"])
	d.Set("range", data["range"])
	d.Set("recursionavailable", data["recursionavailable"])
	d.Set("redirect", data["redirect"])
	d.Set("redirectportrewrite", data["redirectportrewrite"])
	d.Set("redirurl", data["redirurl"])
	d.Set("redirurlflags", data["redirurlflags"])
	d.Set("resrule", data["resrule"])
	d.Set("rtspnat", data["rtspnat"])
	d.Set("rule", data["rule"])
	d.Set("ruletype", data["ruletype"])
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
	d.Set("statechangetimemsec", data["statechangetimemsec"])
	d.Set("statechangetimesec", data["statechangetimesec"])
	d.Set("statechangetimeseconds", data["statechangetimeseconds"])
	d.Set("status", data["status"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("thresholdvalue", data["thresholdvalue"])
	d.Set("tickssincelaststatechange", data["tickssincelaststatechange"])
	d.Set("timeout", data["timeout"])
	d.Set("tosid", data["tosid"])
	d.Set("totalservices", data["totalservices"])
	d.Set("type", data["type"])
	d.Set("v6netmasklen", data["v6netmasklen"])
	d.Set("v6persistmasklen", data["v6persistmasklen"])
	d.Set("value", data["value"])
	d.Set("version", data["version"])
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
	if d.HasChange("vip") {
		log.Printf("[DEBUG] VIP has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ipv46 = d.Get("vip").(string)
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
