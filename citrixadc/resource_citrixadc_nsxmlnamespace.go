package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsxmlnamespace() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsxmlnamespaceFunc,
		Read:          readNsxmlnamespaceFunc,
		Update:        updateNsxmlnamespaceFunc,
		Delete:        deleteNsxmlnamespaceFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsxmlnamespaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Get("prefix").(string)
	nsxmlnamespace := ns.Nsxmlnamespace{
		Description: d.Get("description").(string),
		Namespace:   d.Get("namespace").(string),
		Prefix:      d.Get("prefix").(string),
	}

	_, err := client.AddResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName, &nsxmlnamespace)
	if err != nil {
		return err
	}

	d.SetId(nsxmlnamespaceName)

	err = readNsxmlnamespaceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsxmlnamespace but we can't read it ?? %s", nsxmlnamespaceName)
		return nil
	}
	return nil
}

func readNsxmlnamespaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsxmlnamespace state %s", nsxmlnamespaceName)
	data, err := client.FindResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsxmlnamespace state %s", nsxmlnamespaceName)
		d.SetId("")
		return nil
	}
	d.Set("description", data["description"])
	//d.Set("namespace", data["namespace"])
	d.Set("prefix", data["prefix"])

	return nil

}

func updateNsxmlnamespaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Get("prefix").(string)

	nsxmlnamespace := ns.Nsxmlnamespace{
		Prefix: d.Get("prefix").(string),
	}
	hasChange := false
	if d.HasChange("description") {
		log.Printf("[DEBUG]  citrixadc-provider: Description has changed for nsxmlnamespace %s, starting update", nsxmlnamespaceName)
		nsxmlnamespace.Description = d.Get("description").(string)
		hasChange = true
	}
	if d.HasChange("namespace") {
		log.Printf("[DEBUG]  citrixadc-provider: Namespace has changed for nsxmlnamespace %s, starting update", nsxmlnamespaceName)
		nsxmlnamespace.Namespace = d.Get("namespace").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName, &nsxmlnamespace)
		if err != nil {
			return fmt.Errorf("Error updating nsxmlnamespace %s", nsxmlnamespaceName)
		}
	}
	return readNsxmlnamespaceFunc(d, meta)
}

func deleteNsxmlnamespaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Id()
	err := client.DeleteResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
