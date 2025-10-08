package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcGslbparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createGslbparameterFunc,
		ReadContext:   readGslbparameterFunc,
		UpdateContext: updateGslbparameterFunc,
		DeleteContext: deleteGslbparameterFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"automaticconfigsync": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropldnsreq": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbconfigsyncmonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbsvcstatedelaytime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gslbsyncinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gslbsynclocfiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbsyncmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldnsentrytimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldnsmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldnsprobeorder": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mepkeepalivetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtttolerance": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"svcstatelearningtime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6ldnsmasklen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}
func createGslbparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbparameterName := resource.PrefixedUniqueId("tf-gslbparameter-")

	gslbparameter := gslb.Gslbparameter{
		Automaticconfigsync:   d.Get("automaticconfigsync").(string),
		Dropldnsreq:           d.Get("dropldnsreq").(string),
		Gslbconfigsyncmonitor: d.Get("gslbconfigsyncmonitor").(string),
		Gslbsvcstatedelaytime: d.Get("gslbsvcstatedelaytime").(int),
		Gslbsyncinterval:      d.Get("gslbsyncinterval").(int),
		Gslbsynclocfiles:      d.Get("gslbsynclocfiles").(string),
		Gslbsyncmode:          d.Get("gslbsyncmode").(string),
		Ldnsentrytimeout:      d.Get("ldnsentrytimeout").(int),
		Ldnsmask:              d.Get("ldnsmask").(string),
		Ldnsprobeorder:        toStringList(d.Get("ldnsprobeorder").([]interface{})),
		Mepkeepalivetimeout:   d.Get("mepkeepalivetimeout").(int),
		Rtttolerance:          d.Get("rtttolerance").(int),
		Svcstatelearningtime:  d.Get("svcstatelearningtime").(int),
		V6ldnsmasklen:         d.Get("v6ldnsmasklen").(int),
	}

	err := client.UpdateUnnamedResource(service.Gslbparameter.Type(), &gslbparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(gslbparameterName)
	return readGslbparameterFunc(ctx, d, meta)
}

func readGslbparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading gslbparameter state ")
	data, err := client.FindResource(service.Gslbparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing gslbparameter state ")
		d.SetId("")
		return nil
	}
	d.Set("automaticconfigsync", data["automaticconfigsync"])
	d.Set("dropldnsreq", data["dropldnsreq"])
	d.Set("gslbconfigsyncmonitor", data["gslbconfigsyncmonitor"])
	setToInt("gslbsvcstatedelaytime", d, data["gslbsvcstatedelaytime"])
	setToInt("gslbsyncinterval", d, data["gslbsyncinterval"])
	d.Set("gslbsynclocfiles", data["gslbsynclocfiles"])
	d.Set("gslbsyncmode", data["gslbsyncmode"])
	setToInt("ldnsentrytimeout", d, data["ldnsentrytimeout"])
	d.Set("ldnsmask", data["ldnsmask"])
	d.Set("ldnsprobeorder", data["ldnsprobeorder"])
	setToInt("mepkeepalivetimeout", d, data["mepkeepalivetimeout"])
	setToInt("rtttolerance", d, data["rtttolerance"])
	setToInt("svcstatelearningtime", d, data["svcstatelearningtime"])
	setToInt("v6ldnsmasklen", d, data["v6ldnsmasklen"])

	return nil

}

func updateGslbparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateGslbparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	gslbparameter := gslb.Gslbparameter{}
	hasChange := false
	if d.HasChange("automaticconfigsync") {
		log.Printf("[DEBUG]  citrixadc-provider: Automaticconfigsync has changed for gslbparameter, starting update")
		gslbparameter.Automaticconfigsync = d.Get("automaticconfigsync").(string)
		hasChange = true
	}
	if d.HasChange("dropldnsreq") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropldnsreq has changed for gslbparameter, starting update")
		gslbparameter.Dropldnsreq = d.Get("dropldnsreq").(string)
		hasChange = true
	}
	if d.HasChange("gslbconfigsyncmonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Gslbconfigsyncmonitor has changed for gslbparameter, starting update")
		gslbparameter.Gslbconfigsyncmonitor = d.Get("gslbconfigsyncmonitor").(string)
		hasChange = true
	}
	if d.HasChange("gslbsvcstatedelaytime") {
		log.Printf("[DEBUG]  citrixadc-provider: Gslbsvcstatedelaytime has changed for gslbparameter, starting update")
		gslbparameter.Gslbsvcstatedelaytime = d.Get("gslbsvcstatedelaytime").(int)
		hasChange = true
	}
	if d.HasChange("gslbsyncinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Gslbsyncinterval has changed for gslbparameter, starting update")
		gslbparameter.Gslbsyncinterval = d.Get("gslbsyncinterval").(int)
		hasChange = true
	}
	if d.HasChange("gslbsynclocfiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Gslbsynclocfiles has changed for gslbparameter, starting update")
		gslbparameter.Gslbsynclocfiles = d.Get("gslbsynclocfiles").(string)
		hasChange = true
	}
	if d.HasChange("gslbsyncmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Gslbsyncmode has changed for gslbparameter, starting update")
		gslbparameter.Gslbsyncmode = d.Get("gslbsyncmode").(string)
		hasChange = true
	}
	if d.HasChange("ldnsentrytimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldnsentrytimeout has changed for gslbparameter, starting update")
		gslbparameter.Ldnsentrytimeout = d.Get("ldnsentrytimeout").(int)
		hasChange = true
	}
	if d.HasChange("ldnsmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldnsmask has changed for gslbparameter, starting update")
		gslbparameter.Ldnsmask = d.Get("ldnsmask").(string)
		hasChange = true
	}
	if d.HasChange("ldnsprobeorder") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldnsprobeorder has changed for gslbparameter, starting update")
		gslbparameter.Ldnsprobeorder = toStringList(d.Get("ldnsprobeorder").([]interface{}))
		hasChange = true
	}
	if d.HasChange("mepkeepalivetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Mepkeepalivetimeout has changed for gslbparameter, starting update")
		gslbparameter.Mepkeepalivetimeout = d.Get("mepkeepalivetimeout").(int)
		hasChange = true
	}
	if d.HasChange("rtttolerance") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtttolerance has changed for gslbparameter, starting update")
		gslbparameter.Rtttolerance = d.Get("rtttolerance").(int)
		hasChange = true
	}
	if d.HasChange("svcstatelearningtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Svcstatelearningtime has changed for gslbparameter, starting update")
		gslbparameter.Svcstatelearningtime = d.Get("svcstatelearningtime").(int)
		hasChange = true
	}
	if d.HasChange("v6ldnsmasklen") {
		log.Printf("[DEBUG]  citrixadc-provider: V6ldnsmasklen has changed for gslbparameter, starting update")
		gslbparameter.V6ldnsmasklen = d.Get("v6ldnsmasklen").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Gslbparameter.Type(), &gslbparameter)
		if err != nil {
			return diag.Errorf("Error updating gslbparameter: %s", err.Error())
		}
	}
	return readGslbparameterFunc(ctx, d, meta)
}

func deleteGslbparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbparameterFunc")
	// gslbparameter does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
