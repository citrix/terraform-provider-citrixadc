package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNspartition() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNspartitionFunc,
		ReadContext:   readNspartitionFunc,
		UpdateContext: updateNspartitionFunc,
		DeleteContext: deleteNspartitionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"partitionname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxconn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxmemlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"partitionmac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"save": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNspartitionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Get("partitionname").(string)

	nspartition := make(map[string]interface{})
	if v, ok := d.GetOk("partitionname"); ok {
		nspartition["partitionname"] = v.(string)
	}
	if v, ok := d.GetOkExists("maxbandwidth"); ok {
		nspartition["maxbandwidth"] = v.(int)
	}
	if v, ok := d.GetOkExists("minbandwidth"); ok {
		nspartition["minbandwidth"] = v.(int)
	}
	if v, ok := d.GetOkExists("maxconn"); ok {
		nspartition["maxconn"] = v.(int)
	}
	if v, ok := d.GetOkExists("maxmemlimit"); ok {
		nspartition["maxmemlimit"] = v.(int)
	}
	if v, ok := d.GetOk("partitionmac"); ok {
		nspartition["partitionmac"] = v.(string)
	}
	if v, ok := d.GetOk("force"); ok {
		nspartition["force"] = v.(bool)
	}
	if v, ok := d.GetOk("save"); ok {
		nspartition["save"] = v.(bool)
	}
	_, err := client.AddResource(service.Nspartition.Type(), nspartitionName, &nspartition)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nspartitionName)

	return readNspartitionFunc(ctx, d, meta)
}

func readNspartitionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nspartition state %s", nspartitionName)
	data, err := client.FindResource(service.Nspartition.Type(), nspartitionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nspartition state %s", nspartitionName)
		d.SetId("")
		return nil
	}
	d.Set("force", data["force"])
	setToInt("maxbandwidth", d, data["maxbandwidth"])
	setToInt("maxconn", d, data["maxconn"])
	setToInt("maxmemlimit", d, data["maxmemlimit"])
	// setToInt("minbandwidth", d, data["minbandwidth"])
	d.Set("partitionmac", data["partitionmac"])
	d.Set("partitionname", data["partitionname"])
	d.Set("save", data["save"])

	return nil

}

func updateNspartitionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Get("partitionname").(string)

	nspartition := make(map[string]interface{})
	nspartition["partitionname"] = d.Get("partitionname").(string)
	hasChange := false
	if d.HasChange("force") {
		log.Printf("[DEBUG]  citrixadc-provider: Force has changed for nspartition %s, starting update", nspartitionName)
		nspartition["force"] = d.Get("force").(bool)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for nspartition %s, starting update", nspartitionName)
		nspartition["maxbandwidth"] = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxconn has changed for nspartition %s, starting update", nspartitionName)
		nspartition["maxconn"] = d.Get("maxconn").(int)
		hasChange = true
	}
	if d.HasChange("maxmemlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxmemlimit has changed for nspartition %s, starting update", nspartitionName)
		nspartition["maxmemlimit"] = d.Get("maxmemlimit").(int)
		hasChange = true
	}
	if d.HasChange("minbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Minbandwidth has changed for nspartition %s, starting update", nspartitionName)
		nspartition["minbandwidth"] = d.Get("minbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("partitionmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Partitionmac has changed for nspartition %s, starting update", nspartitionName)
		nspartition["partitionmac"] = d.Get("partitionmac").(string)
		hasChange = true
	}
	if d.HasChange("save") {
		log.Printf("[DEBUG]  citrixadc-provider: Save has changed for nspartition %s, starting update", nspartitionName)
		nspartition["save"] = d.Get("save").(bool)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nspartition.Type(), nspartitionName, &nspartition)
		if err != nil {
			return diag.Errorf("Error updating nspartition %s", nspartitionName)
		}
	}
	return readNspartitionFunc(ctx, d, meta)
}

func deleteNspartitionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Id()
	err := client.DeleteResource(service.Nspartition.Type(), nspartitionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
