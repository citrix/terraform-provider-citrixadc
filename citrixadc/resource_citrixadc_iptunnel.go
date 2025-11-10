package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcIptunnel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIptunnelFunc,
		ReadContext:   readIptunnelFunc,
		UpdateContext: updateIptunnelFunc,
		DeleteContext: deleteIptunnelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vnid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlantagging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tosinherit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"grepayload": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ipsecprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"local": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"remote": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"remotesubnetmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createIptunnelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Get("name").(string)

	iptunnel := network.Iptunnel{
		Grepayload:       d.Get("grepayload").(string),
		Ipsecprofilename: d.Get("ipsecprofilename").(string),
		Local:            d.Get("local").(string),
		Name:             d.Get("name").(string),
		Ownergroup:       d.Get("ownergroup").(string),
		Protocol:         d.Get("protocol").(string),
		Remote:           d.Get("remote").(string),
		Remotesubnetmask: d.Get("remotesubnetmask").(string),
		Tosinherit:       d.Get("tosinherit").(string),
		Vlantagging:      d.Get("vlantagging").(string),
	}
	if raw := d.GetRawConfig().GetAttr("vnid"); !raw.IsNull() {
		iptunnel.Vnid = intPtr(d.Get("vnid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("destport"); !raw.IsNull() {
		iptunnel.Destport = intPtr(d.Get("destport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		iptunnel.Vlan = intPtr(d.Get("vlan").(int))
	}

	_, err := client.AddResource(service.Iptunnel.Type(), iptunnelName, &iptunnel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(iptunnelName)

	return readIptunnelFunc(ctx, d, meta)
}

func readIptunnelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading iptunnel state %s", iptunnelName)
	data, err := client.FindResource(service.Iptunnel.Type(), iptunnelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing iptunnel state %s", iptunnelName)
		d.SetId("")
		return nil
	}
	d.Set("grepayload", data["grepayload"])
	setToInt("vnid", d, data["vnid"])
	d.Set("vlantagging", data["vlantagging"])
	d.Set("tosinherit", data["tosinherit"])
	setToInt("destport", d, data["destport"])
	d.Set("ipsecprofilename", data["ipsecprofilename"])
	d.Set("local", data["local"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("protocol", data["protocol"])
	d.Set("remote", data["remote"])
	d.Set("remotesubnetmask", data["remotesubnetmask"])
	setToInt("vlan", d, data["vlan"])

	return nil

}

func updateIptunnelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Get("name").(string)

	iptunnel := network.Iptunnel{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("tosinherit") {
		log.Printf("[DEBUG]  citrixadc-provider: Tosinherit has changed for iptunnel, starting update")
		iptunnel.Tosinherit = d.Get("tosinherit").(string)
		hasChange = true
	}
	if d.HasChange("vlantagging") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlantagging has changed for iptunnel, starting update")
		iptunnel.Vlantagging = d.Get("vlantagging").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG]  citrixadc-provider: Destport has changed for iptunnel, starting update")
		iptunnel.Destport = intPtr(d.Get("destport").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Iptunnel.Type(), iptunnelName, &iptunnel)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return readIptunnelFunc(ctx, d, meta)
}

func deleteIptunnelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Id()
	err := client.DeleteResource(service.Iptunnel.Type(), iptunnelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
