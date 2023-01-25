package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceCitrixAdcVpnglobal_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_sslcertkey_bindingFunc,
		Read:          readVpnglobal_sslcertkey_bindingFunc,
		Delete:        deleteVpnglobal_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"certkeyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"cacert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"crlcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
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
			"userdataencryptionkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	certkeyname := d.Get("certkeyname").(string)
	vpnglobal_sslcertkey_binding := vpn.Vpnglobalsslcertkeybinding{
		Cacert:                 d.Get("cacert").(string),
		Certkeyname:            d.Get("certkeyname").(string),
		Crlcheck:               d.Get("crlcheck").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Ocspcheck:              d.Get("ocspcheck").(string),
		Userdataencryptionkey:  d.Get("userdataencryptionkey").(string),
	}

	err := client.UpdateUnnamedResource("vpnglobal_sslcertkey_binding", &vpnglobal_sslcertkey_binding)
	if err != nil {
		return err
	}

	d.SetId(certkeyname)

	err = readVpnglobal_sslcertkey_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_sslcertkey_binding but we can't read it ?? %s", certkeyname)
		return nil
	}
	return nil
}

func readVpnglobal_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	certkeyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_sslcertkey_binding state %s", certkeyname)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_sslcertkey_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_sslcertkey_binding state %s", certkeyname)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_sslcertkey_binding state %s", certkeyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("cacert", data["cacert"])
	d.Set("certkeyname", data["certkeyname"])
	d.Set("crlcheck", data["crlcheck"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("ocspcheck", data["ocspcheck"])
	d.Set("userdataencryptionkey", data["userdataencryptionkey"])

	return nil

}

func deleteVpnglobal_sslcertkey_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	certkeyname := d.Id()
	argsMap := make(map[string]string)
	argsMap["certkeyname"] = certkeyname
	if val,ok := d.GetOk("userdataencryptionkey"); ok{
		argsMap["userdataencryptionkey"] = val.(string)
	}
	if val,ok := d.GetOk("cacert"); ok{
		argsMap["cacert"] = val.(string)
	}

	err := client.DeleteResourceWithArgsMap("vpnglobal_sslcertkey_binding", "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
