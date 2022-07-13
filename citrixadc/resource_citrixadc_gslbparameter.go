package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcGslbparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbparameterFunc,
		Read:          readGslbparameterFunc,
		Update:        updateGslbparameterFunc,
		Delete:        deleteGslbparameterFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"automaticconfigsync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropldnsreq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbconfigsyncmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbsvcstatedelaytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gslbsyncinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gslbsynclocfiles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gslbsyncmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldnsentrytimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldnsmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldnsprobeorder": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mepkeepalivetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtttolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"svcstatelearningtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6ldnsmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}
//TODO: I changed create to update basically #124
func createGslbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbparameterName string  //Should I delete this?
	
	gslbparameterName = resource.PrefixedUniqueId("tf-gslbparameter-")

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
		Ldnsprobeorder:        d.Get("ldnsprobeorder").([]interface{}),
		Mepkeepalivetimeout:   d.Get("mepkeepalivetimeout").(int),
		Rtttolerance:          d.Get("rtttolerance").(int),
		Svcstatelearningtime:  d.Get("svcstatelearningtime").(int),
		V6ldnsmasklen:         d.Get("v6ldnsmasklen").(int),
	}

	err := client.UpdateUnnamedResource(service.Gslbparameter.Type(), &gslbparameter)
	if err != nil {
		return err
	}

	d.SetId(gslbparameterName)

	err = readGslbparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just updated this gslbparameter but we can't read it ?? ")
		return nil
	}
	return nil
}

func readGslbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading gslbparameter state ")
	data, err := client.FindResource(service.Gslbparameter.Type(), "") //TODO: is this correct?
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing gslbparameter state ")
		d.SetId("")
		return nil
	}
	d.Set("automaticconfigsync", data["automaticconfigsync"])
	d.Set("dropldnsreq", data["dropldnsreq"])
	d.Set("gslbconfigsyncmonitor", data["gslbconfigsyncmonitor"])
	d.Set("gslbsvcstatedelaytime", data["gslbsvcstatedelaytime"])
	d.Set("gslbsyncinterval", data["gslbsyncinterval"])
	d.Set("gslbsynclocfiles", data["gslbsynclocfiles"])
	d.Set("gslbsyncmode", data["gslbsyncmode"])
	d.Set("ldnsentrytimeout", data["ldnsentrytimeout"])
	d.Set("ldnsmask", data["ldnsmask"])
	d.Set("ldnsprobeorder", data["ldnsprobeorder"])
	d.Set("mepkeepalivetimeout", data["mepkeepalivetimeout"])
	d.Set("rtttolerance", data["rtttolerance"])
	d.Set("svcstatelearningtime", data["svcstatelearningtime"])
	d.Set("v6ldnsmasklen", data["v6ldnsmasklen"])

	return nil

}

func updateGslbparameterFunc(d *schema.ResourceData, meta interface{}) error {
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
		gslbparameter.Ldnsprobeorder = d.Get("ldnsprobeorder").([]interface{}) //TODO: Ask if this is correct
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
		log.Printf("[DEBUG]  citrixadc-provider: Svcstatelearningtime has changed for gslbparameter, starting update")  //TODO: ask about this change "[DEBUG]  citrixadc-provider: Svcstatelearningtime has changed for gslbparameter %s, starting update", gslbparameterName
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
			return fmt.Errorf("Error updating gslbparameter: %s", err.Error())
		}
	}
	return readGslbparameterFunc(d, meta)
}

func deleteGslbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbparameterFunc")
	// gslbparameter does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}