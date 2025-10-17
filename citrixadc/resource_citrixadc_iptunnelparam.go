package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIptunnelparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIptunnelparamFunc,
		ReadContext:   readIptunnelparamFunc,
		UpdateContext: updateIptunnelparamFunc,
		DeleteContext: deleteIptunnelparamFunc,
		Schema: map[string]*schema.Schema{
			"dropfrag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropfragcputhreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enablestrictrx": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablestricttx": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srciproundrobin": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useclientsourceip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIptunnelparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIptunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var iptunnelparamName string
	// there is no primary key in iptunnelparam resource. Hence generate one for terraform state maintenance
	iptunnelparamName = resource.PrefixedUniqueId("tf-iptunnelparam-")
	iptunnelparam := network.Iptunnelparam{
		Dropfrag:          d.Get("dropfrag").(string),
		Enablestrictrx:    d.Get("enablestrictrx").(string),
		Enablestricttx:    d.Get("enablestricttx").(string),
		Mac:               d.Get("mac").(string),
		Srcip:             d.Get("srcip").(string),
		Srciproundrobin:   d.Get("srciproundrobin").(string),
		Useclientsourceip: d.Get("useclientsourceip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("dropfragcputhreshold"); !raw.IsNull() {
		iptunnelparam.Dropfragcputhreshold = intPtr(d.Get("dropfragcputhreshold").(int))
	}

	err := client.UpdateUnnamedResource(service.Iptunnelparam.Type(), &iptunnelparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(iptunnelparamName)

	return readIptunnelparamFunc(ctx, d, meta)
}

func readIptunnelparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIptunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading iptunnelparam state")
	data, err := client.FindResource(service.Iptunnelparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing iptunnelparam state")
		d.SetId("")
		return nil
	}
	d.Set("dropfrag", data["dropfrag"])
	setToInt("dropfragcputhreshold", d, data["dropfragcputhreshold"])
	d.Set("enablestrictrx", data["enablestrictrx"])
	d.Set("enablestricttx", data["enablestricttx"])
	d.Set("mac", data["mac"])
	d.Set("srcip", data["srcip"])
	d.Set("srciproundrobin", data["srciproundrobin"])
	d.Set("useclientsourceip", data["useclientsourceip"])

	return nil

}

func updateIptunnelparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIptunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client

	iptunnelparam := network.Iptunnelparam{}
	hasChange := false
	if d.HasChange("dropfrag") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropfrag has changed for iptunnelparam, starting update")
		iptunnelparam.Dropfrag = d.Get("dropfrag").(string)
		hasChange = true
	}
	if d.HasChange("dropfragcputhreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropfragcputhreshold has changed for iptunnelparam, starting update")
		iptunnelparam.Dropfragcputhreshold = intPtr(d.Get("dropfragcputhreshold").(int))
		hasChange = true
	}
	if d.HasChange("enablestrictrx") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablestrictrx has changed for iptunnelparam, starting update")
		iptunnelparam.Enablestrictrx = d.Get("enablestrictrx").(string)
		hasChange = true
	}
	if d.HasChange("enablestricttx") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablestricttx has changed for iptunnelparam, starting update")
		iptunnelparam.Enablestricttx = d.Get("enablestricttx").(string)
		hasChange = true
	}
	if d.HasChange("mac") {
		log.Printf("[DEBUG]  citrixadc-provider: Mac has changed for iptunnelparam, starting update")
		iptunnelparam.Mac = d.Get("mac").(string)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for iptunnelparam, starting update")
		iptunnelparam.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("srciproundrobin") {
		log.Printf("[DEBUG]  citrixadc-provider: Srciproundrobin has changed for iptunnelparam, starting update")
		iptunnelparam.Srciproundrobin = d.Get("srciproundrobin").(string)
		hasChange = true
	}
	if d.HasChange("useclientsourceip") {
		log.Printf("[DEBUG]  citrixadc-provider: Useclientsourceip has changed for iptunnelparam, starting update")
		iptunnelparam.Useclientsourceip = d.Get("useclientsourceip").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Iptunnelparam.Type(), &iptunnelparam)
		if err != nil {
			return diag.Errorf("Error updating iptunnelparam")
		}
	}
	return readIptunnelparamFunc(ctx, d, meta)
}

func deleteIptunnelparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIptunnelparamFunc")

	d.SetId("")

	return nil
}
