package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsip() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsipFunc,
		ReadContext:   readNsipFunc,
		UpdateContext: updateNsipFunc,
		DeleteContext: deleteNsipFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"arpowner": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"advertiseondefaultpartition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arpresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bgp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrementttl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostroute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostrtgw": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mgmtaccess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"networkroute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospfarea": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ospflsatype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownerdownresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"restrictaccess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"telnet": {
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
			"vrid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserverrhilevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserverrhimode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mptcpadvertise": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsipFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	var ipaddress string
	ipaddress = d.Get("ipaddress").(string)

	nsip := ns.Nsip{
		Advertiseondefaultpartition: d.Get("advertiseondefaultpartition").(string),
		Arp:                         d.Get("arp").(string),
		Arpresponse:                 d.Get("arpresponse").(string),
		Bgp:                         d.Get("bgp").(string),
		Decrementttl:                d.Get("decrementttl").(string),
		Dynamicrouting:              d.Get("dynamicrouting").(string),
		Ftp:                         d.Get("ftp").(string),
		Gui:                         d.Get("gui").(string),
		Hostroute:                   d.Get("hostroute").(string),
		Hostrtgw:                    d.Get("hostrtgw").(string),
		Icmp:                        d.Get("icmp").(string),
		Icmpresponse:                d.Get("icmpresponse").(string),
		Ipaddress:                   ipaddress,
		Mgmtaccess:                  d.Get("mgmtaccess").(string),
		Netmask:                     d.Get("netmask").(string),
		Networkroute:                d.Get("networkroute").(string),
		Ospf:                        d.Get("ospf").(string),
		Ospflsatype:                 d.Get("ospflsatype").(string),
		Ownerdownresponse:           d.Get("ownerdownresponse").(string),
		Restrictaccess:              d.Get("restrictaccess").(string),
		Rip:                         d.Get("rip").(string),
		Snmp:                        d.Get("snmp").(string),
		Ssh:                         d.Get("ssh").(string),
		State:                       d.Get("state").(string),
		Telnet:                      d.Get("telnet").(string),
		Type:                        d.Get("type").(string),
		Vserver:                     d.Get("vserver").(string),
		Vserverrhilevel:             d.Get("vserverrhilevel").(string),
		Mptcpadvertise:              d.Get("mptcpadvertise").(string),
	}
	if raw := d.GetRawConfig().GetAttr("arpowner"); !raw.IsNull() {
		nsip.Arpowner = intPtr(d.Get("arpowner").(int))
	}
	if raw := d.GetRawConfig().GetAttr("metric"); !raw.IsNull() {
		nsip.Metric = intPtr(d.Get("metric").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ospfarea"); !raw.IsNull() {
		nsip.Ospfarea = intPtr(d.Get("ospfarea").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tag"); !raw.IsNull() {
		nsip.Tag = intPtr(d.Get("tag").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nsip.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vrid"); !raw.IsNull() {
		nsip.Vrid = intPtr(d.Get("vrid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ownernode"); !raw.IsNull() {
		nsip.Ownernode = intPtr(d.Get("ownernode").(int))
	}

	_, err := client.AddResource(service.Nsip.Type(), ipaddress, &nsip)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ipaddress)

	return readNsipFunc(ctx, d, meta)
}

func readNsipFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	nsipName := d.Id()
	trafficDomain := 0
	log.Printf("[DEBUG] citrixadc-provider: Reading nsip state %s", nsipName)
	argsMap := make(map[string]string)
	if val, ok := d.GetOk("td"); ok {
		trafficDomain = val.(int)
	}
	argsMap["td"] = fmt.Sprintf("%d", trafficDomain)
	findParams := service.FindParams{
		ResourceType:             service.Nsip.Type(),
		ResourceName:             nsipName,
		ResourceMissingErrorCode: 258,
		ArgsMap:                  argsMap,
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)
	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing nsip state %s", nsipName)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == nsipName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nsip state %s", nsipName)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]
	d.Set("advertiseondefaultpartition", data["advertiseondefaultpartition"])
	setToInt("arpowner", d, data["arpowner"])
	d.Set("arp", data["arp"])
	d.Set("arpresponse", data["arpresponse"])
	d.Set("bgp", data["bgp"])
	d.Set("decrementttl", data["decrementttl"])
	d.Set("dynamicrouting", data["dynamicrouting"])
	d.Set("ftp", data["ftp"])
	d.Set("gui", data["gui"])
	d.Set("hostroute", data["hostroute"])
	d.Set("hostrtgw", data["hostrtgw"])
	d.Set("icmp", data["icmp"])
	d.Set("icmpresponse", data["icmpresponse"])
	d.Set("ipaddress", data["ipaddress"])
	setToInt("metric", d, data["metric"])
	d.Set("mgmtaccess", data["mgmtaccess"])
	d.Set("netmask", data["netmask"])
	d.Set("networkroute", data["networkroute"])
	d.Set("ospf", data["ospf"])
	setToInt("ospfarea", d, data["ospfarea"])
	d.Set("ospflsatype", data["ospflsatype"])
	d.Set("ownerdownresponse", data["ownerdownresponse"])
	setToInt("ownernode", d, data["ownernode"])
	d.Set("restrictaccess", data["restrictaccess"])
	d.Set("rip", data["rip"])
	d.Set("snmp", data["snmp"])
	d.Set("ssh", data["ssh"])
	d.Set("state", data["state"])
	setToInt("tag", d, data["tag"])
	setToInt("td", d, data["td"])
	d.Set("telnet", data["telnet"])
	d.Set("type", data["type"])
	setToInt("vrid", d, data["vrid"])
	d.Set("vserver", data["vserver"])
	d.Set("vserverrhilevel", data["vserverrhilevel"])
	d.Set("mptcpadvertise", data["mptcpadvertise"])

	return nil

}

func updateNsipFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	ipaddress := d.Get("ipaddress").(string)

	nsip := ns.Nsip{
		Ipaddress: d.Get("ipaddress").(string),
	}
	stateChange := false
	hasChange := false
	if d.HasChange("arpowner") {
		log.Printf("[DEBUG]  citrixadc-provider: Arpowner has changed for nsip, starting update")
		nsip.Arpowner = intPtr(d.Get("arpowner").(int))
		hasChange = true
	}
	if d.HasChange("advertiseondefaultpartition") {
		log.Printf("[DEBUG]  citrixadc-provider: Advertiseondefaultpartition has changed for nsip %s, starting update", ipaddress)
		nsip.Advertiseondefaultpartition = d.Get("advertiseondefaultpartition").(string)
		hasChange = true
	}
	if d.HasChange("arp") {
		log.Printf("[DEBUG]  citrixadc-provider: Arp has changed for nsip %s, starting update", ipaddress)
		nsip.Arp = d.Get("arp").(string)
		hasChange = true
	}
	if d.HasChange("arpresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Arpresponse has changed for nsip %s, starting update", ipaddress)
		nsip.Arpresponse = d.Get("arpresponse").(string)
		hasChange = true
	}
	if d.HasChange("bgp") {
		log.Printf("[DEBUG]  citrixadc-provider: Bgp has changed for nsip %s, starting update", ipaddress)
		nsip.Bgp = d.Get("bgp").(string)
		hasChange = true
	}
	if d.HasChange("decrementttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Decrementttl has changed for nsip %s, starting update", ipaddress)
		nsip.Decrementttl = d.Get("decrementttl").(string)
		hasChange = true
	}
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for nsip %s, starting update", ipaddress)
		nsip.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("ftp") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftp has changed for nsip %s, starting update", ipaddress)
		nsip.Ftp = d.Get("ftp").(string)
		hasChange = true
	}
	if d.HasChange("gui") {
		log.Printf("[DEBUG]  citrixadc-provider: Gui has changed for nsip %s, starting update", ipaddress)
		nsip.Gui = d.Get("gui").(string)
		hasChange = true
	}
	if d.HasChange("hostroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostroute has changed for nsip %s, starting update", ipaddress)
		nsip.Hostroute = d.Get("hostroute").(string)
		hasChange = true
	}
	if d.HasChange("hostrtgw") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostrtgw has changed for nsip %s, starting update", ipaddress)
		nsip.Hostrtgw = d.Get("hostrtgw").(string)
		hasChange = true
	}
	if d.HasChange("icmp") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmp has changed for nsip %s, starting update", ipaddress)
		nsip.Icmp = d.Get("icmp").(string)
		hasChange = true
	}
	if d.HasChange("icmpresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpresponse has changed for nsip %s, starting update", ipaddress)
		nsip.Icmpresponse = d.Get("icmpresponse").(string)
		hasChange = true
	}
	if d.HasChange("metric") {
		log.Printf("[DEBUG]  citrixadc-provider: Metric has changed for nsip %s, starting update", ipaddress)
		nsip.Metric = intPtr(d.Get("metric").(int))
		hasChange = true
	}
	if d.HasChange("mgmtaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Mgmtaccess has changed for nsip %s, starting update", ipaddress)
		nsip.Mgmtaccess = d.Get("mgmtaccess").(string)
		hasChange = true
	}
	if d.HasChange("networkroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Networkroute has changed for nsip %s, starting update", ipaddress)
		nsip.Networkroute = d.Get("networkroute").(string)
		hasChange = true
	}
	if d.HasChange("ospf") {
		log.Printf("[DEBUG]  citrixadc-provider: Ospf has changed for nsip %s, starting update", ipaddress)
		nsip.Ospf = d.Get("ospf").(string)
		hasChange = true
	}
	if d.HasChange("ospfarea") {
		log.Printf("[DEBUG]  citrixadc-provider: Ospfarea has changed for nsip %s, starting update", ipaddress)
		nsip.Ospfarea = intPtr(d.Get("ospfarea").(int))
		hasChange = true
	}
	if d.HasChange("ospflsatype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ospflsatype has changed for nsip %s, starting update", ipaddress)
		nsip.Ospflsatype = d.Get("ospflsatype").(string)
		hasChange = true
	}
	if d.HasChange("ownerdownresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownerdownresponse has changed for nsip %s, starting update", ipaddress)
		nsip.Ownerdownresponse = d.Get("ownerdownresponse").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for nsip %s, starting update", ipaddress)
		nsip.Ownernode = intPtr(d.Get("ownernode").(int))
		hasChange = true
	}
	if d.HasChange("restrictaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Restrictaccess has changed for nsip %s, starting update", ipaddress)
		nsip.Restrictaccess = d.Get("restrictaccess").(string)
		hasChange = true
	}
	if d.HasChange("rip") {
		log.Printf("[DEBUG]  citrixadc-provider: Rip has changed for nsip %s, starting update", ipaddress)
		nsip.Rip = d.Get("rip").(string)
		hasChange = true
	}
	if d.HasChange("snmp") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmp has changed for nsip %s, starting update", ipaddress)
		nsip.Snmp = d.Get("snmp").(string)
		hasChange = true
	}
	if d.HasChange("ssh") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssh has changed for nsip %s, starting update", ipaddress)
		nsip.Ssh = d.Get("ssh").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nsip %s, starting update", ipaddress)
		nsip.State = d.Get("state").(string)
		stateChange = true
	}
	if d.HasChange("tag") {
		log.Printf("[DEBUG]  citrixadc-provider: Tag has changed for nsip %s, starting update", ipaddress)
		nsip.Tag = intPtr(d.Get("tag").(int))
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nsip %s, starting update", ipaddress)
		nsip.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("telnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Telnet has changed for nsip %s, starting update", ipaddress)
		nsip.Telnet = d.Get("telnet").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for nsip %s, starting update", ipaddress)
		nsip.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("vrid") {
		log.Printf("[DEBUG]  citrixadc-provider: Vrid has changed for nsip %s, starting update", ipaddress)
		nsip.Vrid = intPtr(d.Get("vrid").(int))
		hasChange = true
	}
	if d.HasChange("vserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserver has changed for nsip %s, starting update", ipaddress)
		nsip.Vserver = d.Get("vserver").(string)
		hasChange = true
	}
	if d.HasChange("vserverrhilevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserverrhilevel has changed for nsip %s, starting update", ipaddress)
		nsip.Vserverrhilevel = d.Get("vserverrhilevel").(string)
		hasChange = true
	}
	if d.HasChange("mptcpadvertise") {
		log.Printf("[DEBUG]  citrixadc-provider: Mptcpadvertise has changed for nsip %s, starting update", ipaddress)
		nsip.Mptcpadvertise = d.Get("mptcpadvertise").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsip.Type(), ipaddress, &nsip)
		if err != nil {
			return diag.Errorf("Error updating nsip %s: %s", ipaddress, err.Error())
		}
	}

	if stateChange {
		err := doNsipStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling nsip %s", ipaddress)
		}
	}
	return readNsipFunc(ctx, d, meta)
}

func deleteNsipFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	ipaddress := d.Id()
	trafficDomain := 0
	if val, ok := d.GetOk("td"); ok {
		trafficDomain = val.(int)
	}
	argsMap := make(map[string]string)
	argsMap["td"] = fmt.Sprintf("%d", trafficDomain)
	err := client.DeleteResourceWithArgsMap(service.Nsip.Type(), ipaddress, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func doNsipStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doNsipStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	nsip := ns.Nsip{
		Ipaddress: d.Get("ipaddress").(string),
		Td:        intPtr(d.Get("td").(int)),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Nsip.Type(), nsip, "enable")
		if err != nil {
			return err
		}
		// Disable action
	} else if newstate == "DISABLED" {
		err := client.ActOnResource(service.Nsip.Type(), nsip, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
