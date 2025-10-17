package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNtpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNtpparamFunc,
		ReadContext:   readNtpparamFunc,
		UpdateContext: updateNtpparamFunc,
		DeleteContext: deleteNtpparamFunc,
		Schema: map[string]*schema.Schema{
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autokeylogsec": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revokelogsec": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trustedkey": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func createNtpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNtpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpparamName := resource.PrefixedUniqueId("tf-ntpparam-")

	ntpparam := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok {
		ntpparam["authentication"] = v.(string)
	}
	if v, ok := d.GetOk("autokeylogsec"); ok {
		ntpparam["autokeylogsec"] = v.(int)
	}
	if v, ok := d.GetOk("revokelogsec"); ok {
		ntpparam["revokelogsec"] = v.(int)
	}
	if v, ok := d.GetOk("trustedkey"); ok {
		ntpparam["trustedkey"] = toIntegerList(v.([]interface{}))
	}

	err := client.UpdateUnnamedResource(service.Ntpparam.Type(), &ntpparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ntpparamName)

	return readNtpparamFunc(ctx, d, meta)
}

func readNtpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNtpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ntpparam state")
	data, err := client.FindResource(service.Ntpparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ntpparam state")
		d.SetId("")
		return nil
	}
	d.Set("authentication", data["authentication"])
	setToInt("autokeylogsec", d, data["autokeylogsec"])
	setToInt("revokelogsec", d, data["revokelogsec"])
	// Convert trustedkey from []string to []int before setting
	if trustedKeys, ok := data["trustedkey"]; ok && trustedKeys != nil {
		d.Set("trustedkey", stringListToIntList(trustedKeys.([]interface{})))
	}

	return nil

}

func updateNtpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNtpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	ntpparam := make(map[string]interface{})
	hasChange := false
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for ntpparam, starting update")
		ntpparam["authentication"] = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("autokeylogsec") {
		log.Printf("[DEBUG]  citrixadc-provider: Autokeylogsec has changed for ntpparam, starting update")
		ntpparam["autokeylogsec"] = intPtr(d.Get("autokeylogsec").(int))
		hasChange = true
	}
	if d.HasChange("revokelogsec") {
		log.Printf("[DEBUG]  citrixadc-provider: Revokelogsec has changed for ntpparam, starting update")
		ntpparam["revokelogsec"] = intPtr(d.Get("revokelogsec").(int))
		hasChange = true
	}
	if d.HasChange("trustedkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Trustedkey has changed for ntpparam, starting update")
		ntpparam["trustedkey"] = toIntegerList(d.Get("trustedkey").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ntpparam.Type(), &ntpparam)
		if err != nil {
			return diag.Errorf("Error updating ntpparam")
		}
	}
	return readNtpparamFunc(ctx, d, meta)
}

func deleteNtpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNtpparamFunc")
	// ntpparam does not support DELETE operation
	d.SetId("")

	return nil
}
