package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	// "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSslcertfile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcertfileFunc,
		Read:          readSslcertfileFunc,
		Delete:        deleteSslcertfileFunc,
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

func createSslcertfileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcertfileFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertfileName := d.Get("name").(string)
	sslcertfile := ssl.Sslcertfile{
		Name: d.Get("name").(string),
		Src:  d.Get("src").(string),
	}

	err := client.ActOnResource(service.Sslcertfile.Type(), &sslcertfile, "import")
	if err != nil {
		return err
	}

	d.SetId(sslcertfileName)

	err = readSslcertfileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslcertfile but we can't read it ?? %s", sslcertfileName)
		return nil
	}
	return nil
}

func readSslcertfileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcertfileFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertfileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcertfile state %s", sslcertfileName)

	dataArr, err := client.FindAllResources(service.Sslcertfile.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcertfile state %s", sslcertfileName)
		d.SetId("")
		return nil
	}
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslcertfile state %s", sslcertfileName)
		d.SetId("")
		return nil
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["name"].(string) == d.Get("name").(string) {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources ipaddress or port not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslcertfile state %s", sslcertfileName)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("name", data["name"])
	// d.Set("src", data["src"])

	return nil

}

func deleteSslcertfileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcertfileFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("name:%v", d.Get("name").(string)))
	err := client.DeleteResourceWithArgs(service.Sslcertfile.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
