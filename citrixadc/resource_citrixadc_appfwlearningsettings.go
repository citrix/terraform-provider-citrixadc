package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppfwlearningsettings() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwlearningsettingsFunc,
		Read:          readAppfwlearningsettingsFunc,
		Update:        updateAppfwlearningsettingsFunc,
		Delete:        deleteAppfwlearningsettingsFunc,// Thought appfwlearningsettings resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"contenttypeautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"contenttypeminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"contenttypepercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookieconsistencyautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookieconsistencyminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookieconsistencypercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creditcardnumberminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creditcardnumberpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csrftagautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csrftagminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"csrftagpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldconsistencyautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldconsistencyminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldconsistencypercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldformatautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldformatminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fieldformatpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"starturlautodeploygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"starturlminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"starturlpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlattachmentminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlattachmentpercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlwsiminthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlwsipercentthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwlearningsettingsFunc(d *schema.ResourceData, meta interface{}) error {
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
		Starturlpercentthreshold:                d.Get("starturlpercentthreshold").(int),
		Xmlattachmentminthreshold:               d.Get("xmlattachmentminthreshold").(int),
		Xmlattachmentpercentthreshold:           d.Get("xmlattachmentpercentthreshold").(int),
		Xmlwsiminthreshold:                      d.Get("xmlwsiminthreshold").(int),
		Xmlwsipercentthreshold:                  d.Get("xmlwsipercentthreshold").(int),
	}

	err := client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
	if err != nil {
		return err
	}

	d.SetId(appfwlearningsettingsName)

	err = readAppfwlearningsettingsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwlearningsettings but we can't read it ?? %s", appfwlearningsettingsName)
		return nil
	}
	return nil
}

func readAppfwlearningsettingsFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("contenttypeautodeploygraceperiod", data["contenttypeautodeploygraceperiod"])
	d.Set("contenttypeminthreshold", data["contenttypeminthreshold"])
	d.Set("contenttypepercentthreshold", data["contenttypepercentthreshold"])
	d.Set("cookieconsistencyautodeploygraceperiod", data["cookieconsistencyautodeploygraceperiod"])
	d.Set("cookieconsistencyminthreshold", data["cookieconsistencyminthreshold"])
	d.Set("cookieconsistencypercentthreshold", data["cookieconsistencypercentthreshold"])
	d.Set("creditcardnumberminthreshold", data["creditcardnumberminthreshold"])
	d.Set("creditcardnumberpercentthreshold", data["creditcardnumberpercentthreshold"])
	d.Set("crosssitescriptingautodeploygraceperiod", data["crosssitescriptingautodeploygraceperiod"])
	d.Set("crosssitescriptingminthreshold", data["crosssitescriptingminthreshold"])
	d.Set("crosssitescriptingpercentthreshold", data["crosssitescriptingpercentthreshold"])
	d.Set("csrftagautodeploygraceperiod", data["csrftagautodeploygraceperiod"])
	d.Set("csrftagminthreshold", data["csrftagminthreshold"])
	d.Set("csrftagpercentthreshold", data["csrftagpercentthreshold"])
	d.Set("fieldconsistencyautodeploygraceperiod", data["fieldconsistencyautodeploygraceperiod"])
	d.Set("fieldconsistencyminthreshold", data["fieldconsistencyminthreshold"])
	d.Set("fieldconsistencypercentthreshold", data["fieldconsistencypercentthreshold"])
	d.Set("fieldformatautodeploygraceperiod", data["fieldformatautodeploygraceperiod"])
	d.Set("fieldformatminthreshold", data["fieldformatminthreshold"])
	d.Set("fieldformatpercentthreshold", data["fieldformatpercentthreshold"])
	d.Set("sqlinjectionautodeploygraceperiod", data["sqlinjectionautodeploygraceperiod"])
	d.Set("sqlinjectionminthreshold", data["sqlinjectionminthreshold"])
	d.Set("sqlinjectionpercentthreshold", data["sqlinjectionpercentthreshold"])
	d.Set("starturlautodeploygraceperiod", data["starturlautodeploygraceperiod"])
	d.Set("starturlminthreshold", data["starturlminthreshold"])
	d.Set("starturlpercentthreshold", data["starturlpercentthreshold"])
	d.Set("xmlattachmentminthreshold", data["xmlattachmentminthreshold"])
	d.Set("xmlattachmentpercentthreshold", data["xmlattachmentpercentthreshold"])
	d.Set("xmlwsiminthreshold", data["xmlwsiminthreshold"])
	d.Set("xmlwsipercentthreshold", data["xmlwsipercentthreshold"])

	return nil

}

func updateAppfwlearningsettingsFunc(d *schema.ResourceData, meta interface{}) error {
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
			return fmt.Errorf("Error updating appfwlearningsettings %s", appfwlearningsettingsName)
		}
	}
	return readAppfwlearningsettingsFunc(d, meta)
}

func deleteAppfwlearningsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwlearningsettingsFunc")
	
	d.SetId("")

	return nil
}
