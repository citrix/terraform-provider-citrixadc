package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSystemparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemparameterFunc,
		Read:          readSystemparameterFunc,
		Update:        updateSystemparameterFunc,
		Delete:        deleteSystemparameterFunc,
		Schema: map[string]*schema.Schema{
			"basicauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cliloglevel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"doppler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fipsusermode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forcepasswordchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"googleanalytics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minpasswordlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"natpcbforceflushlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"natpcbrstontimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"promptstring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rbaonresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reauthonauthparamchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"removesensitivefiles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrictedtimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strongpassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"totalauthtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSystemparameterFunc(d *schema.ResourceData, meta interface{}) error {
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
		Minpasswordlen:          d.Get("minpasswordlen").(int),
		Natpcbforceflushlimit:   d.Get("natpcbforceflushlimit").(int),
		Natpcbrstontimeout:      d.Get("natpcbrstontimeout").(string),
		Promptstring:            d.Get("promptstring").(string),
		Rbaonresponse:           d.Get("rbaonresponse").(string),
		Reauthonauthparamchange: d.Get("reauthonauthparamchange").(string),
		Removesensitivefiles:    d.Get("removesensitivefiles").(string),
		Restrictedtimeout:       d.Get("restrictedtimeout").(string),
		Strongpassword:          d.Get("strongpassword").(string),
		Timeout:                 d.Get("timeout").(int),
		Totalauthtimeout:        d.Get("totalauthtimeout").(int),
	}

	err := client.UpdateUnnamedResource(service.Systemparameter.Type(), &systemparameter)
	if err != nil {
		return err
	}

	d.SetId(systemparameterName)

	err = readSystemparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readSystemparameterFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("minpasswordlen", data["minpasswordlen"])
	d.Set("natpcbforceflushlimit", data["natpcbforceflushlimit"])
	d.Set("natpcbrstontimeout", data["natpcbrstontimeout"])
	d.Set("promptstring", data["promptstring"])
	d.Set("rbaonresponse", data["rbaonresponse"])
	d.Set("reauthonauthparamchange", data["reauthonauthparamchange"])
	d.Set("removesensitivefiles", data["removesensitivefiles"])
	d.Set("restrictedtimeout", data["restrictedtimeout"])
	d.Set("strongpassword", data["strongpassword"])
	d.Set("timeout", data["timeout"])
	d.Set("totalauthtimeout", data["totalauthtimeout"])

	return nil

}

func updateSystemparameterFunc(d *schema.ResourceData, meta interface{}) error {
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
		systemparameter.Minpasswordlen = d.Get("minpasswordlen").(int)
		hasChange = true
	}
	if d.HasChange("natpcbforceflushlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Natpcbforceflushlimit has changed for systemparameter, starting update")
		systemparameter.Natpcbforceflushlimit = d.Get("natpcbforceflushlimit").(int)
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
		systemparameter.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("totalauthtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Totalauthtimeout has changed for systemparameter, starting update")
		systemparameter.Totalauthtimeout = d.Get("totalauthtimeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Systemparameter.Type(), &systemparameter)
		if err != nil {
			return fmt.Errorf("Error updating systemparameters")
		}
	}
	return readSystemparameterFunc(d, meta)
}

func deleteSystemparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemparameterFunc")
	// systemparameter does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
