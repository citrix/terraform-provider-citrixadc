package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/router"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"strings"
)

func resourceCitrixAdcRouterdynamicrouting() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        applyRouterdynamicroutingFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"commandlines": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func applyRouterdynamicroutingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(routerdynamicroutingName)

	return nil
}
