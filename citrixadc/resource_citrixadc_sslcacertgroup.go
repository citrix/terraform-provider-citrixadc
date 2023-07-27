package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslcacertgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcacertgroupFunc,
		Read:          readSslcacertgroupFunc,
		Delete:        deleteSslcacertgroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cacertgroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcacertgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcacertgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcacertgroupName := d.Get("cacertgroupname").(string)

	sslcacertgroup := ssl.Sslcacertgroup{
		Cacertgroupname: sslcacertgroupName,
	}

	_, err := client.AddResource("sslcacertgroup", sslcacertgroupName, &sslcacertgroup)
	if err != nil {
		return err
	}

	d.SetId(sslcacertgroupName)

	err = readSslcacertgroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslcacertgroup but we can't read it ?? %s", sslcacertgroupName)
		return nil
	}
	return nil
}

func readSslcacertgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcacertgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcacertgroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcacertgroup state %s", sslcacertgroupName)
	data, err := client.FindResource("sslcacertgroup", sslcacertgroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcacertgroup state %s", sslcacertgroupName)
		d.SetId("")
		return nil
	}
	d.Set("cacertgroupname", data["cacertgroupname"])

	return nil

}

func deleteSslcacertgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcacertgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcacertgroupName := d.Id()
	err := client.DeleteResource("sslcacertgroup", sslcacertgroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
