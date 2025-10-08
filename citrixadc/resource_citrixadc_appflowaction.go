package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppflowaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppflowactionFunc,
		ReadContext:   readAppflowactionFunc,
		UpdateContext: updateAppflowactionFunc,
		DeleteContext: deleteAppflowactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"botinsight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ciinsight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsidemeasurements": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collectors": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distributionalgorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metricslog": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pagetracking": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"securityinsight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transactionlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"videoanalytics": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webinsight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppflowactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowactionName := d.Get("name").(string)

	appflowaction := appflow.Appflowaction{
		Botinsight:             d.Get("botinsight").(string),
		Ciinsight:              d.Get("ciinsight").(string),
		Clientsidemeasurements: d.Get("clientsidemeasurements").(string),
		Collectors:             toStringList(d.Get("collectors").([]interface{})),
		Comment:                d.Get("comment").(string),
		Distributionalgorithm:  d.Get("distributionalgorithm").(string),
		Metricslog:             d.Get("metricslog").(bool),
		Name:                   d.Get("name").(string),
		Pagetracking:           d.Get("pagetracking").(string),
		Securityinsight:        d.Get("securityinsight").(string),
		Transactionlog:         d.Get("transactionlog").(string),
		Videoanalytics:         d.Get("videoanalytics").(string),
		Webinsight:             d.Get("webinsight").(string),
	}

	_, err := client.AddResource(service.Appflowaction.Type(), appflowactionName, &appflowaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appflowactionName)

	return readAppflowactionFunc(ctx, d, meta)
}

func readAppflowactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowaction state %s", appflowactionName)
	data, err := client.FindResource(service.Appflowaction.Type(), appflowactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowaction state %s", appflowactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("botinsight", data["botinsight"])
	d.Set("ciinsight", data["ciinsight"])
	d.Set("clientsidemeasurements", data["clientsidemeasurements"])
	d.Set("collectors", data["collectors"])
	d.Set("comment", data["comment"])
	d.Set("distributionalgorithm", data["distributionalgorithm"])
	d.Set("metricslog", data["metricslog"])
	d.Set("pagetracking", data["pagetracking"])
	d.Set("securityinsight", data["securityinsight"])
	d.Set("transactionlog", data["transactionlog"])
	d.Set("videoanalytics", data["videoanalytics"])
	d.Set("webinsight", data["webinsight"])

	return nil

}

func updateAppflowactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppflowactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowactionName := d.Get("name").(string)

	appflowaction := appflow.Appflowaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("botinsight") {
		log.Printf("[DEBUG]  citrixadc-provider: Botinsight has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Botinsight = d.Get("botinsight").(string)
		hasChange = true
	}
	if d.HasChange("ciinsight") {
		log.Printf("[DEBUG]  citrixadc-provider: Ciinsight has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Ciinsight = d.Get("ciinsight").(string)
		hasChange = true
	}
	if d.HasChange("clientsidemeasurements") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsidemeasurements has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Clientsidemeasurements = d.Get("clientsidemeasurements").(string)
		hasChange = true
	}
	if d.HasChange("collectors") {
		log.Printf("[DEBUG]  citrixadc-provider: Collectors has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Collectors = toStringList(d.Get("collectors").([]interface{}))
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("distributionalgorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Distributionalgorithm has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Distributionalgorithm = d.Get("distributionalgorithm").(string)
		hasChange = true
	}
	if d.HasChange("metricslog") {
		log.Printf("[DEBUG]  citrixadc-provider: Metricslog has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Metricslog = d.Get("metricslog").(bool)
		hasChange = true
	}
	if d.HasChange("pagetracking") {
		log.Printf("[DEBUG]  citrixadc-provider: Pagetracking has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Pagetracking = d.Get("pagetracking").(string)
		hasChange = true
	}
	if d.HasChange("securityinsight") {
		log.Printf("[DEBUG]  citrixadc-provider: Securityinsight has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Securityinsight = d.Get("securityinsight").(string)
		hasChange = true
	}
	if d.HasChange("transactionlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Transactionlog has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Transactionlog = d.Get("transactionlog").(string)
		hasChange = true
	}
	if d.HasChange("videoanalytics") {
		log.Printf("[DEBUG]  citrixadc-provider: Videoanalytics has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Videoanalytics = d.Get("videoanalytics").(string)
		hasChange = true
	}
	if d.HasChange("webinsight") {
		log.Printf("[DEBUG]  citrixadc-provider: Webinsight has changed for appflowaction %s, starting update", appflowactionName)
		appflowaction.Webinsight = d.Get("webinsight").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appflowaction.Type(), appflowactionName, &appflowaction)
		if err != nil {
			return diag.Errorf("Error updating appflowaction %s", appflowactionName)
		}
	}
	return readAppflowactionFunc(ctx, d, meta)
}

func deleteAppflowactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowactionName := d.Id()
	err := client.DeleteResource(service.Appflowaction.Type(), appflowactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
