package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCrvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCrvserverFunc,
		Read:          readCrvserverFunc,
		Update:        updateCrvserverFunc,
		Delete:        deleteCrvserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backendssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachevserver": &schema.Schema{
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
			"destinationvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsvservername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ghost": &schema.Schema{
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
			"map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"onpolicymatch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"originusip": &schema.Schema{
				Type:     schema.TypeString,
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
			"probeport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"probeprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probesuccessresponsecode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirecturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rhistate": &schema.Schema{
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
			"srcipexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprobeport": &schema.Schema{
				Type:     schema.TypeInt,
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
			"useoriginipportforcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useportrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"via": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCrvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCrvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	crvserverName := d.Get("name").(string)
	crvserver := cr.Crvserver{
		Appflowlog:               d.Get("appflowlog").(string),
		Arp:                      d.Get("arp").(string),
		Backendssl:               d.Get("backendssl").(string),
		Backupvserver:            d.Get("backupvserver").(string),
		Cachetype:                d.Get("cachetype").(string),
		Cachevserver:             d.Get("cachevserver").(string),
		Clttimeout:               d.Get("clttimeout").(int),
		Comment:                  d.Get("comment").(string),
		Destinationvserver:       d.Get("destinationvserver").(string),
		Disableprimaryondown:     d.Get("disableprimaryondown").(string),
		Dnsvservername:           d.Get("dnsvservername").(string),
		Domain:                   d.Get("domain").(string),
		Downstateflush:           d.Get("downstateflush").(string),
		Format:                   d.Get("format").(string),
		Ghost:                    d.Get("ghost").(string),
		Httpprofilename:          d.Get("httpprofilename").(string),
		Icmpvsrresponse:          d.Get("icmpvsrresponse").(string),
		Ipset:                    d.Get("ipset").(string),
		Ipv46:                    d.Get("ipv46").(string),
		L2conn:                   d.Get("l2conn").(string),
		Listenpolicy:             d.Get("listenpolicy").(string),
		Listenpriority:           d.Get("listenpriority").(int),
		Map:                      d.Get("map").(string),
		Name:                     crvserverName,
		Netprofile:               d.Get("netprofile").(string),
		Onpolicymatch:            d.Get("onpolicymatch").(string),
		Originusip:               d.Get("originusip").(string),
		Port:                     d.Get("port").(int),
		Precedence:               d.Get("precedence").(string),
		Probeport:                d.Get("probeport").(int),
		Probeprotocol:            d.Get("probeprotocol").(string),
		Probesuccessresponsecode: d.Get("probesuccessresponsecode").(string),
		Range:                    d.Get("range").(int),
		Redirect:                 d.Get("redirect").(string),
		Redirecturl:              d.Get("redirecturl").(string),
		Reuse:                    d.Get("reuse").(string),
		Rhistate:                 d.Get("rhistate").(string),
		Servicetype:              d.Get("servicetype").(string),
		Sopersistencetimeout:     d.Get("sopersistencetimeout").(int),
		Sothreshold:              d.Get("sothreshold").(int),
		Srcipexpr:                d.Get("srcipexpr").(string),
		State:                    d.Get("state").(string),
		Tcpprobeport:             d.Get("tcpprobeport").(int),
		Tcpprofilename:           d.Get("tcpprofilename").(string),
		Td:                       d.Get("td").(int),
		Useoriginipportforcache:  d.Get("useoriginipportforcache").(string),
		Useportrange:             d.Get("useportrange").(string),
		Via:                      d.Get("via").(string),
	}

	_, err := client.AddResource(service.Crvserver.Type(), crvserverName, &crvserver)
	if err != nil {
		return err
	}

	d.SetId(crvserverName)

	err = readCrvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this crvserver but we can't read it ?? %s", crvserverName)
		return nil
	}
	return nil
}

func readCrvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCrvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	crvserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading crvserver state %s", crvserverName)
	data, err := client.FindResource(service.Crvserver.Type(), crvserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing crvserver state %s", crvserverName)
		d.SetId("")
		return nil
	}
	d.Set("appflowlog", data["appflowlog"])
	d.Set("arp", data["arp"])
	d.Set("backendssl", data["backendssl"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("cachetype", data["cachetype"])
	d.Set("cachevserver", data["cachevserver"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("comment", data["comment"])
	d.Set("destinationvserver", data["destinationvserver"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("dnsvservername", data["dnsvservername"])
	d.Set("domain", data["domain"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("format", data["format"])
	d.Set("ghost", data["ghost"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("ipset", data["ipset"])
	d.Set("ipv46", data["ipv46"])
	d.Set("l2conn", data["l2conn"])
	d.Set("listenpolicy", data["listenpolicy"])
	d.Set("listenpriority", data["listenpriority"])
	d.Set("map", data["map"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("onpolicymatch", data["onpolicymatch"])
	d.Set("originusip", data["originusip"])
	d.Set("port", data["port"])
	d.Set("precedence", data["precedence"])
	d.Set("probeport", data["probeport"])
	d.Set("probeprotocol", data["probeprotocol"])
	d.Set("probesuccessresponsecode", data["probesuccessresponsecode"])
	d.Set("range", data["range"])
	d.Set("redirect", data["redirect"])
	d.Set("redirecturl", data["redirecturl"])
	d.Set("reuse", data["reuse"])
	d.Set("rhistate", data["rhistate"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sopersistencetimeout", data["sopersistencetimeout"])
	d.Set("sothreshold", data["sothreshold"])
	d.Set("srcipexpr", data["srcipexpr"])
	d.Set("state", data["state"])
	d.Set("tcpprobeport", data["tcpprobeport"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("useoriginipportforcache", data["useoriginipportforcache"])
	d.Set("useportrange", data["useportrange"])
	d.Set("via", data["via"])

	return nil

}

func updateCrvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCrvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	crvserverName := d.Get("name").(string)

	crvserver := cr.Crvserver{
		Name: crvserverName,
	}
	hasChange := false
	stateChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for crvserver %s, starting update", crvserverName)
		crvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("arp") {
		log.Printf("[DEBUG]  citrixadc-provider: Arp has changed for crvserver %s, starting update", crvserverName)
		crvserver.Arp = d.Get("arp").(string)
		hasChange = true
	}
	if d.HasChange("backendssl") {
		log.Printf("[DEBUG]  citrixadc-provider: Backendssl has changed for crvserver %s, starting update", crvserverName)
		crvserver.Backendssl = d.Get("backendssl").(string)
		hasChange = true
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Backupvserver has changed for crvserver %s, starting update", crvserverName)
		crvserver.Backupvserver = d.Get("backupvserver").(string)
		hasChange = true
	}
	if d.HasChange("cachetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachetype has changed for crvserver %s, starting update", crvserverName)
		crvserver.Cachetype = d.Get("cachetype").(string)
		hasChange = true
	}
	if d.HasChange("cachevserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachevserver has changed for crvserver %s, starting update", crvserverName)
		crvserver.Cachevserver = d.Get("cachevserver").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clttimeout has changed for crvserver %s, starting update", crvserverName)
		crvserver.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for crvserver %s, starting update", crvserverName)
		crvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("destinationvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Destinationvserver has changed for crvserver %s, starting update", crvserverName)
		crvserver.Destinationvserver = d.Get("destinationvserver").(string)
		hasChange = true
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG]  citrixadc-provider: Disableprimaryondown has changed for crvserver %s, starting update", crvserverName)
		crvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
		hasChange = true
	}
	if d.HasChange("dnsvservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsvservername has changed for crvserver %s, starting update", crvserverName)
		crvserver.Dnsvservername = d.Get("dnsvservername").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for crvserver %s, starting update", crvserverName)
		crvserver.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  citrixadc-provider: Downstateflush has changed for crvserver %s, starting update", crvserverName)
		crvserver.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("format") {
		log.Printf("[DEBUG]  citrixadc-provider: Format has changed for crvserver %s, starting update", crvserverName)
		crvserver.Format = d.Get("format").(string)
		hasChange = true
	}
	if d.HasChange("ghost") {
		log.Printf("[DEBUG]  citrixadc-provider: Ghost has changed for crvserver %s, starting update", crvserverName)
		crvserver.Ghost = d.Get("ghost").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpprofilename has changed for crvserver %s, starting update", crvserverName)
		crvserver.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpvsrresponse has changed for crvserver %s, starting update", crvserverName)
		crvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
		hasChange = true
	}
	if d.HasChange("ipset") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipset has changed for crvserver %s, starting update", crvserverName)
		crvserver.Ipset = d.Get("ipset").(string)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv46 has changed for crvserver %s, starting update", crvserverName)
		crvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG]  citrixadc-provider: L2conn has changed for crvserver %s, starting update", crvserverName)
		crvserver.L2conn = d.Get("l2conn").(string)
		hasChange = true
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG]  citrixadc-provider: Listenpolicy has changed for crvserver %s, starting update", crvserverName)
		crvserver.Listenpolicy = d.Get("listenpolicy").(string)
		hasChange = true
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Listenpriority has changed for crvserver %s, starting update", crvserverName)
		crvserver.Listenpriority = d.Get("listenpriority").(int)
		hasChange = true
	}
	if d.HasChange("map") {
		log.Printf("[DEBUG]  citrixadc-provider: Map has changed for crvserver %s, starting update", crvserverName)
		crvserver.Map = d.Get("map").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for crvserver %s, starting update", crvserverName)
		crvserver.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("onpolicymatch") {
		log.Printf("[DEBUG]  citrixadc-provider: Onpolicymatch has changed for crvserver %s, starting update", crvserverName)
		crvserver.Onpolicymatch = d.Get("onpolicymatch").(string)
		hasChange = true
	}
	if d.HasChange("originusip") {
		log.Printf("[DEBUG]  citrixadc-provider: Originusip has changed for crvserver %s, starting update", crvserverName)
		crvserver.Originusip = d.Get("originusip").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for crvserver %s, starting update", crvserverName)
		crvserver.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("precedence") {
		log.Printf("[DEBUG]  citrixadc-provider: Precedence has changed for crvserver %s, starting update", crvserverName)
		crvserver.Precedence = d.Get("precedence").(string)
		hasChange = true
	}
	if d.HasChange("probeport") {
		log.Printf("[DEBUG]  citrixadc-provider: Probeport has changed for crvserver %s, starting update", crvserverName)
		crvserver.Probeport = d.Get("probeport").(int)
		hasChange = true
	}
	if d.HasChange("probeprotocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Probeprotocol has changed for crvserver %s, starting update", crvserverName)
		crvserver.Probeprotocol = d.Get("probeprotocol").(string)
		hasChange = true
	}
	if d.HasChange("probesuccessresponsecode") {
		log.Printf("[DEBUG]  citrixadc-provider: Probesuccessresponsecode has changed for crvserver %s, starting update", crvserverName)
		crvserver.Probesuccessresponsecode = d.Get("probesuccessresponsecode").(string)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG]  citrixadc-provider: Range has changed for crvserver %s, starting update", crvserverName)
		crvserver.Range = d.Get("range").(int)
		hasChange = true
	}
	if d.HasChange("redirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirect has changed for crvserver %s, starting update", crvserverName)
		crvserver.Redirect = d.Get("redirect").(string)
		hasChange = true
	}
	if d.HasChange("redirecturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirecturl has changed for crvserver %s, starting update", crvserverName)
		crvserver.Redirecturl = d.Get("redirecturl").(string)
		hasChange = true
	}
	if d.HasChange("reuse") {
		log.Printf("[DEBUG]  citrixadc-provider: Reuse has changed for crvserver %s, starting update", crvserverName)
		crvserver.Reuse = d.Get("reuse").(string)
		hasChange = true
	}
	if d.HasChange("rhistate") {
		log.Printf("[DEBUG]  citrixadc-provider: Rhistate has changed for crvserver %s, starting update", crvserverName)
		crvserver.Rhistate = d.Get("rhistate").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for crvserver %s, starting update", crvserverName)
		crvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sopersistencetimeout has changed for crvserver %s, starting update", crvserverName)
		crvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sothreshold has changed for crvserver %s, starting update", crvserverName)
		crvserver.Sothreshold = d.Get("sothreshold").(int)
		hasChange = true
	}
	if d.HasChange("srcipexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipexpr has changed for crvserver %s, starting update", crvserverName)
		crvserver.Srcipexpr = d.Get("srcipexpr").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for crvserver %s, starting update", crvserverName)
		crvserver.State = d.Get("state").(string)
		stateChange = true
	}
	if d.HasChange("tcpprobeport") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprobeport has changed for crvserver %s, starting update", crvserverName)
		crvserver.Tcpprobeport = d.Get("tcpprobeport").(int)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofilename has changed for crvserver %s, starting update", crvserverName)
		crvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for crvserver %s, starting update", crvserverName)
		crvserver.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("useoriginipportforcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Useoriginipportforcache has changed for crvserver %s, starting update", crvserverName)
		crvserver.Useoriginipportforcache = d.Get("useoriginipportforcache").(string)
		hasChange = true
	}
	if d.HasChange("useportrange") {
		log.Printf("[DEBUG]  citrixadc-provider: Useportrange has changed for crvserver %s, starting update", crvserverName)
		crvserver.Useportrange = d.Get("useportrange").(string)
		hasChange = true
	}
	if d.HasChange("via") {
		log.Printf("[DEBUG]  citrixadc-provider: Via has changed for crvserver %s, starting update", crvserverName)
		crvserver.Via = d.Get("via").(string)
		hasChange = true
	}
	if stateChange {
		err := doCrvserverStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling cs vserver %s", crvserverName)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Crvserver.Type(), crvserverName, &crvserver)
		if err != nil {
			return fmt.Errorf("Error updating crvserver %s", crvserverName)
		}
	}
	return readCrvserverFunc(d, meta)
}

func deleteCrvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCrvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	crvserverName := d.Id()
	err := client.DeleteResource(service.Crvserver.Type(), crvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
func doCrvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doLbvserverStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	crvserver := cr.Crvserver{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Crvserver.Type(), crvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Crvserver.Type(), crvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
