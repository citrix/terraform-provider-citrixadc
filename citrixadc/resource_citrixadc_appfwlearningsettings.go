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
		Profilename: d.Get("profilename").(string),
	}

	if raw := d.GetRawConfig().GetAttr("contenttypeautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Contenttypeautodeploygraceperiod = intPtr(d.Get("contenttypeautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("contenttypeminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Contenttypeminthreshold = intPtr(d.Get("contenttypeminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("contenttypepercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Contenttypepercentthreshold = intPtr(d.Get("contenttypepercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("cookieconsistencyautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Cookieconsistencyautodeploygraceperiod = intPtr(d.Get("cookieconsistencyautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("cookieconsistencyminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Cookieconsistencyminthreshold = intPtr(d.Get("cookieconsistencyminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("cookieconsistencypercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Cookieconsistencypercentthreshold = intPtr(d.Get("cookieconsistencypercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("creditcardnumberminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Creditcardnumberminthreshold = intPtr(d.Get("creditcardnumberminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("creditcardnumberpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Creditcardnumberpercentthreshold = intPtr(d.Get("creditcardnumberpercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("crosssitescriptingautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Crosssitescriptingautodeploygraceperiod = intPtr(d.Get("crosssitescriptingautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("crosssitescriptingminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Crosssitescriptingminthreshold = intPtr(d.Get("crosssitescriptingminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("crosssitescriptingpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Crosssitescriptingpercentthreshold = intPtr(d.Get("crosssitescriptingpercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("csrftagautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Csrftagautodeploygraceperiod = intPtr(d.Get("csrftagautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("csrftagminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Csrftagminthreshold = intPtr(d.Get("csrftagminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("csrftagpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Csrftagpercentthreshold = intPtr(d.Get("csrftagpercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("fieldconsistencyautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Fieldconsistencyautodeploygraceperiod = intPtr(d.Get("fieldconsistencyautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("fieldconsistencyminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Fieldconsistencyminthreshold = intPtr(d.Get("fieldconsistencyminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("fieldconsistencypercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Fieldconsistencypercentthreshold = intPtr(d.Get("fieldconsistencypercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("fieldformatautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Fieldformatautodeploygraceperiod = intPtr(d.Get("fieldformatautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("fieldformatminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Fieldformatminthreshold = intPtr(d.Get("fieldformatminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("fieldformatpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Fieldformatpercentthreshold = intPtr(d.Get("fieldformatpercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sqlinjectionautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Sqlinjectionautodeploygraceperiod = intPtr(d.Get("sqlinjectionautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sqlinjectionminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Sqlinjectionminthreshold = intPtr(d.Get("sqlinjectionminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sqlinjectionpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Sqlinjectionpercentthreshold = intPtr(d.Get("sqlinjectionpercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("starturlautodeploygraceperiod"); !raw.IsNull() {
		appfwlearningsettings.Starturlautodeploygraceperiod = intPtr(d.Get("starturlautodeploygraceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("starturlminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Starturlminthreshold = intPtr(d.Get("starturlminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("xmlattachmentminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Xmlattachmentminthreshold = intPtr(d.Get("xmlattachmentminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("xmlattachmentpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Xmlattachmentpercentthreshold = intPtr(d.Get("xmlattachmentpercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("xmlwsiminthreshold"); !raw.IsNull() {
		appfwlearningsettings.Xmlwsiminthreshold = intPtr(d.Get("xmlwsiminthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("xmlwsipercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Xmlwsipercentthreshold = intPtr(d.Get("xmlwsipercentthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("starturlpercentthreshold"); !raw.IsNull() {
		appfwlearningsettings.Starturlpercentthreshold = intPtr(d.Get("starturlpercentthreshold").(int))
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
		appfwlearningsettings.Contenttypeautodeploygraceperiod = intPtr(d.Get("contenttypeautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("contenttypeminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Contenttypeminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Contenttypeminthreshold = intPtr(d.Get("contenttypeminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("contenttypepercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Contenttypepercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Contenttypepercentthreshold = intPtr(d.Get("contenttypepercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("cookieconsistencyautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencyautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Cookieconsistencyautodeploygraceperiod = intPtr(d.Get("cookieconsistencyautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("cookieconsistencyminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencyminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Cookieconsistencyminthreshold = intPtr(d.Get("cookieconsistencyminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("cookieconsistencypercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencypercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Cookieconsistencypercentthreshold = intPtr(d.Get("cookieconsistencypercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("creditcardnumberminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardnumberminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Creditcardnumberminthreshold = intPtr(d.Get("creditcardnumberminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("creditcardnumberpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardnumberpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Creditcardnumberpercentthreshold = intPtr(d.Get("creditcardnumberpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("crosssitescriptingautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Crosssitescriptingautodeploygraceperiod = intPtr(d.Get("crosssitescriptingautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("crosssitescriptingminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Crosssitescriptingminthreshold = intPtr(d.Get("crosssitescriptingminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("crosssitescriptingpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Crosssitescriptingpercentthreshold = intPtr(d.Get("crosssitescriptingpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("csrftagautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Csrftagautodeploygraceperiod = intPtr(d.Get("csrftagautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("csrftagminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Csrftagminthreshold = intPtr(d.Get("csrftagminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("csrftagpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Csrftagpercentthreshold = intPtr(d.Get("csrftagpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("fieldconsistencyautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencyautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldconsistencyautodeploygraceperiod = intPtr(d.Get("fieldconsistencyautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("fieldconsistencyminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencyminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldconsistencyminthreshold = intPtr(d.Get("fieldconsistencyminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("fieldconsistencypercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencypercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldconsistencypercentthreshold = intPtr(d.Get("fieldconsistencypercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("fieldformatautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformatautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldformatautodeploygraceperiod = intPtr(d.Get("fieldformatautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("fieldformatminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformatminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldformatminthreshold = intPtr(d.Get("fieldformatminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("fieldformatpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformatpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Fieldformatpercentthreshold = intPtr(d.Get("fieldformatpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("sqlinjectionautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Sqlinjectionautodeploygraceperiod = intPtr(d.Get("sqlinjectionautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("sqlinjectionminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Sqlinjectionminthreshold = intPtr(d.Get("sqlinjectionminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("sqlinjectionpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Sqlinjectionpercentthreshold = intPtr(d.Get("sqlinjectionpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("starturlautodeploygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlautodeploygraceperiod has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Starturlautodeploygraceperiod = intPtr(d.Get("starturlautodeploygraceperiod").(int))
		hasChange = true
	}
	if d.HasChange("starturlminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Starturlminthreshold = intPtr(d.Get("starturlminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("starturlpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Starturlpercentthreshold = intPtr(d.Get("starturlpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("xmlattachmentminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlattachmentminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlattachmentminthreshold = intPtr(d.Get("xmlattachmentminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("xmlattachmentpercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlattachmentpercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlattachmentpercentthreshold = intPtr(d.Get("xmlattachmentpercentthreshold").(int))
		hasChange = true
	}
	if d.HasChange("xmlwsiminthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlwsiminthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlwsiminthreshold = intPtr(d.Get("xmlwsiminthreshold").(int))
		hasChange = true
	}
	if d.HasChange("xmlwsipercentthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlwsipercentthreshold has changed for appfwlearningsettings %s, starting update", appfwlearningsettingsName)
		appfwlearningsettings.Xmlwsipercentthreshold = intPtr(d.Get("xmlwsipercentthreshold").(int))
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
