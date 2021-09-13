package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcPolicymap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicymapFunc,
		Read:          readPolicymapFunc,
		Delete:        deletePolicymapFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"mappolicyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sd": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"su": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicymapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicymapFunc")
	client := meta.(*NetScalerNitroClient).client
	var policymapName string
	if v, ok := d.GetOk("mappolicyname"); ok {
		policymapName = v.(string)
	} else {
		policymapName = resource.PrefixedUniqueId("tf-policymap-")
		d.Set("mappolicyname", policymapName)
	}
	policymap := policy.Policymap{
		Mappolicyname: d.Get("mappolicyname").(string),
		Sd:            d.Get("sd").(string),
		Su:            d.Get("su").(string),
		Td:            d.Get("td").(string),
		Tu:            d.Get("tu").(string),
	}

	_, err := client.AddResource(service.Policymap.Type(), policymapName, &policymap)
	if err != nil {
		return err
	}

	d.SetId(policymapName)

	err = readPolicymapFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policymap but we can't read it ?? %s", policymapName)
		return nil
	}
	return nil
}

func readPolicymapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicymapFunc")
	client := meta.(*NetScalerNitroClient).client
	policymapName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policymap state %s", policymapName)
	data, err := client.FindResource(service.Policymap.Type(), policymapName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policymap state %s", policymapName)
		d.SetId("")
		return nil
	}
	d.Set("mappolicyname", data["mappolicyname"])
	d.Set("mappolicyname", data["mappolicyname"])
	d.Set("sd", data["sd"])
	d.Set("su", data["su"])
	d.Set("td", data["td"])
	d.Set("tu", data["tu"])

	return nil

}

func deletePolicymapFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicymapFunc")
	client := meta.(*NetScalerNitroClient).client
	policymapName := d.Id()
	err := client.DeleteResource(service.Policymap.Type(), policymapName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
