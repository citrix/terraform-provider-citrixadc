package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbsipparameters() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbsipparametersFunc,
		Read:          readLbsipparametersFunc,
		Update:        updateLbsipparametersFunc,
		Delete:        deleteLbsipparametersFunc, // Thought lbsipparameters resource does not have a DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"addrportvip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retrydur": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatdstport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatsecuredstport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatsecuresrcport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatsrcport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sip503ratethreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbsipparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbsipparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	var lbsipparametersName string

	// there is no primary key in lbsipparameters resource. Hence generate one for terraform state maintenance
	lbsipparametersName = resource.PrefixedUniqueId("tf-lbsipparameters-")

	lbsipparameters := lb.Lbsipparameters{
		Addrportvip:         d.Get("addrportvip").(string),
		Retrydur:            d.Get("retrydur").(int),
		Rnatdstport:         d.Get("rnatdstport").(int),
		Rnatsecuredstport:   d.Get("rnatsecuredstport").(int),
		Rnatsecuresrcport:   d.Get("rnatsecuresrcport").(int),
		Rnatsrcport:         d.Get("rnatsrcport").(int),
		Sip503ratethreshold: d.Get("sip503ratethreshold").(int),
	}

	err := client.UpdateUnnamedResource(service.Lbsipparameters.Type(), &lbsipparameters)
	if err != nil {
		return fmt.Errorf("Error updating lbsipparameters")
	}

	d.SetId(lbsipparametersName)

	err = readLbsipparametersFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just updated the lbsipparameters but we can't read it ??")
		return nil
	}
	return nil
}

func readLbsipparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbsipparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading lbsipparameters state")
	data, err := client.FindResource(service.Lbsipparameters.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbsipparameters state")
		d.SetId("")
		return nil
	}
	d.Set("addrportvip", data["addrportvip"])
	d.Set("retrydur", data["retrydur"])
	d.Set("rnatdstport", data["rnatdstport"])
	d.Set("rnatsecuredstport", data["rnatsecuredstport"])
	d.Set("rnatsecuresrcport", data["rnatsecuresrcport"])
	d.Set("rnatsrcport", data["rnatsrcport"])
	d.Set("sip503ratethreshold", data["sip503ratethreshold"])

	return nil
}

func updateLbsipparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbsipparametersFunc")
	client := meta.(*NetScalerNitroClient).client

	lbsipparameters := lb.Lbsipparameters{}
	hasChange := false

	if d.HasChange("addrportvip") {
		log.Printf("[DEBUG]  citrixadc-provider: Addrportvip has changed for lbsipparameters, starting update")
		lbsipparameters.Addrportvip = d.Get("addrportvip").(string)
		hasChange = true
	}
	if d.HasChange("retrydur") {
		log.Printf("[DEBUG]  citrixadc-provider: Retrydur has changed for lbsipparameters, starting update")
		lbsipparameters.Retrydur = d.Get("retrydur").(int)
		hasChange = true
	}
	if d.HasChange("rnatdstport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatdstport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatdstport = d.Get("rnatdstport").(int)
		hasChange = true
	}
	if d.HasChange("rnatsecuredstport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatsecuredstport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatsecuredstport = d.Get("rnatsecuredstport").(int)
		hasChange = true
	}
	if d.HasChange("rnatsecuresrcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatsecuresrcport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatsecuresrcport = d.Get("rnatsecuresrcport").(int)
		hasChange = true
	}
	if d.HasChange("rnatsrcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatsrcport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatsrcport = d.Get("rnatsrcport").(int)
		hasChange = true
	}
	if d.HasChange("sip503ratethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sip503ratethreshold has changed for lbsipparameters, starting update")
		lbsipparameters.Sip503ratethreshold = d.Get("sip503ratethreshold").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Lbsipparameters.Type(), &lbsipparameters)
		if err != nil {
			return fmt.Errorf("Error updating lbsipparameters: %s", err.Error())
		}
	}
	return readLbsipparametersFunc(d, meta)
}

func deleteLbsipparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbsipparametersFunc")
	// lbsipparameters do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
