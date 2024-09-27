package citrixadc

import (
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLbmonitor() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbmonitorFunc,
		Read:          readLbmonitorFunc,
		Update:        updateLbmonitorFunc,
		Delete:        deleteLbmonitorFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alertretries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"application": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"basedn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"binddn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customheaders": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deviation": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dispatcherip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dispatcherport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downtime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"evalrule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failureretries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmwarerevision": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httprequest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbandsecurityid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"iptunnel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lasversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logonpointname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lrtm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxforwards": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metric": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrictable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metricthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metricweight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitorname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mssqlprotocolversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oraclesid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"originhost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"originrealm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"productname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"querytype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radaccountsession": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radaccounttype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radapn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radframedip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radkey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radmsisdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resptimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resptimeoutthresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reverse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtsprequest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scriptargs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scriptname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondarypassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sipmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipreguri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipuri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitepath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpcommunity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpoid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpthreshold": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlquery": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storedb": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storefrontacctservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storefrontcheckbackendservices": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"successretries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tos": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tosid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transparent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trofscode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trofsstring": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"units1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units4": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validatecred": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vendorspecificvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"respcode": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: false,
			},
		},
	}
}

func createLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In createLbmonitorFunc")
	client := meta.(*NetScalerNitroClient).client

	meta.(*NetScalerNitroClient).lock.Lock()
	defer meta.(*NetScalerNitroClient).lock.Unlock()

	var lbmonitorName string
	if v, ok := d.GetOk("monitorname"); ok {
		lbmonitorName = v.(string)
	} else {
		lbmonitorName = resource.PrefixedUniqueId("tf-lbmonitor-")
		d.Set("monitorname", lbmonitorName)
	}
	lbmonitor := lb.Lbmonitor{
		Action:                         d.Get("action").(string),
		Alertretries:                   d.Get("alertretries").(int),
		Application:                    d.Get("application").(string),
		Attribute:                      d.Get("attribute").(string),
		Basedn:                         d.Get("basedn").(string),
		Binddn:                         d.Get("binddn").(string),
		Customheaders:                  d.Get("customheaders").(string),
		Database:                       d.Get("database").(string),
		Destip:                         d.Get("destip").(string),
		Destport:                       d.Get("destport").(int),
		Deviation:                      d.Get("deviation").(int),
		Dispatcherip:                   d.Get("dispatcherip").(string),
		Dispatcherport:                 d.Get("dispatcherport").(int),
		Domain:                         d.Get("domain").(string),
		Downtime:                       d.Get("downtime").(int),
		Evalrule:                       d.Get("evalrule").(string),
		Failureretries:                 d.Get("failureretries").(int),
		Filename:                       d.Get("filename").(string),
		Filter:                         d.Get("filter").(string),
		Firmwarerevision:               d.Get("firmwarerevision").(int),
		Group:                          d.Get("group").(string),
		Hostipaddress:                  d.Get("hostipaddress").(string),
		Hostname:                       d.Get("hostname").(string),
		Httprequest:                    d.Get("httprequest").(string),
		Inbandsecurityid:               d.Get("inbandsecurityid").(string),
		Interval:                       d.Get("interval").(int),
		Ipaddress:                      toStringList(d.Get("ipaddress").([]interface{})),
		Iptunnel:                       d.Get("iptunnel").(string),
		Kcdaccount:                     d.Get("kcdaccount").(string),
		Lasversion:                     d.Get("lasversion").(string),
		Logonpointname:                 d.Get("logonpointname").(string),
		Lrtm:                           d.Get("lrtm").(string),
		Maxforwards:                    d.Get("maxforwards").(int),
		Metric:                         d.Get("metric").(string),
		Metrictable:                    d.Get("metrictable").(string),
		Metricthreshold:                d.Get("metricthreshold").(int),
		Metricweight:                   d.Get("metricweight").(int),
		Monitorname:                    d.Get("monitorname").(string),
		Mssqlprotocolversion:           d.Get("mssqlprotocolversion").(string),
		Netprofile:                     d.Get("netprofile").(string),
		Oraclesid:                      d.Get("oraclesid").(string),
		Originhost:                     d.Get("originhost").(string),
		Originrealm:                    d.Get("originrealm").(string),
		Password:                       d.Get("password").(string),
		Productname:                    d.Get("productname").(string),
		Query:                          d.Get("query").(string),
		Querytype:                      d.Get("querytype").(string),
		Radaccountsession:              d.Get("radaccountsession").(string),
		Radaccounttype:                 d.Get("radaccounttype").(int),
		Radapn:                         d.Get("radapn").(string),
		Radframedip:                    d.Get("radframedip").(string),
		Radkey:                         d.Get("radkey").(string),
		Radmsisdn:                      d.Get("radmsisdn").(string),
		Radnasid:                       d.Get("radnasid").(string),
		Radnasip:                       d.Get("radnasip").(string),
		Recv:                           d.Get("recv").(string),
		Resptimeout:                    d.Get("resptimeout").(int),
		Resptimeoutthresh:              d.Get("resptimeoutthresh").(int),
		Retries:                        d.Get("retries").(int),
		Reverse:                        d.Get("reverse").(string),
		Rtsprequest:                    d.Get("rtsprequest").(string),
		Scriptargs:                     d.Get("scriptargs").(string),
		Scriptname:                     d.Get("scriptname").(string),
		Secondarypassword:              d.Get("secondarypassword").(string),
		Secure:                         d.Get("secure").(string),
		Send:                           d.Get("send").(string),
		Servicegroupname:               d.Get("servicegroupname").(string),
		Servicename:                    d.Get("servicename").(string),
		Sipmethod:                      d.Get("sipmethod").(string),
		Sipreguri:                      d.Get("sipreguri").(string),
		Sipuri:                         d.Get("sipuri").(string),
		Sitepath:                       d.Get("sitepath").(string),
		Snmpcommunity:                  d.Get("snmpcommunity").(string),
		Snmpoid:                        d.Get("snmpoid").(string),
		Snmpthreshold:                  d.Get("snmpthreshold").(string),
		Snmpversion:                    d.Get("snmpversion").(string),
		Sqlquery:                       d.Get("sqlquery").(string),
		Sslprofile:                     d.Get("sslprofile").(string),
		State:                          d.Get("state").(string),
		Storedb:                        d.Get("storedb").(string),
		Storefrontacctservice:          d.Get("storefrontacctservice").(string),
		Storefrontcheckbackendservices: d.Get("storefrontcheckbackendservices").(string),
		Storename:                      d.Get("storename").(string),
		Successretries:                 d.Get("successretries").(int),
		Tos:                            d.Get("tos").(string),
		Tosid:                          d.Get("tosid").(int),
		Transparent:                    d.Get("transparent").(string),
		Trofscode:                      d.Get("trofscode").(int),
		Trofsstring:                    d.Get("trofsstring").(string),
		Type:                           d.Get("type").(string),
		Units1:                         d.Get("units1").(string),
		Units2:                         d.Get("units2").(string),
		Units3:                         d.Get("units3").(string),
		Units4:                         d.Get("units4").(string),
		Username:                       d.Get("username").(string),
		Validatecred:                   d.Get("validatecred").(string),
		Vendorid:                       d.Get("vendorid").(int),
		Vendorspecificvendorid:         d.Get("vendorspecificvendorid").(int),
		Respcode:                       toStringList(d.Get("respcode").([]interface{})),
	}

	_, err := client.AddResource(service.Lbmonitor.Type(), lbmonitorName, &lbmonitor)
	if err != nil {
		return err
	}

	d.SetId(lbmonitorName)

	err = readLbmonitorFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbmonitor but we can't read it ?? %s", lbmonitorName)
		return nil
	}
	return nil
}

func readLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readLbmonitorFunc")
	client := meta.(*NetScalerNitroClient).client

	lbmonitorName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading lbmonitor state %s", lbmonitorName)
	data, err := client.FindResource(service.Lbmonitor.Type(), lbmonitorName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing lbmonitor state %s", lbmonitorName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("alertretries", data["alertretries"])
	d.Set("application", data["application"])
	d.Set("attribute", data["attribute"])
	d.Set("basedn", data["basedn"])
	d.Set("binddn", data["binddn"])
	d.Set("customheaders", data["customheaders"])
	d.Set("database", data["database"])
	d.Set("destip", data["destip"])
	d.Set("destport", data["destport"])
	d.Set("deviation", data["deviation"])
	d.Set("dispatcherip", data["dispatcherip"])
	d.Set("dispatcherport", data["dispatcherport"])
	d.Set("domain", data["domain"])
	d.Set("downtime", data["downtime"])
	d.Set("evalrule", data["evalrule"])
	d.Set("failureretries", data["failureretries"])
	d.Set("filename", data["filename"])
	d.Set("filter", data["filter"])
	d.Set("firmwarerevision", data["firmwarerevision"])
	d.Set("group", data["group"])
	d.Set("hostipaddress", data["hostipaddress"])
	d.Set("hostname", data["hostname"])
	d.Set("httprequest", data["httprequest"])
	d.Set("inbandsecurityid", data["inbandsecurityid"])
	// d.Set("interval", data["interval"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("iptunnel", data["iptunnel"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("lasversion", data["lasversion"])
	d.Set("logonpointname", data["logonpointname"])
	d.Set("lrtm", data["lrtm"])
	d.Set("maxforwards", data["maxforwards"])
	d.Set("metric", data["metric"])
	d.Set("metrictable", data["metrictable"])
	d.Set("metricthreshold", data["metricthreshold"])
	d.Set("metricweight", data["metricweight"])
	d.Set("monitorname", data["monitorname"])
	d.Set("mssqlprotocolversion", data["mssqlprotocolversion"])
	d.Set("netprofile", data["netprofile"])
	d.Set("oraclesid", data["oraclesid"])
	d.Set("originhost", data["originhost"])
	d.Set("originrealm", data["originrealm"])
	// d.Set("password", data["password"])	// We get the hash value from the NetScaler, which creates terraform to update the resource attribute on our next terraform apply command
	d.Set("productname", data["productname"])
	d.Set("query", data["query"])
	d.Set("querytype", data["querytype"])
	d.Set("radaccountsession", data["radaccountsession"])
	d.Set("radaccounttype", data["radaccounttype"])
	d.Set("radapn", data["radapn"])
	d.Set("radframedip", data["radframedip"])
	// d.Set("radkey", data["radkey"]) // We get the hash value from the NetScaler, which creates terraform to update the resource attribute on our next terraform apply command
	d.Set("radmsisdn", data["radmsisdn"])
	d.Set("radnasid", data["radnasid"])
	d.Set("radnasip", data["radnasip"])
	d.Set("recv", data["recv"])
	d.Set("resptimeout", data["resptimeout"])
	d.Set("resptimeoutthresh", data["resptimeoutthresh"])
	d.Set("retries", data["retries"])
	d.Set("reverse", data["reverse"])
	d.Set("rtsprequest", data["rtsprequest"])
	d.Set("scriptargs", data["scriptargs"])
	d.Set("scriptname", data["scriptname"])
	// d.Set("secondarypassword", data["secondarypassword"]) // We get the hash value from the NetScaler, which creates terraform to update the resource attribute on our next terraform apply command
	d.Set("secure", data["secure"])
	d.Set("send", data["send"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("servicename", data["servicename"])
	d.Set("sipmethod", data["sipmethod"])
	d.Set("sipreguri", data["sipreguri"])
	d.Set("sipuri", data["sipuri"])
	d.Set("sitepath", data["sitepath"])
	d.Set("snmpcommunity", data["snmpcommunity"])
	d.Set("snmpoid", data["snmpoid"])
	d.Set("snmpthreshold", data["snmpthreshold"])
	d.Set("snmpversion", data["snmpversion"])
	d.Set("sqlquery", data["sqlquery"])
	d.Set("sslprofile", data["sslprofile"])
	d.Set("state", data["state"])
	d.Set("storedb", data["storedb"])
	d.Set("storefrontacctservice", data["storefrontacctservice"])
	d.Set("storefrontcheckbackendservices", data["storefrontcheckbackendservices"])
	d.Set("storename", data["storename"])
	d.Set("successretries", data["successretries"])
	d.Set("tos", data["tos"])
	d.Set("tosid", data["tosid"])
	d.Set("transparent", data["transparent"])
	d.Set("trofscode", data["trofscode"])
	d.Set("trofsstring", data["trofsstring"])
	d.Set("type", data["type"])
	d.Set("units1", data["units1"])
	d.Set("units2", data["units2"])
	// d.Set("units3", data["units3"])
	d.Set("units4", data["units4"])
	d.Set("username", data["username"])
	d.Set("validatecred", data["validatecred"])
	d.Set("vendorid", data["vendorid"])
	d.Set("vendorspecificvendorid", data["vendorspecificvendorid"])
	// d.Set("respcode", data["respcode"]) // we receive different value from NetScaler

	// FIXME: in lbmonitor, for `interval=60`, the `units3` will wrongly be set to `MIN` by the NetScaler.
	// Hence, we will set it to `SEC` to make it idempotent
	// Refer Issue: #324 (https://github.com/netscaler/ansible-collection-netscaleradc/issues/324) in ansible-collection-netscaleradc
	// Refer Issue: #1165 (https://github.com/citrix/terraform-provider-citrixadc/issues/1165) in terraform-provider-citrixadc
	if val, ok := d.GetOk("units3"); !ok || val.(string) == "SEC" {
		if existingUnits3, exists := data["units3"]; exists && existingUnits3.(string) == "MIN" {
			if interval, intervalExists := data["interval"]; intervalExists {
				intervalInt := int(interval.(float64))
				data["interval"] = strconv.Itoa(intervalInt * 60)
				log.Println("[DEBUG] netscaler-provider:  interval is in MIN, converting it to SEC")
			}
			data["units3"] = "SEC"
		}
	} else {
		d.Set("units3", data["units3"])
		d.Set("interval", data["interval"])
	}

	return nil

}

func updateLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateLbmonitorFunc")
	client := meta.(*NetScalerNitroClient).client

	meta.(*NetScalerNitroClient).lock.Lock()
	defer meta.(*NetScalerNitroClient).lock.Unlock()

	lbmonitorName := d.Get("monitorname").(string)

	lbmonitor := lb.Lbmonitor{
		Monitorname: d.Get("monitorname").(string),
		Type:        d.Get("type").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG] netscaler-provider:  Action has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("alertretries") {
		log.Printf("[DEBUG] netscaler-provider:  Alertretries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Alertretries = d.Get("alertretries").(int)
		hasChange = true
	}
	if d.HasChange("application") {
		log.Printf("[DEBUG] netscaler-provider:  Application has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Application = d.Get("application").(string)
		hasChange = true
	}
	if d.HasChange("attribute") {
		log.Printf("[DEBUG] netscaler-provider:  Attribute has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Attribute = d.Get("attribute").(string)
		hasChange = true
	}
	if d.HasChange("basedn") {
		log.Printf("[DEBUG] netscaler-provider:  Basedn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Basedn = d.Get("basedn").(string)
		hasChange = true
	}
	if d.HasChange("binddn") {
		log.Printf("[DEBUG] netscaler-provider:  Binddn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Binddn = d.Get("binddn").(string)
		hasChange = true
	}
	if d.HasChange("customheaders") {
		log.Printf("[DEBUG] netscaler-provider:  Customheaders has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Customheaders = d.Get("customheaders").(string)
		hasChange = true
	}
	if d.HasChange("database") {
		log.Printf("[DEBUG] netscaler-provider:  Database has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Database = d.Get("database").(string)
		hasChange = true
	}
	if d.HasChange("destip") {
		log.Printf("[DEBUG] netscaler-provider:  Destip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Destip = d.Get("destip").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG] netscaler-provider:  Destport has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Destport = d.Get("destport").(int)
		hasChange = true
	}
	if d.HasChange("deviation") {
		log.Printf("[DEBUG] netscaler-provider:  Deviation has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Deviation = d.Get("deviation").(int)
		hasChange = true
	}
	if d.HasChange("dispatcherip") {
		log.Printf("[DEBUG] netscaler-provider:  Dispatcherip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Dispatcherip = d.Get("dispatcherip").(string)
		hasChange = true
	}
	if d.HasChange("dispatcherport") {
		log.Printf("[DEBUG] netscaler-provider:  Dispatcherport has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Dispatcherport = d.Get("dispatcherport").(int)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG] netscaler-provider:  Domain has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("downtime") {
		log.Printf("[DEBUG] netscaler-provider:  Downtime has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Downtime = d.Get("downtime").(int)
		hasChange = true
	}
	if d.HasChange("evalrule") {
		log.Printf("[DEBUG] netscaler-provider:  Evalrule has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Evalrule = d.Get("evalrule").(string)
		hasChange = true
	}
	if d.HasChange("failureretries") {
		log.Printf("[DEBUG] netscaler-provider:  Failureretries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Failureretries = d.Get("failureretries").(int)
		hasChange = true
	}
	if d.HasChange("filename") {
		log.Printf("[DEBUG] netscaler-provider:  Filename has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Filename = d.Get("filename").(string)
		hasChange = true
	}
	if d.HasChange("filter") {
		log.Printf("[DEBUG] netscaler-provider:  Filter has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Filter = d.Get("filter").(string)
		hasChange = true
	}
	if d.HasChange("firmwarerevision") {
		log.Printf("[DEBUG] netscaler-provider:  Firmwarerevision has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Firmwarerevision = d.Get("firmwarerevision").(int)
		hasChange = true
	}
	if d.HasChange("group") {
		log.Printf("[DEBUG] netscaler-provider:  Group has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Group = d.Get("group").(string)
		hasChange = true
	}
	if d.HasChange("hostipaddress") {
		log.Printf("[DEBUG] netscaler-provider:  Hostipaddress has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Hostipaddress = d.Get("hostipaddress").(string)
		hasChange = true
	}
	if d.HasChange("hostname") {
		log.Printf("[DEBUG] netscaler-provider:  Hostname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Hostname = d.Get("hostname").(string)
		hasChange = true
	}
	if d.HasChange("httprequest") {
		log.Printf("[DEBUG] netscaler-provider:  Httprequest has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Httprequest = d.Get("httprequest").(string)
		hasChange = true
	}
	if d.HasChange("inbandsecurityid") {
		log.Printf("[DEBUG] netscaler-provider:  Inbandsecurityid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Inbandsecurityid = d.Get("inbandsecurityid").(string)
		hasChange = true
	}
	if d.HasChange("interval") {
		log.Printf("[DEBUG] netscaler-provider:  Interval has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Interval = d.Get("interval").(int)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG] netscaler-provider:  Iptunnel has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Ipaddress = toStringList(d.Get("ipaddress").([]interface{}))
		hasChange = true
	}
	if d.HasChange("iptunnel") {
		log.Printf("[DEBUG] netscaler-provider:  Iptunnel has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Iptunnel = d.Get("iptunnel").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG] netscaler-provider:  Kcdaccount has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("lasversion") {
		log.Printf("[DEBUG] netscaler-provider:  Lasversion has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Lasversion = d.Get("lasversion").(string)
		hasChange = true
	}
	if d.HasChange("logonpointname") {
		log.Printf("[DEBUG] netscaler-provider:  Logonpointname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Logonpointname = d.Get("logonpointname").(string)
		hasChange = true
	}
	if d.HasChange("lrtm") {
		log.Printf("[DEBUG] netscaler-provider:  Lrtm has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Lrtm = d.Get("lrtm").(string)
		hasChange = true
	}
	if d.HasChange("maxforwards") {
		log.Printf("[DEBUG] netscaler-provider:  Maxforwards has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Maxforwards = d.Get("maxforwards").(int)
		hasChange = true
	}
	if d.HasChange("metric") {
		log.Printf("[DEBUG] netscaler-provider:  Metric has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metric = d.Get("metric").(string)
		hasChange = true
	}
	if d.HasChange("metrictable") {
		log.Printf("[DEBUG] netscaler-provider:  Metrictable has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metrictable = d.Get("metrictable").(string)
		hasChange = true
	}
	if d.HasChange("metricthreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Metricthreshold has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metricthreshold = d.Get("metricthreshold").(int)
		hasChange = true
	}
	if d.HasChange("metricweight") {
		log.Printf("[DEBUG] netscaler-provider:  Metricweight has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metricweight = d.Get("metricweight").(int)
		hasChange = true
	}
	if d.HasChange("monitorname") {
		log.Printf("[DEBUG] netscaler-provider:  Monitorname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Monitorname = d.Get("monitorname").(string)
		hasChange = true
	}
	if d.HasChange("mssqlprotocolversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mssqlprotocolversion has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Mssqlprotocolversion = d.Get("mssqlprotocolversion").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Netprofile has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("oraclesid") {
		log.Printf("[DEBUG]  netscaler-provider: Oraclesid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Oraclesid = d.Get("oraclesid").(string)
		hasChange = true
	}
	if d.HasChange("originhost") {
		log.Printf("[DEBUG] netscaler-provider:  Originhost has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Originhost = d.Get("originhost").(string)
		hasChange = true
	}
	if d.HasChange("originrealm") {
		log.Printf("[DEBUG] netscaler-provider:  Originrealm has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Originrealm = d.Get("originrealm").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG] netscaler-provider:  Password has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("productname") {
		log.Printf("[DEBUG] netscaler-provider:  Productname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Productname = d.Get("productname").(string)
		hasChange = true
	}
	if d.HasChange("query") {
		log.Printf("[DEBUG] netscaler-provider:  Query has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Query = d.Get("query").(string)
		hasChange = true
	}
	if d.HasChange("querytype") {
		log.Printf("[DEBUG] netscaler-provider:  Querytype has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Querytype = d.Get("querytype").(string)
		hasChange = true
	}
	if d.HasChange("radaccountsession") {
		log.Printf("[DEBUG] netscaler-provider:  Radaccountsession has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radaccountsession = d.Get("radaccountsession").(string)
		hasChange = true
	}
	if d.HasChange("radaccounttype") {
		log.Printf("[DEBUG] netscaler-provider:  Radaccounttype has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radaccounttype = d.Get("radaccounttype").(int)
		hasChange = true
	}
	if d.HasChange("radapn") {
		log.Printf("[DEBUG] netscaler-provider:  Radapn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radapn = d.Get("radapn").(string)
		hasChange = true
	}
	if d.HasChange("radframedip") {
		log.Printf("[DEBUG] netscaler-provider:  Radframedip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radframedip = d.Get("radframedip").(string)
		hasChange = true
	}
	if d.HasChange("radkey") {
		log.Printf("[DEBUG] netscaler-provider:  Radkey has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radkey = d.Get("radkey").(string)
		hasChange = true
	}
	if d.HasChange("radmsisdn") {
		log.Printf("[DEBUG] netscaler-provider:  Radmsisdn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radmsisdn = d.Get("radmsisdn").(string)
		hasChange = true
	}
	if d.HasChange("radnasid") {
		log.Printf("[DEBUG] netscaler-provider:  Radnasid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radnasid = d.Get("radnasid").(string)
		hasChange = true
	}
	if d.HasChange("radnasip") {
		log.Printf("[DEBUG] netscaler-provider:  Radnasip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radnasip = d.Get("radnasip").(string)
		hasChange = true
	}
	if d.HasChange("recv") {
		log.Printf("[DEBUG] netscaler-provider:  Recv has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Recv = d.Get("recv").(string)
		hasChange = true
	}
	if d.HasChange("resptimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Resptimeout has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Resptimeout = d.Get("resptimeout").(int)
		hasChange = true
	}
	if d.HasChange("resptimeoutthresh") {
		log.Printf("[DEBUG] netscaler-provider:  Resptimeoutthresh has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Resptimeoutthresh = d.Get("resptimeoutthresh").(int)
		hasChange = true
	}
	if d.HasChange("retries") {
		log.Printf("[DEBUG] netscaler-provider:  Retries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Retries = d.Get("retries").(int)
		hasChange = true
	}
	if d.HasChange("reverse") {
		log.Printf("[DEBUG] netscaler-provider:  Reverse has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Reverse = d.Get("reverse").(string)
		hasChange = true
	}
	if d.HasChange("rtsprequest") {
		log.Printf("[DEBUG] netscaler-provider:  Rtsprequest has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Rtsprequest = d.Get("rtsprequest").(string)
		hasChange = true
	}
	if d.HasChange("scriptargs") {
		log.Printf("[DEBUG] netscaler-provider:  Scriptargs has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Scriptargs = d.Get("scriptargs").(string)
		hasChange = true
	}
	if d.HasChange("scriptname") {
		log.Printf("[DEBUG] netscaler-provider:  Scriptname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Scriptname = d.Get("scriptname").(string)
		hasChange = true
	}
	if d.HasChange("secondarypassword") {
		log.Printf("[DEBUG] netscaler-provider:  Secondarypassword has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Secondarypassword = d.Get("secondarypassword").(string)
		hasChange = true
	}
	if d.HasChange("secure") {
		log.Printf("[DEBUG] netscaler-provider:  Secure has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Secure = d.Get("secure").(string)
		hasChange = true
	}
	if d.HasChange("send") {
		log.Printf("[DEBUG] netscaler-provider:  Send has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Send = d.Get("send").(string)
		hasChange = true
	}
	if d.HasChange("servicegroupname") {
		log.Printf("[DEBUG] netscaler-provider:  Servicegroupname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Servicegroupname = d.Get("servicegroupname").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG] netscaler-provider:  Servicename has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("sipmethod") {
		log.Printf("[DEBUG] netscaler-provider:  Sipmethod has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sipmethod = d.Get("sipmethod").(string)
		hasChange = true
	}
	if d.HasChange("sipreguri") {
		log.Printf("[DEBUG] netscaler-provider:  Sipreguri has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sipreguri = d.Get("sipreguri").(string)
		hasChange = true
	}
	if d.HasChange("sipuri") {
		log.Printf("[DEBUG] netscaler-provider:  Sipuri has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sipuri = d.Get("sipuri").(string)
		hasChange = true
	}
	if d.HasChange("sitepath") {
		log.Printf("[DEBUG] netscaler-provider:  Sitepath has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sitepath = d.Get("sitepath").(string)
		hasChange = true
	}
	if d.HasChange("snmpcommunity") {
		log.Printf("[DEBUG] netscaler-provider:  Snmpcommunity has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpcommunity = d.Get("snmpcommunity").(string)
		hasChange = true
	}
	if d.HasChange("snmpoid") {
		log.Printf("[DEBUG] netscaler-provider:  Snmpoid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpoid = d.Get("snmpoid").(string)
		hasChange = true
	}
	if d.HasChange("snmpthreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Snmpthreshold has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpthreshold = d.Get("snmpthreshold").(string)
		hasChange = true
	}
	if d.HasChange("snmpversion") {
		log.Printf("[DEBUG] netscaler-provider:  Snmpversion has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpversion = d.Get("snmpversion").(string)
		hasChange = true
	}
	if d.HasChange("sqlquery") {
		log.Printf("[DEBUG] netscaler-provider:  Sqlquery has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sqlquery = d.Get("sqlquery").(string)
		hasChange = true
	}
	if d.HasChange("sslprofile") {
		log.Printf("[DEBUG]  netscaler-provider: Sslprofile has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sslprofile = d.Get("sslprofile").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] netscaler-provider:  State has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("storedb") {
		log.Printf("[DEBUG] netscaler-provider:  Storedb has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storedb = d.Get("storedb").(string)
		hasChange = true
	}
	if d.HasChange("storefrontacctservice") {
		log.Printf("[DEBUG] netscaler-provider:  Storefrontacctservice has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storefrontacctservice = d.Get("storefrontacctservice").(string)
		hasChange = true
	}
	if d.HasChange("storefrontcheckbackendservices") {
		log.Printf("[DEBUG]  netscaler-provider: Storefrontcheckbackendservices has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storefrontcheckbackendservices = d.Get("storefrontcheckbackendservices").(string)
		hasChange = true
	}
	if d.HasChange("storename") {
		log.Printf("[DEBUG] netscaler-provider:  Storename has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storename = d.Get("storename").(string)
		hasChange = true
	}
	if d.HasChange("successretries") {
		log.Printf("[DEBUG] netscaler-provider:  Successretries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Successretries = d.Get("successretries").(int)
		hasChange = true
	}
	if d.HasChange("tos") {
		log.Printf("[DEBUG] netscaler-provider:  Tos has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Tos = d.Get("tos").(string)
		hasChange = true
	}
	if d.HasChange("tosid") {
		log.Printf("[DEBUG] netscaler-provider:  Tosid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Tosid = d.Get("tosid").(int)
		hasChange = true
	}
	if d.HasChange("transparent") {
		log.Printf("[DEBUG] netscaler-provider:  Transparent has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Transparent = d.Get("transparent").(string)
		hasChange = true
	}
	if d.HasChange("trofscode") {
		log.Printf("[DEBUG]  netscaler-provider: Trofscode has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Trofscode = d.Get("trofscode").(int)
		hasChange = true
	}
	if d.HasChange("trofsstring") {
		log.Printf("[DEBUG]  netscaler-provider: Trofsstring has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Trofsstring = d.Get("trofsstring").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG] netscaler-provider:  Type has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("units1") {
		log.Printf("[DEBUG] netscaler-provider:  Units1 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units1 = d.Get("units1").(string)
		hasChange = true
	}
	if d.HasChange("units2") {
		log.Printf("[DEBUG] netscaler-provider:  Units2 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units2 = d.Get("units2").(string)
		hasChange = true
	}
	if d.HasChange("units3") {
		log.Printf("[DEBUG] netscaler-provider:  Units3 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units3 = d.Get("units3").(string)
		lbmonitor.Interval = d.Get("interval").(int)
		hasChange = true
	}
	if d.HasChange("units4") {
		log.Printf("[DEBUG] netscaler-provider:  Units4 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units4 = d.Get("units4").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG] netscaler-provider:  Username has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Username = d.Get("username").(string)
		hasChange = true
	}
	if d.HasChange("validatecred") {
		log.Printf("[DEBUG] netscaler-provider:  Validatecred has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Validatecred = d.Get("validatecred").(string)
		hasChange = true
	}
	if d.HasChange("vendorid") {
		log.Printf("[DEBUG] netscaler-provider:  Vendorid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Vendorid = d.Get("vendorid").(int)
		hasChange = true
	}
	if d.HasChange("vendorspecificvendorid") {
		log.Printf("[DEBUG] netscaler-provider:  Vendorspecificvendorid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Vendorspecificvendorid = d.Get("vendorspecificvendorid").(int)
		hasChange = true
	}

	if d.HasChange("respcode") {
		log.Printf("[DEBUG] netscaler-provider:  Respcode has changed for lbmonitor %s, starting update", lbmonitorName)
		_, ok := d.GetOk("respcode")
		respcode_val := toStringList(d.Get("respcode").([]interface{}))

		if ok {
			lbmonitor.Respcode = respcode_val
			hasChange = true
		}
	}

	if hasChange {
		_, err := client.UpdateResource(service.Lbmonitor.Type(), lbmonitorName, &lbmonitor)
		if err != nil {
			return err
		}
	}
	return readLbmonitorFunc(d, meta)
}

func deleteLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In deleteLbmonitorFunc")
	client := meta.(*NetScalerNitroClient).client

	meta.(*NetScalerNitroClient).lock.Lock()
	defer meta.(*NetScalerNitroClient).lock.Unlock()

	lbmonitorName := d.Id()
	args := make([]string, 1, 1)
	args[0] = "type:" + d.Get("type").(string)
	err := client.DeleteResourceWithArgs(service.Lbmonitor.Type(), lbmonitorName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
