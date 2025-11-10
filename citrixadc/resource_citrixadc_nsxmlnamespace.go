package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsxmlnamespace() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsxmlnamespaceFunc,
		ReadContext:   readNsxmlnamespaceFunc,
		UpdateContext: updateNsxmlnamespaceFunc,
		DeleteContext: deleteNsxmlnamespaceFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsxmlnamespaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Get("prefix").(string)
	nsxmlnamespace := ns.Nsxmlnamespace{
		Description: d.Get("description").(string),
		Namespace:   d.Get("namespace").(string),
		Prefix:      d.Get("prefix").(string),
	}

	_, err := client.AddResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName, &nsxmlnamespace)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsxmlnamespaceName)

	return readNsxmlnamespaceFunc(ctx, d, meta)
}

func readNsxmlnamespaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsxmlnamespace state %s", nsxmlnamespaceName)
	data, err := client.FindResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsxmlnamespace state %s", nsxmlnamespaceName)
		d.SetId("")
		return nil
	}
	d.Set("description", data["description"])
	//d.Set("namespace", data["namespace"])
	d.Set("prefix", data["prefix"])

	return nil

}

func updateNsxmlnamespaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Get("prefix").(string)

	nsxmlnamespace := ns.Nsxmlnamespace{
		Prefix: d.Get("prefix").(string),
	}
	hasChange := false
	if d.HasChange("description") {
		log.Printf("[DEBUG]  citrixadc-provider: Description has changed for nsxmlnamespace %s, starting update", nsxmlnamespaceName)
		nsxmlnamespace.Description = d.Get("description").(string)
		hasChange = true
	}
	if d.HasChange("namespace") {
		log.Printf("[DEBUG]  citrixadc-provider: Namespace has changed for nsxmlnamespace %s, starting update", nsxmlnamespaceName)
		nsxmlnamespace.Namespace = d.Get("namespace").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName, &nsxmlnamespace)
		if err != nil {
			return diag.Errorf("Error updating nsxmlnamespace %s", nsxmlnamespaceName)
		}
	}
	return readNsxmlnamespaceFunc(ctx, d, meta)
}

func deleteNsxmlnamespaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsxmlnamespaceFunc")
	client := meta.(*NetScalerNitroClient).client
	nsxmlnamespaceName := d.Id()
	err := client.DeleteResource(service.Nsxmlnamespace.Type(), nsxmlnamespaceName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
