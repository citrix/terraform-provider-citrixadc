package citrixadc

import (
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"errors"
	"fmt"
	"log"
	"net/url"
)

// nsip6 struct is defined here to add MPTCPadvertise support.
// Once this attribute available in the main builds, respective go-notro file will be taken care.
type nsip6 struct {
	Advertiseondefaultpartition string      `json:"advertiseondefaultpartition,omitempty"`
	Curstate                    string      `json:"curstate,omitempty"`
	Decrementhoplimit           string      `json:"decrementhoplimit,omitempty"`
	Dynamicrouting              string      `json:"dynamicrouting,omitempty"`
	Ftp                         string      `json:"ftp,omitempty"`
	Gui                         string      `json:"gui,omitempty"`
	Hostroute                   string      `json:"hostroute,omitempty"`
	Icmp                        string      `json:"icmp,omitempty"`
	Ip6hostrtgw                 string      `json:"ip6hostrtgw,omitempty"`
	Iptype                      interface{} `json:"iptype,omitempty"`
	Ipv6address                 string      `json:"ipv6address,omitempty"`
	Map                         string      `json:"map,omitempty"`
	Metric                      int         `json:"metric,omitempty"`
	Mgmtaccess                  string      `json:"mgmtaccess,omitempty"`
	Nd                          string      `json:"nd,omitempty"`
	Networkroute                string      `json:"networkroute,omitempty"`
	Ospf6lsatype                string      `json:"ospf6lsatype,omitempty"`
	Ospfarea                    int         `json:"ospfarea,omitempty"`
	Ownerdownresponse           string      `json:"ownerdownresponse,omitempty"`
	Ownernode                   int         `json:"ownernode,omitempty"`
	Restrictaccess              string      `json:"restrictaccess,omitempty"`
	Scope                       string      `json:"scope,omitempty"`
	Snmp                        string      `json:"snmp,omitempty"`
	Ssh                         string      `json:"ssh,omitempty"`
	State                       string      `json:"state,omitempty"`
	Systemtype                  string      `json:"systemtype,omitempty"`
	Tag                         int         `json:"tag,omitempty"`
	Td                          int         `json:"td,omitempty"`
	Telnet                      string      `json:"telnet,omitempty"`
	Type                        string      `json:"type,omitempty"`
	Viprtadv2bsd                bool        `json:"viprtadv2bsd,omitempty"`
	Vipvsercount                int         `json:"vipvsercount,omitempty"`
	Vipvserdowncount            int         `json:"vipvserdowncount,omitempty"`
	Vlan                        int         `json:"vlan,omitempty"`
	Vrid6                       int         `json:"vrid6,omitempty"`
	Vserver                     string      `json:"vserver,omitempty"`
	Vserverrhilevel             string      `json:"vserverrhilevel,omitempty"`
	Mptcpadvertise              string      `json:"mptcpadvertise,omitempty"`
}

func resourceCitrixAdcNsip6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsip6Func,
		Read:          readNsip6Func,
		Update:        updateNsip6Func,
		Delete:        deleteNsip6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"advertiseondefaultpartition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrementhoplimit": &schema.Schema{
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
			"icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6hostrtgw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"nd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"networkroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6lsatype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospfarea": &schema.Schema{
				Type:     schema.TypeInt,
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
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vrid6": &schema.Schema{
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
			"mptcpadvertise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsip6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsip6Func")
	client := meta.(*NetScalerNitroClient).client
	ipv6address := d.Get("ipv6address").(string)
	nsip6 := nsip6{
		Advertiseondefaultpartition: d.Get("advertiseondefaultpartition").(string),
		Decrementhoplimit:           d.Get("decrementhoplimit").(string),
		Dynamicrouting:              d.Get("dynamicrouting").(string),
		Ftp:                         d.Get("ftp").(string),
		Gui:                         d.Get("gui").(string),
		Hostroute:                   d.Get("hostroute").(string),
		Icmp:                        d.Get("icmp").(string),
		Ip6hostrtgw:                 d.Get("ip6hostrtgw").(string),
		Ipv6address:                 d.Get("ipv6address").(string),
		Map:                         d.Get("map").(string),
		Metric:                      d.Get("metric").(int),
		Mgmtaccess:                  d.Get("mgmtaccess").(string),
		Nd:                          d.Get("nd").(string),
		Networkroute:                d.Get("networkroute").(string),
		Ospf6lsatype:                d.Get("ospf6lsatype").(string),
		Ospfarea:                    d.Get("ospfarea").(int),
		Ownerdownresponse:           d.Get("ownerdownresponse").(string),
		Ownernode:                   d.Get("ownernode").(int),
		Restrictaccess:              d.Get("restrictaccess").(string),
		Scope:                       d.Get("scope").(string),
		Snmp:                        d.Get("snmp").(string),
		Ssh:                         d.Get("ssh").(string),
		State:                       d.Get("state").(string),
		Tag:                         d.Get("tag").(int),
		Td:                          d.Get("td").(int),
		Telnet:                      d.Get("telnet").(string),
		Type:                        d.Get("type").(string),
		Vlan:                        d.Get("vlan").(int),
		Vrid6:                       d.Get("vrid6").(int),
		Vserver:                     d.Get("vserver").(string),
		Vserverrhilevel:             d.Get("vserverrhilevel").(string),
		Mptcpadvertise:              d.Get("mptcpadvertise").(string),
	}

	_, err := client.AddResource(netscaler.Nsip6.Type(), ipv6address, &nsip6)
	if err != nil {
		return err
	}

	d.SetId(ipv6address)

	err = readNsip6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsip6 but we can't read it ?? %s", ipv6address)
		return nil
	}
	return nil
}

func readNsip6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsip6Func")
	client := meta.(*NetScalerNitroClient).client
	ipv6address := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsip6 state %s", ipv6address)
	array, _ := client.FindAllResources(netscaler.Nsip6.Type())

	// Iterate over the retrieved addresses to find the particular ipv6address
	foundAddress := false
	foundIndex := -1
	for i, item := range array {
		if item["ipv6address"] == ipv6address {
			foundAddress = true
			foundIndex = i
			break
		}
	}
	if !foundAddress {
		log.Printf("[WARN] citrixadc-provider: Clearing nsip6 state %s", ipv6address)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := array[foundIndex]

	d.Set("advertiseondefaultpartition", data["advertiseondefaultpartition"])
	d.Set("decrementhoplimit", data["decrementhoplimit"])
	d.Set("dynamicrouting", data["dynamicrouting"])
	d.Set("ftp", data["ftp"])
	d.Set("gui", data["gui"])
	d.Set("hostroute", data["hostroute"])
	d.Set("icmp", data["icmp"])
	d.Set("ip6hostrtgw", data["ip6hostrtgw"])
	d.Set("ipv6address", data["ipv6address"])
	d.Set("map", data["map"])
	d.Set("metric", data["metric"])
	d.Set("mgmtaccess", data["mgmtaccess"])
	d.Set("nd", data["nd"])
	d.Set("networkroute", data["networkroute"])
	d.Set("ospf6lsatype", data["ospf6lsatype"])
	d.Set("ospfarea", data["ospfarea"])
	d.Set("ownerdownresponse", data["ownerdownresponse"])
	d.Set("ownernode", data["ownernode"])
	d.Set("restrictaccess", data["restrictaccess"])
	d.Set("scope", data["scope"])
	d.Set("snmp", data["snmp"])
	d.Set("ssh", data["ssh"])
	d.Set("state", data["state"])
	d.Set("tag", data["tag"])
	d.Set("td", data["td"])
	d.Set("telnet", data["telnet"])

	// Type is a special case
	// Need to add sanity check to make sure we don't parse the wrong value
	iptype := data["iptype"].([]interface{})
	if len(iptype) > 1 {
		return errors.New("Found iptype to contain more than one ip type")
	}
	d.Set("type", iptype[0].(string))

	d.Set("vlan", data["vlan"])
	d.Set("vrid6", data["vrid6"])
	d.Set("vserver", data["vserver"])
	d.Set("vserverrhilevel", data["vserverrhilevel"])
	d.Set("mptcpadvertise", data["mptcpadvertise"])

	return nil

}

func updateNsip6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsip6Func")
	client := meta.(*NetScalerNitroClient).client
	ipv6address := d.Get("ipv6address").(string)

	nsip6 := nsip6{
		Ipv6address: d.Get("ipv6address").(string),
	}
	hasChange := false
	if d.HasChange("advertiseondefaultpartition") {
		log.Printf("[DEBUG]  citrixadc-provider: Advertiseondefaultpartition has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Advertiseondefaultpartition = d.Get("advertiseondefaultpartition").(string)
		hasChange = true
	}
	if d.HasChange("decrementhoplimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Decrementhoplimit has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Decrementhoplimit = d.Get("decrementhoplimit").(string)
		hasChange = true
	}
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("ftp") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftp has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ftp = d.Get("ftp").(string)
		hasChange = true
	}
	if d.HasChange("gui") {
		log.Printf("[DEBUG]  citrixadc-provider: Gui has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Gui = d.Get("gui").(string)
		hasChange = true
	}
	if d.HasChange("hostroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostroute has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Hostroute = d.Get("hostroute").(string)
		hasChange = true
	}
	if d.HasChange("icmp") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmp has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Icmp = d.Get("icmp").(string)
		hasChange = true
	}
	if d.HasChange("ip6hostrtgw") {
		log.Printf("[DEBUG]  citrixadc-provider: Ip6hostrtgw has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ip6hostrtgw = d.Get("ip6hostrtgw").(string)
		hasChange = true
	}
	if d.HasChange("map") {
		log.Printf("[DEBUG]  citrixadc-provider: Map has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Map = d.Get("map").(string)
		hasChange = true
	}
	if d.HasChange("metric") {
		log.Printf("[DEBUG]  citrixadc-provider: Metric has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Metric = d.Get("metric").(int)
		hasChange = true
	}
	if d.HasChange("mgmtaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Mgmtaccess has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Mgmtaccess = d.Get("mgmtaccess").(string)
		hasChange = true
	}
	if d.HasChange("nd") {
		log.Printf("[DEBUG]  citrixadc-provider: Nd has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Nd = d.Get("nd").(string)
		hasChange = true
	}
	if d.HasChange("networkroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Networkroute has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Networkroute = d.Get("networkroute").(string)
		hasChange = true
	}
	if d.HasChange("ospf6lsatype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ospf6lsatype has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ospf6lsatype = d.Get("ospf6lsatype").(string)
		hasChange = true
	}
	if d.HasChange("ospfarea") {
		log.Printf("[DEBUG]  citrixadc-provider: Ospfarea has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ospfarea = d.Get("ospfarea").(int)
		hasChange = true
	}
	if d.HasChange("ownerdownresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownerdownresponse has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ownerdownresponse = d.Get("ownerdownresponse").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}
	if d.HasChange("restrictaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Restrictaccess has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Restrictaccess = d.Get("restrictaccess").(string)
		hasChange = true
	}
	if d.HasChange("scope") {
		log.Printf("[DEBUG]  citrixadc-provider: Scope has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Scope = d.Get("scope").(string)
		hasChange = true
	}
	if d.HasChange("snmp") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmp has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Snmp = d.Get("snmp").(string)
		hasChange = true
	}
	if d.HasChange("ssh") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssh has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Ssh = d.Get("ssh").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nsip6 %s, starting update", ipv6address)
		nsip6.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("tag") {
		log.Printf("[DEBUG]  citrixadc-provider: Tag has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Tag = d.Get("tag").(int)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("telnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Telnet has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Telnet = d.Get("telnet").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	if d.HasChange("vrid6") {
		log.Printf("[DEBUG]  citrixadc-provider: Vrid6 has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Vrid6 = d.Get("vrid6").(int)
		hasChange = true
	}
	if d.HasChange("vserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserver has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Vserver = d.Get("vserver").(string)
		hasChange = true
	}
	if d.HasChange("vserverrhilevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserverrhilevel has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Vserverrhilevel = d.Get("vserverrhilevel").(string)
		hasChange = true
	}
	if d.HasChange("mptcpadvertise") {
		log.Printf("[DEBUG]  citrixadc-provider: Mptcpadvertise has changed for nsip6 %s, starting update", ipv6address)
		nsip6.Mptcpadvertise = d.Get("mptcpadvertise").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Nsip6.Type(), "", &nsip6)
		if err != nil {
			return fmt.Errorf("Error updating nsip6 %s", ipv6address)
		}
	}
	return readNsip6Func(d, meta)
}

func deleteNsip6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsip6Func")
	client := meta.(*NetScalerNitroClient).client
	ipv6address := d.Id()
	argsMap := make(map[string]string)
	argsMap["ipv6address"] = url.QueryEscape(ipv6address)
	if val, ok := d.GetOk("td"); ok {
		argsMap["td"] = val.(string)
	}
	err := client.DeleteResourceWithArgsMap(netscaler.Nsip6.Type(), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
