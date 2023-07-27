package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslservice_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslservice_sslcertkey_bindingFunc,
		Read:          readSslservice_sslcertkey_bindingFunc,
		Delete:        deleteSslservice_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ca": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certkeyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crlcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocspcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"skipcaname": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"snicert": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservice_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservice_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicename := d.Get("servicename")
	certkeyname := d.Get("certkeyname")
	bindingId := fmt.Sprintf("%s,%s", servicename, certkeyname)
	sslservice_sslcertkey_binding := ssl.Sslservicesslcertkeybinding{
		Ca:          d.Get("ca").(bool),
		Certkeyname: d.Get("certkeyname").(string),
		Crlcheck:    d.Get("crlcheck").(string),
		Ocspcheck:   d.Get("ocspcheck").(string),
		Servicename: d.Get("servicename").(string),
		Skipcaname:  d.Get("skipcaname").(bool),
		Snicert:     d.Get("snicert").(bool),
	}

	err := client.UpdateUnnamedResource(service.Sslservice_sslcertkey_binding.Type(), &sslservice_sslcertkey_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslservice_sslcertkey_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslservice_sslcertkey_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslservice_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservice_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	certkeyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslservice_sslcertkey_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslservice_sslcertkey_binding",
		ResourceName:             servicename,
		ResourceMissingErrorCode: 258,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["certkeyname"].(string) == certkeyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams certkeyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ca", data["ca"])
	d.Set("certkeyname", data["certkeyname"])
	d.Set("crlcheck", data["crlcheck"])
	d.Set("ocspcheck", data["ocspcheck"])
	d.Set("servicename", data["servicename"])
	d.Set("skipcaname", data["skipcaname"])
	d.Set("snicert", data["snicert"])

	return nil

}

func deleteSslservice_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservice_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	certkeyname := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["certkeyname"] = url.QueryEscape(certkeyname)

	if v, ok := d.GetOk("ca"); ok {
		argsMap["ca"] = url.QueryEscape(fmt.Sprintf("%v", v))

	}

	if v, ok := d.GetOk("crlcheck"); ok {
		argsMap["crlcheck"] = url.QueryEscape(v.(string))
	}

	if v, ok := d.GetOk("snicert"); ok {
		argsMap["snicert"] = url.QueryEscape(v.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Sslservice_sslcertkey_binding.Type(), servicename, argsMap)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
