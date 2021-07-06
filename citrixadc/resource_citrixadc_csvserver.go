package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

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
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				ForceNew: true,
			},
			"netprofile": &schema.Schema{
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
				ForceNew: true,
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
				ForceNew: true,
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
				ForceNew: true,
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
				ForceNew: true,
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
				ForceNew: true,
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
			"sslcertkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snisslcertkeys": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sslprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ciphers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ciphersuites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lbvserverbinding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslpolicybinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      sslpolicybindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gotopriorityexpression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invoke": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"labelname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labeltype": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policyname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In createCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var csvserverName string
	if v, ok := d.GetOk("name"); ok {
		csvserverName = v.(string)
	} else {
		csvserverName = resource.PrefixedUniqueId("tf-csvserver-")
		d.Set("name", csvserverName)
	}
	sslcertkey, sok := d.GetOk("sslcertkey")
	if sok {
		exists := client.ResourceExists(service.Sslcertkey.Type(), sslcertkey.(string))
		if !exists {
			return fmt.Errorf("[ERROR] netscaler-provider: Specified ssl cert key does not exist on netscaler!")
		}
	}

	snisslcertkeys, sniok := d.GetOk("snisslcertkeys")

	if sniok {
		exists_err := snisslcertkeysExist(snisslcertkeys, meta)
		if exists_err != nil {
			return exists_err
		}
	}

	csvserver := cs.Csvserver{
		Name:                    csvserverName,
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
		Netprofile:              d.Get("netprofile").(string),
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

	_, err := client.AddResource(service.Csvserver.Type(), csvserverName, &csvserver)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: could not add resource %s of type %s", service.Csvserver.Type(), csvserverName)
		return err
	}
	if sok { //ssl cert is specified
		binding := ssl.Sslvservercertkeybinding{
			Vservername: csvserverName,
			Certkeyname: sslcertkey.(string),
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl cert %s to csvserver %s", sslcertkey, csvserverName)
		err = client.BindResource(service.Sslvserver.Type(), csvserverName, service.Sslcertkey.Type(), sslcertkey.(string), &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to csvserver %s", sslcertkey, csvserverName)
			err2 := client.DeleteResource(service.Csvserver.Type(), csvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete csvserver %s after bind to ssl cert failed", csvserverName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete csvserver %s after bind to ssl cert failed", csvserverName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to csvserver %s", sslcertkey, csvserverName)
		}
	}

	if sniok {
		err := syncSnisslcert(d, meta, csvserverName)
		if err != nil {
			return err
		}
	}

	// Ignore for standalone
	if isTargetAdcCluster(client) {
		if err := syncCiphers(d, meta, csvserverName); err != nil {
			return err
		}
	}

	if err := syncCiphersuites(d, meta, csvserverName); err != nil {
		return err
	}

	sslprofile, spok := d.GetOk("sslprofile")
	if spok { //ssl profile is specified
		sslvserver := ssl.Sslvserver{
			Vservername: csvserverName,
			Sslprofile:  sslprofile.(string),
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl profile %s to csvserver %s", sslprofile, csvserverName)
		_, err := client.UpdateResource(service.Sslvserver.Type(), csvserverName, &sslvserver)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to csvserver %s", sslprofile, csvserverName)
			err2 := client.DeleteResource(service.Csvserver.Type(), csvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete csvserver %s after bind to ssl profile failed", csvserverName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete csvserver %s after bind to ssl profile failed", csvserverName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to csvserver %s", sslprofile, csvserverName)
		}
	}

	lbVserver, lbok := d.GetOk("lbvserverbinding")
	if lbok { //LBvserver binding is specified
		lbVserverName := lbVserver.(string)
		log.Printf("Adding binding to lbvserver %s", lbVserverName)

		bindingStruct := cs.Csvservervserverbinding{
			Name:      d.Get("name").(string),
			Lbvserver: lbVserverName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding lbvserver %s to csvserver %s", lbVserverName, csvserverName)

		err := client.BindResource(service.Csvserver.Type(), csvserverName, service.Lbvserver.Type(), lbVserverName, &bindingStruct)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind lbvserver %s to csvserver %s", lbVserverName, csvserverName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind lbvserver %s to csvserver %s", lbVserverName, csvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: lbvserver %s has been bound to csvserver %s", lbVserverName, csvserverName)
	}

	// update sslpolicy bindings
	if err := updateSslpolicyBindings(d, meta, csvserverName); err != nil {
		return err
	}

	d.SetId(csvserverName)

	err = readCsvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: we just created this csvserver but we can't read it ?? %s", csvserverName)
		return nil
	}
	return nil
}

func readCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: csvserver state %s", csvserverName)
	data, err := client.FindResource(service.Csvserver.Type(), csvserverName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: csvserver state %s", csvserverName)
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
	d.Set("stateupdate", data["stateupdate"])
	d.Set("targettype", data["targettype"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("ttl", data["ttl"])
	d.Set("vipheader", data["vipheader"])

	_, sslok := d.GetOk("sslcertkey")
	_, sniok := d.GetOk("snisslcertkeys")
	if sslok || sniok {
		if err := readSslcerts(d, meta, csvserverName); err != nil {
			return err
		}
	}

	if err := readSslpolicyBindings(d, meta, csvserverName); err != nil {
		return err
	}

	dataSsl, _ := client.FindResource(service.Sslvserver.Type(), csvserverName)
	d.Set("sslprofile", dataSsl["sslprofile"])

	// Avoid duplicate listing of ciphersuites in standalone
	if isTargetAdcCluster(client) {
		setCipherData(d, meta, csvserverName)
	}

	setCiphersuiteData(d, meta, csvserverName)

	// Read Lbvserver binding
	lbbinding, _ := client.FindResource("csvserver_lbvserver_binding", csvserverName)
	log.Printf("binding %v\n", lbbinding)
	d.Set("lbvserverbinding", lbbinding["lbvserver"])

	return nil

}

func updateCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Get("name").(string)

	csvserver := cs.Csvserver{
		Name: d.Get("name").(string),
	}
	stateChange := false
	hasChange := false
	sslcertkeyChanged := false
	sslprofileChanged := false
	snisslcertkeysChanged := false
	ciphersChanged := false
	ciphersuitesChanged := false
	lbvserverbindingChanged := false

	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG] netscaler-provider:  Appflowlog has changed for csvserver %s, starting update", csvserverName)
		csvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG] netscaler-provider:  Authentication has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationhost") {
		log.Printf("[DEBUG] netscaler-provider:  Authenticationhost has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authenticationhost = d.Get("authenticationhost").(string)
		hasChange = true
	}
	if d.HasChange("authn401") {
		log.Printf("[DEBUG] netscaler-provider:  Authn401 has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authn401 = d.Get("authn401").(string)
		hasChange = true
	}
	if d.HasChange("authnprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Authnprofile has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authnprofile = d.Get("authnprofile").(string)
		hasChange = true
	}
	if d.HasChange("authnvsname") {
		log.Printf("[DEBUG] netscaler-provider:  Authnvsname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authnvsname = d.Get("authnvsname").(string)
		hasChange = true
	}
	if d.HasChange("backupip") {
		log.Printf("[DEBUG]  netscaler-provider: Backupip has changed for csvserver %s, starting update", csvserverName)
		csvserver.Backupip = d.Get("backupip").(string)
		hasChange = true
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG] netscaler-provider:  Backupvserver has changed for csvserver %s, starting update", csvserverName)
		csvserver.Backupvserver = d.Get("backupvserver").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG] netscaler-provider:  Cacheable has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("casesensitive") {
		log.Printf("[DEBUG] netscaler-provider:  Casesensitive has changed for csvserver %s, starting update", csvserverName)
		csvserver.Casesensitive = d.Get("casesensitive").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Clttimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG] netscaler-provider:  Comment has changed for csvserver %s, starting update", csvserverName)
		csvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookiedomain") {
		log.Printf("[DEBUG]  netscaler-provider: Cookiedomain has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cookiedomain = d.Get("cookiedomain").(string)
		hasChange = true
	}
	if d.HasChange("cookietimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Cookietimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cookietimeout = d.Get("cookietimeout").(int)
		hasChange = true
	}
	if d.HasChange("dbprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Dbprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dbprofilename = d.Get("dbprofilename").(string)
		hasChange = true
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG] netscaler-provider:  Disableprimaryondown has changed for csvserver %s, starting update", csvserverName)
		csvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
		hasChange = true
	}
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Dnsprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("dnsrecordtype") {
		log.Printf("[DEBUG]  netscaler-provider: Dnsrecordtype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dnsrecordtype = d.Get("dnsrecordtype").(string)
		hasChange = true
	}
	if d.HasChange("domainname") {
		log.Printf("[DEBUG]  netscaler-provider: Domainname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Domainname = d.Get("domainname").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG] netscaler-provider:  Downstateflush has changed for csvserver %s, starting update", csvserverName)
		csvserver.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Httpprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG] netscaler-provider:  Icmpvsrresponse has changed for csvserver %s, starting update", csvserverName)
		csvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
		hasChange = true
	}
	if d.HasChange("insertvserveripport") {
		log.Printf("[DEBUG] netscaler-provider:  Insertvserveripport has changed for csvserver %s, starting update", csvserverName)
		csvserver.Insertvserveripport = d.Get("insertvserveripport").(string)
		hasChange = true
	}
	if d.HasChange("ipmask") {
		log.Printf("[DEBUG] netscaler-provider:  Ipmask has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipmask = d.Get("ipmask").(string)
		hasChange = true
	}
	if d.HasChange("ippattern") {
		log.Printf("[DEBUG] netscaler-provider:  Ippattern has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ippattern = d.Get("ippattern").(string)
		hasChange = true
	}
	if d.HasChange("ipset") {
		log.Printf("[DEBUG]  netscaler-provider: Ipset has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipset = d.Get("ipset").(string)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG] netscaler-provider:  Ipv46 has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG] netscaler-provider:  L2conn has changed for csvserver %s, starting update", csvserverName)
		csvserver.L2conn = d.Get("l2conn").(string)
		hasChange = true
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG] netscaler-provider:  Listenpolicy has changed for csvserver %s, starting update", csvserverName)
		csvserver.Listenpolicy = d.Get("listenpolicy").(string)
		hasChange = true
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG] netscaler-provider:  Listenpriority has changed for csvserver %s, starting update", csvserverName)
		csvserver.Listenpriority = d.Get("listenpriority").(int)
		hasChange = true
	}
	if d.HasChange("mssqlserverversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mssqlserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mssqlserverversion = d.Get("mssqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("mysqlcharacterset") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlcharacterset has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlcharacterset = d.Get("mysqlcharacterset").(int)
		hasChange = true
	}
	if d.HasChange("mysqlprotocolversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlprotocolversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlprotocolversion = d.Get("mysqlprotocolversion").(int)
		hasChange = true
	}
	if d.HasChange("mysqlservercapabilities") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlservercapabilities has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlservercapabilities = d.Get("mysqlservercapabilities").(int)
		hasChange = true
	}
	if d.HasChange("mysqlserverversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlserverversion = d.Get("mysqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG] netscaler-provider:  Name has changed for csvserver %s, starting update", csvserverName)
		csvserver.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Netprofile has changed for csvserver %s, starting update", csvserverName)
		csvserver.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("oracleserverversion") {
		log.Printf("[DEBUG]  netscaler-provider: Oracleserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Oracleserverversion = d.Get("oracleserverversion").(string)
		hasChange = true
	}
	if d.HasChange("persistenceid") {
		log.Printf("[DEBUG]  netscaler-provider: Persistenceid has changed for csvserver %s, starting update", csvserverName)
		csvserver.Persistenceid = d.Get("persistenceid").(int)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG] netscaler-provider:  Port has changed for csvserver %s, starting update", csvserverName)
		csvserver.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("precedence") {
		log.Printf("[DEBUG] netscaler-provider:  Precedence has changed for csvserver %s, starting update", csvserverName)
		csvserver.Precedence = d.Get("precedence").(string)
		hasChange = true
	}
	if d.HasChange("push") {
		log.Printf("[DEBUG] netscaler-provider:  Push has changed for csvserver %s, starting update", csvserverName)
		csvserver.Push = d.Get("push").(string)
		hasChange = true
	}
	if d.HasChange("pushlabel") {
		log.Printf("[DEBUG] netscaler-provider:  Pushlabel has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushlabel = d.Get("pushlabel").(string)
		hasChange = true
	}
	if d.HasChange("pushmulticlients") {
		log.Printf("[DEBUG] netscaler-provider:  Pushmulticlients has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushmulticlients = d.Get("pushmulticlients").(string)
		hasChange = true
	}
	if d.HasChange("pushvserver") {
		log.Printf("[DEBUG] netscaler-provider:  Pushvserver has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushvserver = d.Get("pushvserver").(string)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG] netscaler-provider:  Range has changed for csvserver %s, starting update", csvserverName)
		csvserver.Range = d.Get("range").(int)
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG] netscaler-provider:  Redirectportrewrite has changed for csvserver %s, starting update", csvserverName)
		csvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("redirecturl") {
		log.Printf("[DEBUG] netscaler-provider:  Redirecturl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Redirecturl = d.Get("redirecturl").(string)
		hasChange = true
	}
	if d.HasChange("rhistate") {
		log.Printf("[DEBUG]  netscaler-provider: Rhistate has changed for csvserver %s, starting update", csvserverName)
		csvserver.Rhistate = d.Get("rhistate").(string)
		hasChange = true
	}
	if d.HasChange("rtspnat") {
		log.Printf("[DEBUG] netscaler-provider:  Rtspnat has changed for csvserver %s, starting update", csvserverName)
		csvserver.Rtspnat = d.Get("rtspnat").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG] netscaler-provider:  Servicetype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitedomainttl") {
		log.Printf("[DEBUG]  netscaler-provider: Sitedomainttl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sitedomainttl = d.Get("sitedomainttl").(int)
		hasChange = true
	}
	if d.HasChange("sobackupaction") {
		log.Printf("[DEBUG] netscaler-provider:  Sobackupaction has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sobackupaction = d.Get("sobackupaction").(string)
		hasChange = true
	}
	if d.HasChange("somethod") {
		log.Printf("[DEBUG] netscaler-provider:  Somethod has changed for csvserver %s, starting update", csvserverName)
		csvserver.Somethod = d.Get("somethod").(string)
		hasChange = true
	}
	if d.HasChange("sopersistence") {
		log.Printf("[DEBUG] netscaler-provider:  Sopersistence has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sopersistence = d.Get("sopersistence").(string)
		hasChange = true
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Sopersistencetimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Sothreshold has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sothreshold = d.Get("sothreshold").(int)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] netscaler-provider:  State has changed for csvserver %s, starting update", csvserverName)
		stateChange = true
	}
	if d.HasChange("stateupdate") {
		log.Printf("[DEBUG] netscaler-provider:  Stateupdate has changed for csvserver %s, starting update", csvserverName)
		csvserver.Stateupdate = d.Get("stateupdate").(string)
		hasChange = true
	}
	if d.HasChange("targettype") {
		log.Printf("[DEBUG]  netscaler-provider: Targettype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Targettype = d.Get("targettype").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Tcpprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG] netscaler-provider:  Td has changed for csvserver %s, starting update", csvserverName)
		csvserver.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  netscaler-provider: Ttl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("vipheader") {
		log.Printf("[DEBUG] netscaler-provider:  Vipheader has changed for csvserver %s, starting update", csvserverName)
		csvserver.Vipheader = d.Get("vipheader").(string)
		hasChange = true
	}
	if d.HasChange("sslcertkey") {
		log.Printf("[DEBUG] netscaler-provider:  ssl certkey has changed for csvserver %s, starting update", csvserverName)
		sslcertkeyChanged = true
	}
	if d.HasChange("snisslcertkeys") {
		log.Printf("[DEBUG] netscaler-provider:  sni ssl certkeys has changed for lbvserver %s, starting update", csvserverName)
		snisslcertkeysChanged = true
	}
	if d.HasChange("sslprofile") {
		log.Printf("[DEBUG] netscaler-provider:  ssl profile has changed for csvserver %s, starting update", csvserverName)
		sslprofileChanged = true
	}
	if d.HasChange("ciphers") {
		log.Printf("[DEBUG] netscaler-provider:  ciphers have changed %s, starting update", csvserverName)
		ciphersChanged = true
	}
	if d.HasChange("ciphersuites") {
		log.Printf("[DEBUG] netscaler-provider:  ciphersuites have changed %s, starting update", csvserverName)
		ciphersuitesChanged = true
	}
	if d.HasChange("lbvserverbinding") {
		log.Printf("[DEBUG] netscaler-provider:  LB Vserver binding has changed for csvserver %s, starting update", csvserverName)
		lbvserverbindingChanged = true
	}

	sslcertkey := d.Get("sslcertkey")
	sslcertkeyName := sslcertkey.(string)
	if sslcertkeyChanged {
		//Binding has to be updated
		//First we unbind from cs vserver
		oldSslcertkey, _ := d.GetChange("sslcertkey")
		oldSslcertkeyName := oldSslcertkey.(string)
		if oldSslcertkeyName != "" {
			err := client.UnbindResource(service.Sslvserver.Type(), csvserverName, service.Sslcertkey.Type(), oldSslcertkeyName, "certkeyname")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding sslcertkey from csvserver %s", oldSslcertkeyName)
			}
			log.Printf("[DEBUG] netscaler-provider: sslcertkey has been unbound from csvserver for sslcertkey %s ", oldSslcertkeyName)
		}
	}
	if lbvserverbindingChanged {
		// Binding need to be updated
		// first unbind old lbvserver from csvserver
		oldLbVserver, _ := d.GetChange("lbvserverbinding")
		oldLbVserverName := oldLbVserver.(string)

		if oldLbVserverName != "" {
			err := client.UnbindResource(service.Csvserver.Type(), csvserverName, service.Lbvserver.Type(), oldLbVserverName, "lbvserver")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding lbvserver %s from csvserver %s", oldLbVserverName, csvserverName)
			}
			log.Printf("[DEBUG] netscaler-provider: lbvserver %s has been unbound from csvserver %s ", oldLbVserverName, csvserverName)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Csvserver.Type(), csvserverName, &csvserver)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error updating csvserver %s", csvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: csvserver has been updated  csvserver %s ", csvserverName)
	}

	if sslcertkeyChanged && sslcertkeyName != "" {
		//Binding has to be updated
		//rebind
		binding := ssl.Sslvservercertkeybinding{
			Vservername: csvserverName,
			Certkeyname: sslcertkeyName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl cert %s to csvserver %s", sslcertkeyName, csvserverName)
		err := client.BindResource(service.Sslvserver.Type(), csvserverName, service.Sslcertkey.Type(), sslcertkeyName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to csvserver %s", sslcertkeyName, csvserverName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to csvserver %s", sslcertkeyName, csvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: new ssl cert has been bound to csvserver  sslcertkey %s csvserver %s", sslcertkeyName, csvserverName)
	}
	newLbvserver := d.Get("lbvserverbinding")
	newLbvserverName := newLbvserver.(string)
	if lbvserverbindingChanged && newLbvserverName != "" {
		//Binding has to be updated
		//rebind
		bindingStruct := cs.Csvservervserverbinding{
			Name:      csvserverName,
			Lbvserver: newLbvserverName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding lbvserver %s to csvserver %s", newLbvserverName, csvserverName)
		err := client.BindResource(service.Csvserver.Type(), csvserverName, service.Lbvserver.Type(), newLbvserverName, &bindingStruct)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind lbvserver %s to csvserver %s", newLbvserverName, csvserverName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind lbvserver %s to csvserver %s", newLbvserverName, csvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: new lbvserver %s has been bound to csvserver %s", newLbvserverName, csvserverName)
	}
	sslprofile := d.Get("sslprofile")
	if sslprofileChanged {
		sslprofileName := sslprofile.(string)

		if sslprofileName == "" {
			sslvserver := ssl.Sslvserver{
				Vservername: csvserverName,
				Sslprofile:  "true",
			}
			err := client.ActOnResource(service.Sslvserver.Type(), &sslvserver, "unset")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding ssl profile from csvserver %s", csvserverName)
			}
		} else {
			sslvserver := ssl.Sslvserver{
				Vservername: csvserverName,
				Sslprofile:  sslprofileName,
			}
			log.Printf("[INFO] netscaler-provider:  Binding ssl profile %s to csvserver %s", sslprofileName, csvserverName)
			_, err := client.UpdateResource(service.Sslvserver.Type(), csvserverName, &sslvserver)
			if err != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to csvserver %s", sslprofileName, csvserverName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to csvserver %s", sslprofileName, csvserverName)
			}
			log.Printf("[DEBUG] netscaler-provider: new ssl profile has been bound to csvserver  sslprofile %s csvserver %s", sslprofileName, csvserverName)
		}
	}

	if snisslcertkeysChanged {
		err := syncSnisslcert(d, meta, csvserverName)
		if err != nil {
			return err
		}
	}

	// Ignore for standalone
	if ciphersChanged && isTargetAdcCluster(client) {
		if err := syncCiphers(d, meta, csvserverName); err != nil {
			return err
		}
	}

	if ciphersuitesChanged {
		if err := syncCiphersuites(d, meta, csvserverName); err != nil {
			return err
		}
	}

	if d.HasChange("sslpolicybinding") {
		if err := updateSslpolicyBindings(d, meta, csvserverName); err != nil {
			return err
		}
	}

	if stateChange {
		err := doCsvserverStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling cs vserver %s", csvserverName)
		}
	}

	return readCsvserverFunc(d, meta)
}

func deleteCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In deleteCsvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Id()
	err := client.DeleteResource(service.Csvserver.Type(), csvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func doCsvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doLbvserverStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	csvserver := cs.Csvserver{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Csvserver.Type(), csvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Csvserver.Type(), csvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
