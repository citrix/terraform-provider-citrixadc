package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslvserver_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslvserver_sslcertkey_bindingFunc,
		Read:          readSslvserver_sslcertkey_bindingFunc,
		Delete:        deleteSslvserver_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certkeyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crlcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocspcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"skipcaname": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"snicert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vservername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslvserver_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslvserver_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vservername := d.Get("vservername").(string)
	certkeyname := d.Get("certkeyname").(string)
	bindingId := fmt.Sprintf("%s,%s", vservername, certkeyname)

	sslvserver_sslcertkey_binding := ssl.Sslvservercertkeybinding{
		Ca:          d.Get("ca").(bool),
		Certkeyname: d.Get("certkeyname").(string),
		Crlcheck:    d.Get("crlcheck").(string),
		Ocspcheck:   d.Get("ocspcheck").(string),
		Skipcaname:  d.Get("skipcaname").(bool),
		Snicert:     d.Get("snicert").(bool),
		Vservername: d.Get("vservername").(string),
	}

	err := client.UpdateUnnamedResource(service.Sslvserver_sslcertkey_binding.Type(), &sslvserver_sslcertkey_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslvserver_sslcertkey_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslvserver_sslcertkey_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslvserver_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslvserver_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vservername := idSlice[0]
	certkeyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslvserver_sslcertkey_binding state %s", bindingId)
	findParams := service.FindParams{
		ResourceType:             "sslvserver_sslcertkey_binding",
		ResourceName:             vservername,
		ResourceMissingErrorCode: 461,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right certkeyname
	foundIndex := -1
	for i, v := range dataArr {
		if v["certkeyname"].(string) == certkeyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ca", data["ca"])
	d.Set("certkeyname", data["certkeyname"])
	d.Set("crlcheck", data["crlcheck"])
	d.Set("ocspcheck", data["ocspcheck"])
	d.Set("skipcaname", data["skipcaname"])
	d.Set("snicert", data["snicert"])
	d.Set("vservername", data["vservername"])

	return nil

}

func deleteSslvserver_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslvserver_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	vservername := idSlice[0]
	certkeyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("certkeyname:%v", certkeyname))

	if v, ok := d.GetOk("ca"); ok {
		args = append(args, fmt.Sprintf("ca:%v", v))
	}

	if v, ok := d.GetOk("crlcheck"); ok {
		args = append(args, fmt.Sprintf("crlcheck:%v", v))
	}

	if v, ok := d.GetOk("snicert"); ok {
		args = append(args, fmt.Sprintf("snicert:%v", v))
	}

	if v, ok := d.GetOk("ocspcheck"); ok {
		args = append(args, fmt.Sprintf("ocspcheck:%v", v))
	}

	err := client.DeleteResourceWithArgs(service.Sslvserver_sslcertkey_binding.Type(), vservername, args)
	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}
