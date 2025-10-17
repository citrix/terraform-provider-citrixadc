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

func resourceCitrixAdcSslcertkey_sslocspresponder_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcertkey_sslocspresponder_bindingFunc,
		ReadContext:   readSslcertkey_sslocspresponder_bindingFunc,
		DeleteContext: deleteSslcertkey_sslocspresponder_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"certkey": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ocspresponder": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcertkey_sslocspresponder_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcertkey_sslocspresponder_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	certkey := d.Get("certkey").(string)
	ocspresponder := d.Get("ocspresponder").(string)
	bindingId := fmt.Sprintf("%s,%s", certkey, ocspresponder)
	sslcertkey_sslocspresponder_binding := ssl.Sslcertkeysslocspresponderbinding{
		Certkey:       d.Get("certkey").(string),
		Ocspresponder: d.Get("ocspresponder").(string),
	}

	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		sslcertkey_sslocspresponder_binding.Priority = intPtr(d.Get("priority").(int))
	}

	_, err := client.AddResource(service.Sslcertkey_sslocspresponder_binding.Type(), bindingId, &sslcertkey_sslocspresponder_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslcertkey_sslocspresponder_bindingFunc(ctx, d, meta)
}

func readSslcertkey_sslocspresponder_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcertkey_sslocspresponder_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	certkey := idSlice[0]
	ocspresponder := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslcertkey_sslocspresponder_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslcertkey_sslocspresponder_binding",
		ResourceName:             certkey,
		ResourceMissingErrorCode: 258,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslcertkey_sslocspresponder_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ocspresponder"].(string) == ocspresponder {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslcertkey_sslocspresponder_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("certkey", data["certkey"])
	d.Set("ocspresponder", data["ocspresponder"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteSslcertkey_sslocspresponder_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcertkey_sslocspresponder_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ocspresponder := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ocspresponder:%s", ocspresponder))

	err := client.DeleteResourceWithArgs(service.Sslcertkey_sslocspresponder_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
