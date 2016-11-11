package netscaler

import (
	"github.com/chiradeep/go-nitro/config/lb"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerLbmonitor() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbmonitorFunc,
		Read:          readLbmonitorFunc,
		Update:        updateLbmonitorFunc,
		Delete:        deleteLbmonitorFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alertretries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"basedn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"binddn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customheaders": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deviation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dispatcherip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dispatcherport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"evalrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failureretries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmwarerevision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httprequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbandsecurityid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iptunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lasversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logonpointname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lrtm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxforwards": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metricthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metricweight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mssqlprotocolversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"originhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"originrealm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"productname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"querytype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radaccountsession": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radaccounttype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radapn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radframedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radmsisdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resptimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resptimeoutthresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtsprequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scriptargs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scriptname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondarypassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipreguri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipuri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitepath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpcommunity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpoid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpthreshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storedb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storefrontacctservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"successretries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transparent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validatecred": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vendorspecificvendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	var lbmonitorName string
	if v, ok := d.GetOk("monitorname"); ok {
		lbmonitorName = v.(string)
	} else {
		lbmonitorName = resource.PrefixedUniqueId("tf-lbmonitor-")
		d.Set("monitorname", lbmonitorName)
	}
	lbmonitor := lb.Lbmonitor{
		Action:                 d.Get("action").(string),
		Alertretries:           d.Get("alertretries").(int),
		Application:            d.Get("application").(string),
		Attribute:              d.Get("attribute").(string),
		Basedn:                 d.Get("basedn").(string),
		Binddn:                 d.Get("binddn").(string),
		Customheaders:          d.Get("customheaders").(string),
		Database:               d.Get("database").(string),
		Destip:                 d.Get("destip").(string),
		Destport:               d.Get("destport").(int),
		Deviation:              d.Get("deviation").(int),
		Dispatcherip:           d.Get("dispatcherip").(string),
		Dispatcherport:         d.Get("dispatcherport").(int),
		Domain:                 d.Get("domain").(string),
		Downtime:               d.Get("downtime").(int),
		Evalrule:               d.Get("evalrule").(string),
		Failureretries:         d.Get("failureretries").(int),
		Filename:               d.Get("filename").(string),
		Filter:                 d.Get("filter").(string),
		Firmwarerevision:       d.Get("firmwarerevision").(int),
		Group:                  d.Get("group").(string),
		Hostipaddress:          d.Get("hostipaddress").(string),
		Hostname:               d.Get("hostname").(string),
		Httprequest:            d.Get("httprequest").(string),
		Inbandsecurityid:       d.Get("inbandsecurityid").(string),
		Interval:               d.Get("interval").(int),
		Iptunnel:               d.Get("iptunnel").(string),
		Kcdaccount:             d.Get("kcdaccount").(string),
		Lasversion:             d.Get("lasversion").(string),
		Logonpointname:         d.Get("logonpointname").(string),
		Lrtm:                   d.Get("lrtm").(string),
		Maxforwards:            d.Get("maxforwards").(int),
		Metric:                 d.Get("metric").(string),
		Metrictable:            d.Get("metrictable").(string),
		Metricthreshold:        d.Get("metricthreshold").(int),
		Metricweight:           d.Get("metricweight").(int),
		Monitorname:            d.Get("monitorname").(string),
		Mssqlprotocolversion:   d.Get("mssqlprotocolversion").(string),
		Netprofile:             d.Get("netprofile").(string),
		Originhost:             d.Get("originhost").(string),
		Originrealm:            d.Get("originrealm").(string),
		Password:               d.Get("password").(string),
		Productname:            d.Get("productname").(string),
		Query:                  d.Get("query").(string),
		Querytype:              d.Get("querytype").(string),
		Radaccountsession:      d.Get("radaccountsession").(string),
		Radaccounttype:         d.Get("radaccounttype").(int),
		Radapn:                 d.Get("radapn").(string),
		Radframedip:            d.Get("radframedip").(string),
		Radkey:                 d.Get("radkey").(string),
		Radmsisdn:              d.Get("radmsisdn").(string),
		Radnasid:               d.Get("radnasid").(string),
		Radnasip:               d.Get("radnasip").(string),
		Recv:                   d.Get("recv").(string),
		Resptimeout:            d.Get("resptimeout").(int),
		Resptimeoutthresh:      d.Get("resptimeoutthresh").(int),
		Retries:                d.Get("retries").(int),
		Reverse:                d.Get("reverse").(string),
		Rtsprequest:            d.Get("rtsprequest").(string),
		Scriptargs:             d.Get("scriptargs").(string),
		Scriptname:             d.Get("scriptname").(string),
		Secondarypassword:      d.Get("secondarypassword").(string),
		Secure:                 d.Get("secure").(string),
		Send:                   d.Get("send").(string),
		Servicegroupname:       d.Get("servicegroupname").(string),
		Servicename:            d.Get("servicename").(string),
		Sipmethod:              d.Get("sipmethod").(string),
		Sipreguri:              d.Get("sipreguri").(string),
		Sipuri:                 d.Get("sipuri").(string),
		Sitepath:               d.Get("sitepath").(string),
		Snmpcommunity:          d.Get("snmpcommunity").(string),
		Snmpoid:                d.Get("snmpoid").(string),
		Snmpthreshold:          d.Get("snmpthreshold").(string),
		Snmpversion:            d.Get("snmpversion").(string),
		Sqlquery:               d.Get("sqlquery").(string),
		State:                  d.Get("state").(string),
		Storedb:                d.Get("storedb").(string),
		Storefrontacctservice:  d.Get("storefrontacctservice").(string),
		Storename:              d.Get("storename").(string),
		Successretries:         d.Get("successretries").(int),
		Tos:                    d.Get("tos").(string),
		Tosid:                  d.Get("tosid").(int),
		Transparent:            d.Get("transparent").(string),
		Type:                   d.Get("type").(string),
		Units1:                 d.Get("units1").(string),
		Units2:                 d.Get("units2").(string),
		Units3:                 d.Get("units3").(string),
		Units4:                 d.Get("units4").(string),
		Username:               d.Get("username").(string),
		Validatecred:           d.Get("validatecred").(string),
		Vendorid:               d.Get("vendorid").(int),
		Vendorspecificvendorid: d.Get("vendorspecificvendorid").(int),
	}

	_, err := client.AddResource(netscaler.Lbmonitor.Type(), lbmonitorName, &lbmonitor)
	if err != nil {
		return err
	}

	d.SetId(lbmonitorName)

	err = readLbmonitorFunc(d, meta)
	if err != nil {
		log.Printf("?? we just created this lbmonitor but we can't read it ?? %s", lbmonitorName)
		return nil
	}
	return nil
}

func readLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	lbmonitorName := d.Id()
	log.Printf("Reading lbmonitor state %s", lbmonitorName)
	data, err := client.FindResource(netscaler.Lbmonitor.Type(), lbmonitorName)
	if err != nil {
		log.Printf("Clearing lbmonitor state %s", lbmonitorName)
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
	d.Set("interval", data["interval"])
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
	d.Set("originhost", data["originhost"])
	d.Set("originrealm", data["originrealm"])
	d.Set("password", data["password"])
	d.Set("productname", data["productname"])
	d.Set("query", data["query"])
	d.Set("querytype", data["querytype"])
	d.Set("radaccountsession", data["radaccountsession"])
	d.Set("radaccounttype", data["radaccounttype"])
	d.Set("radapn", data["radapn"])
	d.Set("radframedip", data["radframedip"])
	d.Set("radkey", data["radkey"])
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
	d.Set("secondarypassword", data["secondarypassword"])
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
	d.Set("state", data["state"])
	d.Set("storedb", data["storedb"])
	d.Set("storefrontacctservice", data["storefrontacctservice"])
	d.Set("storename", data["storename"])
	d.Set("successretries", data["successretries"])
	d.Set("tos", data["tos"])
	d.Set("tosid", data["tosid"])
	d.Set("transparent", data["transparent"])
	d.Set("type", data["type"])
	d.Set("units1", data["units1"])
	d.Set("units2", data["units2"])
	d.Set("units3", data["units3"])
	d.Set("units4", data["units4"])
	d.Set("username", data["username"])
	d.Set("validatecred", data["validatecred"])
	d.Set("vendorid", data["vendorid"])
	d.Set("vendorspecificvendorid", data["vendorspecificvendorid"])

	return nil

}

func updateLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] In update func")
	client := meta.(*NetScalerNitroClient).client
	lbmonitorName := d.Get("monitorname").(string)

	lbmonitor := lb.Lbmonitor{
		Monitorname: d.Get("monitorname").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG] Action has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("alertretries") {
		log.Printf("[DEBUG] Alertretries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Alertretries = d.Get("alertretries").(int)
		hasChange = true
	}
	if d.HasChange("application") {
		log.Printf("[DEBUG] Application has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Application = d.Get("application").(string)
		hasChange = true
	}
	if d.HasChange("attribute") {
		log.Printf("[DEBUG] Attribute has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Attribute = d.Get("attribute").(string)
		hasChange = true
	}
	if d.HasChange("basedn") {
		log.Printf("[DEBUG] Basedn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Basedn = d.Get("basedn").(string)
		hasChange = true
	}
	if d.HasChange("binddn") {
		log.Printf("[DEBUG] Binddn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Binddn = d.Get("binddn").(string)
		hasChange = true
	}
	if d.HasChange("customheaders") {
		log.Printf("[DEBUG] Customheaders has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Customheaders = d.Get("customheaders").(string)
		hasChange = true
	}
	if d.HasChange("database") {
		log.Printf("[DEBUG] Database has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Database = d.Get("database").(string)
		hasChange = true
	}
	if d.HasChange("destip") {
		log.Printf("[DEBUG] Destip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Destip = d.Get("destip").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG] Destport has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Destport = d.Get("destport").(int)
		hasChange = true
	}
	if d.HasChange("deviation") {
		log.Printf("[DEBUG] Deviation has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Deviation = d.Get("deviation").(int)
		hasChange = true
	}
	if d.HasChange("dispatcherip") {
		log.Printf("[DEBUG] Dispatcherip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Dispatcherip = d.Get("dispatcherip").(string)
		hasChange = true
	}
	if d.HasChange("dispatcherport") {
		log.Printf("[DEBUG] Dispatcherport has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Dispatcherport = d.Get("dispatcherport").(int)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG] Domain has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("downtime") {
		log.Printf("[DEBUG] Downtime has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Downtime = d.Get("downtime").(int)
		hasChange = true
	}
	if d.HasChange("evalrule") {
		log.Printf("[DEBUG] Evalrule has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Evalrule = d.Get("evalrule").(string)
		hasChange = true
	}
	if d.HasChange("failureretries") {
		log.Printf("[DEBUG] Failureretries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Failureretries = d.Get("failureretries").(int)
		hasChange = true
	}
	if d.HasChange("filename") {
		log.Printf("[DEBUG] Filename has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Filename = d.Get("filename").(string)
		hasChange = true
	}
	if d.HasChange("filter") {
		log.Printf("[DEBUG] Filter has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Filter = d.Get("filter").(string)
		hasChange = true
	}
	if d.HasChange("firmwarerevision") {
		log.Printf("[DEBUG] Firmwarerevision has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Firmwarerevision = d.Get("firmwarerevision").(int)
		hasChange = true
	}
	if d.HasChange("group") {
		log.Printf("[DEBUG] Group has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Group = d.Get("group").(string)
		hasChange = true
	}
	if d.HasChange("hostipaddress") {
		log.Printf("[DEBUG] Hostipaddress has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Hostipaddress = d.Get("hostipaddress").(string)
		hasChange = true
	}
	if d.HasChange("hostname") {
		log.Printf("[DEBUG] Hostname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Hostname = d.Get("hostname").(string)
		hasChange = true
	}
	if d.HasChange("httprequest") {
		log.Printf("[DEBUG] Httprequest has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Httprequest = d.Get("httprequest").(string)
		hasChange = true
	}
	if d.HasChange("inbandsecurityid") {
		log.Printf("[DEBUG] Inbandsecurityid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Inbandsecurityid = d.Get("inbandsecurityid").(string)
		hasChange = true
	}
	if d.HasChange("interval") {
		log.Printf("[DEBUG] Interval has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Interval = d.Get("interval").(int)
		hasChange = true
	}
	if d.HasChange("iptunnel") {
		log.Printf("[DEBUG] Iptunnel has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Iptunnel = d.Get("iptunnel").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG] Kcdaccount has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("lasversion") {
		log.Printf("[DEBUG] Lasversion has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Lasversion = d.Get("lasversion").(string)
		hasChange = true
	}
	if d.HasChange("logonpointname") {
		log.Printf("[DEBUG] Logonpointname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Logonpointname = d.Get("logonpointname").(string)
		hasChange = true
	}
	if d.HasChange("lrtm") {
		log.Printf("[DEBUG] Lrtm has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Lrtm = d.Get("lrtm").(string)
		hasChange = true
	}
	if d.HasChange("maxforwards") {
		log.Printf("[DEBUG] Maxforwards has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Maxforwards = d.Get("maxforwards").(int)
		hasChange = true
	}
	if d.HasChange("metric") {
		log.Printf("[DEBUG] Metric has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metric = d.Get("metric").(string)
		hasChange = true
	}
	if d.HasChange("metrictable") {
		log.Printf("[DEBUG] Metrictable has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metrictable = d.Get("metrictable").(string)
		hasChange = true
	}
	if d.HasChange("metricthreshold") {
		log.Printf("[DEBUG] Metricthreshold has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metricthreshold = d.Get("metricthreshold").(int)
		hasChange = true
	}
	if d.HasChange("metricweight") {
		log.Printf("[DEBUG] Metricweight has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Metricweight = d.Get("metricweight").(int)
		hasChange = true
	}
	if d.HasChange("monitorname") {
		log.Printf("[DEBUG] Monitorname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Monitorname = d.Get("monitorname").(string)
		hasChange = true
	}
	if d.HasChange("mssqlprotocolversion") {
		log.Printf("[DEBUG] Mssqlprotocolversion has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Mssqlprotocolversion = d.Get("mssqlprotocolversion").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] Netprofile has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("originhost") {
		log.Printf("[DEBUG] Originhost has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Originhost = d.Get("originhost").(string)
		hasChange = true
	}
	if d.HasChange("originrealm") {
		log.Printf("[DEBUG] Originrealm has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Originrealm = d.Get("originrealm").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG] Password has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("productname") {
		log.Printf("[DEBUG] Productname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Productname = d.Get("productname").(string)
		hasChange = true
	}
	if d.HasChange("query") {
		log.Printf("[DEBUG] Query has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Query = d.Get("query").(string)
		hasChange = true
	}
	if d.HasChange("querytype") {
		log.Printf("[DEBUG] Querytype has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Querytype = d.Get("querytype").(string)
		hasChange = true
	}
	if d.HasChange("radaccountsession") {
		log.Printf("[DEBUG] Radaccountsession has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radaccountsession = d.Get("radaccountsession").(string)
		hasChange = true
	}
	if d.HasChange("radaccounttype") {
		log.Printf("[DEBUG] Radaccounttype has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radaccounttype = d.Get("radaccounttype").(int)
		hasChange = true
	}
	if d.HasChange("radapn") {
		log.Printf("[DEBUG] Radapn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radapn = d.Get("radapn").(string)
		hasChange = true
	}
	if d.HasChange("radframedip") {
		log.Printf("[DEBUG] Radframedip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radframedip = d.Get("radframedip").(string)
		hasChange = true
	}
	if d.HasChange("radkey") {
		log.Printf("[DEBUG] Radkey has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radkey = d.Get("radkey").(string)
		hasChange = true
	}
	if d.HasChange("radmsisdn") {
		log.Printf("[DEBUG] Radmsisdn has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radmsisdn = d.Get("radmsisdn").(string)
		hasChange = true
	}
	if d.HasChange("radnasid") {
		log.Printf("[DEBUG] Radnasid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radnasid = d.Get("radnasid").(string)
		hasChange = true
	}
	if d.HasChange("radnasip") {
		log.Printf("[DEBUG] Radnasip has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Radnasip = d.Get("radnasip").(string)
		hasChange = true
	}
	if d.HasChange("recv") {
		log.Printf("[DEBUG] Recv has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Recv = d.Get("recv").(string)
		hasChange = true
	}
	if d.HasChange("resptimeout") {
		log.Printf("[DEBUG] Resptimeout has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Resptimeout = d.Get("resptimeout").(int)
		hasChange = true
	}
	if d.HasChange("resptimeoutthresh") {
		log.Printf("[DEBUG] Resptimeoutthresh has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Resptimeoutthresh = d.Get("resptimeoutthresh").(int)
		hasChange = true
	}
	if d.HasChange("retries") {
		log.Printf("[DEBUG] Retries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Retries = d.Get("retries").(int)
		hasChange = true
	}
	if d.HasChange("reverse") {
		log.Printf("[DEBUG] Reverse has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Reverse = d.Get("reverse").(string)
		hasChange = true
	}
	if d.HasChange("rtsprequest") {
		log.Printf("[DEBUG] Rtsprequest has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Rtsprequest = d.Get("rtsprequest").(string)
		hasChange = true
	}
	if d.HasChange("scriptargs") {
		log.Printf("[DEBUG] Scriptargs has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Scriptargs = d.Get("scriptargs").(string)
		hasChange = true
	}
	if d.HasChange("scriptname") {
		log.Printf("[DEBUG] Scriptname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Scriptname = d.Get("scriptname").(string)
		hasChange = true
	}
	if d.HasChange("secondarypassword") {
		log.Printf("[DEBUG] Secondarypassword has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Secondarypassword = d.Get("secondarypassword").(string)
		hasChange = true
	}
	if d.HasChange("secure") {
		log.Printf("[DEBUG] Secure has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Secure = d.Get("secure").(string)
		hasChange = true
	}
	if d.HasChange("send") {
		log.Printf("[DEBUG] Send has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Send = d.Get("send").(string)
		hasChange = true
	}
	if d.HasChange("servicegroupname") {
		log.Printf("[DEBUG] Servicegroupname has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Servicegroupname = d.Get("servicegroupname").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG] Servicename has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("sipmethod") {
		log.Printf("[DEBUG] Sipmethod has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sipmethod = d.Get("sipmethod").(string)
		hasChange = true
	}
	if d.HasChange("sipreguri") {
		log.Printf("[DEBUG] Sipreguri has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sipreguri = d.Get("sipreguri").(string)
		hasChange = true
	}
	if d.HasChange("sipuri") {
		log.Printf("[DEBUG] Sipuri has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sipuri = d.Get("sipuri").(string)
		hasChange = true
	}
	if d.HasChange("sitepath") {
		log.Printf("[DEBUG] Sitepath has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sitepath = d.Get("sitepath").(string)
		hasChange = true
	}
	if d.HasChange("snmpcommunity") {
		log.Printf("[DEBUG] Snmpcommunity has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpcommunity = d.Get("snmpcommunity").(string)
		hasChange = true
	}
	if d.HasChange("snmpoid") {
		log.Printf("[DEBUG] Snmpoid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpoid = d.Get("snmpoid").(string)
		hasChange = true
	}
	if d.HasChange("snmpthreshold") {
		log.Printf("[DEBUG] Snmpthreshold has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpthreshold = d.Get("snmpthreshold").(string)
		hasChange = true
	}
	if d.HasChange("snmpversion") {
		log.Printf("[DEBUG] Snmpversion has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Snmpversion = d.Get("snmpversion").(string)
		hasChange = true
	}
	if d.HasChange("sqlquery") {
		log.Printf("[DEBUG] Sqlquery has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Sqlquery = d.Get("sqlquery").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] State has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("storedb") {
		log.Printf("[DEBUG] Storedb has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storedb = d.Get("storedb").(string)
		hasChange = true
	}
	if d.HasChange("storefrontacctservice") {
		log.Printf("[DEBUG] Storefrontacctservice has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storefrontacctservice = d.Get("storefrontacctservice").(string)
		hasChange = true
	}
	if d.HasChange("storename") {
		log.Printf("[DEBUG] Storename has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Storename = d.Get("storename").(string)
		hasChange = true
	}
	if d.HasChange("successretries") {
		log.Printf("[DEBUG] Successretries has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Successretries = d.Get("successretries").(int)
		hasChange = true
	}
	if d.HasChange("tos") {
		log.Printf("[DEBUG] Tos has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Tos = d.Get("tos").(string)
		hasChange = true
	}
	if d.HasChange("tosid") {
		log.Printf("[DEBUG] Tosid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Tosid = d.Get("tosid").(int)
		hasChange = true
	}
	if d.HasChange("transparent") {
		log.Printf("[DEBUG] Transparent has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Transparent = d.Get("transparent").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG] Type has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("units1") {
		log.Printf("[DEBUG] Units1 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units1 = d.Get("units1").(string)
		hasChange = true
	}
	if d.HasChange("units2") {
		log.Printf("[DEBUG] Units2 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units2 = d.Get("units2").(string)
		hasChange = true
	}
	if d.HasChange("units3") {
		log.Printf("[DEBUG] Units3 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units3 = d.Get("units3").(string)
		hasChange = true
	}
	if d.HasChange("units4") {
		log.Printf("[DEBUG] Units4 has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Units4 = d.Get("units4").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG] Username has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Username = d.Get("username").(string)
		hasChange = true
	}
	if d.HasChange("validatecred") {
		log.Printf("[DEBUG] Validatecred has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Validatecred = d.Get("validatecred").(string)
		hasChange = true
	}
	if d.HasChange("vendorid") {
		log.Printf("[DEBUG] Vendorid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Vendorid = d.Get("vendorid").(int)
		hasChange = true
	}
	if d.HasChange("vendorspecificvendorid") {
		log.Printf("[DEBUG] Vendorspecificvendorid has changed for lbmonitor %s, starting update", lbmonitorName)
		lbmonitor.Vendorspecificvendorid = d.Get("vendorspecificvendorid").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Lbmonitor.Type(), lbmonitorName, &lbmonitor)
		if err != nil {
			return fmt.Errorf("Error updating lbmonitor %s", lbmonitorName)
		}
	}
	return readLbmonitorFunc(d, meta)
}

func deleteLbmonitorFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	lbmonitorName := d.Id()
	args := make([]string, 1, 1)
	args[0] = "type:" + d.Get("type").(string)
	err := client.DeleteResourceWithArgs(netscaler.Lbmonitor.Type(), lbmonitorName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
