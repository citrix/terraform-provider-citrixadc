package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"log"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNshttpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNshttpparamFunc,
		ReadContext:   readNshttpparamFunc,
		UpdateContext: updateNshttpparamFunc,
		DeleteContext: deleteNshttpparamFunc,
		Schema: map[string]*schema.Schema{
			"conmultiplex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropinvalreqs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2serverside": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignoreconnectcodingscheme": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insnssrvrhdr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logerrresp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"markconnreqinval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"markhttp09inval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxreusepool": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nssrvrhdr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNshttpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNshttpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var nshttpparamName string
	// there is no primary key in nshttpparam resource. Hence generate one for terraform state maintenance
	nshttpparamName = resource.PrefixedUniqueId("tf-nshttpparam-")
	nshttpparam := ns.Nshttpparam{
		Conmultiplex:              d.Get("conmultiplex").(string),
		Dropinvalreqs:             d.Get("dropinvalreqs").(string),
		Http2serverside:           d.Get("http2serverside").(string),
		Ignoreconnectcodingscheme: d.Get("ignoreconnectcodingscheme").(string),
		Insnssrvrhdr:              d.Get("insnssrvrhdr").(string),
		Logerrresp:                d.Get("logerrresp").(string),
		Markconnreqinval:          d.Get("markconnreqinval").(string),
		Markhttp09inval:           d.Get("markhttp09inval").(string),
		Nssrvrhdr:                 d.Get("nssrvrhdr").(string),
	}

	if raw := d.GetRawConfig().GetAttr("maxreusepool"); !raw.IsNull() {
		nshttpparam.Maxreusepool = intPtr(d.Get("maxreusepool").(int))
	}

	err := client.UpdateUnnamedResource(service.Nshttpparam.Type(), &nshttpparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nshttpparamName)

	return readNshttpparamFunc(ctx, d, meta)
}

func readNshttpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNshttpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nshttpparam state")
	data, err := client.FindResource(service.Nshttpparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nshttpparam state")
		d.SetId("")
		return nil
	}
	d.Set("conmultiplex", data["conmultiplex"])
	d.Set("dropinvalreqs", data["dropinvalreqs"])
	d.Set("http2serverside", data["http2serverside"])
	d.Set("ignoreconnectcodingscheme", data["ignoreconnectcodingscheme"])
	d.Set("insnssrvrhdr", data["insnssrvrhdr"])
	d.Set("logerrresp", data["logerrresp"])
	d.Set("markconnreqinval", data["markconnreqinval"])
	d.Set("markhttp09inval", data["markhttp09inval"])
	val, _ := strconv.Atoi(data["maxreusepool"].(string))
	d.Set("maxreusepool", val)
	d.Set("nssrvrhdr", data["nssrvrhdr"])

	return nil

}

func updateNshttpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNshttpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	nshttpparam := ns.Nshttpparam{}

	if raw := d.GetRawConfig().GetAttr("maxreusepool"); !raw.IsNull() {
		nshttpparam.Maxreusepool = intPtr(d.Get("maxreusepool").(int))
	}
	hasChange := false
	if d.HasChange("conmultiplex") {
		log.Printf("[DEBUG]  citrixadc-provider: Conmultiplex has changed for nshttpparam ,starting update")
		nshttpparam.Conmultiplex = d.Get("conmultiplex").(string)
		hasChange = true
	}
	if d.HasChange("dropinvalreqs") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropinvalreqs has changed for nshttpparam ,starting update")
		nshttpparam.Dropinvalreqs = d.Get("dropinvalreqs").(string)
		hasChange = true
	}
	if d.HasChange("http2serverside") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2serverside has changed for nshttpparam ,starting update")
		nshttpparam.Http2serverside = d.Get("http2serverside").(string)
		hasChange = true
	}
	if d.HasChange("ignoreconnectcodingscheme") {
		log.Printf("[DEBUG]  citrixadc-provider: Ignoreconnectcodingscheme has changed for nshttpparam ,starting update")
		nshttpparam.Ignoreconnectcodingscheme = d.Get("ignoreconnectcodingscheme").(string)
		hasChange = true
	}
	if d.HasChange("insnssrvrhdr") {
		log.Printf("[DEBUG]  citrixadc-provider: Insnssrvrhdr has changed for nshttpparam ,starting update")
		nshttpparam.Insnssrvrhdr = d.Get("insnssrvrhdr").(string)
		hasChange = true
	}
	if d.HasChange("logerrresp") {
		log.Printf("[DEBUG]  citrixadc-provider: Logerrresp has changed for nshttpparam ,starting update")
		nshttpparam.Logerrresp = d.Get("logerrresp").(string)
		hasChange = true
	}
	if d.HasChange("markconnreqinval") {
		log.Printf("[DEBUG]  citrixadc-provider: Markconnreqinval has changed for nshttpparam ,starting update")
		nshttpparam.Markconnreqinval = d.Get("markconnreqinval").(string)
		hasChange = true
	}
	if d.HasChange("markhttp09inval") {
		log.Printf("[DEBUG]  citrixadc-provider: Markhttp09inval has changed for nshttpparam ,starting update")
		nshttpparam.Markhttp09inval = d.Get("markhttp09inval").(string)
		hasChange = true
	}
	if d.HasChange("nssrvrhdr") {
		log.Printf("[DEBUG]  citrixadc-provider: Nssrvrhdr has changed for nshttpparam ,starting update")
		nshttpparam.Nssrvrhdr = d.Get("nssrvrhdr").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nshttpparam.Type(), &nshttpparam)
		if err != nil {
			return diag.Errorf("Error updating nshttpparam")
		}
	}
	return readNshttpparamFunc(ctx, d, meta)
}

func deleteNshttpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNshttpparamFunc")

	d.SetId("")

	return nil
}
