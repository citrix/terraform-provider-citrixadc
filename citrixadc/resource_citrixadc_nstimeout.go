package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNstimeout() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNstimeoutFunc,
		ReadContext:   readNstimeoutFunc,
		UpdateContext: updateNstimeoutFunc,
		DeleteContext: deleteNstimeoutFunc, // Thought nstimeout resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"anyclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anyserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anytcpclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anytcpserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"client": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"halfclose": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newconnidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nontcpzombie": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reducedfintimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reducedrsttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zombie": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNstimeoutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstimeoutFunc")
	client := meta.(*NetScalerNitroClient).client
	nstimeoutName := resource.PrefixedUniqueId("tf-nstimeout-")

	nstimeout := ns.Nstimeout{}

	if raw := d.GetRawConfig().GetAttr("anyclient"); !raw.IsNull() {
		nstimeout.Anyclient = intPtr(d.Get("anyclient").(int))
	}
	if raw := d.GetRawConfig().GetAttr("anyserver"); !raw.IsNull() {
		nstimeout.Anyserver = intPtr(d.Get("anyserver").(int))
	}
	if raw := d.GetRawConfig().GetAttr("anytcpclient"); !raw.IsNull() {
		nstimeout.Anytcpclient = intPtr(d.Get("anytcpclient").(int))
	}
	if raw := d.GetRawConfig().GetAttr("anytcpserver"); !raw.IsNull() {
		nstimeout.Anytcpserver = intPtr(d.Get("anytcpserver").(int))
	}
	if raw := d.GetRawConfig().GetAttr("client"); !raw.IsNull() {
		nstimeout.Client = intPtr(d.Get("client").(int))
	}
	if raw := d.GetRawConfig().GetAttr("halfclose"); !raw.IsNull() {
		nstimeout.Halfclose = intPtr(d.Get("halfclose").(int))
	}
	if raw := d.GetRawConfig().GetAttr("httpclient"); !raw.IsNull() {
		nstimeout.Httpclient = intPtr(d.Get("httpclient").(int))
	}
	if raw := d.GetRawConfig().GetAttr("httpserver"); !raw.IsNull() {
		nstimeout.Httpserver = intPtr(d.Get("httpserver").(int))
	}
	if raw := d.GetRawConfig().GetAttr("newconnidletimeout"); !raw.IsNull() {
		nstimeout.Newconnidletimeout = intPtr(d.Get("newconnidletimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("nontcpzombie"); !raw.IsNull() {
		nstimeout.Nontcpzombie = intPtr(d.Get("nontcpzombie").(int))
	}
	if raw := d.GetRawConfig().GetAttr("reducedfintimeout"); !raw.IsNull() {
		nstimeout.Reducedfintimeout = intPtr(d.Get("reducedfintimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("reducedrsttimeout"); !raw.IsNull() {
		nstimeout.Reducedrsttimeout = intPtr(d.Get("reducedrsttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("server"); !raw.IsNull() {
		nstimeout.Server = intPtr(d.Get("server").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpclient"); !raw.IsNull() {
		nstimeout.Tcpclient = intPtr(d.Get("tcpclient").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpserver"); !raw.IsNull() {
		nstimeout.Tcpserver = intPtr(d.Get("tcpserver").(int))
	}
	if raw := d.GetRawConfig().GetAttr("zombie"); !raw.IsNull() {
		nstimeout.Zombie = intPtr(d.Get("zombie").(int))
	}

	err := client.UpdateUnnamedResource(service.Nstimeout.Type(), &nstimeout)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nstimeoutName)

	return readNstimeoutFunc(ctx, d, meta)
}

func readNstimeoutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstimeoutFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nstimeout state")
	data, err := client.FindResource(service.Nstimeout.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstimeout state")
		d.SetId("")
		return nil
	}
	setToInt("anyclient", d, data["anyclient"])
	setToInt("anyserver", d, data["anyserver"])
	setToInt("anytcpclient", d, data["anytcpclient"])
	setToInt("anytcpserver", d, data["anytcpserver"])
	setToInt("client", d, data["client"])
	setToInt("halfclose", d, data["halfclose"])
	setToInt("httpclient", d, data["httpclient"])
	setToInt("httpserver", d, data["httpserver"])
	setToInt("newconnidletimeout", d, data["newconnidletimeout"])
	setToInt("nontcpzombie", d, data["nontcpzombie"])
	setToInt("reducedfintimeout", d, data["reducedfintimeout"])
	setToInt("reducedrsttimeout", d, data["reducedrsttimeout"])
	setToInt("server", d, data["server"])
	setToInt("tcpclient", d, data["tcpclient"])
	setToInt("tcpserver", d, data["tcpserver"])
	setToInt("zombie", d, data["zombie"])

	return nil

}

func updateNstimeoutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNstimeoutFunc")
	client := meta.(*NetScalerNitroClient).client

	nstimeout := ns.Nstimeout{}

	hasChange := false
	if d.HasChange("anyclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Anyclient has changed for nstimeout, starting update")
		nstimeout.Anyclient = intPtr(d.Get("anyclient").(int))
		hasChange = true
	}
	if d.HasChange("anyserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Anyserver has changed for nstimeout, starting update")
		nstimeout.Anyserver = intPtr(d.Get("anyserver").(int))
		hasChange = true
	}
	if d.HasChange("anytcpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Anytcpclient has changed for nstimeout, starting update")
		nstimeout.Anytcpclient = intPtr(d.Get("anytcpclient").(int))
		hasChange = true
	}
	if d.HasChange("anytcpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Anytcpserver has changed for nstimeout, starting update")
		nstimeout.Anytcpserver = intPtr(d.Get("anytcpserver").(int))
		hasChange = true
	}
	if d.HasChange("client") {
		log.Printf("[DEBUG]  citrixadc-provider: Client has changed for nstimeout, starting update")
		nstimeout.Client = intPtr(d.Get("client").(int))
		hasChange = true
	}
	if d.HasChange("halfclose") {
		log.Printf("[DEBUG]  citrixadc-provider: Halfclose has changed for nstimeout, starting update")
		nstimeout.Halfclose = intPtr(d.Get("halfclose").(int))
		hasChange = true
	}
	if d.HasChange("httpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpclient has changed for nstimeout, starting update")
		nstimeout.Httpclient = intPtr(d.Get("httpclient").(int))
		hasChange = true
	}
	if d.HasChange("httpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpserver has changed for nstimeout, starting update")
		nstimeout.Httpserver = intPtr(d.Get("httpserver").(int))
		hasChange = true
	}
	if d.HasChange("newconnidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Newconnidletimeout has changed for nstimeout, starting update")
		nstimeout.Newconnidletimeout = intPtr(d.Get("newconnidletimeout").(int))
		hasChange = true
	}
	if d.HasChange("nontcpzombie") {
		log.Printf("[DEBUG]  citrixadc-provider: Nontcpzombie has changed for nstimeout, starting update")
		nstimeout.Nontcpzombie = intPtr(d.Get("nontcpzombie").(int))
		hasChange = true
	}
	if d.HasChange("reducedfintimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reducedfintimeout has changed for nstimeout, starting update")
		nstimeout.Reducedfintimeout = intPtr(d.Get("reducedfintimeout").(int))
		hasChange = true
	}
	if d.HasChange("reducedrsttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reducedrsttimeout has changed for nstimeout, starting update")
		nstimeout.Reducedrsttimeout = intPtr(d.Get("reducedrsttimeout").(int))
		hasChange = true
	}
	if d.HasChange("server") {
		log.Printf("[DEBUG]  citrixadc-provider: Server has changed for nstimeout, starting update")
		nstimeout.Server = intPtr(d.Get("server").(int))
		hasChange = true
	}
	if d.HasChange("tcpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpclient has changed for nstimeout, starting update")
		nstimeout.Tcpclient = intPtr(d.Get("tcpclient").(int))
		hasChange = true
	}
	if d.HasChange("tcpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpserver has changed for nstimeout, starting update")
		nstimeout.Tcpserver = intPtr(d.Get("tcpserver").(int))
		hasChange = true
	}
	if d.HasChange("zombie") {
		log.Printf("[DEBUG]  citrixadc-provider: Zombie has changed for nstimeout, starting update")
		nstimeout.Zombie = intPtr(d.Get("zombie").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nstimeout.Type(), &nstimeout)
		if err != nil {
			return diag.Errorf("Error updating nstimeout")
		}
	}
	return readNstimeoutFunc(ctx, d, meta)
}

func deleteNstimeoutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstimeoutFunc")
	// nstimeout do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
