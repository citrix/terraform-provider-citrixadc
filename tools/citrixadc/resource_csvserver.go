package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/cs"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCsvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCsvserverFunc,
		Read:          readCsvserverFunc,
		Update:        updateCsvserverFunc,
		Delete:        deleteCsvserverFunc,
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
			"backupip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"casesensitive": &schema.Schema{
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
			"cookiedomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsrecordtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domainname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
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
			"ipset": &schema.Schema{
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
			"oracleserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistenceid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"precedence": &schema.Schema{
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
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirecturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rhistate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtspnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitedomainttl": &schema.Schema{
				Type:     schema.TypeInt,
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
			"stateupdate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"targettype": &schema.Schema{
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
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var csvserverName string
	if v, ok := d.GetOk("name"); ok {
		csvserverName = v.(string)
	} else {
		csvserverName = resource.PrefixedUniqueId("tf-csvserver-")
		d.Set("name", csvserverName)
	}
	csvserver := cs.Csvserver{
		Appflowlog:              d.Get("appflowlog").(string),
		Authentication:          d.Get("authentication").(string),
		Authenticationhost:      d.Get("authenticationhost").(string),
		Authn401:                d.Get("authn401").(string),
		Authnprofile:            d.Get("authnprofile").(string),
		Authnvsname:             d.Get("authnvsname").(string),
		Backupip:                d.Get("backupip").(string),
		Backupvserver:           d.Get("backupvserver").(string),
		Cacheable:               d.Get("cacheable").(string),
		Casesensitive:           d.Get("casesensitive").(string),
		Clttimeout:              d.Get("clttimeout").(int),
		Comment:                 d.Get("comment").(string),
		Cookiedomain:            d.Get("cookiedomain").(string),
		Cookietimeout:           d.Get("cookietimeout").(int),
		Dbprofilename:           d.Get("dbprofilename").(string),
		Disableprimaryondown:    d.Get("disableprimaryondown").(string),
		Dnsprofilename:          d.Get("dnsprofilename").(string),
		Dnsrecordtype:           d.Get("dnsrecordtype").(string),
		Domainname:              d.Get("domainname").(string),
		Downstateflush:          d.Get("downstateflush").(string),
		Httpprofilename:         d.Get("httpprofilename").(string),
		Icmpvsrresponse:         d.Get("icmpvsrresponse").(string),
		Insertvserveripport:     d.Get("insertvserveripport").(string),
		Ipmask:                  d.Get("ipmask").(string),
		Ippattern:               d.Get("ippattern").(string),
		Ipset:                   d.Get("ipset").(string),
		Ipv46:                   d.Get("ipv46").(string),
		L2conn:                  d.Get("l2conn").(string),
		Listenpolicy:            d.Get("listenpolicy").(string),
		Listenpriority:          d.Get("listenpriority").(int),
		Mssqlserverversion:      d.Get("mssqlserverversion").(string),
		Mysqlcharacterset:       d.Get("mysqlcharacterset").(int),
		Mysqlprotocolversion:    d.Get("mysqlprotocolversion").(int),
		Mysqlservercapabilities: d.Get("mysqlservercapabilities").(int),
		Mysqlserverversion:      d.Get("mysqlserverversion").(string),
		Name:                    d.Get("name").(string),
		Netprofile:              d.Get("netprofile").(string),
		Newname:                 d.Get("newname").(string),
		Oracleserverversion:     d.Get("oracleserverversion").(string),
		Persistenceid:           d.Get("persistenceid").(int),
		Port:                    d.Get("port").(int),
		Precedence:              d.Get("precedence").(string),
		Push:                    d.Get("push").(string),
		Pushlabel:               d.Get("pushlabel").(string),
		Pushmulticlients:        d.Get("pushmulticlients").(string),
		Pushvserver:             d.Get("pushvserver").(string),
		Range:                   d.Get("range").(int),
		Redirectportrewrite:     d.Get("redirectportrewrite").(string),
		Redirecturl:             d.Get("redirecturl").(string),
		Rhistate:                d.Get("rhistate").(string),
		Rtspnat:                 d.Get("rtspnat").(string),
		Servicetype:             d.Get("servicetype").(string),
		Sitedomainttl:           d.Get("sitedomainttl").(int),
		Sobackupaction:          d.Get("sobackupaction").(string),
		Somethod:                d.Get("somethod").(string),
		Sopersistence:           d.Get("sopersistence").(string),
		Sopersistencetimeout:    d.Get("sopersistencetimeout").(int),
		Sothreshold:             d.Get("sothreshold").(int),
		State:                   d.Get("state").(string),
		Stateupdate:             d.Get("stateupdate").(string),
		Targettype:              d.Get("targettype").(string),
		Tcpprofilename:          d.Get("tcpprofilename").(string),
		Td:                      d.Get("td").(int),
		Ttl:                     d.Get("ttl").(int),
		Vipheader:               d.Get("vipheader").(string),
	}

	_, err := client.AddResource(netscaler.Csvserver.Type(), csvserverName, &csvserver)
	if err != nil {
		return err
	}

	d.SetId(csvserverName)

	err = readCsvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this csvserver but we can't read it ?? %s", csvserverName)
		return nil
	}
	return nil
}

func readCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver state %s", csvserverName)
	data, err := client.FindResource(netscaler.Csvserver.Type(), csvserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver state %s", csvserverName)
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
	d.Set("backupip", data["backupip"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("cacheable", data["cacheable"])
	d.Set("casesensitive", data["casesensitive"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("comment", data["comment"])
	d.Set("cookiedomain", data["cookiedomain"])
	d.Set("cookietimeout", data["cookietimeout"])
	d.Set("dbprofilename", data["dbprofilename"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("dnsrecordtype", data["dnsrecordtype"])
	d.Set("domainname", data["domainname"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("insertvserveripport", data["insertvserveripport"])
	d.Set("ipmask", data["ipmask"])
	d.Set("ippattern", data["ippattern"])
	d.Set("ipset", data["ipset"])
	d.Set("ipv46", data["ipv46"])
	d.Set("l2conn", data["l2conn"])
	d.Set("listenpolicy", data["listenpolicy"])
	d.Set("listenpriority", data["listenpriority"])
	d.Set("mssqlserverversion", data["mssqlserverversion"])
	d.Set("mysqlcharacterset", data["mysqlcharacterset"])
	d.Set("mysqlprotocolversion", data["mysqlprotocolversion"])
	d.Set("mysqlservercapabilities", data["mysqlservercapabilities"])
	d.Set("mysqlserverversion", data["mysqlserverversion"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("newname", data["newname"])
	d.Set("oracleserverversion", data["oracleserverversion"])
	d.Set("persistenceid", data["persistenceid"])
	d.Set("port", data["port"])
	d.Set("precedence", data["precedence"])
	d.Set("push", data["push"])
	d.Set("pushlabel", data["pushlabel"])
	d.Set("pushmulticlients", data["pushmulticlients"])
	d.Set("pushvserver", data["pushvserver"])
	d.Set("range", data["range"])
	d.Set("redirectportrewrite", data["redirectportrewrite"])
	d.Set("redirecturl", data["redirecturl"])
	d.Set("rhistate", data["rhistate"])
	d.Set("rtspnat", data["rtspnat"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sitedomainttl", data["sitedomainttl"])
	d.Set("sobackupaction", data["sobackupaction"])
	d.Set("somethod", data["somethod"])
	d.Set("sopersistence", data["sopersistence"])
	d.Set("sopersistencetimeout", data["sopersistencetimeout"])
	d.Set("sothreshold", data["sothreshold"])
	d.Set("state", data["state"])
	d.Set("stateupdate", data["stateupdate"])
	d.Set("targettype", data["targettype"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("ttl", data["ttl"])
	d.Set("vipheader", data["vipheader"])

	return nil

}

func updateCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Get("name").(string)

	csvserver := cs.Csvserver{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for csvserver %s, starting update", csvserverName)
		csvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationhost") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationhost has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authenticationhost = d.Get("authenticationhost").(string)
		hasChange = true
	}
	if d.HasChange("authn401") {
		log.Printf("[DEBUG]  citrixadc-provider: Authn401 has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authn401 = d.Get("authn401").(string)
		hasChange = true
	}
	if d.HasChange("authnprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Authnprofile has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authnprofile = d.Get("authnprofile").(string)
		hasChange = true
	}
	if d.HasChange("authnvsname") {
		log.Printf("[DEBUG]  citrixadc-provider: Authnvsname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authnvsname = d.Get("authnvsname").(string)
		hasChange = true
	}
	if d.HasChange("backupip") {
		log.Printf("[DEBUG]  citrixadc-provider: Backupip has changed for csvserver %s, starting update", csvserverName)
		csvserver.Backupip = d.Get("backupip").(string)
		hasChange = true
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Backupvserver has changed for csvserver %s, starting update", csvserverName)
		csvserver.Backupvserver = d.Get("backupvserver").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheable has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("casesensitive") {
		log.Printf("[DEBUG]  citrixadc-provider: Casesensitive has changed for csvserver %s, starting update", csvserverName)
		csvserver.Casesensitive = d.Get("casesensitive").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clttimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for csvserver %s, starting update", csvserverName)
		csvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookiedomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookiedomain has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cookiedomain = d.Get("cookiedomain").(string)
		hasChange = true
	}
	if d.HasChange("cookietimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookietimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cookietimeout = d.Get("cookietimeout").(int)
		hasChange = true
	}
	if d.HasChange("dbprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dbprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dbprofilename = d.Get("dbprofilename").(string)
		hasChange = true
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG]  citrixadc-provider: Disableprimaryondown has changed for csvserver %s, starting update", csvserverName)
		csvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
		hasChange = true
	}
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("dnsrecordtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsrecordtype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dnsrecordtype = d.Get("dnsrecordtype").(string)
		hasChange = true
	}
	if d.HasChange("domainname") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Domainname = d.Get("domainname").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  citrixadc-provider: Downstateflush has changed for csvserver %s, starting update", csvserverName)
		csvserver.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpvsrresponse has changed for csvserver %s, starting update", csvserverName)
		csvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
		hasChange = true
	}
	if d.HasChange("insertvserveripport") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertvserveripport has changed for csvserver %s, starting update", csvserverName)
		csvserver.Insertvserveripport = d.Get("insertvserveripport").(string)
		hasChange = true
	}
	if d.HasChange("ipmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipmask has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipmask = d.Get("ipmask").(string)
		hasChange = true
	}
	if d.HasChange("ippattern") {
		log.Printf("[DEBUG]  citrixadc-provider: Ippattern has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ippattern = d.Get("ippattern").(string)
		hasChange = true
	}
	if d.HasChange("ipset") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipset has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipset = d.Get("ipset").(string)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv46 has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG]  citrixadc-provider: L2conn has changed for csvserver %s, starting update", csvserverName)
		csvserver.L2conn = d.Get("l2conn").(string)
		hasChange = true
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG]  citrixadc-provider: Listenpolicy has changed for csvserver %s, starting update", csvserverName)
		csvserver.Listenpolicy = d.Get("listenpolicy").(string)
		hasChange = true
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Listenpriority has changed for csvserver %s, starting update", csvserverName)
		csvserver.Listenpriority = d.Get("listenpriority").(int)
		hasChange = true
	}
	if d.HasChange("mssqlserverversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Mssqlserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mssqlserverversion = d.Get("mssqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("mysqlcharacterset") {
		log.Printf("[DEBUG]  citrixadc-provider: Mysqlcharacterset has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlcharacterset = d.Get("mysqlcharacterset").(int)
		hasChange = true
	}
	if d.HasChange("mysqlprotocolversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Mysqlprotocolversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlprotocolversion = d.Get("mysqlprotocolversion").(int)
		hasChange = true
	}
	if d.HasChange("mysqlservercapabilities") {
		log.Printf("[DEBUG]  citrixadc-provider: Mysqlservercapabilities has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlservercapabilities = d.Get("mysqlservercapabilities").(int)
		hasChange = true
	}
	if d.HasChange("mysqlserverversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Mysqlserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlserverversion = d.Get("mysqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for csvserver %s, starting update", csvserverName)
		csvserver.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for csvserver %s, starting update", csvserverName)
		csvserver.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("oracleserverversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Oracleserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Oracleserverversion = d.Get("oracleserverversion").(string)
		hasChange = true
	}
	if d.HasChange("persistenceid") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistenceid has changed for csvserver %s, starting update", csvserverName)
		csvserver.Persistenceid = d.Get("persistenceid").(int)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for csvserver %s, starting update", csvserverName)
		csvserver.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("precedence") {
		log.Printf("[DEBUG]  citrixadc-provider: Precedence has changed for csvserver %s, starting update", csvserverName)
		csvserver.Precedence = d.Get("precedence").(string)
		hasChange = true
	}
	if d.HasChange("push") {
		log.Printf("[DEBUG]  citrixadc-provider: Push has changed for csvserver %s, starting update", csvserverName)
		csvserver.Push = d.Get("push").(string)
		hasChange = true
	}
	if d.HasChange("pushlabel") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushlabel has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushlabel = d.Get("pushlabel").(string)
		hasChange = true
	}
	if d.HasChange("pushmulticlients") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushmulticlients has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushmulticlients = d.Get("pushmulticlients").(string)
		hasChange = true
	}
	if d.HasChange("pushvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushvserver has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushvserver = d.Get("pushvserver").(string)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG]  citrixadc-provider: Range has changed for csvserver %s, starting update", csvserverName)
		csvserver.Range = d.Get("range").(int)
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectportrewrite has changed for csvserver %s, starting update", csvserverName)
		csvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("redirecturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirecturl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Redirecturl = d.Get("redirecturl").(string)
		hasChange = true
	}
	if d.HasChange("rhistate") {
		log.Printf("[DEBUG]  citrixadc-provider: Rhistate has changed for csvserver %s, starting update", csvserverName)
		csvserver.Rhistate = d.Get("rhistate").(string)
		hasChange = true
	}
	if d.HasChange("rtspnat") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtspnat has changed for csvserver %s, starting update", csvserverName)
		csvserver.Rtspnat = d.Get("rtspnat").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitedomainttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Sitedomainttl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sitedomainttl = d.Get("sitedomainttl").(int)
		hasChange = true
	}
	if d.HasChange("sobackupaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Sobackupaction has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sobackupaction = d.Get("sobackupaction").(string)
		hasChange = true
	}
	if d.HasChange("somethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Somethod has changed for csvserver %s, starting update", csvserverName)
		csvserver.Somethod = d.Get("somethod").(string)
		hasChange = true
	}
	if d.HasChange("sopersistence") {
		log.Printf("[DEBUG]  citrixadc-provider: Sopersistence has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sopersistence = d.Get("sopersistence").(string)
		hasChange = true
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sopersistencetimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sothreshold has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sothreshold = d.Get("sothreshold").(int)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for csvserver %s, starting update", csvserverName)
		csvserver.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("stateupdate") {
		log.Printf("[DEBUG]  citrixadc-provider: Stateupdate has changed for csvserver %s, starting update", csvserverName)
		csvserver.Stateupdate = d.Get("stateupdate").(string)
		hasChange = true
	}
	if d.HasChange("targettype") {
		log.Printf("[DEBUG]  citrixadc-provider: Targettype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Targettype = d.Get("targettype").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for csvserver %s, starting update", csvserverName)
		csvserver.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("vipheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Vipheader has changed for csvserver %s, starting update", csvserverName)
		csvserver.Vipheader = d.Get("vipheader").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Csvserver.Type(), csvserverName, &csvserver)
		if err != nil {
			return fmt.Errorf("Error updating csvserver %s", csvserverName)
		}
	}
	return readCsvserverFunc(d, meta)
}

func deleteCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Id()
	err := client.DeleteResource(netscaler.Csvserver.Type(), csvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
