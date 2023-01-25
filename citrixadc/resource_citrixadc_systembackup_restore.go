package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSystembackupRestore() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystembackupRestoreFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"skipbackup": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSystembackupRestoreFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystembackupRestoreFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName :=  resource.PrefixedUniqueId(d.Get("filename").(string) + "-")

	systembackup := system.Systembackup{
		Filename:         d.Get("filename").(string),
		Skipbackup: 	  d.Get("skipbackup").(bool),
	}

	err := client.ActOnResource(service.Systembackup.Type(), &systembackup, "restore")
	if err != nil {
		return err
	}

	d.SetId(systembackupName)
	return nil
}