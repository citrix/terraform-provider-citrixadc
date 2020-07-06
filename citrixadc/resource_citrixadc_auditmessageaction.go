package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/audit"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuditmessageaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditmessageactionFunc,
		Read:          readAuditmessageactionFunc,
		Update:        updateAuditmessageactionFunc,
		Delete:        deleteAuditmessageactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bypasssafetycheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loglevel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtonewnslog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stringbuilderexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuditmessageactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client

	auditmessageactionName := d.Get("name").(string)
	d.Set("name", auditmessageactionName)

	auditmessageaction := audit.Auditmessageaction{
		Bypasssafetycheck: d.Get("bypasssafetycheck").(string),
		Loglevel:          d.Get("loglevel").(string),
		Logtonewnslog:     d.Get("logtonewnslog").(string),
		Name:              d.Get("name").(string),
		Stringbuilderexpr: d.Get("stringbuilderexpr").(string),
	}

	_, err := client.AddResource(netscaler.Auditmessageaction.Type(), auditmessageactionName, &auditmessageaction)
	if err != nil {
		return err
	}

	d.SetId(auditmessageactionName)

	err = readAuditmessageactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditmessageaction but we can't read it ?? %s", auditmessageactionName)
		return nil
	}
	return nil
}

func readAuditmessageactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditmessageactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading auditmessageaction state %s", auditmessageactionName)
	data, err := client.FindResource(netscaler.Auditmessageaction.Type(), auditmessageactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditmessageaction state %s", auditmessageactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("bypasssafetycheck", data["bypasssafetycheck"])
	d.Set("loglevel", data["loglevel1"])
	d.Set("logtonewnslog", data["logtonewnslog"])
	d.Set("name", data["name"])
	d.Set("stringbuilderexpr", data["stringbuilderexpr"])

	return nil

}

func updateAuditmessageactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditmessageactionName := d.Get("name").(string)

	auditmessageaction := audit.Auditmessageaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bypasssafetycheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Bypasssafetycheck has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Bypasssafetycheck = d.Get("bypasssafetycheck").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Loglevel = d.Get("loglevel").(string)
		hasChange = true
	}
	if d.HasChange("logtonewnslog") {
		log.Printf("[DEBUG]  citrixadc-provider: Logtonewnslog has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Logtonewnslog = d.Get("logtonewnslog").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("stringbuilderexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Stringbuilderexpr has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Stringbuilderexpr = d.Get("stringbuilderexpr").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Auditmessageaction.Type(), auditmessageactionName, &auditmessageaction)
		if err != nil {
			return fmt.Errorf("Error updating auditmessageaction %s", auditmessageactionName)
		}
	}
	return readAuditmessageactionFunc(d, meta)
}

func deleteAuditmessageactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditmessageactionName := d.Id()
	err := client.DeleteResource(netscaler.Auditmessageaction.Type(), auditmessageactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
