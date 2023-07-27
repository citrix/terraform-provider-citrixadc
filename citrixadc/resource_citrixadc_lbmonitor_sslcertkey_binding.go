package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcLbmonitor_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbmonitor_sslcertkey_bindingFunc,
		Read:          readLbmonitor_sslcertkey_bindingFunc,
		Delete:        deleteLbmonitor_sslcertkey_bindingFunc,
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
			"monitorname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ocspcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbmonitor_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbmonitor_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	monitorName := d.Get("monitorname").(string)
	sslcertkeyName := d.Get("certkeyname").(string)
	bindingId := fmt.Sprintf("%s,%s", monitorName, sslcertkeyName)

	lbmonitor_sslcertkey_binding := lb.Lbmonitorsslcertkeybinding{
		Ca:          d.Get("ca").(bool),
		Certkeyname: sslcertkeyName,
		Crlcheck:    d.Get("crlcheck").(string),
		Monitorname: monitorName,
		Ocspcheck:   d.Get("ocspcheck").(string),
	}

	_, err := client.AddResource("lbmonitor_sslcertkey_binding", monitorName, &lbmonitor_sslcertkey_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLbmonitor_sslcertkey_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbmonitor_sslcertkey_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLbmonitor_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbmonitor_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	monitorName := idSlice[0]
	sslcertkeyName := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lbmonitor_sslcertkey_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lbmonitor_sslcertkey_binding",
		ResourceName:             monitorName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbmonitor_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["certkeyname"].(string) == sslcertkeyName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams sslcertkeyName not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbmonitor_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ca", data["ca"])
	d.Set("certkeyname", data["certkeyname"])
	d.Set("crlcheck", data["crlcheck"])
	d.Set("monitorname", data["monitorname"])
	d.Set("ocspcheck", data["ocspcheck"])

	return nil

}

func deleteLbmonitor_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbmonitor_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	monitorName := idSlice[0]
	sslcertkeyName := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["certkeyname"] = url.QueryEscape(sslcertkeyName)

	if v, ok := d.GetOk("bindpoint"); ok {
		argsMap["bindpoint"] = url.QueryEscape(v.(string))
	}

	err := client.DeleteResourceWithArgsMap("lbmonitor_sslcertkey_binding", monitorName, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
