package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ntp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNtpsync() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNtpsyncFunc,
		UpdateContext: updateNtpsyncFunc,
		ReadContext:   readNtpsyncFunc,
		DeleteContext: deleteNtpsyncFunc,
		Schema: map[string]*schema.Schema{
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createNtpsyncFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNtpsyncFunc")
	ntpsyncName := resource.PrefixedUniqueId("tf-ntpsync-")
	client := meta.(*NetScalerNitroClient).client

	err := doNtpsyncChange(d, client)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ntpsyncName)

	return readNtpsyncFunc(ctx, d, meta)
}

func readNtpsyncFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNtpsyncFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ntpsync state")
	data, err := client.FindResource(service.Ntpsync.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ntpsync state")
		d.SetId("")
		return nil
	}
	d.Set("state", data["state"].(string))

	return nil

}

func updateNtpsyncFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNtpSyncFunc")
	client := meta.(*NetScalerNitroClient).client

	hasChange := false
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: state has changed for ntpsync, starting update")
		hasChange = true
	}

	if hasChange {
		err := doNtpsyncChange(d, client)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return readNtpsyncFunc(ctx, d, meta)
}
func doNtpsyncChange(d *schema.ResourceData, client *service.NitroClient) error {
	ntpsync := ntp.Ntpsync{}

	newstate := d.Get("state").(string)

	var err error
	// Enable action
	if newstate == "ENABLED" {
		err = client.ActOnResource(service.Ntpsync.Type(), &ntpsync, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err = client.ActOnResource(service.Ntpsync.Type(), &ntpsync, "disable")

	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}

	if err != nil {
		return err
	}
	return nil
}

func deleteNtpsyncFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNtpsyncFunc")
	// ntpsync does not support DELETE operation
	d.SetId("")

	return nil
}
