package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcSystembackup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystembackupFunc,
		Read:          schema.Noop,
		Delete:        deleteSystembackupFunc,
		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSystembackupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName :=  resource.PrefixedUniqueId(d.Get("filename").(string) + "-")

	systembackup := system.Systembackup{
		Filename:         d.Get("filename").(string),
	}

	_, err := client.AddResource(service.Systembackup.Type(), systembackupName, &systembackup, )
	if err != nil {
		return err
	}

	d.SetId(systembackupName)
	return nil
}

func deleteSystembackupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := d.Get("filename").(string)
	err := client.DeleteResource(service.Systembackup.Type(), systembackupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}