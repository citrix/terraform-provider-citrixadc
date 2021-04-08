package citrixadc

import (
	"net/url"

	"github.com/chiradeep/go-nitro/config/appfw"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwprofile_cookieconsistency_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_cookieconsistency_bindingFunc,
		Read:          readAppfwprofile_cookieconsistency_bindingFunc,
		Delete:        deleteAppfwprofile_cookieconsistency_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cookieconsistency": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alertonly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isautodeployed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isregex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwprofile_cookieconsistency_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_cookieconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofile_cookieconsistency_bindingName := d.Get("name").(string)
	appfwprofile_cookieconsistency_binding := appfw.Appfwprofilecookieconsistencybinding{
		Alertonly:         d.Get("alertonly").(string),
		Comment:           d.Get("comment").(string),
		Cookieconsistency: d.Get("cookieconsistency").(string),
		Isautodeployed:    d.Get("isautodeployed").(string),
		Isregex:           d.Get("isregex").(string),
		Name:              d.Get("name").(string),
		State:             d.Get("state").(string),
	}

	_, err := client.AddResource(netscaler.Appfwprofile_cookieconsistency_binding.Type(), appfwprofile_cookieconsistency_bindingName, &appfwprofile_cookieconsistency_binding)
	if err != nil {
		return err
	}

	d.SetId(appfwprofile_cookieconsistency_bindingName)

	err = readAppfwprofile_cookieconsistency_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_cookieconsistency_binding but we can't read it ?? %s", appfwprofile_cookieconsistency_bindingName)
		return nil
	}
	return nil
}
func readAppfwprofile_cookieconsistency_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_cookieconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofile_cookieconsistency_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_cookieconsistency_binding state %s", appfwprofile_cookieconsistency_bindingName)
	data, err := client.FindResource(netscaler.Appfwprofile_cookieconsistency_binding.Type(), appfwprofile_cookieconsistency_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cookieconsistency_binding state %s", appfwprofile_cookieconsistency_bindingName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("cookieconsistency", data["cookieconsistency"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex", data["isregex"])
	d.Set("name", data["name"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_cookieconsistency_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_cookieconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	cookieConsistencyString := d.Get("cookieconsistency").(string)
	appFwName := d.Get("name").(string)
	args := make(map[string]string)
	args["cookieconsistency"] = url.QueryEscape(cookieConsistencyString)
	err := client.DeleteResourceWithArgsMap(netscaler.Appfwprofile_cookieconsistency_binding.Type(), appFwName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
