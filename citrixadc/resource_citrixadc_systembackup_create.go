package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSystemCreatebackup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystembackupCreateFunc,
		Read:          readSystembackupCreateFunc,
		Delete:        deleteSystembackupCreateFunc,
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"includekernel": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"level": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"uselocaltimezone": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSystembackupCreateFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := resource.PrefixedUniqueId(d.Get("filename").(string) + "-")

	systembackup := system.Systembackup{
		Filename:         d.Get("filename").(string),
		Uselocaltimezone: d.Get("uselocaltimezone").(bool),
		Level:            d.Get("level").(string),
		Includekernel:    d.Get("includekernel").(string),
		Comment:          d.Get("comment").(string),
	}

	err := client.ActOnResource(service.Systembackup.Type(), &systembackup, "create")
	if err != nil {
		return err
	}

	d.SetId(systembackupName)

	err = readSystembackupCreateFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systembackup but we can't read it ?? %s", systembackupName)
		return nil
	}
	return nil
}

func readSystembackupCreateFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systembackup state %s", systembackupName)
	data, err := client.FindResource(service.Systembackup.Type(), d.Get("filename").(string)+".tgz")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systembackup state %s", systembackupName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	//d.Set("filename", data["filename"])
	//d.Set("includekernel", data["includekernel"])
	//d.Set("level", data["level"])
	//d.Set("skipbackup", data["skipbackup"])
	d.Set("uselocaltimezone", data["uselocaltimezone"])

	return nil

}

func deleteSystembackupCreateFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := d.Get("filename").(string) + ".tgz"
	err := client.DeleteResource(service.Systembackup.Type(), systembackupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
