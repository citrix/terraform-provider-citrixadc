package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsip() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsipFunc,
		Read:          readNsipFunc,
		Update:        updateNsipFunc,
		Delete:        deleteNsipFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"advertiseondefaultpartition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arpresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bgp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrementttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamicrouting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostrtgw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mgmtaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"networkroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospfarea": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ospflsatype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownerdownresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownernode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"restrictaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"telnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserverrhilevel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserverrhimode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsipFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	var ipaddress string
	if v, ok := d.GetOk("ipaddress"); ok {
		ipaddress = v.(string)
	} else {
		ipaddress = resource.PrefixedUniqueId("tf-nsip-")
		d.Set("ipaddress", ipaddress)
	}
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
		Metric:                      d.Get("metric").(int),
		Mgmtaccess:                  d.Get("mgmtaccess").(string),
		Netmask:                     d.Get("netmask").(string),
		Networkroute:                d.Get("networkroute").(string),
		Ospf:                        d.Get("ospf").(string),
		Ospfarea:                    d.Get("ospfarea").(int),
		Ospflsatype:                 d.Get("ospflsatype").(string),
		Ownerdownresponse:           d.Get("ownerdownresponse").(string),
		Ownernode:                   d.Get("ownernode").(int),
		Restrictaccess:              d.Get("restrictaccess").(string),
		Rip:                         d.Get("rip").(string),
		Snmp:                        d.Get("snmp").(string),
		Ssh:                         d.Get("ssh").(string),
		State:                       d.Get("state").(string),
		Tag:                         d.Get("tag").(int),
		Td:                          d.Get("td").(int),
		Telnet:                      d.Get("telnet").(string),
		Type:                        d.Get("type").(string),
		Vrid:                        d.Get("vrid").(int),
		Vserver:                     d.Get("vserver").(string),
		Vserverrhilevel:             d.Get("vserverrhilevel").(string),
		Vserverrhimode:              d.Get("vserverrhimode").(string),
	}

	_, err := client.AddResource(netscaler.Nsip.Type(), ipaddress, &nsip)
	if err != nil {
		return err
	}

	d.SetId(ipaddress)

	err = readNsipFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsip but we can't read it ?? %s", ipaddress)
		return nil
	}
	return nil
}

func readNsipFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	nsipName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsip state %s", nsipName)
	data, err := client.FindResource(netscaler.Nsip.Type(), nsipName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsip state %s", nsipName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("advertiseondefaultpartition", data["advertiseondefaultpartition"])
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
	d.Set("metric", data["metric"])
	d.Set("mgmtaccess", data["mgmtaccess"])
	d.Set("netmask", data["netmask"])
	d.Set("networkroute", data["networkroute"])
	d.Set("ospf", data["ospf"])
	d.Set("ospfarea", data["ospfarea"])
	d.Set("ospflsatype", data["ospflsatype"])
	d.Set("ownerdownresponse", data["ownerdownresponse"])
	d.Set("ownernode", data["ownernode"])
	d.Set("restrictaccess", data["restrictaccess"])
	d.Set("rip", data["rip"])
	d.Set("snmp", data["snmp"])
	d.Set("ssh", data["ssh"])
	d.Set("state", data["state"])
	d.Set("tag", data["tag"])
	d.Set("td", data["td"])
	d.Set("telnet", data["telnet"])
	d.Set("type", data["type"])
	d.Set("vrid", data["vrid"])
	d.Set("vserver", data["vserver"])
	d.Set("vserverrhilevel", data["vserverrhilevel"])
	d.Set("vserverrhimode", data["vserverrhimode"])

	return nil

}

func updateNsipFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	ipaddress := d.Get("ipaddress").(string)

	nsip := ns.Nsip{
		Ipaddress: d.Get("ipaddress").(string),
	}
	stateChange := false
	hasChange := false
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
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for nsip %s, starting update", ipaddress)
		nsip.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("metric") {
		log.Printf("[DEBUG]  citrixadc-provider: Metric has changed for nsip %s, starting update", ipaddress)
		nsip.Metric = d.Get("metric").(int)
		hasChange = true
	}
	if d.HasChange("mgmtaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Mgmtaccess has changed for nsip %s, starting update", ipaddress)
		nsip.Mgmtaccess = d.Get("mgmtaccess").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for nsip %s, starting update", ipaddress)
		nsip.Netmask = d.Get("netmask").(string)
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
		nsip.Ospfarea = d.Get("ospfarea").(int)
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
		nsip.Ownernode = d.Get("ownernode").(int)
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
		nsip.Tag = d.Get("tag").(int)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nsip %s, starting update", ipaddress)
		nsip.Td = d.Get("td").(int)
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
		nsip.Vrid = d.Get("vrid").(int)
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
	if d.HasChange("vserverrhimode") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserverrhimode has changed for nsip %s, starting update", ipaddress)
		nsip.Vserverrhimode = d.Get("vserverrhimode").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Nsip.Type(), ipaddress, &nsip)
		if err != nil {
			return fmt.Errorf("Error updating nsip %s: %s", ipaddress, err.Error())
		}
	}

	if stateChange {
		err := doNsipStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling nsip %s", ipaddress)
		}
	}
	return readNsipFunc(d, meta)
}

func deleteNsipFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsipFunc")
	client := meta.(*NetScalerNitroClient).client
	ipaddress := d.Id()
	err := client.DeleteResource(netscaler.Nsip.Type(), ipaddress)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func doNsipStateChange(d *schema.ResourceData, client *netscaler.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doNsipStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	nsip := ns.Nsip{
		Ipaddress: d.Get("ipaddress").(string),
		Td:        d.Get("td").(int),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(netscaler.Nsip.Type(), nsip, "enable")
		if err != nil {
			return err
		}
		// Disable action
	} else if newstate == "DISABLED" {
		err := client.ActOnResource(netscaler.Nsip.Type(), nsip, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
