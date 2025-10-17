package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCrvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCrvserverFunc,
		ReadContext:   readCrvserverFunc,
		UpdateContext: updateCrvserverFunc,
		DeleteContext: deleteCrvserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicetype": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backendssl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachevserver": {
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
			"destinationvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsvservername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ghost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpvsrresponse": {
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
			"map": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"onpolicymatch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"originusip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"precedence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"probesuccessresponsecode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"redirect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirecturl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rhistate": {
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
			"srcipexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprobeport": {
				Type:     schema.TypeInt,
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
			},
			"useoriginipportforcache": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useportrange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"via": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCrvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		Map:                      d.Get("map").(string),
		Name:                     crvserverName,
		Netprofile:               d.Get("netprofile").(string),
		Onpolicymatch:            d.Get("onpolicymatch").(string),
		Originusip:               d.Get("originusip").(string),
		Precedence:               d.Get("precedence").(string),
		Probeprotocol:            d.Get("probeprotocol").(string),
		Probesuccessresponsecode: d.Get("probesuccessresponsecode").(string),
		Redirect:                 d.Get("redirect").(string),
		Redirecturl:              d.Get("redirecturl").(string),
		Reuse:                    d.Get("reuse").(string),
		Rhistate:                 d.Get("rhistate").(string),
		Servicetype:              d.Get("servicetype").(string),
		Srcipexpr:                d.Get("srcipexpr").(string),
		State:                    d.Get("state").(string),
		Tcpprofilename:           d.Get("tcpprofilename").(string),
		Useoriginipportforcache:  d.Get("useoriginipportforcache").(string),
		Useportrange:             d.Get("useportrange").(string),
		Via:                      d.Get("via").(string),
	}

	if raw := d.GetRawConfig().GetAttr("clttimeout"); !raw.IsNull() {
		crvserver.Clttimeout = intPtr(d.Get("clttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("listenpriority"); !raw.IsNull() {
		crvserver.Listenpriority = intPtr(d.Get("listenpriority").(int))
	}
	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		crvserver.Port = intPtr(d.Get("port").(int))
	}
	if raw := d.GetRawConfig().GetAttr("probeport"); !raw.IsNull() {
		crvserver.Probeport = intPtr(d.Get("probeport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("range"); !raw.IsNull() {
		crvserver.Range = intPtr(d.Get("range").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sopersistencetimeout"); !raw.IsNull() {
		crvserver.Sopersistencetimeout = intPtr(d.Get("sopersistencetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sothreshold"); !raw.IsNull() {
		crvserver.Sothreshold = intPtr(d.Get("sothreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpprobeport"); !raw.IsNull() {
		crvserver.Tcpprobeport = intPtr(d.Get("tcpprobeport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		crvserver.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Crvserver.Type(), crvserverName, &crvserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(crvserverName)

	return readCrvserverFunc(ctx, d, meta)
}

func readCrvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	setToInt("clttimeout", d, data["clttimeout"])
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
	setToInt("listenpriority", d, data["listenpriority"])
	d.Set("map", data["map"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("onpolicymatch", data["onpolicymatch"])
	d.Set("originusip", data["originusip"])
	setToInt("port", d, data["port"])
	d.Set("precedence", data["precedence"])
	setToInt("probeport", d, data["probeport"])
	d.Set("probeprotocol", data["probeprotocol"])
	d.Set("probesuccessresponsecode", data["probesuccessresponsecode"])
	setToInt("range", d, data["range"])
	d.Set("redirect", data["redirect"])
	d.Set("redirecturl", data["redirecturl"])
	d.Set("reuse", data["reuse"])
	d.Set("rhistate", data["rhistate"])
	d.Set("servicetype", data["servicetype"])
	setToInt("sopersistencetimeout", d, data["sopersistencetimeout"])
	setToInt("sothreshold", d, data["sothreshold"])
	d.Set("srcipexpr", data["srcipexpr"])
	d.Set("state", data["state"])
	setToInt("tcpprobeport", d, data["tcpprobeport"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	setToInt("td", d, data["td"])
	d.Set("useoriginipportforcache", data["useoriginipportforcache"])
	d.Set("useportrange", data["useportrange"])
	d.Set("via", data["via"])

	return nil

}

func updateCrvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		crvserver.Clttimeout = intPtr(d.Get("clttimeout").(int))
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
		crvserver.Listenpriority = intPtr(d.Get("listenpriority").(int))
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
		crvserver.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}
	if d.HasChange("precedence") {
		log.Printf("[DEBUG]  citrixadc-provider: Precedence has changed for crvserver %s, starting update", crvserverName)
		crvserver.Precedence = d.Get("precedence").(string)
		hasChange = true
	}
	if d.HasChange("probeport") {
		log.Printf("[DEBUG]  citrixadc-provider: Probeport has changed for crvserver %s, starting update", crvserverName)
		crvserver.Probeport = intPtr(d.Get("probeport").(int))
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
		crvserver.Range = intPtr(d.Get("range").(int))
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
		crvserver.Sopersistencetimeout = intPtr(d.Get("sopersistencetimeout").(int))
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sothreshold has changed for crvserver %s, starting update", crvserverName)
		crvserver.Sothreshold = intPtr(d.Get("sothreshold").(int))
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
		crvserver.Tcpprobeport = intPtr(d.Get("tcpprobeport").(int))
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofilename has changed for crvserver %s, starting update", crvserverName)
		crvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for crvserver %s, starting update", crvserverName)
		crvserver.Td = intPtr(d.Get("td").(int))
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
			return diag.Errorf("Error enabling/disabling cs vserver %s", crvserverName)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Crvserver.Type(), crvserverName, &crvserver)
		if err != nil {
			return diag.Errorf("Error updating crvserver %s", crvserverName)
		}
	}
	return readCrvserverFunc(ctx, d, meta)
}

func deleteCrvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCrvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	crvserverName := d.Id()
	err := client.DeleteResource(service.Crvserver.Type(), crvserverName)
	if err != nil {
		return diag.FromErr(err)
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
