package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func resourceCitrixAdcInterfacepair() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createInterfacepairFunc,
		ReadContext:   readInterfacepairFunc,
		DeleteContext: deleteInterfacepairFunc,
		Schema: map[string]*schema.Schema{
			"interface_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ifnum": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createInterfacepairFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createInterfacepairFunc")
	client := meta.(*NetScalerNitroClient).client
	interfacepairName := strconv.Itoa(d.Get("interface_id").(int))

	interfacepair := network.Interfacepair{
		Ifnum: toStringList(d.Get("ifnum").([]interface{})),
	}

	if raw := d.GetRawConfig().GetAttr("interface_id"); !raw.IsNull() {
		interfacepair.Id = intPtr(d.Get("interface_id").(int))
	}

	_, err := client.AddResource(service.Interfacepair.Type(), interfacepairName, &interfacepair)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(interfacepairName)

	return readInterfacepairFunc(ctx, d, meta)
}

func readInterfacepairFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readInterfacepairFunc")
	client := meta.(*NetScalerNitroClient).client
	interfacepairName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading interfacepair state %s", interfacepairName)
	data, err := client.FindResource(service.Interfacepair.Type(), interfacepairName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing interfacepair state %s", interfacepairName)
		d.SetId("")
		return nil
	}
	d.Set("id", data["id"])
	//d.Set("ifnum", data["ifnum"])

	return nil

}

func deleteInterfacepairFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteInterfacepairFunc")
	client := meta.(*NetScalerNitroClient).client
	interfacepairName := d.Id()
	err := client.DeleteResource(service.Interfacepair.Type(), interfacepairName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
