package citrixadc

import (
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcHafailover() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createHafailoverFunc,
		Read:          readHafailoverFunc,
		Delete:        deleteHafailoverFunc,
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

func createHafailoverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createHafailoverFunc")
	client := meta.(*NetScalerNitroClient).client
	hafailoverName := resource.PrefixedUniqueId("tf-hafailover-")
	hafailover := ha.Hafailover{
		Force: d.Get("force").(bool),
	}

	curState, err := readHaNodeState(d, meta)
	if err != nil {
		return err
	}

	if curState != d.Get("state").(string) {
		err := client.ActOnResource("hafailover", &hafailover, "Force")
		if err != nil {
			return err
		}
	}

	d.SetId(hafailoverName)

	err = readHafailoverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this hafailover but we can't read it ?? %s", hafailoverName)
		return nil
	}

	return nil
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

func readHafailoverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readHafailoverFunc")

	state, err := readHaNodeState(d, meta)
	if err != nil {
		return err
	}

	d.Set("state", state)

	return nil

}

func deleteHafailoverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteHafailoverFunc")

	d.SetId("")

	return nil
}
