package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTmsessionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTmsessionactionFunc,
		Read:          readTmsessionactionFunc,
		Update:        updateTmsessionactionFunc,
		Delete:        deleteTmsessionactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"defaultauthorizationaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"homepage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httponlycookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookievalidity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssocredential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssodomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTmsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Get("name").(string)
	
	tmsessionaction := tm.Tmsessionaction{
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Homepage:                   d.Get("homepage").(string),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Name:                       d.Get("name").(string),
		Persistentcookie:           d.Get("persistentcookie").(string),
		Persistentcookievalidity:   d.Get("persistentcookievalidity").(int),
		Sesstimeout:                d.Get("sesstimeout").(int),
		Sso:                        d.Get("sso").(string),
		Ssocredential:              d.Get("ssocredential").(string),
		Ssodomain:                  d.Get("ssodomain").(string),
	}

	_, err := client.AddResource(service.Tmsessionaction.Type(), tmsessionactionName, &tmsessionaction)
	if err != nil {
		return err
	}

	d.SetId(tmsessionactionName)

	err = readTmsessionactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this tmsessionaction but we can't read it ?? %s", tmsessionactionName)
		return nil
	}
	return nil
}

func readTmsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tmsessionaction state %s", tmsessionactionName)
	data, err := client.FindResource(service.Tmsessionaction.Type(), tmsessionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmsessionaction state %s", tmsessionactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthorizationaction", data["defaultauthorizationaction"])
	d.Set("homepage", data["homepage"])
	d.Set("httponlycookie", data["httponlycookie"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("name", data["name"])
	d.Set("persistentcookie", data["persistentcookie"])
	d.Set("persistentcookievalidity", data["persistentcookievalidity"])
	d.Set("sesstimeout", data["sesstimeout"])
	d.Set("sso", data["sso"])
	d.Set("ssocredential", data["ssocredential"])
	d.Set("ssodomain", data["ssodomain"])

	return nil

}

func updateTmsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Get("name").(string)

	tmsessionaction := tm.Tmsessionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthorizationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthorizationaction has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Defaultauthorizationaction = d.Get("defaultauthorizationaction").(string)
		hasChange = true
	}
	if d.HasChange("homepage") {
		log.Printf("[DEBUG]  citrixadc-provider: Homepage has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Homepage = d.Get("homepage").(string)
		hasChange = true
	}
	if d.HasChange("httponlycookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookie has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Httponlycookie = d.Get("httponlycookie").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookie has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Persistentcookie = d.Get("persistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookievalidity") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookievalidity has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Persistentcookievalidity = d.Get("persistentcookievalidity").(int)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Sesstimeout = d.Get("sesstimeout").(int)
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("ssocredential") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredential has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Ssocredential = d.Get("ssocredential").(string)
		hasChange = true
	}
	if d.HasChange("ssodomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssodomain has changed for tmsessionaction %s, starting update", tmsessionactionName)
		tmsessionaction.Ssodomain = d.Get("ssodomain").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmsessionaction.Type(), &tmsessionaction)
		if err != nil {
			return fmt.Errorf("Error updating tmsessionaction %s", tmsessionactionName)
		}
	}
	return readTmsessionactionFunc(d, meta)
}

func deleteTmsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionactionName := d.Id()
	err := client.DeleteResource(service.Tmsessionaction.Type(), tmsessionactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
