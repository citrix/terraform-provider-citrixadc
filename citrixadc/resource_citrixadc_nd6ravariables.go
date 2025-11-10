package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNd6ravariables() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNd6ravariablesFunc,
		ReadContext:   readNd6ravariablesFunc,
		UpdateContext: updateNd6ravariablesFunc,
		DeleteContext: deleteNd6ravariablesFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vlan": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ceaserouteradv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"currhoplimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultlifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"linkmtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"managedaddrconfig": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxrtadvinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minrtadvinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"onlyunicastrtadvresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"otheraddrconfig": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reachabletime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retranstime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sendrouteradv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srclinklayeraddroption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNd6ravariablesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNd6ravariablesFunc")
	client := meta.(*NetScalerNitroClient).client
	nd6ravariablesName := strconv.Itoa(d.Get("vlan").(int))
	nd6ravariables := network.Nd6ravariables{
		Ceaserouteradv:           d.Get("ceaserouteradv").(string),
		Managedaddrconfig:        d.Get("managedaddrconfig").(string),
		Onlyunicastrtadvresponse: d.Get("onlyunicastrtadvresponse").(string),
		Otheraddrconfig:          d.Get("otheraddrconfig").(string),
		Sendrouteradv:            d.Get("sendrouteradv").(string),
		Srclinklayeraddroption:   d.Get("srclinklayeraddroption").(string),
	}

	if raw := d.GetRawConfig().GetAttr("currhoplimit"); !raw.IsNull() {
		nd6ravariables.Currhoplimit = intPtr(d.Get("currhoplimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("defaultlifetime"); !raw.IsNull() {
		nd6ravariables.Defaultlifetime = intPtr(d.Get("defaultlifetime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("linkmtu"); !raw.IsNull() {
		nd6ravariables.Linkmtu = intPtr(d.Get("linkmtu").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxrtadvinterval"); !raw.IsNull() {
		nd6ravariables.Maxrtadvinterval = intPtr(d.Get("maxrtadvinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minrtadvinterval"); !raw.IsNull() {
		nd6ravariables.Minrtadvinterval = intPtr(d.Get("minrtadvinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("reachabletime"); !raw.IsNull() {
		nd6ravariables.Reachabletime = intPtr(d.Get("reachabletime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("retranstime"); !raw.IsNull() {
		nd6ravariables.Retranstime = intPtr(d.Get("retranstime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		nd6ravariables.Vlan = intPtr(d.Get("vlan").(int))
	}

	err := client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nd6ravariablesName)

	return readNd6ravariablesFunc(ctx, d, meta)
}

func readNd6ravariablesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNd6ravariablesFunc")
	client := meta.(*NetScalerNitroClient).client
	nd6ravariablesName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nd6ravariables state %s", nd6ravariablesName)
	data, err := client.FindResource(service.Nd6ravariables.Type(), nd6ravariablesName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nd6ravariables state %s", nd6ravariablesName)
		d.SetId("")
		return nil
	}
	vlan_int, _ := strconv.Atoi(data["vlan"].(string))
	d.Set("vlan", vlan_int)
	d.Set("ceaserouteradv", data["ceaserouteradv"])
	setToInt("currhoplimit", d, data["currhoplimit"])
	setToInt("defaultlifetime", d, data["defaultlifetime"])
	setToInt("linkmtu", d, data["linkmtu"])
	d.Set("managedaddrconfig", data["managedaddrconfig"])
	setToInt("maxrtadvinterval", d, data["maxrtadvinterval"])
	setToInt("minrtadvinterval", d, data["minrtadvinterval"])
	d.Set("onlyunicastrtadvresponse", data["onlyunicastrtadvresponse"])
	d.Set("otheraddrconfig", data["otheraddrconfig"])
	setToInt("reachabletime", d, data["reachabletime"])
	setToInt("retranstime", d, data["retranstime"])
	d.Set("sendrouteradv", data["sendrouteradv"])
	d.Set("srclinklayeraddroption", data["srclinklayeraddroption"])
	setToInt("vlan", d, data["vlan"])

	return nil

}

func updateNd6ravariablesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNd6ravariablesFunc")
	client := meta.(*NetScalerNitroClient).client
	nd6ravariablesName := strconv.Itoa(d.Get("vlan").(int))

	nd6ravariables := network.Nd6ravariables{}

	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		nd6ravariables.Vlan = intPtr(d.Get("vlan").(int))
	}
	hasChange := false
	if d.HasChange("ceaserouteradv") {
		log.Printf("[DEBUG]  citrixadc-provider: Ceaserouteradv has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Ceaserouteradv = d.Get("ceaserouteradv").(string)
		hasChange = true
	}
	if d.HasChange("currhoplimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Currhoplimit has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Currhoplimit = intPtr(d.Get("currhoplimit").(int))
		hasChange = true
	}
	if d.HasChange("defaultlifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultlifetime has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Defaultlifetime = intPtr(d.Get("defaultlifetime").(int))
		hasChange = true
	}
	if d.HasChange("linkmtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkmtu has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Linkmtu = intPtr(d.Get("linkmtu").(int))
		hasChange = true
	}
	if d.HasChange("managedaddrconfig") {
		log.Printf("[DEBUG]  citrixadc-provider: Managedaddrconfig has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Managedaddrconfig = d.Get("managedaddrconfig").(string)
		hasChange = true
	}
	if d.HasChange("maxrtadvinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxrtadvinterval has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Maxrtadvinterval = intPtr(d.Get("maxrtadvinterval").(int))
		hasChange = true
	}
	if d.HasChange("minrtadvinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Minrtadvinterval has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Minrtadvinterval = intPtr(d.Get("minrtadvinterval").(int))
		hasChange = true
	}
	if d.HasChange("onlyunicastrtadvresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Onlyunicastrtadvresponse has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Onlyunicastrtadvresponse = d.Get("onlyunicastrtadvresponse").(string)
		hasChange = true
	}
	if d.HasChange("otheraddrconfig") {
		log.Printf("[DEBUG]  citrixadc-provider: Otheraddrconfig has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Otheraddrconfig = d.Get("otheraddrconfig").(string)
		hasChange = true
	}
	if d.HasChange("reachabletime") {
		log.Printf("[DEBUG]  citrixadc-provider: Reachabletime has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Reachabletime = intPtr(d.Get("reachabletime").(int))
		hasChange = true
	}
	if d.HasChange("retranstime") {
		log.Printf("[DEBUG]  citrixadc-provider: Retranstime has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Retranstime = intPtr(d.Get("retranstime").(int))
		hasChange = true
	}
	if d.HasChange("sendrouteradv") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendrouteradv has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Sendrouteradv = d.Get("sendrouteradv").(string)
		hasChange = true
	}
	if d.HasChange("srclinklayeraddroption") {
		log.Printf("[DEBUG]  citrixadc-provider: Srclinklayeraddroption has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Srclinklayeraddroption = d.Get("srclinklayeraddroption").(string)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Vlan = intPtr(d.Get("vlan").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
		if err != nil {
			return diag.Errorf("Error updating nd6ravariables %s", nd6ravariablesName)
		}
	}
	return readNd6ravariablesFunc(ctx, d, meta)
}

func deleteNd6ravariablesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNd6ravariablesFunc")
	// nd6ravariables does not support DELETE operation
	d.SetId("")

	return nil
}
