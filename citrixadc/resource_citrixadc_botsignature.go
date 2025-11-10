package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcBotsignature() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotsignatureFunc,
		ReadContext:   readBotsignatureFunc,
		DeleteContext: deleteBotsignatureFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"src": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotsignatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotsignatureFunc")
	client := meta.(*NetScalerNitroClient).client

	botsignatureName := d.Get("name").(string)

	botsignature := bot.Botsignature{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource("botsignature", &botsignature, "Import")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(botsignatureName)

	return readBotsignatureFunc(ctx, d, meta)
}

func readBotsignatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotsignatureFunc")
	client := meta.(*NetScalerNitroClient).client
	botsignatureName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading botsignature state %s", botsignatureName)
	data, err := client.FindResource("botsignature", botsignatureName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing botsignature state %s", botsignatureName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteBotsignatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotsignatureFunc")
	client := meta.(*NetScalerNitroClient).client
	botsignatureName := d.Id()
	err := client.DeleteResource("botsignature", botsignatureName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
