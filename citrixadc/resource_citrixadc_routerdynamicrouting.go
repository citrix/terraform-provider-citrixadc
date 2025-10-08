package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/router"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRouterdynamicrouting() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: applyRouterdynamicroutingFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"commandlines": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func applyRouterdynamicroutingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRouterdynamicroutingFunc")
	client := meta.(*NetScalerNitroClient).client
	routerdynamicroutingName := resource.PrefixedUniqueId("tf-routerdynamicrouting-")

	lines := d.Get("commandlines").([]interface{})
	stringArray := make([]string, 0, len(lines))
	for _, line := range lines {
		stringArray = append(stringArray, line.(string))
	}

	cmdString := strings.Join(stringArray, "\n")

	routerdynamicrouting := router.Routerdynamicrouting{
		Commandstring: cmdString,
		Nodeid:        d.Get("nodeid").(int),
	}

	err := client.ActOnResource("routerdynamicrouting", &routerdynamicrouting, "apply")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(routerdynamicroutingName)

	return nil
}
