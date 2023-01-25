package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppqoecustomresp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppqoecustomrespFunc,
		Read:          readAppqoecustomrespFunc,
		Delete:        deleteAppqoecustomrespFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAppqoecustomrespFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppqoecustomrespFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoecustomrespName := d.Get("name").(string)
	appqoecustomresp := appqoe.Appqoecustomresp{
		Name: d.Get("name").(string),
		Src:  d.Get("src").(string),
	}

	err := client.ActOnResource(service.Appqoecustomresp.Type(), &appqoecustomresp, "import")
	if err != nil {
		return err
	}

	d.SetId(appqoecustomrespName)

	err = readAppqoecustomrespFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appqoecustomresp but we can't read it ?? %s", appqoecustomrespName)
		return nil
	}
	return nil
}

func readAppqoecustomrespFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppqoecustomrespFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoecustomrespName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appqoecustomresp state %s", appqoecustomrespName)
	dataArr, err := client.FindAllResources(service.Appqoecustomresp.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appqoecustomresp state %s", appqoecustomrespName)
		d.SetId("")
		return nil
	}
	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: appqoecustomresp does not exist. Clearing state.")
		d.SetId("")
		return nil
	}
	
	foundIndex := -1
	for i, v := range dataArr {
		if v["name"].(string) == appqoecustomrespName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllresources appqoecustomresp not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appqoecustomresp state %s", appqoecustomrespName)
		d.SetId("")
		return nil
	}
	data := dataArr[foundIndex]
	d.Set("name", data["name"])
	//d.Set("src", data["src"])

	return nil

}

func deleteAppqoecustomrespFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppqoecustomrespFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoecustomrespName := d.Id()
	err := client.DeleteResource(service.Appqoecustomresp.Type(), appqoecustomrespName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
