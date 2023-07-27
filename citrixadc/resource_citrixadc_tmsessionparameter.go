package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTmsessionparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTmsessionparameterFunc,
		Read:          readTmsessionparameterFunc,
		Update:        updateTmsessionparameterFunc,
		Delete:        deleteTmsessionparameterFunc,
		Schema: map[string]*schema.Schema{
			"defaultauthorizationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"homepage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httponlycookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookievalidity": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sso": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssocredential": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssodomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTmsessionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmsessionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionparameterName := resource.PrefixedUniqueId("tf-tmsessionparameter-")

	tmsessionparameter := tm.Tmsessionparameter{
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Homepage:                   d.Get("homepage").(string),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Persistentcookie:           d.Get("persistentcookie").(string),
		Persistentcookievalidity:   d.Get("persistentcookievalidity").(int),
		Sesstimeout:                d.Get("sesstimeout").(int),
		Sso:                        d.Get("sso").(string),
		Ssocredential:              d.Get("ssocredential").(string),
		Ssodomain:                  d.Get("ssodomain").(string),
	}

	err := client.UpdateUnnamedResource(service.Tmsessionparameter.Type(), &tmsessionparameter)
	if err != nil {
		return err
	}

	d.SetId(tmsessionparameterName)

	err = readTmsessionparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this tmsessionparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readTmsessionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmsessionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading tmsessionparameter state")
	data, err := client.FindResource(service.Tmsessionparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmsessionparameter state")
		d.SetId("")
		return nil
	}
	d.Set("defaultauthorizationaction", data["defaultauthorizationaction"])
	d.Set("homepage", data["homepage"])
	d.Set("httponlycookie", data["httponlycookie"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("persistentcookie", data["persistentcookie"])
	d.Set("persistentcookievalidity", data["persistentcookievalidity"])
	d.Set("sesstimeout", data["sesstimeout"])
	d.Set("sso", data["sso"])
	d.Set("ssocredential", data["ssocredential"])
	d.Set("ssodomain", data["ssodomain"])

	return nil

}

func updateTmsessionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmsessionparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	tmsessionparameter := tm.Tmsessionparameter{}
	hasChange := false
	if d.HasChange("defaultauthorizationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthorizationaction has changed for tmsessionparameter, starting update")
		tmsessionparameter.Defaultauthorizationaction = d.Get("defaultauthorizationaction").(string)
		hasChange = true
	}
	if d.HasChange("homepage") {
		log.Printf("[DEBUG]  citrixadc-provider: Homepage has changed for tmsessionparameter, starting update")
		tmsessionparameter.Homepage = d.Get("homepage").(string)
		hasChange = true
	}
	if d.HasChange("httponlycookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookie has changed for tmsessionparameter, starting update")
		tmsessionparameter.Httponlycookie = d.Get("httponlycookie").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for tmsessionparameter, starting update")
		tmsessionparameter.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookie has changed for tmsessionparameter, starting update")
		tmsessionparameter.Persistentcookie = d.Get("persistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookievalidity") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookievalidity has changed for tmsessionparameter, starting update")
		tmsessionparameter.Persistentcookievalidity = d.Get("persistentcookievalidity").(int)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for tmsessionparameter, starting update")
		tmsessionparameter.Sesstimeout = d.Get("sesstimeout").(int)
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for tmsessionparameter, starting update")
		tmsessionparameter.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("ssocredential") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredential has changed for tmsessionparameter, starting update")
		tmsessionparameter.Ssocredential = d.Get("ssocredential").(string)
		hasChange = true
	}
	if d.HasChange("ssodomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssodomain has changed for tmsessionparameter, starting update")
		tmsessionparameter.Ssodomain = d.Get("ssodomain").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmsessionparameter.Type(), &tmsessionparameter)
		if err != nil {
			return fmt.Errorf("Error updating tmsessionparameter")
		}
	}
	return readTmsessionparameterFunc(d, meta)
}

func deleteTmsessionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmsessionparameterFunc")
	// tmsessionparameter does not support DELETE operation
	d.SetId("")

	return nil
}
