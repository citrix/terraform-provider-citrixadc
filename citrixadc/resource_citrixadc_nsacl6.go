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

func resourceCitrixAdcNsacl6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsacl6Func,
		ReadContext:   readNsacl6Func,
		UpdateContext: updateNsacl6Func,
		DeleteContext: deleteNsacl6Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"acl6name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"acl6action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"aclaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destipv6val": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": {
				Type:     schema.TypeBool,
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
			"dfdprefix": {
				Type:     schema.TypeInt,
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
			"srcipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcipv6val": {
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
			"srcport": {
				Type:     schema.TypeBool,
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
		},
	}
}

func createNsacl6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Get("acl6name").(string)
	nsacl6 := ns.Nsacl6{
		Acl6action:  d.Get("acl6action").(string),
		Acl6name:    d.Get("acl6name").(string),
		Aclaction:   d.Get("aclaction").(string),
		Destipop:    d.Get("destipop").(string),
		Destipv6:    d.Get("destipv6").(bool),
		Destipv6val: d.Get("destipv6val").(string),
		Destport:    d.Get("destport").(bool),
		Destportop:  d.Get("destportop").(string),
		Destportval: d.Get("destportval").(string),
		Dfdhash:     d.Get("dfdhash").(string),
		Established: d.Get("established").(bool),
		Interface:   d.Get("interface").(string),
		Logstate:    d.Get("logstate").(string),
		Protocol:    d.Get("protocol").(string),
		Srcipop:     d.Get("srcipop").(string),
		Srcipv6:     d.Get("srcipv6").(bool),
		Srcipv6val:  d.Get("srcipv6val").(string),
		Srcmac:      d.Get("srcmac").(string),
		Srcmacmask:  d.Get("srcmacmask").(string),
		Srcport:     d.Get("srcport").(bool),
		Srcportop:   d.Get("srcportop").(string),
		Srcportval:  d.Get("srcportval").(string),
		State:       d.Get("state").(string),
		Stateful:    d.Get("stateful").(string),
		Type:        d.Get("type").(string),
	}
	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		nsacl6.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("dfdprefix"); !raw.IsNull() {
		nsacl6.Dfdprefix = intPtr(d.Get("dfdprefix").(int))
	}
	if raw := d.GetRawConfig().GetAttr("icmpcode"); !raw.IsNull() {
		nsacl6.Icmpcode = intPtr(d.Get("icmpcode").(int))
	}
	if raw := d.GetRawConfig().GetAttr("icmptype"); !raw.IsNull() {
		nsacl6.Icmptype = intPtr(d.Get("icmptype").(int))
	}
	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		nsacl6.Priority = intPtr(d.Get("priority").(int))
	}
	if raw := d.GetRawConfig().GetAttr("protocolnumber"); !raw.IsNull() {
		nsacl6.Protocolnumber = intPtr(d.Get("protocolnumber").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ratelimit"); !raw.IsNull() {
		nsacl6.Ratelimit = intPtr(d.Get("ratelimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nsacl6.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		nsacl6.Ttl = intPtr(d.Get("ttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		nsacl6.Vlan = intPtr(d.Get("vlan").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vxlan"); !raw.IsNull() {
		nsacl6.Vxlan = intPtr(d.Get("vxlan").(int))
	}

	_, err := client.AddResource(service.Nsacl6.Type(), nsacl6Name, &nsacl6)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsacl6Name)

	return readNsacl6Func(ctx, d, meta)
}

func readNsacl6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsacl6 state %s", nsacl6Name)
	data, err := client.FindResource(service.Nsacl6.Type(), nsacl6Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsacl6 state %s", nsacl6Name)
		d.SetId("")
		return nil
	}
	d.Set("acl6action", data["acl6action"])
	setToInt("nodeid", d, data["nodeid"])
	d.Set("acl6name", data["acl6name"])
	d.Set("aclaction", data["aclaction"])
	d.Set("destipop", data["destipop"])
	d.Set("destipv6", data["destipv6"])
	d.Set("destipv6val", data["destipv6val"])
	d.Set("destport", data["destport"])
	d.Set("destportop", data["destportop"])
	d.Set("destportval", data["destportval"])
	d.Set("dfdhash", data["dfdhash"])
	setToInt("dfdprefix", d, data["dfdprefix"])
	d.Set("established", data["established"])
	setToInt("icmpcode", d, data["icmpcode"])
	setToInt("icmptype", d, data["icmptype"])
	d.Set("interface", data["interface"])
	d.Set("logstate", data["logstate"])
	setToInt("priority", d, data["priority"])
	d.Set("protocol", data["protocol"])
	setToInt("protocolnumber", d, data["protocolnumber"])
	setToInt("ratelimit", d, data["ratelimit"])
	d.Set("srcipop", data["srcipop"])
	d.Set("srcipv6", data["srcipv6"])
	d.Set("srcipv6val", data["srcipv6val"])
	d.Set("srcmac", data["srcmac"])
	d.Set("srcmacmask", data["srcmacmask"])
	d.Set("srcport", data["srcport"])
	d.Set("srcportop", data["srcportop"])
	d.Set("srcportval", data["srcportval"])
	d.Set("state", data["state"])
	d.Set("stateful", data["stateful"])
	setToInt("td", d, data["td"])
	setToInt("ttl", d, data["ttl"])
	d.Set("type", data["type"])
	setToInt("vlan", d, data["vlan"])
	setToInt("vxlan", d, data["vxlan"])

	return nil

}

func updateNsacl6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Get("acl6name").(string)

	nsacl6 := ns.Nsacl6{
		Acl6name: d.Get("acl6name").(string),
	}
	hasChange := false
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for nsacl6, starting update")
		nsacl6.Nodeid = intPtr(d.Get("nodeid").(int))
		hasChange = true
	}
	stateChange := false
	if d.HasChange("acl6action") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl6action has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Acl6action = d.Get("acl6action").(string)
		hasChange = true
	}
	if d.HasChange("aclaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Aclaction has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Aclaction = d.Get("aclaction").(string)
		hasChange = true
	}
	if d.HasChange("destipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destipop = d.Get("destipop").(string)
		hasChange = true
	}
	if d.HasChange("destipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipv6 has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destipv6 = d.Get("destipv6").(bool)
		hasChange = true
	}
	if d.HasChange("destipv6val") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipv6val has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destipv6val = d.Get("destipv6val").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG]  citrixadc-provider: Destport has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destport = d.Get("destport").(bool)
		hasChange = true
	}
	if d.HasChange("destportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destportop = d.Get("destportop").(string)
		hasChange = true
	}
	if d.HasChange("destportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportval has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destportval = d.Get("destportval").(string)
		hasChange = true
	}
	if d.HasChange("dfdhash") {
		log.Printf("[DEBUG]  citrixadc-provider: Dfdhash has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Dfdhash = d.Get("dfdhash").(string)
		hasChange = true
	}
	if d.HasChange("dfdprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Dfdprefix has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Dfdprefix = intPtr(d.Get("dfdprefix").(int))
		hasChange = true
	}
	if d.HasChange("established") {
		log.Printf("[DEBUG]  citrixadc-provider: Established has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Established = d.Get("established").(bool)
		hasChange = true
	}
	if d.HasChange("icmpcode") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpcode has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Icmpcode = intPtr(d.Get("icmpcode").(int))
		hasChange = true
	}
	if d.HasChange("icmptype") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmptype has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Icmptype = intPtr(d.Get("icmptype").(int))
		hasChange = true
	}
	if d.HasChange("interface") {
		log.Printf("[DEBUG]  citrixadc-provider: Interface has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Interface = d.Get("interface").(string)
		hasChange = true
	}
	if d.HasChange("logstate") {
		log.Printf("[DEBUG]  citrixadc-provider: Logstate has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Logstate = d.Get("logstate").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Priority = intPtr(d.Get("priority").(int))
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocolnumber has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Protocolnumber = intPtr(d.Get("protocolnumber").(int))
		hasChange = true
	}
	if d.HasChange("ratelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Ratelimit has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Ratelimit = intPtr(d.Get("ratelimit").(int))
		hasChange = true
	}
	if d.HasChange("srcipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcipop = d.Get("srcipop").(string)
		hasChange = true
	}
	if d.HasChange("srcipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipv6 has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcipv6 = d.Get("srcipv6").(bool)
		hasChange = true
	}
	if d.HasChange("srcipv6val") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipv6val has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcipv6val = d.Get("srcipv6val").(string)
		hasChange = true
	}
	if d.HasChange("srcmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmac has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcmac = d.Get("srcmac").(string)
		hasChange = true
	}
	if d.HasChange("srcmacmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmacmask has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcmacmask = d.Get("srcmacmask").(string)
		hasChange = true
	}
	if d.HasChange("srcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcport has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcport = d.Get("srcport").(bool)
		hasChange = true
	}
	if d.HasChange("srcportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcportop = d.Get("srcportop").(string)
		hasChange = true
	}
	if d.HasChange("srcportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportval has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcportval = d.Get("srcportval").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nsacl6 %s, starting update", nsacl6Name)
		stateChange = true
	}
	if d.HasChange("stateful") {
		log.Printf("[DEBUG]  citrixadc-provider: Stateful has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Stateful = d.Get("stateful").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Ttl = intPtr(d.Get("ttl").(int))
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlan has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Vxlan = intPtr(d.Get("vxlan").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsacl6.Type(), nsacl6Name, &nsacl6)
		if err != nil {
			return diag.Errorf("Error updating nsacl6 %s", nsacl6Name)
		}
	}

	if stateChange {
		err := doNsacl6StateSchange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling Nsacl6 %s", nsacl6Name)
		}
	}
	return readNsacl6Func(ctx, d, meta)
}

func deleteNsacl6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Id()
	err := client.DeleteResource(service.Nsacl6.Type(), nsacl6Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func doNsacl6StateSchange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doNsacl6StateSchange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	nsacl6 := ns.Nsacl6{
		Acl6name: d.Get("acl6name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Nsacl6.Type(), nsacl6, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Nsacl6.Type(), nsacl6, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
