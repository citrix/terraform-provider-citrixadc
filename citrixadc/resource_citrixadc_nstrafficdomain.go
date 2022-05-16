package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"strconv"
)

func resourceCitrixAdcNstrafficdomain() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstrafficdomainFunc,
		Read:          readNstrafficdomainFunc,
		Delete:        deleteNstrafficdomainFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aliasname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNstrafficdomainFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstrafficdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	td_Id := d.Get("td").(int)
	nstrafficdomain := ns.Nstrafficdomain{
		Aliasname: d.Get("aliasname").(string),
		Td:        d.Get("td").(int),
		Vmac:      d.Get("vmac").(string),
	}
	td_IdStr := strconv.Itoa(td_Id)

	_, err := client.AddResource(service.Nstrafficdomain.Type(), td_IdStr, &nstrafficdomain)
	if err != nil {
		return err
	}

	d.SetId(td_IdStr)

	err = readNstrafficdomainFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstrafficdomain but we can't read it ?? %s", td_IdStr)
		return nil
	}
	return nil
}

func readNstrafficdomainFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstrafficdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	td_IdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nstrafficdomain state %s", td_IdStr)
	data, err := client.FindResource(service.Nstrafficdomain.Type(), td_IdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain state %s", td_IdStr)
		d.SetId("")
		return nil
	}
	d.Set("aliasname", data["aliasname"])
	d.Set("td", data["td"])
	d.Set("vmac", data["vmac"])

	return nil

}

func deleteNstrafficdomainFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstrafficdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	nstrafficdomainName := d.Id()
	err := client.DeleteResource(service.Nstrafficdomain.Type(), nstrafficdomainName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
