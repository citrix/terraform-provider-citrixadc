package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcSslvserver_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslvserver_sslcertkey_bindingFunc,
		ReadContext:   readSslvserver_sslcertkey_bindingFunc,
		DeleteContext: deleteSslvserver_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
			"skipcaname": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"snicert": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"vservername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslvserver_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslvserver_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vservername := d.Get("vservername").(string)
	certkeyname := d.Get("certkeyname").(string)
	snicert := d.Get("snicert").(bool)
	ca := d.Get("ca").(bool)
	bindingId := fmt.Sprintf("%s,%s,%t,%t", vservername, certkeyname, snicert, ca)

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
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslvserver_sslcertkey_bindingFunc(ctx, d, meta)
}

func readSslvserver_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslvserver_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	vservername := idSlice[0]
	certkeyname := idSlice[1]
	snicert := false
	ca := false
	if len(idSlice) > 2 {
		snicert = idSlice[2] == "true"
		if len(idSlice) > 3 {
			ca = idSlice[3] == "true"
		} else {
			ca = d.Get("ca").(bool)
			bindingId = fmt.Sprintf("%s,%s,%t,%t", vservername, certkeyname, snicert, ca)
			d.SetId(bindingId)
		}
	} else {
		snicert = d.Get("snicert").(bool)
		ca = d.Get("ca").(bool)
		bindingId = fmt.Sprintf("%s,%t,%t", bindingId, snicert, ca)
		d.SetId(bindingId)
	}

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
		return diag.FromErr(err)
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
		if v["certkeyname"].(string) == certkeyname && v["snicert"].(bool) == snicert && v["ca"].(bool) == ca {
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

func deleteSslvserver_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslvserver_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	vservername := idSlice[0]
	certkeyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("certkeyname:%v", certkeyname))

	if v, ok := d.GetOk("ca"); ok && v.(bool) {
		args = append(args, fmt.Sprintf("ca:%v", v))
	}

	if v, ok := d.GetOk("snicert"); ok && v.(bool) {
		args = append(args, fmt.Sprintf("snicert:%v", v))
	}

	err := client.DeleteResourceWithArgs(service.Sslvserver_sslcertkey_binding.Type(), vservername, args)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}
