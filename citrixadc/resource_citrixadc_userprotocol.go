package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/user"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcUserprotocol() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createUserprotocolFunc,
		ReadContext:   readUserprotocolFunc,
		UpdateContext: updateUserprotocolFunc,
		DeleteContext: deleteUserprotocolFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"extension": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transport": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createUserprotocolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Get("name").(string)
	userprotocol := user.Userprotocol{
		Comment:   d.Get("comment").(string),
		Extension: d.Get("extension").(string),
		Name:      d.Get("name").(string),
		Transport: d.Get("transport").(string),
	}

	_, err := client.AddResource("userprotocol", userprotocolName, &userprotocol)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(userprotocolName)

	return readUserprotocolFunc(ctx, d, meta)
}

func readUserprotocolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading userprotocol state %s", userprotocolName)
	data, err := client.FindResource("userprotocol", userprotocolName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing userprotocol state %s", userprotocolName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("extension", data["extension"])
	d.Set("transport", data["transport"])

	return nil

}

func updateUserprotocolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Get("name").(string)

	userprotocol := user.Userprotocol{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for userprotocol %s, starting update", userprotocolName)
		userprotocol.Comment = d.Get("comment").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("userprotocol", &userprotocol)
		if err != nil {
			return diag.Errorf("Error updating userprotocol %s", userprotocolName)
		}
	}
	return readUserprotocolFunc(ctx, d, meta)
}

func deleteUserprotocolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteUserprotocolFunc")
	client := meta.(*NetScalerNitroClient).client
	userprotocolName := d.Id()
	err := client.DeleteResource("userprotocol", userprotocolName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
