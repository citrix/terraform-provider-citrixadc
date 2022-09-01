package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/db"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDbdbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDbdbprofileFunc,
		Read:          readDbdbprofileFunc,
		Update:        updateDbdbprofileFunc,
		Delete:        deleteDbdbprofileFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablecachingconmuxoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interpretquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stickiness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDbdbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Get("name").(string)
	dbdbprofile := db.Dbdbprofile{
		Conmultiplex:           d.Get("conmultiplex").(string),
		Enablecachingconmuxoff: d.Get("enablecachingconmuxoff").(string),
		Interpretquery:         d.Get("interpretquery").(string),
		Kcdaccount:             d.Get("kcdaccount").(string),
		Name:                   d.Get("name").(string),
		Stickiness:             d.Get("stickiness").(string),
	}

	_, err := client.AddResource(service.Dbdbprofile.Type(), dbdbprofileName, &dbdbprofile)
	if err != nil {
		return err
	}

	d.SetId(dbdbprofileName)

	err = readDbdbprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dbdbprofile but we can't read it ?? %s", dbdbprofileName)
		return nil
	}
	return nil
}

func readDbdbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dbdbprofile state %s", dbdbprofileName)
	data, err := client.FindResource(service.Dbdbprofile.Type(), dbdbprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dbdbprofile state %s", dbdbprofileName)
		d.SetId("")
		return nil
	}
	d.Set("conmultiplex", data["conmultiplex"])
	d.Set("enablecachingconmuxoff", data["enablecachingconmuxoff"])
	d.Set("interpretquery", data["interpretquery"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("name", data["name"])
	d.Set("stickiness", data["stickiness"])

	return nil

}

func updateDbdbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Get("name").(string)

	dbdbprofile := db.Dbdbprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("conmultiplex") {
		log.Printf("[DEBUG]  citrixadc-provider: Conmultiplex has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Conmultiplex = d.Get("conmultiplex").(string)
		hasChange = true
	}
	if d.HasChange("enablecachingconmuxoff") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablecachingconmuxoff has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Enablecachingconmuxoff = d.Get("enablecachingconmuxoff").(string)
		hasChange = true
	}
	if d.HasChange("interpretquery") {
		log.Printf("[DEBUG]  citrixadc-provider: Interpretquery has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Interpretquery = d.Get("interpretquery").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("stickiness") {
		log.Printf("[DEBUG]  citrixadc-provider: Stickiness has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Stickiness = d.Get("stickiness").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Dbdbprofile.Type(), &dbdbprofile)
		if err != nil {
			return fmt.Errorf("Error updating dbdbprofile %s", dbdbprofileName)
		}
	}
	return readDbdbprofileFunc(d, meta)
}

func deleteDbdbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Id()
	err := client.DeleteResource(service.Dbdbprofile.Type(), dbdbprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
