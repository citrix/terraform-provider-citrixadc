package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaatacacsparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaatacacsparamsFunc,
		Read:          readAaatacacsparamsFunc,
		Update:        updateAaatacacsparamsFunc,
		Delete:        deleteAaatacacsparamsFunc,
		Schema: map[string]*schema.Schema{
			"accounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auditfailedcmds": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupattrname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tacacssecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaatacacsparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaatacacsparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	aaatacacsparamsName := resource.PrefixedUniqueId("tf-aaatacacsparams-")

	aaatacacsparams := aaa.Aaatacacsparams{
		Accounting:                 d.Get("accounting").(string),
		Auditfailedcmds:            d.Get("auditfailedcmds").(string),
		Authorization:              d.Get("authorization").(string),
		Authtimeout:                d.Get("authtimeout").(int),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Groupattrname:              d.Get("groupattrname").(string),
		Serverip:                   d.Get("serverip").(string),
		Serverport:                 d.Get("serverport").(int),
		Tacacssecret:               d.Get("tacacssecret").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaatacacsparams.Type(), &aaatacacsparams)
	if err != nil {
		return err
	}

	d.SetId(aaatacacsparamsName)

	err = readAaatacacsparamsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaatacacsparams but we can't read it ??")
		return nil
	}
	return nil
}

func readAaatacacsparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaatacacsparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaatacacsparams state")
	data, err := client.FindResource(service.Aaatacacsparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaatacacsparams state")
		d.SetId("")
		return nil
	}
	d.Set("accounting", data["accounting"])
	d.Set("auditfailedcmds", data["auditfailedcmds"])
	d.Set("authorization", data["authorization"])
	d.Set("authtimeout", data["authtimeout"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("groupattrname", data["groupattrname"])
	d.Set("serverip", data["serverip"])
	d.Set("serverport", data["serverport"])
	d.Set("tacacssecret", data["tacacssecret"])

	return nil

}

func updateAaatacacsparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaatacacsparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	aaatacacsparams := aaa.Aaatacacsparams{}
	hasChange := false
	if d.HasChange("accounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Accounting has changed for aaatacacsparams, starting update")
		aaatacacsparams.Accounting = d.Get("accounting").(string)
		hasChange = true
	}
	if d.HasChange("auditfailedcmds") {
		log.Printf("[DEBUG]  citrixadc-provider: Auditfailedcmds has changed for aaatacacsparams, starting update")
		aaatacacsparams.Auditfailedcmds = d.Get("auditfailedcmds").(string)
		hasChange = true
	}
	if d.HasChange("authorization") {
		log.Printf("[DEBUG]  citrixadc-provider: Authorization has changed for aaatacacsparams, starting update")
		aaatacacsparams.Authorization = d.Get("authorization").(string)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for aaatacacsparams, starting update")
		aaatacacsparams.Authtimeout = d.Get("authtimeout").(int)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for aaatacacsparams, starting update")
		aaatacacsparams.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("groupattrname") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupattrname has changed for aaatacacsparams, starting update")
		aaatacacsparams.Groupattrname = d.Get("groupattrname").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for aaatacacsparams, starting update")
		aaatacacsparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for aaatacacsparams, starting update")
		aaatacacsparams.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("tacacssecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Tacacssecret has changed for aaatacacsparams, starting update")
		aaatacacsparams.Tacacssecret = d.Get("tacacssecret").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaatacacsparams.Type(), &aaatacacsparams)
		if err != nil {
			return fmt.Errorf("Error updating aaatacacsparams")
		}
	}
	return readAaatacacsparamsFunc(d, meta)
}

func deleteAaatacacsparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaatacacsparamsFunc")
	// aaatacacsparams does not support DELETE operation
	d.SetId("")

	return nil
}
