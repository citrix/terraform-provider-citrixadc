package citrixadc

import (
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSnmpcommunity() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpcommunityFunc,
		Read:          readSnmpcommunityFunc,
		Delete:        deleteSnmpcommunityFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"communityname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permissions": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validatePermissions,
				StateFunc: func(v interface{}) string {
					return strings.ToUpper(v.(string))
				},
			},
		},
	}
}

func createSnmpcommunityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpcommunityFunc")
	client := meta.(*NetScalerNitroClient).client
	communityname := d.Get("communityname").(string)
	snmpcommunity := snmp.Snmpcommunity{
		Communityname: d.Get("communityname").(string),
		Permissions:   d.Get("permissions").(string),
	}

	_, err := client.AddResource(service.Snmpcommunity.Type(), communityname, &snmpcommunity)
	if err != nil {
		return err
	}

	d.SetId(communityname)

	err = readSnmpcommunityFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpcommunity but we can't read it ?? %s", communityname)
		return nil
	}
	return nil
}

func readSnmpcommunityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpcommunityFunc")
	client := meta.(*NetScalerNitroClient).client
	communityname := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpcommunity state %s", communityname)
	data, err := client.FindResource(service.Snmpcommunity.Type(), communityname)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpcommunity state %s", communityname)
		d.SetId("")
		return nil
	}
	d.Set("communityname", data["communityname"])
	d.Set("permissions", data["permissions"])

	return nil

}

func deleteSnmpcommunityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpcommunityFunc")
	client := meta.(*NetScalerNitroClient).client
	communityname := d.Id()
	err := client.DeleteResource(service.Snmpcommunity.Type(), communityname)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func validatePermissions(v interface{}, k string) (ws []string, errors []error) {
	allowedValues := []string{"GET", "GET_NEXT", "GET_BULK", "SET", "ALL"}

	value := v.(string)

	for _, allowed := range allowedValues {
		if strings.EqualFold(value, allowed) {
			return
		}
	}

	errors = append(errors, fmt.Errorf("%q must be one of %q", k, allowedValues))
	return
}
