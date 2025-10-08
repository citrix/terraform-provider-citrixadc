package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAppfwlearningsettings() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwlearningsettingsFunc,
		ReadContext:   readAppfwlearningsettingsFunc,
		UpdateContext: updateAppfwlearningsettingsFunc,
		DeleteContext: deleteAppfwlearningsettingsFunc, // Thought appfwlearningsettings resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"profilename": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"contenttypeautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"contenttypeminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"contenttypepercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookieconsistencyautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookieconsistencyminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookieconsistencypercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creditcardnumberminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creditcardnumberpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csrftagautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csrftagminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csrftagpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldconsistencyautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldconsistencyminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldconsistencypercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldformatautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldformatminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldformatpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"starturlautodeploygraceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"starturlminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"starturlpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlattachmentminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlattachmentpercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlwsiminthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlwsipercentthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwlearningsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwlearningsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwlearningsettingsName := d.Get("profilename").(string)

	appfwlearningsettings := appfw.Appfwlearningsettings{
		Contenttypeautodeploygraceperiod:        d.Get("contenttypeautodeploygraceperiod").(int),
		Contenttypeminthreshold:                 d.Get("contenttypeminthreshold").(int),
		Contenttypepercentthreshold:             d.Get("contenttypepercentthreshold").(int),
		Cookieconsistencyautodeploygraceperiod:  d.Get("cookieconsistencyautodeploygraceperiod").(int),
		Cookieconsistencyminthreshold:           d.Get("cookieconsistencyminthreshold").(int),
		Cookieconsistencypercentthreshold:       d.Get("cookieconsistencypercentthreshold").(int),
		Creditcardnumberminthreshold:            d.Get("creditcardnumberminthreshold").(int),
		Creditcardnumberpercentthreshold:        d.Get("creditcardnumberpercentthreshold").(int),
		Crosssitescriptingautodeploygraceperiod: d.Get("crosssitescriptingautodeploygraceperiod").(int),
		Crosssitescriptingminthreshold:          d.Get("crosssitescriptingminthreshold").(int),
		Crosssitescriptingpercentthreshold:      d.Get("crosssitescriptingpercentthreshold").(int),
		Csrftagautodeploygraceperiod:            d.Get("csrftagautodeploygraceperiod").(int),
		Csrftagminthreshold:                     d.Get("csrftagminthreshold").(int),
		Csrftagpercentthreshold:                 d.Get("csrftagpercentthreshold").(int),
		Fieldconsistencyautodeploygraceperiod:   d.Get("fieldconsistencyautodeploygraceperiod").(int),
		Fieldconsistencyminthreshold:            d.Get("fieldconsistencyminthreshold").(int),
		Fieldconsistencypercentthreshold:        d.Get("fieldconsistencypercentthreshold").(int),
		Fieldformatautodeploygraceperiod:        d.Get("fieldformatautodeploygraceperiod").(int),
		Fieldformatminthreshold:                 d.Get("fieldformatminthreshold").(int),
		Fieldformatpercentthreshold:             d.Get("fieldformatpercentthreshold").(int),
		Profilename:                             d.Get("profilename").(string),
		Sqlinjectionautodeploygraceperiod:       d.Get("sqlinjectionautodeploygraceperiod").(int),
		Sqlinjectionminthreshold:                d.Get("sqlinjectionminthreshold").(int),
		Sqlinjectionpercentthreshold:            d.Get("sqlinjectionpercentthreshold").(int),
		Starturlautodeploygraceperiod:           d.Get("starturlautodeploygraceperiod").(int),
		Starturlminthreshold:                    d.Get("starturlminthreshold").(int),
		Xmlattachmentminthreshold:               d.Get("xmlattachmentminthreshold").(int),
		Xmlattachmentpercentthreshold:           d.Get("xmlattachmentpercentthreshold").(int),
		Xmlwsiminthreshold:                      d.Get("xmlwsiminthreshold").(int),
		Xmlwsipercentthreshold:                  d.Get("xmlwsipercentthreshold").(int),
		Starturlpercentthreshold:                d.Get("starturlpercentthreshold").(int),
	}

	err := client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwlearningsettingsName)

	return readAppfwlearningsettingsFunc(ctx, d, meta)
}

func readAppfwlearningsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwlearningsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwlearningsettingsName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwlearningsettings state %s", appfwlearningsettingsName)
	data, err := client.FindResource(service.Appfwlearningsettings.Type(), appfwlearningsettingsName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwlearningsettings state %s", appfwlearningsettingsName)
		d.SetId("")
		return nil
	}
	d.Set("profilename", data["profilename"])
	setToInt("contenttypeautodeploygraceperiod", d, data["contenttypeautodeploygraceperiod"])
	setToInt("contenttypeminthreshold", d, data["contenttypeminthreshold"])
	setToInt("contenttypepercentthreshold", d, data["contenttypepercentthreshold"])
	setToInt("cookieconsistencyautodeploygraceperiod", d, data["cookieconsistencyautodeploygraceperiod"])
	setToInt("cookieconsistencyminthreshold", d, data["cookieconsistencyminthreshold"])
	setToInt("cookieconsistencypercentthreshold", d, data["cookieconsistencypercentthreshold"])
	setToInt("creditcardnumberminthreshold", d, data["creditcardnumberminthreshold"])
	setToInt("creditcardnumberpercentthreshold", d, data["creditcardnumberpercentthreshold"])
	setToInt("crosssitescriptingautodeploygraceperiod", d, data["crosssitescriptingautodeploygraceperiod"])
	setToInt("crosssitescriptingminthreshold", d, data["crosssitescriptingminthreshold"])
	setToInt("crosssitescriptingpercentthreshold", d, data["crosssitescriptingpercentthreshold"])
	setToInt("csrftagautodeploygraceperiod", d, data["csrftagautodeploygraceperiod"])
	setToInt("csrftagminthreshold", d, data["csrftagminthreshold"])
	setToInt("csrftagpercentthreshold", d, data["csrftagpercentthreshold"])
	setToInt("fieldconsistencyautodeploygraceperiod", d, data["fieldconsistencyautodeploygraceperiod"])
	setToInt("fieldconsistencyminthreshold", d, data["fieldconsistencyminthreshold"])
	setToInt("fieldconsistencypercentthreshold", d, data["fieldconsistencypercentthreshold"])
	setToInt("fieldformatautodeploygraceperiod", d, data["fieldformatautodeploygraceperiod"])
	setToInt("fieldformatminthreshold", d, data["fieldformatminthreshold"])
	setToInt("fieldformatpercentthreshold", d, data["fieldformatpercentthreshold"])
	setToInt("sqlinjectionautodeploygraceperiod", d, data["sqlinjectionautodeploygraceperiod"])
	setToInt("sqlinjectionminthreshold", d, data["sqlinjectionminthreshold"])
	setToInt("sqlinjectionpercentthreshold", d, data["sqlinjectionpercentthreshold"])
	setToInt("starturlautodeploygraceperiod", d, data["starturlautodeploygraceperiod"])
	setToInt("starturlminthreshold", d, data["starturlminthreshold"])
	setToInt("starturlpercentthreshold", d, data["starturlpercentthreshold"])
	setToInt("xmlattachmentminthreshold", d, data["xmlattachmentminthreshold"])
	setToInt("xmlattachmentpercentthreshold", d, data["xmlattachmentpercentthreshold"])
	setToInt("xmlwsiminthreshold", d, data["xmlwsiminthreshold"])
	setToInt("xmlwsipercentthreshold", d, data["xmlwsipercentthreshold"])

	return nil

}

func updateAppfwlearningsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwlearningsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwlearningsettingsName := d.Get("profilename").(string)

	appfwlearningsettings := appfw.Appfwlearningsettings{
		Profilename: d.Get("profilename").(string),
	}
	hasChange := false
	if d.HasChange("contenttypeautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Contenttypeautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Contenttypeautodeploygraceperiod = d.Get("contenttypeautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("contenttypeminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Contenttypeminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Contenttypeminthreshold = d.Get("contenttypeminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("contenttypepercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Contenttypepercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Contenttypepercentthreshold = d.Get("contenttypepercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("cookieconsistencyautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencyautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Cookieconsistencyautodeploygraceperiod = d.Get("cookieconsistencyautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("cookieconsistencyminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencyminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Cookieconsistencyminthreshold = d.Get("cookieconsistencyminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("cookieconsistencypercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencypercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Cookieconsistencypercentthreshold = d.Get("cookieconsistencypercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("creditcardnumberminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardnumberminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Creditcardnumberminthreshold = d.Get("creditcardnumberminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("creditcardnumberpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardnumberpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Creditcardnumberpercentthreshold = d.Get("creditcardnumberpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("crosssitescriptingautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Crosssitescriptingautodeploygraceperiod = d.Get("crosssitescriptingautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("crosssitescriptingminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Crosssitescriptingminthreshold = d.Get("crosssitescriptingminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("crosssitescriptingpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Crosssitescriptingpercentthreshold = d.Get("crosssitescriptingpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("csrftagautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Csrftagautodeploygraceperiod = d.Get("csrftagautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("csrftagminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Csrftagminthreshold = d.Get("csrftagminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("csrftagpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Csrftagpercentthreshold = d.Get("csrftagpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("fieldconsistencyautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencyautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldconsistencyautodeploygraceperiod = d.Get("fieldconsistencyautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("fieldconsistencyminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencyminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldconsistencyminthreshold = d.Get("fieldconsistencyminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("fieldconsistencypercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencypercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldconsistencypercentthreshold = d.Get("fieldconsistencypercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("fieldformatautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformatautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldformatautodeploygraceperiod = d.Get("fieldformatautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("fieldformatminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformatminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldformatminthreshold = d.Get("fieldformatminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("fieldformatpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformatpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldformatpercentthreshold = d.Get("fieldformatpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("sqlinjectionautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Sqlinjectionautodeploygraceperiod = d.Get("sqlinjectionautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("sqlinjectionminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Sqlinjectionminthreshold = d.Get("sqlinjectionminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("sqlinjectionpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Sqlinjectionpercentthreshold = d.Get("sqlinjectionpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("starturlautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Starturlautodeploygraceperiod = d.Get("starturlautodeploygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("starturlminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Starturlminthreshold = d.Get("starturlminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("starturlpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Starturlpercentthreshold = d.Get("starturlpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("xmlattachmentminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlattachmentminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlattachmentminthreshold = d.Get("xmlattachmentminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("xmlattachmentpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlattachmentpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlattachmentpercentthreshold = d.Get("xmlattachmentpercentthreshold").(int)
		hasChange = true
	}
	if d.HasChange("xmlwsiminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlwsiminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlwsiminthreshold = d.Get("xmlwsiminthreshold").(int)
		hasChange = true
	}
	if d.HasChange("xmlwsipercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlwsipercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlwsipercentthreshold = d.Get("xmlwsipercentthreshold").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
		if err != nil {
			return diag.Errorf("Error updating appfwlearningsettings %s", appfwlearningsettingsName)
		}
	}
	return readAppfwlearningsettingsFunc(ctx, d, meta)
}

func deleteAppfwlearningsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwlearningsettingsFunc")

	d.SetId("")

	return nil
}
