package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsacl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsaclFunc,
		ReadContext:   readNsaclFunc,
		UpdateContext: updateNsaclFunc,
		DeleteContext: deleteNsaclFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"aclaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aclname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfdhash": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"established": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"icmpcode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmptype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocolnumber": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ratelimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmacmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stateful": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcport": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destport": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if d.Get("destportop") != nil && d.Get("destipval") == nil {
		return diag.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if d.Get("srcipop") != nil && d.Get("srcipval") == nil {
		return diag.Errorf("Error in nsacl spec %s cannot have srcipop without srcipval", nsaclName)
	}
	if d.Get("srcportop") != nil && d.Get("srcportval") == nil {
		return diag.Errorf("Error in nsacl spec %s cannot have srcportop without srcportval", nsaclName)
	}

	nsacl := ns.Nsacl{
		Aclaction:   d.Get("aclaction").(string),
		Aclname:     d.Get("aclname").(string),
		Destip:      destip,
		Destipop:    d.Get("destipop").(string),
		Destipval:   d.Get("destipval").(string),
		Destport:    destport,
		Destportop:  d.Get("destportop").(string),
		Destportval: d.Get("destportval").(string),
		Dfdhash:     d.Get("dfdhash").(string),
		Established: d.Get("established").(bool),
		Interface:   d.Get("interface").(string),
		Logstate:    d.Get("logstate").(string),
		Protocol:    d.Get("protocol").(string),
		Srcip:       srcip,
		Srcipop:     d.Get("srcipop").(string),
		Srcipval:    d.Get("srcipval").(string),
		Srcmac:      d.Get("srcmac").(string),
		Srcmacmask:  d.Get("srcmacmask").(string),
		Srcport:     srcport,
		Srcportop:   d.Get("srcportop").(string),
		Srcportval:  d.Get("srcportval").(string),
		State:       d.Get("state").(string),
		Stateful:    d.Get("stateful").(string),
	}

	if raw := d.GetRawConfig().GetAttr("icmpcode"); !raw.IsNull() {
		nsacl.Icmpcode = intPtr(d.Get("icmpcode").(int))
	}
	if raw := d.GetRawConfig().GetAttr("icmptype"); !raw.IsNull() {
		nsacl.Icmptype = intPtr(d.Get("icmptype").(int))
	}
	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		nsacl.Priority = intPtr(d.Get("priority").(int))
	}
	if raw := d.GetRawConfig().GetAttr("protocolnumber"); !raw.IsNull() {
		nsacl.Protocolnumber = intPtr(d.Get("protocolnumber").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ratelimit"); !raw.IsNull() {
		nsacl.Ratelimit = intPtr(d.Get("ratelimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nsacl.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		nsacl.Ttl = intPtr(d.Get("ttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		nsacl.Vlan = intPtr(d.Get("vlan").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vxlan"); !raw.IsNull() {
		nsacl.Vxlan = intPtr(d.Get("vxlan").(int))
	}

	_, err := client.AddResource(service.Nsacl.Type(), nsaclName, &nsacl)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsaclName)

	return readNsaclFunc(ctx, d, meta)
}

func readNsaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading nsacl state %s", nsaclName)
	data, err := client.FindResource(service.Nsacl.Type(), nsaclName)
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
	setToInt("icmpcode", d, data["icmpcode"])
	setToInt("icmptype", d, data["icmptype"])
	d.Set("interface", data["interface"])
	d.Set("logstate", data["logstate"])
	setToInt("priority", d, data["priority"])
	d.Set("protocol", data["protocol"])
	setToInt("protocolnumber", d, data["protocolnumber"])
	setToInt("ratelimit", d, data["ratelimit"])
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
	setToInt("td", d, data["td"])
	setToInt("ttl", d, data["ttl"])
	setToInt("vlan", d, data["vlan"])
	setToInt("vxlan", d, data["vxlan"])

	return nil

}

func updateNsaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In updateNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := d.Get("aclname").(string)

	nsacl := ns.Nsacl{
		Aclname: d.Get("aclname").(string),
	}
	stateChange := false
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
		nsacl.Icmpcode = intPtr(d.Get("icmpcode").(int))
		hasChange = true
	}
	if d.HasChange("icmptype") {
		log.Printf("[DEBUG]  netscaler-provider: Icmptype has changed for nsacl %s, starting update", nsaclName)
		nsacl.Icmptype = intPtr(d.Get("icmptype").(int))
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
		nsacl.Priority = intPtr(d.Get("priority").(int))
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  netscaler-provider: Protocol has changed for nsacl %s, starting update", nsaclName)
		nsacl.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  netscaler-provider: Protocolnumber has changed for nsacl %s, starting update", nsaclName)
		nsacl.Protocolnumber = intPtr(d.Get("protocolnumber").(int))
		hasChange = true
	}
	if d.HasChange("ratelimit") {
		log.Printf("[DEBUG]  netscaler-provider: Ratelimit has changed for nsacl %s, starting update", nsaclName)
		nsacl.Ratelimit = intPtr(d.Get("ratelimit").(int))
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
		stateChange = true
	}
	if d.HasChange("stateful") {
		log.Printf("[DEBUG]  netscaler-provider: Stateful has changed for nsacl %s, starting update", nsaclName)
		nsacl.Stateful = d.Get("stateful").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  netscaler-provider: Td has changed for nsacl %s, starting update", nsaclName)
		nsacl.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  netscaler-provider: Ttl has changed for nsacl %s, starting update", nsaclName)
		nsacl.Ttl = intPtr(d.Get("ttl").(int))
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  netscaler-provider: Vlan has changed for nsacl %s, starting update", nsaclName)
		nsacl.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  netscaler-provider: Vxlan has changed for nsacl %s, starting update", nsaclName)
		nsacl.Vxlan = intPtr(d.Get("vxlan").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsacl.Type(), nsaclName, &nsacl)
		if err != nil {
			return diag.Errorf("Error updating nsacl %s", nsaclName)
		}
	}

	if stateChange {
		err := doNsaclStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling nsacl %s", nsaclName)
		}
	}
	return readNsaclFunc(ctx, d, meta)
}

func deleteNsaclFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteNsaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nsaclName := d.Id()
	err := client.DeleteResource(service.Nsacl.Type(), nsaclName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func doNsaclStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doServerStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	nsacl := ns.Nsacl{
		Aclname: d.Get("aclname").(string),
	}

	newstate := d.Get("state")

	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Nsacl.Type(), nsacl, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		err := client.ActOnResource(service.Nsacl.Type(), nsacl, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
