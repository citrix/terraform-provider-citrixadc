package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSystemparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystemparameterFunc,
		ReadContext:   readSystemparameterFunc,
		UpdateContext: updateSystemparameterFunc,
		DeleteContext: deleteSystemparameterFunc,
		Schema: map[string]*schema.Schema{
			"basicauth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cliloglevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"doppler": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fipsusermode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forcepasswordchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"googleanalytics": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localauth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxclient": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minpasswordlen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"natpcbforceflushlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"natpcbrstontimeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"promptstring": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rbaonresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reauthonauthparamchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"removesensitivefiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrictedtimeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strongpassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"totalauthtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSystemparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	systemparameterName := resource.PrefixedUniqueId("tf-systemparameter-")

	systemparameter := system.Systemparameter{
		Basicauth:               d.Get("basicauth").(string),
		Cliloglevel:             d.Get("cliloglevel").(string),
		Doppler:                 d.Get("doppler").(string),
		Fipsusermode:            d.Get("fipsusermode").(string),
		Forcepasswordchange:     d.Get("forcepasswordchange").(string),
		Googleanalytics:         d.Get("googleanalytics").(string),
		Localauth:               d.Get("localauth").(string),
		Maxclient:               d.Get("maxclient").(string),
		Natpcbrstontimeout:      d.Get("natpcbrstontimeout").(string),
		Promptstring:            d.Get("promptstring").(string),
		Rbaonresponse:           d.Get("rbaonresponse").(string),
		Reauthonauthparamchange: d.Get("reauthonauthparamchange").(string),
		Removesensitivefiles:    d.Get("removesensitivefiles").(string),
		Restrictedtimeout:       d.Get("restrictedtimeout").(string),
		Strongpassword:          d.Get("strongpassword").(string),
	}

	if raw := d.GetRawConfig().GetAttr("minpasswordlen"); !raw.IsNull() {
		systemparameter.Minpasswordlen = intPtr(d.Get("minpasswordlen").(int))
	}
	if raw := d.GetRawConfig().GetAttr("natpcbforceflushlimit"); !raw.IsNull() {
		systemparameter.Natpcbforceflushlimit = intPtr(d.Get("natpcbforceflushlimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("timeout"); !raw.IsNull() {
		systemparameter.Timeout = intPtr(d.Get("timeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("totalauthtimeout"); !raw.IsNull() {
		systemparameter.Totalauthtimeout = intPtr(d.Get("totalauthtimeout").(int))
	}

	err := client.UpdateUnnamedResource(service.Systemparameter.Type(), &systemparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(systemparameterName)

	return readSystemparameterFunc(ctx, d, meta)
}

func readSystemparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading systemparameter state")
	data, err := client.FindResource(service.Systemparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systemparameter state")
		d.SetId("")
		return nil
	}
	d.Set("basicauth", data["basicauth"])
	d.Set("cliloglevel", data["cliloglevel"])
	d.Set("doppler", data["doppler"])
	d.Set("fipsusermode", data["fipsusermode"])
	d.Set("forcepasswordchange", data["forcepasswordchange"])
	d.Set("googleanalytics", data["googleanalytics"])
	d.Set("localauth", data["localauth"])
	setToInt("minpasswordlen", d, data["minpasswordlen"])
	d.Set("maxclient", data["maxclient"])
	setToInt("natpcbforceflushlimit", d, data["natpcbforceflushlimit"])
	d.Set("natpcbrstontimeout", data["natpcbrstontimeout"])
	d.Set("promptstring", data["promptstring"])
	d.Set("rbaonresponse", data["rbaonresponse"])
	d.Set("reauthonauthparamchange", data["reauthonauthparamchange"])
	d.Set("removesensitivefiles", data["removesensitivefiles"])
	d.Set("restrictedtimeout", data["restrictedtimeout"])
	d.Set("strongpassword", data["strongpassword"])
	setToInt("timeout", d, data["timeout"])
	setToInt("totalauthtimeout", d, data["totalauthtimeout"])

	return nil

}

func updateSystemparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	systemparameter := system.Systemparameter{}
	hasChange := false
	if d.HasChange("basicauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Basicauth has changed for systemparameter, starting update")
		systemparameter.Basicauth = d.Get("basicauth").(string)
		hasChange = true
	}
	if d.HasChange("cliloglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Cliloglevel has changed for systemparameter, starting update")
		systemparameter.Cliloglevel = d.Get("cliloglevel").(string)
		hasChange = true
	}
	if d.HasChange("doppler") {
		log.Printf("[DEBUG]  citrixadc-provider: Doppler has changed for systemparameter, starting update")
		systemparameter.Doppler = d.Get("doppler").(string)
		hasChange = true
	}
	if d.HasChange("fipsusermode") {
		log.Printf("[DEBUG]  citrixadc-provider: Fipsusermode has changed for systemparameter, starting update")
		systemparameter.Fipsusermode = d.Get("fipsusermode").(string)
		hasChange = true
	}
	if d.HasChange("forcepasswordchange") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcepasswordchange has changed for systemparameter, starting update")
		systemparameter.Forcepasswordchange = d.Get("forcepasswordchange").(string)
		hasChange = true
	}
	if d.HasChange("googleanalytics") {
		log.Printf("[DEBUG]  citrixadc-provider: Googleanalytics has changed for systemparameter, starting update")
		systemparameter.Googleanalytics = d.Get("googleanalytics").(string)
		hasChange = true
	}
	if d.HasChange("localauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Localauth has changed for systemparameter, starting update")
		systemparameter.Localauth = d.Get("localauth").(string)
		hasChange = true
	}
	if d.HasChange("minpasswordlen") {
		log.Printf("[DEBUG]  citrixadc-provider: Minpasswordlen has changed for systemparameter, starting update")
		systemparameter.Minpasswordlen = intPtr(d.Get("minpasswordlen").(int))
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Minpasswordlen has changed for systemparameter, starting update")
		systemparameter.Maxclient = d.Get("maxclient").(string)
		hasChange = true
	}
	if d.HasChange("natpcbforceflushlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Natpcbforceflushlimit has changed for systemparameter, starting update")
		systemparameter.Natpcbforceflushlimit = intPtr(d.Get("natpcbforceflushlimit").(int))
		hasChange = true
	}
	if d.HasChange("natpcbrstontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Natpcbrstontimeout has changed for systemparameter, starting update")
		systemparameter.Natpcbrstontimeout = d.Get("natpcbrstontimeout").(string)
		hasChange = true
	}
	if d.HasChange("promptstring") {
		log.Printf("[DEBUG]  citrixadc-provider: Promptstring has changed for systemparameter, starting update")
		systemparameter.Promptstring = d.Get("promptstring").(string)
		hasChange = true
	}
	if d.HasChange("rbaonresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Rbaonresponse has changed for systemparameter, starting update")
		systemparameter.Rbaonresponse = d.Get("rbaonresponse").(string)
		hasChange = true
	}
	if d.HasChange("reauthonauthparamchange") {
		log.Printf("[DEBUG]  citrixadc-provider: Reauthonauthparamchange has changed for systemparameter, starting update")
		systemparameter.Reauthonauthparamchange = d.Get("reauthonauthparamchange").(string)
		hasChange = true
	}
	if d.HasChange("removesensitivefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Removesensitivefiles has changed for systemparameter, starting update")
		systemparameter.Removesensitivefiles = d.Get("removesensitivefiles").(string)
		hasChange = true
	}
	if d.HasChange("restrictedtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Restrictedtimeout has changed for systemparameter, starting update")
		systemparameter.Restrictedtimeout = d.Get("restrictedtimeout").(string)
		hasChange = true
	}
	if d.HasChange("strongpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Strongpassword has changed for systemparameter, starting update")
		systemparameter.Strongpassword = d.Get("strongpassword").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for systemparameter, starting update")
		systemparameter.Timeout = intPtr(d.Get("timeout").(int))
		hasChange = true
	}
	if d.HasChange("totalauthtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Totalauthtimeout has changed for systemparameter, starting update")
		systemparameter.Totalauthtimeout = intPtr(d.Get("totalauthtimeout").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Systemparameter.Type(), &systemparameter)
		if err != nil {
			return diag.Errorf("Error updating systemparameters")
		}
	}
	return readSystemparameterFunc(ctx, d, meta)
}

func deleteSystemparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemparameterFunc")
	// systemparameter does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
