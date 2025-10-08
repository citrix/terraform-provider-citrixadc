package citrixadc

import (
	"context"
	"log"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsconfigUpdate() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsconfigUpdateFunc,
		ReadContext:   readNsconfigUpdateFunc,
		UpdateContext: updateNsconfigUpdateFunc,
		Delete:        schema.Noop, // Should we call `unset ns config` here?
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsvlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ifnum": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tagged": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsconfigUpdateFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsconfigUpdateFunc")
	client := meta.(*NetScalerNitroClient).client
	nsconfigName := resource.PrefixedUniqueId("tf-nsconfig-update")

	nsconfig := ns.Nsconfig{}
	nsconfig.Ipaddress = d.Get("ipaddress").(string)
	nsconfig.Netmask = d.Get("netmask").(string)
	nsconfig.Nsvlan = d.Get("nsvlan").(int)
	nsconfig.Ifnum = toStringList(getIfnumValue(d))
	nsconfig.Tagged = d.Get("tagged").(string)

	err := client.UpdateUnnamedResource(service.Nsconfig.Type(), &nsconfig)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsconfigName)

	return readNsconfigUpdateFunc(ctx, d, meta)
}

func readNsconfigUpdateFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsconfigFunc")
	client := meta.(*NetScalerNitroClient).client
	nsconfigName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsconfig state %s", nsconfigName)
	data, err := client.FindResource(service.Nsconfig.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsconfig state %s", nsconfigName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("netmask", data["netmask"])
	setToInt("nsvlan", d, data["nsvlan"])
	d.Set("ifnum", data["ifnum"])

	d.Set("tagged", data["tagged"])

	return nil

}

func updateNsconfigUpdateFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsconfigUpdateFunc")
	client := meta.(*NetScalerNitroClient).client

	nsconfigName := d.Id()

	nsconfig := ns.Nsconfig{}
	hasIPChanged := false
	hasNetmaskChanged := false
	hasNsvlanChanged := false
	hasIfnumChanged := false
	hasTaggedChanged := false
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for nsconfig %s, starting update", nsconfigName)
		hasIPChanged = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for nsconfig %s, starting update", nsconfigName)
		hasNetmaskChanged = true
	}
	if d.HasChange("nsvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Nsvlan has changed for nsconfig %s, starting update", nsconfigName)
		hasNsvlanChanged = true
	}
	if d.HasChange("ifnum") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifnum has changed for nsconfig %s, starting update", nsconfigName)
		hasIfnumChanged = true
	}
	if d.HasChange("tagged") {
		log.Printf("[DEBUG]  citrixadc-provider: Tagged has changed for nsconfig %s, starting update", nsconfigName)
		hasTaggedChanged = true
	}

	if hasIPChanged || hasNetmaskChanged {
		nsconfig.Ipaddress = d.Get("ipaddress").(string)
		nsconfig.Netmask = d.Get("netmask").(string)
	}
	if hasNsvlanChanged || hasIfnumChanged || hasTaggedChanged {
		nsconfig.Nsvlan = d.Get("nsvlan").(int)
		nsconfig.Ifnum = toStringList(getIfnumValue(d))
		nsconfig.Tagged = d.Get("tagged").(string)
	}

	if hasIPChanged || hasNetmaskChanged || hasNsvlanChanged || hasIfnumChanged || hasTaggedChanged {
		err := client.UpdateUnnamedResource(service.Nsconfig.Type(), &nsconfig)
		if err != nil {
			return diag.Errorf("Error updating nsconfig %s", nsconfigName)
		}
	}
	return readNsconfigUpdateFunc(ctx, d, meta)
}

func getIfnumValue(d *schema.ResourceData) []interface{} {
	if val, ok := d.GetOk("ifnum"); ok {
		return val.(*schema.Set).List()
	} else {
		return make([]interface{}, 0, 0)
	}
}
