package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsweblogparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsweblogparamFunc,
		ReadContext:   readNsweblogparamFunc,
		UpdateContext: updateNsweblogparamFunc,
		DeleteContext: deleteNsweblogparamFunc,
		Schema: map[string]*schema.Schema{
			"buffersizemb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"customreqhdrs": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"customrsphdrs": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsweblogparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsweblogparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsweblogparamName string
	// there is no primary key in nsweblogparam resource. Hence generate one for terraform state maintenance
	nsweblogparamName = resource.PrefixedUniqueId("tf-nsweblogparam-")
	nsweblogparam := ns.Nsweblogparam{
		Buffersizemb:  d.Get("buffersizemb").(int),
		Customreqhdrs: toStringList(d.Get("customreqhdrs").([]interface{})),
		Customrsphdrs: toStringList(d.Get("customrsphdrs").([]interface{})),
	}

	err := client.UpdateUnnamedResource(service.Nsweblogparam.Type(), &nsweblogparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsweblogparamName)

	return readNsweblogparamFunc(ctx, d, meta)
}

func readNsweblogparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsweblogparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsweblogparam state")
	data, err := client.FindResource(service.Nsweblogparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsweblogparam state")
		d.SetId("")
		return nil
	}
	value, _ := strconv.Atoi(data["buffersizemb"].(string))
	d.Set("buffersizemb", value)
	d.Set("customreqhdrs", data["customreqhdrs"])
	d.Set("customrsphdrs", data["customrsphdrs"])

	return nil

}

func updateNsweblogparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsweblogparamFunc")
	client := meta.(*NetScalerNitroClient).client

	nsweblogparam := ns.Nsweblogparam{}
	hasChange := false
	if d.HasChange("buffersizemb") {
		log.Printf("[DEBUG]  citrixadc-provider: Buffersizemb has changed for nsweblogparam , starting update")
		nsweblogparam.Buffersizemb = d.Get("buffersizemb").(int)
		hasChange = true
	}
	if d.HasChange("customreqhdrs") {
		log.Printf("[DEBUG]  citrixadc-provider: Customreqhdrs has changed for nsweblogparam , starting update")
		nsweblogparam.Customreqhdrs = toStringList(d.Get("customreqhdrs").([]interface{}))
		hasChange = true
	}
	if d.HasChange("customrsphdrs") {
		log.Printf("[DEBUG]  citrixadc-provider: Customrsphdrs has changed for nsweblogparam , starting update")
		nsweblogparam.Customrsphdrs = toStringList(d.Get("customrsphdrs").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsweblogparam.Type(), &nsweblogparam)
		if err != nil {
			return diag.Errorf("Error updating nsweblogparam")
		}
	}
	return readNsweblogparamFunc(ctx, d, meta)
}

func deleteNsweblogparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsweblogparamFunc")

	d.SetId("")

	return nil
}
