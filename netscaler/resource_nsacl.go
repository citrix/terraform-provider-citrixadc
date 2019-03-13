package netscaler

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerNsacl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsaclFunc,
		Read:          readNsaclFunc,
		Update:        updateNsaclFunc,
		Delete:        deleteNsaclFunc,
		Schema: map[string]*schema.Schema{
			"aclaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aclname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfdhash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"established": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocolnumber": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ratelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcipop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmacmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stateful": &schema.Schema{
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
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsaclName string
	if v, ok := d.GetOk("aclname"); ok {
		nsaclName = v.(string)
	} else {
		nsaclName = resource.PrefixedUniqueId("tf-nsacl-")
		d.Set("aclname", nsaclName)
	}
	destip := false
	destport := false
	srcip := false
	srcport := false
	if d.Get("destipval") != nil && d.Get("destipval") != "" {
		destip = true
	}
	if d.Get("destportval") != nil && d.Get("destportval") != "" {
		destport = true
	}
	if d.Get("srcipval") != nil && d.Get("srcipval") != "" {
		srcip = true
	}
	if d.Get("srcportval") != nil && d.Get("srcportval") != "" {
		srcport = true
	}

	if d.Get("destipop") != nil && d.Get("destipval") == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if d.Get("destportop") != nil && d.Get("destipval") == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if d.Get("srcipop") != nil && d.Get("srcipval") == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcipop without srcipval", nsaclName)
	}
	if d.Get("srcportop") != nil && d.Get("srcportval") == nil {
		return fmt.Errorf("Error in nsacl spec %s cannot have srcportop without srcportval", nsaclName)
	}

	nsacl := ns.Nsacl{
		Aclaction:      d.Get("aclaction").(string),
		Aclname:        d.Get("aclname").(string),
		Destip:         destip,
		Destipop:       d.Get("destipop").(string),
		Destipval:      d.Get("destipval").(string),
		Destport:       destport,
		Destportop:     d.Get("destportop").(string),
		Destportval:    d.Get("destportval").(string),
		Dfdhash:        d.Get("dfdhash").(string),
		Established:    d.Get("established").(bool),
		Icmpcode:       d.Get("icmpcode").(int),
		Icmptype:       d.Get("icmptype").(int),
		Interface:      d.Get("interface").(string),
		Logstate:       d.Get("logstate").(string),
		Priority:       d.Get("priority").(int),
		Protocol:       d.Get("protocol").(string),
		Protocolnumber: d.Get("protocolnumber").(int),
		Ratelimit:      d.Get("ratelimit").(int),
		Srcip:          srcip,
		Srcipop:        d.Get("srcipop").(string),
		Srcipval:       d.Get("srcipval").(string),
		Srcmac:         d.Get("srcmac").(string),
		Srcmacmask:     d.Get("srcmacmask").(string),
		Srcport:        srcport,
		Srcportop:      d.Get("srcportop").(string),
		Srcportval:     d.Get("srcportval").(string),
		State:          d.Get("state").(string),
		Stateful:       d.Get("stateful").(string),
		Td:             d.Get("td").(int),
		Ttl:            d.Get("ttl").(int),
		Vlan:           d.Get("vlan").(int),
		Vxlan:          d.Get("vxlan").(int),
	}

	_, err := client.AddResource(netscaler.Nsacl.Type(), nsaclName, &nsacl)
	if err != nil {
		return err
	}

	d.SetId(nsaclName)

	err = readNsaclFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsacl but we can't read it ?? %s", nsaclName)
		return nil
	}
	return nil
}

func readNsaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading nsacl state %s", nsaclName)
	data, err := client.FindResource(netscaler.Nsacl.Type(), nsaclName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing nsacl state %s", nsaclName)
		d.SetId("")
		return nil
	}
	d.Set("aclname", data["aclname"])
	d.Set("aclaction", data["aclaction"])
	d.Set("aclname", data["aclname"])
	d.Set("destip", data["destip"])
	d.Set("destipop", data["destipop"])
	d.Set("destipval", data["destipval"])
	d.Set("destport", data["destport"])
	d.Set("destportop", data["destportop"])
	d.Set("destportval", data["destportval"])
	d.Set("dfdhash", data["dfdhash"])
	d.Set("established", data["established"])
	d.Set("icmpcode", data["icmpcode"])
	d.Set("icmptype", data["icmptype"])
	d.Set("interface", data["interface"])
	d.Set("logstate", data["logstate"])
	d.Set("priority", data["priority"])
	d.Set("protocol", data["protocol"])
	d.Set("protocolnumber", data["protocolnumber"])
	d.Set("ratelimit", data["ratelimit"])
	d.Set("srcip", data["srcip"])
	d.Set("srcipop", data["srcipop"])
	d.Set("srcipval", data["srcipval"])
	d.Set("srcmac", data["srcmac"])
	d.Set("srcmacmask", data["srcmacmask"])
	d.Set("srcport", data["srcport"])
	d.Set("srcportop", data["srcportop"])
	d.Set("srcportval", data["srcportval"])
	d.Set("state", data["state"])
	d.Set("stateful", data["stateful"])
	d.Set("td", data["td"])
	d.Set("ttl", data["ttl"])
	d.Set("vlan", data["vlan"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func updateNsaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := d.Get("aclname").(string)

	nsacl := ns.Nsacl{
		Aclname: d.Get("aclname").(string),
	}
	hasChange := false
	if d.HasChange("aclaction") {
		log.Printf("[DEBUG]  netscaler-provider: Aclaction has changed for nsacl %s, starting update", nsaclName)
		nsacl.Aclaction = d.Get("aclaction").(string)
		hasChange = true
	}
	if d.HasChange("aclname") {
		log.Printf("[DEBUG]  netscaler-provider: Aclname has changed for nsacl %s, starting update", nsaclName)
		nsacl.Aclname = d.Get("aclname").(string)
		hasChange = true
	}
	if d.HasChange("destipop") {
		log.Printf("[DEBUG]  netscaler-provider: Destipop has changed for nsacl %s, starting update", nsaclName)
		nsacl.Destipop = d.Get("destipop").(string)
		nsacl.Destip = true
		hasChange = true
	}
	if d.HasChange("destipval") {
		log.Printf("[DEBUG]  netscaler-provider: Destipval has changed for nsacl %s, starting update", nsaclName)
		nsacl.Destipval = d.Get("destipval").(string)
		nsacl.Destip = true
		hasChange = true
	}
	if d.HasChange("destportop") {
		log.Printf("[DEBUG]  netscaler-provider: Destportop has changed for nsacl %s, starting update", nsaclName)
		nsacl.Destportop = d.Get("destportop").(string)
		nsacl.Destport = true
		hasChange = true
	}
	if d.HasChange("destportval") {
		log.Printf("[DEBUG]  netscaler-provider: Destportval has changed for nsacl %s, starting update", nsaclName)
		nsacl.Destportval = d.Get("destportval").(string)
		nsacl.Destport = true
		hasChange = true
	}
	if d.HasChange("dfdhash") {
		log.Printf("[DEBUG]  netscaler-provider: Dfdhash has changed for nsacl %s, starting update", nsaclName)
		nsacl.Dfdhash = d.Get("dfdhash").(string)
		hasChange = true
	}
	if d.HasChange("established") {
		log.Printf("[DEBUG]  netscaler-provider: Established has changed for nsacl %s, starting update", nsaclName)
		nsacl.Established = d.Get("established").(bool)
		hasChange = true
	}
	if d.HasChange("icmpcode") {
		log.Printf("[DEBUG]  netscaler-provider: Icmpcode has changed for nsacl %s, starting update", nsaclName)
		nsacl.Icmpcode = d.Get("icmpcode").(int)
		hasChange = true
	}
	if d.HasChange("icmptype") {
		log.Printf("[DEBUG]  netscaler-provider: Icmptype has changed for nsacl %s, starting update", nsaclName)
		nsacl.Icmptype = d.Get("icmptype").(int)
		hasChange = true
	}
	if d.HasChange("interface") {
		log.Printf("[DEBUG]  netscaler-provider: Interface has changed for nsacl %s, starting update", nsaclName)
		nsacl.Interface = d.Get("interface").(string)
		hasChange = true
	}
	if d.HasChange("logstate") {
		log.Printf("[DEBUG]  netscaler-provider: Logstate has changed for nsacl %s, starting update", nsaclName)
		nsacl.Logstate = d.Get("logstate").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  netscaler-provider: Priority has changed for nsacl %s, starting update", nsaclName)
		nsacl.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  netscaler-provider: Protocol has changed for nsacl %s, starting update", nsaclName)
		nsacl.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  netscaler-provider: Protocolnumber has changed for nsacl %s, starting update", nsaclName)
		nsacl.Protocolnumber = d.Get("protocolnumber").(int)
		hasChange = true
	}
	if d.HasChange("ratelimit") {
		log.Printf("[DEBUG]  netscaler-provider: Ratelimit has changed for nsacl %s, starting update", nsaclName)
		nsacl.Ratelimit = d.Get("ratelimit").(int)
		hasChange = true
	}
	if d.HasChange("srcipop") {
		log.Printf("[DEBUG]  netscaler-provider: Srcipop has changed for nsacl %s, starting update", nsaclName)
		nsacl.Srcipop = d.Get("srcipop").(string)
		nsacl.Srcip = true
		hasChange = true
	}
	if d.HasChange("srcipval") {
		log.Printf("[DEBUG]  netscaler-provider: Srcipval has changed for nsacl %s, starting update", nsaclName)
		nsacl.Srcipval = d.Get("srcipval").(string)
		nsacl.Srcip = true
		hasChange = true
	}
	if d.HasChange("srcmac") {
		log.Printf("[DEBUG]  netscaler-provider: Srcmac has changed for nsacl %s, starting update", nsaclName)
		nsacl.Srcmac = d.Get("srcmac").(string)
		hasChange = true
	}
	if d.HasChange("srcmacmask") {
		log.Printf("[DEBUG]  netscaler-provider: Srcmacmask has changed for nsacl %s, starting update", nsaclName)
		nsacl.Srcmacmask = d.Get("srcmacmask").(string)
		hasChange = true
	}
	if d.HasChange("srcportop") {
		log.Printf("[DEBUG]  netscaler-provider: Srcportop has changed for nsacl %s, starting update", nsaclName)
		nsacl.Srcportop = d.Get("srcportop").(string)
		nsacl.Srcport = true
		hasChange = true
	}
	if d.HasChange("srcportval") {
		log.Printf("[DEBUG]  netscaler-provider: Srcportval has changed for nsacl %s, starting update", nsaclName)
		nsacl.Srcportval = d.Get("srcportval").(string)
		nsacl.Srcport = true
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for nsacl %s, starting update", nsaclName)
		nsacl.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("stateful") {
		log.Printf("[DEBUG]  netscaler-provider: Stateful has changed for nsacl %s, starting update", nsaclName)
		nsacl.Stateful = d.Get("stateful").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  netscaler-provider: Td has changed for nsacl %s, starting update", nsaclName)
		nsacl.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  netscaler-provider: Ttl has changed for nsacl %s, starting update", nsaclName)
		nsacl.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  netscaler-provider: Vlan has changed for nsacl %s, starting update", nsaclName)
		nsacl.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  netscaler-provider: Vxlan has changed for nsacl %s, starting update", nsaclName)
		nsacl.Vxlan = d.Get("vxlan").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Nsacl.Type(), nsaclName, &nsacl)
		if err != nil {
			return fmt.Errorf("Error updating nsacl %s", nsaclName)
		}
	}
	return readNsaclFunc(d, meta)
}

func deleteNsaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := d.Id()
	err := client.DeleteResource(netscaler.Nsacl.Type(), nsaclName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
