package citrixadc

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"log"
	"strconv"
)

func dataSourceCitrixAdcNsversion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCitrixAdcNsversionRead,
		Schema: map[string]*schema.Schema{
			"installedversion": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceCitrixAdcNsversionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcNsversionRead")
	client := meta.(*NetScalerNitroClient).client
	findParams := service.FindParams{
		ResourceType: "nsversion",
	}
	if val, exists := d.GetOkExists("installedversion"); exists {
		argsMap := make(map[string]string)
		argsMap["installedversion"] = fmt.Sprintf("%v", val.(bool))
		findParams.ArgsMap = argsMap
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[ERROR] citrixadc-provider: Error during read %s", err)
		return diag.FromErr(err)
	}
	if len(dataArr) != 1 {
		return diag.Errorf("Unexpected length of nsversion response: %v", dataArr)
	}

	data := dataArr[0]

	d.SetId(data["version"].(string))
	d.Set("version", data["version"].(string))
	if val, ok := data["mode"]; ok {
		intVal, err := strconv.Atoi(val.(string))
		if err != nil {
			return diag.Errorf("Error during Atoi for mode: %s", err.Error())
		}
		d.Set("mode", intVal)
	}

	return diags

}
