package citrixadc

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcHafailover() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createHafailoverFunc,
		ReadContext:   readHafailoverFunc,
		DeleteContext: deleteHafailoverFunc,
		Schema: map[string]*schema.Schema{
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createHafailoverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createHafailoverFunc")
	client := meta.(*NetScalerNitroClient).client
	hafailoverName := resource.PrefixedUniqueId("tf-hafailover-")
	hafailover := ha.Hafailover{
		Force: d.Get("force").(bool),
	}

	curState, err := readHaNodeState(d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if curState != d.Get("state").(string) {
		err := client.ActOnResource("hafailover", &hafailover, "Force")
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(hafailoverName)

	return readHafailoverFunc(ctx, d, meta)
}

func readHaNodeState(d *schema.ResourceData, meta interface{}) (string, error) {
	log.Printf("[DEBUG] citrixadc-provider:  In readHafailoverFunc")
	client := meta.(*NetScalerNitroClient).client
	hafailoverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading hafailover state %s", hafailoverName)
	findParams := service.FindParams{
		ResourceType:             "hanode",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return "", err
	}
	log.Printf("[DEBUG] citrixadc-provider: ha node dataArr %v", dataArr)

	found := false
	var retval string
	for _, v := range dataArr {
		if v["ipaddress"] == d.Get("ipaddress").(string) {
			found = true
			retval = v["state"].(string)
		}
	}

	if !found {
		return "", fmt.Errorf("Cannot find ipaddress %v in ha node", d.Get("ipaddress").(string))
	}

	return retval, nil

}

func readHafailoverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readHafailoverFunc")

	state, err := readHaNodeState(d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("state", state)

	return nil

}

func deleteHafailoverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteHafailoverFunc")

	d.SetId("")

	return nil
}
