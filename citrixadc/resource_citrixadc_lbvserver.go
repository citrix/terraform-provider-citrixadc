package citrixadc

import (
	"context"
	"fmt"
	"log"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLbvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbvserverFunc,
		ReadContext:   readLbvserverFunc,
		UpdateContext: updateLbvserverFunc,
		DeleteContext: deleteLbvserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"toggleorder": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quicprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probesuccessresponsecode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistavpno": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"orderthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dnsoverhttps": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apiprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adfsproxyprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationhost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authn401": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnvsname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backuplbmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backuppersistencetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"backupvserver": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bypassaaaa": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connfailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datalength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dataoffset": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dbprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dbslb": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns64": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healththreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpsredirecturl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpvsrresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertvserveripport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippattern": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipset": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv46": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2conn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpolicy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"m": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macmoderetainvlan": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxautoscalemembers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minautoscalemembers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mssqlserverversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mysqlcharacterset": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlprotocolversion": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlservercapabilities": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlserverversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newservicerequest": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newservicerequestincrementinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newservicerequestunit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oracleserverversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistencebackup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistencetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"probeport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"probeprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"push": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushlabel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushmulticlients": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"recursionavailable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectfromport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"redirectportrewrite": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirurlflags": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resrule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retainconnectionsoncluster": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rhistate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtspnat": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sessionless": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skippersistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sobackupaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"somethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistencetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sothreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: false,
				Default:  2,
			},
			"tosid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trofspersistence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"v6netmasklen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6persistmasklen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vipheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslcertkey": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snisslcertkeys": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sslprofile": {
				Type:     schema.TypeString,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply,
				Optional: true,
			},
			"quicbridgeprofilename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ciphers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ciphersuites": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sslpolicybinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      sslpolicybindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gotopriorityexpression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invoke": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"labelname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labeltype": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policyname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": {
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

func createLbvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In createLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var lbvserverName string
	if v, ok := d.GetOk("name"); ok {
		lbvserverName = v.(string)
	} else {
		lbvserverName = resource.PrefixedUniqueId("tf-lbvserver-")
		d.Set("name", lbvserverName)
	}
	sslcertkey, sok := d.GetOk("sslcertkey")
	if sok {
		exists := client.ResourceExists(service.Sslcertkey.Type(), sslcertkey.(string))
		if !exists {
			return diag.Errorf("[ERROR] netscaler-provider: Specified ssl cert key does not exist on netscaler!")
		}
	}

	snisslcertkeys, sniok := d.GetOk("snisslcertkeys")

	if sniok {
		exists_err := snisslcertkeysExist(snisslcertkeys, meta)
		if exists_err != nil {
			return diag.FromErr(exists_err)
		}
	}

	lbvserver := lb.Lbvserver{
		Name:                       lbvserverName,
		Appflowlog:                 d.Get("appflowlog").(string),
		Authentication:             d.Get("authentication").(string),
		Authenticationhost:         d.Get("authenticationhost").(string),
		Authn401:                   d.Get("authn401").(string),
		Authnprofile:               d.Get("authnprofile").(string),
		Authnvsname:                d.Get("authnvsname").(string),
		Backuplbmethod:             d.Get("backuplbmethod").(string),
		Backupvserver:              d.Get("backupvserver").(string),
		Bypassaaaa:                 d.Get("bypassaaaa").(string),
		Cacheable:                  d.Get("cacheable").(string),
		Comment:                    d.Get("comment").(string),
		Connfailover:               d.Get("connfailover").(string),
		Cookiename:                 d.Get("cookiename").(string),
		Dbprofilename:              d.Get("dbprofilename").(string),
		Dbslb:                      d.Get("dbslb").(string),
		Disableprimaryondown:       d.Get("disableprimaryondown").(string),
		Dns64:                      d.Get("dns64").(string),
		Dnsprofilename:             d.Get("dnsprofilename").(string),
		Downstateflush:             d.Get("downstateflush").(string),
		Httpprofilename:            d.Get("httpprofilename").(string),
		Httpsredirecturl:           d.Get("httpsredirecturl").(string),
		Icmpvsrresponse:            d.Get("icmpvsrresponse").(string),
		Insertvserveripport:        d.Get("insertvserveripport").(string),
		Ipmask:                     d.Get("ipmask").(string),
		Ippattern:                  d.Get("ippattern").(string),
		Ipset:                      d.Get("ipset").(string),
		Ipv46:                      d.Get("ipv46").(string),
		L2conn:                     d.Get("l2conn").(string),
		Lbmethod:                   d.Get("lbmethod").(string),
		Lbprofilename:              d.Get("lbprofilename").(string),
		Listenpolicy:               d.Get("listenpolicy").(string),
		M:                          d.Get("m").(string),
		Macmoderetainvlan:          d.Get("macmoderetainvlan").(string),
		Mssqlserverversion:         d.Get("mssqlserverversion").(string),
		Mysqlserverversion:         d.Get("mysqlserverversion").(string),
		Netmask:                    d.Get("netmask").(string),
		Netprofile:                 d.Get("netprofile").(string),
		Newservicerequestunit:      d.Get("newservicerequestunit").(string),
		Oracleserverversion:        d.Get("oracleserverversion").(string),
		Persistencebackup:          d.Get("persistencebackup").(string),
		Persistencetype:            d.Get("persistencetype").(string),
		Persistmask:                d.Get("persistmask").(string),
		Processlocal:               d.Get("processlocal").(string),
		Push:                       d.Get("push").(string),
		Pushlabel:                  d.Get("pushlabel").(string),
		Pushmulticlients:           d.Get("pushmulticlients").(string),
		Pushvserver:                d.Get("pushvserver").(string),
		Recursionavailable:         d.Get("recursionavailable").(string),
		Redirectportrewrite:        d.Get("redirectportrewrite").(string),
		Redirurl:                   d.Get("redirurl").(string),
		Redirurlflags:              d.Get("redirurlflags").(bool),
		Resrule:                    d.Get("resrule").(string),
		Retainconnectionsoncluster: d.Get("retainconnectionsoncluster").(string),
		Rhistate:                   d.Get("rhistate").(string),
		Rtspnat:                    d.Get("rtspnat").(string),
		Rule:                       d.Get("rule").(string),
		Servicename:                d.Get("servicename").(string),
		Servicetype:                d.Get("servicetype").(string),
		Sessionless:                d.Get("sessionless").(string),
		Skippersistency:            d.Get("skippersistency").(string),
		Sobackupaction:             d.Get("sobackupaction").(string),
		Somethod:                   d.Get("somethod").(string),
		Sopersistence:              d.Get("sopersistence").(string),
		State:                      d.Get("state").(string),
		Tcpprofilename:             d.Get("tcpprofilename").(string),
		Trofspersistence:           d.Get("trofspersistence").(string),
		Vipheader:                  d.Get("vipheader").(string),
		Quicbridgeprofilename:      d.Get("quicbridgeprofilename").(string),
		Probeprotocol:              d.Get("probeprotocol").(string),
		Adfsproxyprofile:           d.Get("adfsproxyprofile").(string),
		Apiprofile:                 d.Get("apiprofile").(string),
		Dnsoverhttps:               d.Get("dnsoverhttps").(string),
		Persistavpno:               toIntegerList(d.Get("persistavpno").([]interface{})),
		Probesuccessresponsecode:   d.Get("probesuccessresponsecode").(string),
		Quicprofilename:            d.Get("quicprofilename").(string),
		Toggleorder:                d.Get("toggleorder").(string),
	}
	if raw := d.GetRawConfig().GetAttr("orderthreshold"); !raw.IsNull() {
		lbvserver.Orderthreshold = intPtr(d.Get("orderthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("backuppersistencetimeout"); !raw.IsNull() {
		lbvserver.Backuppersistencetimeout = intPtr(d.Get("backuppersistencetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("clttimeout"); !raw.IsNull() {
		lbvserver.Clttimeout = intPtr(d.Get("clttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("datalength"); !raw.IsNull() {
		lbvserver.Datalength = intPtr(d.Get("datalength").(int))
	}
	if raw := d.GetRawConfig().GetAttr("dataoffset"); !raw.IsNull() {
		lbvserver.Dataoffset = intPtr(d.Get("dataoffset").(int))
	}
	if raw := d.GetRawConfig().GetAttr("hashlength"); !raw.IsNull() {
		lbvserver.Hashlength = intPtr(d.Get("hashlength").(int))
	}
	if raw := d.GetRawConfig().GetAttr("healththreshold"); !raw.IsNull() {
		lbvserver.Healththreshold = intPtr(d.Get("healththreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("listenpriority"); !raw.IsNull() {
		lbvserver.Listenpriority = intPtr(d.Get("listenpriority").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxautoscalemembers"); !raw.IsNull() {
		lbvserver.Maxautoscalemembers = intPtr(d.Get("maxautoscalemembers").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minautoscalemembers"); !raw.IsNull() {
		lbvserver.Minautoscalemembers = intPtr(d.Get("minautoscalemembers").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mysqlcharacterset"); !raw.IsNull() {
		lbvserver.Mysqlcharacterset = intPtr(d.Get("mysqlcharacterset").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mysqlprotocolversion"); !raw.IsNull() {
		lbvserver.Mysqlprotocolversion = intPtr(d.Get("mysqlprotocolversion").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mysqlservercapabilities"); !raw.IsNull() {
		lbvserver.Mysqlservercapabilities = intPtr(d.Get("mysqlservercapabilities").(int))
	}
	if raw := d.GetRawConfig().GetAttr("newservicerequest"); !raw.IsNull() {
		lbvserver.Newservicerequest = intPtr(d.Get("newservicerequest").(int))
	}
	if raw := d.GetRawConfig().GetAttr("newservicerequestincrementinterval"); !raw.IsNull() {
		lbvserver.Newservicerequestincrementinterval = intPtr(d.Get("newservicerequestincrementinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		lbvserver.Port = intPtr(d.Get("port").(int))
	}
	if raw := d.GetRawConfig().GetAttr("range"); !raw.IsNull() {
		lbvserver.Range = intPtr(d.Get("range").(int))
	}
	if raw := d.GetRawConfig().GetAttr("redirectfromport"); !raw.IsNull() {
		lbvserver.Redirectfromport = intPtr(d.Get("redirectfromport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sopersistencetimeout"); !raw.IsNull() {
		lbvserver.Sopersistencetimeout = intPtr(d.Get("sopersistencetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sothreshold"); !raw.IsNull() {
		lbvserver.Sothreshold = intPtr(d.Get("sothreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		lbvserver.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("timeout"); !raw.IsNull() {
		lbvserver.Timeout = intPtr(d.Get("timeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tosid"); !raw.IsNull() {
		lbvserver.Tosid = intPtr(d.Get("tosid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("v6netmasklen"); !raw.IsNull() {
		lbvserver.V6netmasklen = intPtr(d.Get("v6netmasklen").(int))
	}
	if raw := d.GetRawConfig().GetAttr("v6persistmasklen"); !raw.IsNull() {
		lbvserver.V6persistmasklen = intPtr(d.Get("v6persistmasklen").(int))
	}
	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		lbvserver.Weight = intPtr(d.Get("weight").(int))
	}
	if raw := d.GetRawConfig().GetAttr("probeport"); !raw.IsNull() {
		lbvserver.Probeport = intPtr(d.Get("probeport").(int))
	}

	_, err := client.AddResource(service.Lbvserver.Type(), lbvserverName, &lbvserver)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: could not add resource %s of type %s", service.Lbvserver.Type(), lbvserverName)
		return diag.FromErr(err)
	}
	if sok { //ssl cert is specified
		binding := ssl.Sslvservercertkeybinding{
			Vservername: lbvserverName,
			Certkeyname: sslcertkey.(string),
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl cert %s to lbvserver %s", sslcertkey, lbvserverName)
		err = client.BindResource(service.Sslvserver.Type(), lbvserverName, service.Sslcertkey.Type(), sslcertkey.(string), &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkey, lbvserverName)
			err2 := client.DeleteResource(service.Lbvserver.Type(), lbvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete lbvserver %s after bind to ssl cert failed", lbvserverName)
				return diag.Errorf("[ERROR] netscaler-provider:  Failed to delete lbvserver %s after bind to ssl cert failed", lbvserverName)
			}
			return diag.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkey, lbvserverName)
		}
	}

	if sniok {
		err := syncSnisslcert(d, meta, lbvserverName)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Ignore for standalone
	if isTargetAdcCluster(client) {
		if err := syncCiphers(d, meta, lbvserverName); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := syncCiphersuites(d, meta, lbvserverName); err != nil {
		return diag.FromErr(err)
	}

	sslprofile, spok := d.GetOk("sslprofile")
	if spok { //ssl profile is specified
		sslvserver := ssl.Sslvserver{
			Vservername: lbvserverName,
			Sslprofile:  sslprofile.(string),
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl profile %s to lbvserver %s", sslprofile, lbvserverName)
		_, err := client.UpdateResource(service.Sslvserver.Type(), lbvserverName, &sslvserver)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to lbvserver %s", sslprofile, lbvserverName)
			err2 := client.DeleteResource(service.Lbvserver.Type(), lbvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete lbvserver %s after bind to ssl profile failed", lbvserverName)
				return diag.Errorf("[ERROR] netscaler-provider:  Failed to delete lbvserver %s after bind to ssl profile failed", lbvserverName)
			}
			return diag.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to lbvserver %s", sslprofile, lbvserverName)
		}
	}

	// update sslpolicy bindings
	if err := updateSslpolicyBindings(d, meta, lbvserverName); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lbvserverName)

	return readLbvserverFunc(ctx, d, meta)
}

func readLbvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading lbvserver state %s", lbvserverName)
	data, err := client.FindResource(service.Lbvserver.Type(), lbvserverName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing lbvserver state %s", lbvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("toggleorder", data["toggleorder"])
	d.Set("quicprofilename", data["quicprofilename"])
	d.Set("probesuccessresponsecode", data["probesuccessresponsecode"])
	d.Set("persistavpno", data["persistavpno"])
	setToInt("orderthreshold", d, data["orderthreshold"])
	d.Set("dnsoverhttps", data["dnsoverhttps"])
	d.Set("apiprofile", data["apiprofile"])
	d.Set("adfsproxyprofile", data["adfsproxyprofile"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("authentication", data["authentication"])
	d.Set("authenticationhost", data["authenticationhost"])
	d.Set("authn401", data["authn401"])
	d.Set("authnprofile", data["authnprofile"])
	d.Set("authnvsname", data["authnvsname"])
	d.Set("backuplbmethod", data["backuplbmethod"])
	setToInt("backuppersistencetimeout", d, data["backuppersistencetimeout"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("bypassaaaa", data["bypassaaaa"])
	d.Set("cacheable", data["cacheable"])
	setToInt("clttimeout", d, data["clttimeout"])
	d.Set("comment", data["comment"])
	d.Set("connfailover", data["connfailover"])
	d.Set("cookiename", data["cookiename"])
	setToInt("datalength", d, data["datalength"])
	setToInt("dataoffset", d, data["dataoffset"])
	d.Set("dbprofilename", data["dbprofilename"])
	d.Set("dbslb", data["dbslb"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("dns64", data["dns64"])
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("downstateflush", data["downstateflush"])
	setToInt("hashlength", d, data["hashlength"])
	setToInt("healththreshold", d, data["healththreshold"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("httpsredirecturl", data["httpsredirecturl"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("insertvserveripport", data["insertvserveripport"])
	d.Set("ipmask", data["ipmask"])
	d.Set("ippattern", data["ippattern"])
	d.Set("ipset", data["ipset"])
	d.Set("ipv46", data["ipv46"])
	d.Set("l2conn", data["l2conn"])
	d.Set("lbmethod", data["lbmethod"])
	d.Set("lbprofilename", data["lbprofilename"])
	d.Set("listenpolicy", data["listenpolicy"])
	setToInt("listenpriority", d, data["listenpriority"])
	d.Set("m", data["m"])
	d.Set("macmoderetainvlan", data["macmoderetainvlan"])
	setToInt("maxautoscalemembers", d, data["maxautoscalemembers"])
	setToInt("minautoscalemembers", d, data["minautoscalemembers"])
	d.Set("mssqlserverversion", data["mssqlserverversion"])
	setToInt("mysqlcharacterset", d, data["mysqlcharacterset"])
	setToInt("mysqlprotocolversion", d, data["mysqlprotocolversion"])
	setToInt("mysqlservercapabilities", d, data["mysqlservercapabilities"])
	d.Set("mysqlserverversion", data["mysqlserverversion"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("netprofile", data["netprofile"])
	setToInt("newservicerequest", d, data["newservicerequest"])
	setToInt("newservicerequestincrementinterval", d, data["newservicerequestincrementinterval"])
	d.Set("newservicerequestunit", data["newservicerequestunit"])
	d.Set("oracleserverversion", data["oracleserverversion"])
	d.Set("persistencebackup", data["persistencebackup"])
	d.Set("persistencetype", data["persistencetype"])
	d.Set("persistmask", data["persistmask"])
	setToInt("port", d, data["port"])
	d.Set("processlocal", data["processlocal"])
	d.Set("push", data["push"])
	d.Set("pushlabel", data["pushlabel"])
	d.Set("pushmulticlients", data["pushmulticlients"])
	d.Set("pushvserver", data["pushvserver"])
	setToInt("range", d, data["range"])
	d.Set("recursionavailable", data["recursionavailable"])
	setToInt("redirectfromport", d, data["redirectfromport"])
	d.Set("redirectportrewrite", data["redirectportrewrite"])
	d.Set("redirurl", data["redirurl"])
	d.Set("redirurlflags", data["redirurlflags"])
	d.Set("resrule", data["resrule"])
	d.Set("retainconnectionsoncluster", data["retainconnectionsoncluster"])
	d.Set("rhistate", data["rhistate"])
	d.Set("rtspnat", data["rtspnat"])
	d.Set("rule", data["rule"])
	d.Set("servicename", data["servicename"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sessionless", data["sessionless"])
	d.Set("skippersistency", data["skippersistency"])
	d.Set("sobackupaction", data["sobackupaction"])
	d.Set("somethod", data["somethod"])
	d.Set("sopersistence", data["sopersistence"])
	setToInt("sopersistencetimeout", d, data["sopersistencetimeout"])
	setToInt("sothreshold", d, data["sothreshold"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	setToInt("td", d, data["td"])
	setToInt("timeout", d, data["timeout"])
	setToInt("tosid", d, data["tosid"])
	d.Set("trofspersistence", data["trofspersistence"])
	setToInt("v6netmasklen", d, data["v6netmasklen"])
	setToInt("v6persistmasklen", d, data["v6persistmasklen"])
	d.Set("vipheader", data["vipheader"])
	setToInt("weight", d, data["weight"])
	d.Set("quicbridgeprofilename", data["quicbridgeprofilename"])
	setToInt("probeport", d, data["probeport"])
	d.Set("probeprotocol", data["probeprotocol"])

	_, sslok := d.GetOk("sslcertkey")
	_, sniok := d.GetOk("snisslcertkeys")
	if sslok || sniok {
		if err := readSslcerts(d, meta, lbvserverName); err != nil {
			return diag.FromErr(err)
		}
	}
	// Set state according to curstate
	if data["curstate"] == "OUT OF SERVICE" {
		d.Set("state", "DISABLED")
	} else {
		d.Set("state", "ENABLED")
	}

	if err := readSslpolicyBindings(d, meta, lbvserverName); err != nil {
		return diag.FromErr(err)
	}

	dataSsl, _ := client.FindResource(service.Sslvserver.Type(), lbvserverName)
	d.Set("sslprofile", dataSsl["sslprofile"])

	// Avoid duplicate listing of ciphersuites in standalone
	if isTargetAdcCluster(client) {
		setCipherData(d, meta, lbvserverName)
	}
	setCiphersuiteData(d, meta, lbvserverName)

	return nil

}

func updateLbvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In updateLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Get("name").(string)

	lbvserver := lb.Lbvserver{
		Name: d.Get("name").(string),
	}
	stateChange := false
	hasChange := false
	if d.HasChange("toggleorder") {
		log.Printf("[DEBUG]  citrixadc-provider: Toggleorder has changed for lbvserver, starting update")
		lbvserver.Toggleorder = d.Get("toggleorder").(string)
		hasChange = true
	}
	if d.HasChange("quicprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Quicprofilename has changed for lbvserver, starting update")
		lbvserver.Quicprofilename = d.Get("quicprofilename").(string)
		hasChange = true
	}
	if d.HasChange("probesuccessresponsecode") {
		log.Printf("[DEBUG]  citrixadc-provider: Probesuccessresponsecode has changed for lbvserver, starting update")
		lbvserver.Probesuccessresponsecode = d.Get("probesuccessresponsecode").(string)
		hasChange = true
	}
	if d.HasChange("persistavpno") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistavpno has changed for lbvserver, starting update")
		lbvserver.Persistavpno = toIntegerList(d.Get("persistavpno").([]interface{}))
		hasChange = true
	}
	if d.HasChange("orderthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Orderthreshold has changed for lbvserver, starting update")
		lbvserver.Orderthreshold = intPtr(d.Get("orderthreshold").(int))
		hasChange = true
	}
	if d.HasChange("dnsoverhttps") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsoverhttps has changed for lbvserver, starting update")
		lbvserver.Dnsoverhttps = d.Get("dnsoverhttps").(string)
		hasChange = true
	}
	if d.HasChange("apiprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Apiprofile has changed for lbvserver, starting update")
		lbvserver.Apiprofile = d.Get("apiprofile").(string)
		hasChange = true
	}
	if d.HasChange("adfsproxyprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Adfsproxyprofile has changed for lbvserver, starting update")
		lbvserver.Adfsproxyprofile = d.Get("adfsproxyprofile").(string)
		hasChange = true
	}
	sslcertkeyChanged := false
	sslprofileChanged := false
	snisslcertkeysChanged := false
	ciphersChanged := false
	ciphersuitesChanged := false

	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG] netscaler-provider:  Appflowlog has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG] netscaler-provider:  Authentication has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationhost") {
		log.Printf("[DEBUG] netscaler-provider:  Authenticationhost has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authenticationhost = d.Get("authenticationhost").(string)
		hasChange = true
	}
	if d.HasChange("authn401") {
		log.Printf("[DEBUG] netscaler-provider:  Authn401 has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authn401 = d.Get("authn401").(string)
		hasChange = true
	}
	if d.HasChange("authnprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Authnprofile has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authnprofile = d.Get("authnprofile").(string)
		hasChange = true
	}
	if d.HasChange("authnvsname") {
		log.Printf("[DEBUG] netscaler-provider:  Authnvsname has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authnvsname = d.Get("authnvsname").(string)
		hasChange = true
	}
	if d.HasChange("backuplbmethod") {
		log.Printf("[DEBUG]  netscaler-provider: Backuplbmethod has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Backuplbmethod = d.Get("backuplbmethod").(string)
		hasChange = true
	}
	if d.HasChange("backuppersistencetimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Backuppersistencetimeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Backuppersistencetimeout = intPtr(d.Get("backuppersistencetimeout").(int))
		hasChange = true
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG] netscaler-provider:  Backupvserver has changed for lbvserver %s, starting update", lbvserverName)
		oldBackupvserver, newBackupvserver := d.GetChange("backupvserver")
		oldBackupvserverStr := oldBackupvserver.(string)
		newBackupvserverStr := newBackupvserver.(string)

		if oldBackupvserverStr != "" && newBackupvserverStr == "" {
			// Changed from a value to empty - need to unset
			log.Printf("[DEBUG] netscaler-provider:  Unsetting backupvserver for lbvserver %s", lbvserverName)
			lbvserverUnset := lb.Lbvserver{
				Name:          lbvserverName,
				Backupvserver: "true",
			}
			err := client.ActOnResource(service.Lbvserver.Type(), &lbvserverUnset, "unset")
			if err != nil {
				return diag.Errorf("[ERROR] netscaler-provider: Error unsetting backupvserver from lbvserver %s", lbvserverName)
			}
			log.Printf("[DEBUG] netscaler-provider: backupvserver has been unset from lbvserver %s", lbvserverName)
		} else {
			// Normal update with a new value
			lbvserver.Backupvserver = newBackupvserverStr
			hasChange = true
		}
	}

	if d.HasChange("bypassaaaa") {
		log.Printf("[DEBUG] netscaler-provider:  Bypassaaaa has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Bypassaaaa = d.Get("bypassaaaa").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG] netscaler-provider:  Cacheable has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Clttimeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Clttimeout = intPtr(d.Get("clttimeout").(int))
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG] netscaler-provider:  Comment has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG] netscaler-provider:  Connfailover has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("cookiename") {
		log.Printf("[DEBUG] netscaler-provider:  Cookiename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Cookiename = d.Get("cookiename").(string)
		hasChange = true
	}
	if d.HasChange("datalength") {
		log.Printf("[DEBUG] netscaler-provider:  Datalength has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Datalength = intPtr(d.Get("datalength").(int))
		hasChange = true
	}
	if d.HasChange("dataoffset") {
		log.Printf("[DEBUG] netscaler-provider:  Dataoffset has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dataoffset = intPtr(d.Get("dataoffset").(int))
		hasChange = true
	}
	if d.HasChange("dbprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Dbprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dbprofilename = d.Get("dbprofilename").(string)
		hasChange = true
	}
	if d.HasChange("dbslb") {
		log.Printf("[DEBUG] netscaler-provider:  Dbslb has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dbslb = d.Get("dbslb").(string)
		hasChange = true
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG] netscaler-provider:  Disableprimaryondown has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
		hasChange = true
	}
	if d.HasChange("dns64") {
		log.Printf("[DEBUG] netscaler-provider:  Dns64 has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dns64 = d.Get("dns64").(string)
		hasChange = true
	}
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Dnsprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG] netscaler-provider:  Downstateflush has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("hashlength") {
		log.Printf("[DEBUG] netscaler-provider:  Hashlength has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Hashlength = intPtr(d.Get("hashlength").(int))
		hasChange = true
	}
	if d.HasChange("healththreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Healththreshold has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Healththreshold = intPtr(d.Get("healththreshold").(int))
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Httpprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("httpsredirecturl") {
		log.Printf("[DEBUG]  netscaler-provider: Httpsredirecturl has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Httpsredirecturl = d.Get("httpsredirecturl").(string)
		hasChange = true
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG] netscaler-provider:  Icmpvsrresponse has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
		hasChange = true
	}
	if d.HasChange("insertvserveripport") {
		log.Printf("[DEBUG] netscaler-provider:  Insertvserveripport has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Insertvserveripport = d.Get("insertvserveripport").(string)
		hasChange = true
	}
	if d.HasChange("ipmask") {
		log.Printf("[DEBUG] netscaler-provider:  Ipmask has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ipmask = d.Get("ipmask").(string)
		hasChange = true
	}
	if d.HasChange("ippattern") {
		log.Printf("[DEBUG] netscaler-provider:  Ippattern has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ippattern = d.Get("ippattern").(string)
		hasChange = true
	}
	if d.HasChange("ipset") {
		log.Printf("[DEBUG]  netscaler-provider: Ipset has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ipset = d.Get("ipset").(string)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG] netscaler-provider:  Ipv46 has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG] netscaler-provider:  L2conn has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.L2conn = d.Get("l2conn").(string)
		hasChange = true
	}
	if d.HasChange("lbmethod") {
		log.Printf("[DEBUG] netscaler-provider:  Lbmethod has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Lbmethod = d.Get("lbmethod").(string)
		hasChange = true
	}
	if d.HasChange("lbprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Lbprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Lbprofilename = d.Get("lbprofilename").(string)
		hasChange = true
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG] netscaler-provider:  Listenpolicy has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Listenpolicy = d.Get("listenpolicy").(string)
		hasChange = true
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG] netscaler-provider:  Listenpriority has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Listenpriority = intPtr(d.Get("listenpriority").(int))
		hasChange = true
	}
	if d.HasChange("m") {
		log.Printf("[DEBUG] netscaler-provider:  M has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.M = d.Get("m").(string)
		hasChange = true
	}
	if d.HasChange("macmoderetainvlan") {
		log.Printf("[DEBUG] netscaler-provider:  Macmoderetainvlan has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Macmoderetainvlan = d.Get("macmoderetainvlan").(string)
		hasChange = true
	}
	if d.HasChange("maxautoscalemembers") {
		log.Printf("[DEBUG] netscaler-provider:  Maxautoscalemembers has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Maxautoscalemembers = intPtr(d.Get("maxautoscalemembers").(int))
		hasChange = true
	}
	if d.HasChange("minautoscalemembers") {
		log.Printf("[DEBUG] netscaler-provider:  Minautoscalemembers has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Minautoscalemembers = intPtr(d.Get("minautoscalemembers").(int))
		hasChange = true
	}
	if d.HasChange("mssqlserverversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mssqlserverversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mssqlserverversion = d.Get("mssqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("mysqlcharacterset") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlcharacterset has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlcharacterset = intPtr(d.Get("mysqlcharacterset").(int))
		hasChange = true
	}
	if d.HasChange("mysqlprotocolversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlprotocolversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlprotocolversion = intPtr(d.Get("mysqlprotocolversion").(int))
		hasChange = true
	}
	if d.HasChange("mysqlservercapabilities") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlservercapabilities has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlservercapabilities = intPtr(d.Get("mysqlservercapabilities").(int))
		hasChange = true
	}
	if d.HasChange("mysqlserverversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlserverversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlserverversion = d.Get("mysqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG] netscaler-provider:  Name has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG] netscaler-provider:  Netmask has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Netprofile has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("newservicerequest") {
		log.Printf("[DEBUG] netscaler-provider:  Newservicerequest has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newservicerequest = intPtr(d.Get("newservicerequest").(int))
		hasChange = true
	}
	if d.HasChange("newservicerequestincrementinterval") {
		log.Printf("[DEBUG] netscaler-provider:  Newservicerequestincrementinterval has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newservicerequestincrementinterval = intPtr(d.Get("newservicerequestincrementinterval").(int))
		hasChange = true
	}
	if d.HasChange("newservicerequestunit") {
		log.Printf("[DEBUG] netscaler-provider:  Newservicerequestunit has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newservicerequestunit = d.Get("newservicerequestunit").(string)
		hasChange = true
	}
	if d.HasChange("oracleserverversion") {
		log.Printf("[DEBUG]  netscaler-provider: Oracleserverversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Oracleserverversion = d.Get("oracleserverversion").(string)
		hasChange = true
	}
	if d.HasChange("persistencebackup") {
		log.Printf("[DEBUG] netscaler-provider:  Persistencebackup has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Persistencebackup = d.Get("persistencebackup").(string)
		hasChange = true
	}
	if d.HasChange("persistencetype") {
		log.Printf("[DEBUG] netscaler-provider:  Persistencetype has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Persistencetype = d.Get("persistencetype").(string)
		hasChange = true
	}
	if d.HasChange("persistmask") {
		log.Printf("[DEBUG] netscaler-provider:  Persistmask has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Persistmask = d.Get("persistmask").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG] netscaler-provider:  Port has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}
	if d.HasChange("probeport") {
		log.Printf("[DEBUG] netscaler-provider:  probeport has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Probeport = intPtr(d.Get("probeport").(int))
		hasChange = true
	}
	if d.HasChange("probeprotocol") {
		log.Printf("[DEBUG] netscaler-provider:  probeprotocol has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Probeprotocol = d.Get("probeprotocol").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  netscaler-provider: Processlocal has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("push") {
		log.Printf("[DEBUG] netscaler-provider:  Push has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Push = d.Get("push").(string)
		hasChange = true
	}
	if d.HasChange("pushlabel") {
		log.Printf("[DEBUG] netscaler-provider:  Pushlabel has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pushlabel = d.Get("pushlabel").(string)
		hasChange = true
	}
	if d.HasChange("pushmulticlients") {
		log.Printf("[DEBUG] netscaler-provider:  Pushmulticlients has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pushmulticlients = d.Get("pushmulticlients").(string)
		hasChange = true
	}
	if d.HasChange("pushvserver") {
		log.Printf("[DEBUG] netscaler-provider:  Pushvserver has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pushvserver = d.Get("pushvserver").(string)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG] netscaler-provider:  Range has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Range = intPtr(d.Get("range").(int))
		hasChange = true
	}
	if d.HasChange("recursionavailable") {
		log.Printf("[DEBUG] netscaler-provider:  Recursionavailable has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Recursionavailable = d.Get("recursionavailable").(string)
		hasChange = true
	}
	if d.HasChange("redirectfromport") {
		log.Printf("[DEBUG]  netscaler-provider: Redirectfromport has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirectfromport = intPtr(d.Get("redirectfromport").(int))
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG] netscaler-provider:  Redirectportrewrite has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("redirurl") {
		log.Printf("[DEBUG] netscaler-provider:  Redirurl has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirurl = d.Get("redirurl").(string)
		hasChange = true
	}
	if d.HasChange("redirurlflags") {
		log.Printf("[DEBUG] netscaler-provider:  Redirurlflags has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirurlflags = d.Get("redirurlflags").(bool)
		hasChange = true
	}
	if d.HasChange("resrule") {
		log.Printf("[DEBUG] netscaler-provider:  Resrule has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Resrule = d.Get("resrule").(string)
		hasChange = true
	}
	if d.HasChange("retainconnectionsoncluster") {
		log.Printf("[DEBUG]  netscaler-provider: Retainconnectionsoncluster has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Retainconnectionsoncluster = d.Get("retainconnectionsoncluster").(string)
		hasChange = true
	}
	if d.HasChange("rhistate") {
		log.Printf("[DEBUG]  netscaler-provider: Rhistate has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Rhistate = d.Get("rhistate").(string)
		hasChange = true
	}
	if d.HasChange("rtspnat") {
		log.Printf("[DEBUG] netscaler-provider:  Rtspnat has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Rtspnat = d.Get("rtspnat").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG] netscaler-provider:  Rule has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG] netscaler-provider:  Servicename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG] netscaler-provider:  Servicetype has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sessionless") {
		log.Printf("[DEBUG] netscaler-provider:  Sessionless has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sessionless = d.Get("sessionless").(string)
		hasChange = true
	}
	if d.HasChange("skippersistency") {
		log.Printf("[DEBUG] netscaler-provider:  Skippersistency has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Skippersistency = d.Get("skippersistency").(string)
		hasChange = true
	}
	if d.HasChange("sobackupaction") {
		log.Printf("[DEBUG] netscaler-provider:  Sobackupaction has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sobackupaction = d.Get("sobackupaction").(string)
		hasChange = true
	}
	if d.HasChange("somethod") {
		log.Printf("[DEBUG] netscaler-provider:  Somethod has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Somethod = d.Get("somethod").(string)
		hasChange = true
	}
	if d.HasChange("sopersistence") {
		log.Printf("[DEBUG] netscaler-provider:  Sopersistence has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sopersistence = d.Get("sopersistence").(string)
		hasChange = true
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Sopersistencetimeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sopersistencetimeout = intPtr(d.Get("sopersistencetimeout").(int))
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Sothreshold has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sothreshold = intPtr(d.Get("sothreshold").(int))
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] netscaler-provider:  State has changed for lbvserver %s, starting update", lbvserverName)
		stateChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Tcpprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG] netscaler-provider:  Td has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG] netscaler-provider:  Timeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Timeout = intPtr(d.Get("timeout").(int))
		hasChange = true
	}
	if d.HasChange("tosid") {
		log.Printf("[DEBUG] netscaler-provider:  Tosid has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Tosid = intPtr(d.Get("tosid").(int))
		hasChange = true
	}
	if d.HasChange("trofspersistence") {
		log.Printf("[DEBUG]  netscaler-provider: Trofspersistence has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Trofspersistence = d.Get("trofspersistence").(string)
		hasChange = true
	}
	if d.HasChange("v6netmasklen") {
		log.Printf("[DEBUG] netscaler-provider:  V6netmasklen has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.V6netmasklen = intPtr(d.Get("v6netmasklen").(int))
		hasChange = true
	}
	if d.HasChange("v6persistmasklen") {
		log.Printf("[DEBUG] netscaler-provider:  V6persistmasklen has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.V6persistmasklen = intPtr(d.Get("v6persistmasklen").(int))
		hasChange = true
	}
	if d.HasChange("vipheader") {
		log.Printf("[DEBUG] netscaler-provider:  Vipheader has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Vipheader = d.Get("vipheader").(string)
		hasChange = true
	}
	if d.HasChange("quicbridgeprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Quicbridgeprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Quicbridgeprofilename = d.Get("quicbridgeprofilename").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG] netscaler-provider:  Weight has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Weight = intPtr(d.Get("weight").(int))
		hasChange = true
	}
	if d.HasChange("sslcertkey") {
		log.Printf("[DEBUG] netscaler-provider:  ssl certkey has changed for lbvserver %s, starting update", lbvserverName)
		sslcertkeyChanged = true
	}
	if d.HasChange("snisslcertkeys") {
		log.Printf("[DEBUG] netscaler-provider:  sni ssl certkeys has changed for lbvserver %s, starting update", lbvserverName)
		snisslcertkeysChanged = true
	}
	if d.HasChange("sslprofile") {
		log.Printf("[DEBUG] netscaler-provider:  ssl profile has changed for lbvserver %s, starting update", lbvserverName)
		sslprofileChanged = true
	}
	if d.HasChange("ciphers") {
		log.Printf("[DEBUG] netscaler-provider:  ciphers have changed %s, starting update", lbvserverName)
		ciphersChanged = true
	}

	if d.HasChange("ciphersuites") {
		log.Printf("[DEBUG] netscaler-provider:  ciphers have changed %s, starting update", lbvserverName)
		ciphersuitesChanged = true
	}

	sslcertkey := d.Get("sslcertkey")
	sslcertkeyName := sslcertkey.(string)
	if sslcertkeyChanged {
		//Binding has to be updated
		//First we unbind from lb vserver
		oldSslcertkey, _ := d.GetChange("sslcertkey")
		oldSslcertkeyName := oldSslcertkey.(string)
		if oldSslcertkeyName != "" {
			err := client.UnbindResource(service.Sslvserver.Type(), lbvserverName, service.Sslcertkey.Type(), oldSslcertkeyName, "certkeyname")
			if err != nil {
				return diag.Errorf("[ERROR] netscaler-provider: Error unbinding sslcertkey from lbvserver %s", oldSslcertkeyName)
			}
			log.Printf("[DEBUG] netscaler-provider: sslcertkey has been unbound from lbvserver for sslcertkey %s ", oldSslcertkeyName)
		}
	}

	if hasChange {
		_, err := client.UpdateResource(service.Lbvserver.Type(), lbvserverName, &lbvserver)
		if err != nil {
			return diag.Errorf("[ERROR] netscaler-provider: Error updating lbvserver %s", lbvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: lbvserver has been updated  lbvserver %s ", lbvserverName)
	}

	if sslcertkeyChanged && sslcertkeyName != "" {
		//Binding has to be updated
		//rebind
		binding := ssl.Sslvservercertkeybinding{
			Vservername: lbvserverName,
			Certkeyname: sslcertkeyName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl cert %s to lbvserver %s", sslcertkeyName, lbvserverName)
		err := client.BindResource(service.Sslvserver.Type(), lbvserverName, service.Sslcertkey.Type(), sslcertkeyName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkeyName, lbvserverName)
			return diag.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkeyName, lbvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: new ssl cert has been bound to lbvserver  sslcertkey %s lbvserver %s", sslcertkeyName, lbvserverName)
	}

	if snisslcertkeysChanged {
		err := syncSnisslcert(d, meta, lbvserverName)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Ignore for standalone
	if ciphersChanged && isTargetAdcCluster(client) {
		if err := syncCiphers(d, meta, lbvserverName); err != nil {
			return diag.FromErr(err)
		}
	}

	if ciphersuitesChanged {
		if err := syncCiphersuites(d, meta, lbvserverName); err != nil {
			return diag.FromErr(err)
		}
	}

	sslprofile := d.Get("sslprofile")
	if sslprofileChanged {
		sslprofileName := sslprofile.(string)

		if sslprofileName == "" {
			sslvserver := ssl.Sslvserver{
				Vservername: lbvserverName,
				Sslprofile:  "true",
			}
			err := client.ActOnResource(service.Sslvserver.Type(), &sslvserver, "unset")
			if err != nil {
				return diag.Errorf("[ERROR] netscaler-provider: Error unbinding ssl profile from lbvserver %s", lbvserverName)
			}
		} else {
			sslvserver := ssl.Sslvserver{
				Vservername: lbvserverName,
				Sslprofile:  sslprofileName,
			}
			log.Printf("[INFO] netscaler-provider:  Binding ssl profile %s to lbvserver %s", sslprofileName, lbvserverName)
			_, err := client.UpdateResource(service.Sslvserver.Type(), lbvserverName, &sslvserver)
			if err != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to lbvserver %s", sslprofileName, lbvserverName)
				return diag.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl profile %s to lbvserver %s", sslprofileName, lbvserverName)
			}
			log.Printf("[DEBUG] netscaler-provider: new ssl profile has been bound to lbvserver  sslprofile %s lbvserver %s", sslprofileName, lbvserverName)
		}
	}

	if d.HasChange("sslpolicybinding") {
		if err := updateSslpolicyBindings(d, meta, lbvserverName); err != nil {
			return diag.FromErr(err)
		}
	}

	if stateChange {
		err := doLbvserverStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling lbvserver %s", lbvserverName)
		}
	}
	return readLbvserverFunc(ctx, d, meta)
}

func deleteLbvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In deleteLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Id()
	err := client.DeleteResource(service.Lbvserver.Type(), lbvserverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func doLbvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doLbvserverStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	type lbvserverStateChange struct {
		Name string `json:"name,omitempty"`
	}
	lbvserver := lbvserverStateChange{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Lbvserver.Type(), lbvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Lbvserver.Type(), lbvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
