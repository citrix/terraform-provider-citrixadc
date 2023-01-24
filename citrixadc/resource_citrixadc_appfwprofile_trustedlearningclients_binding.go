package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcAppfwprofile_trustedlearningclients_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_trustedlearningclients_bindingFunc,
		Read:          readAppfwprofile_trustedlearningclients_bindingFunc,
		Delete:        deleteAppfwprofile_trustedlearningclients_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"trustedlearningclients": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_trustedlearningclients_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_trustedlearningclients_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	trustedlearningclients := d.Get("trustedlearningclients")
	bindingId := fmt.Sprintf("%s,%s", name, trustedlearningclients)
	appfwprofile_trustedlearningclients_binding := appfw.Appfwprofiletrustedlearningclientsbinding{
		Alertonly:              d.Get("alertonly").(string),
		Comment:                d.Get("comment").(string),
		Isautodeployed:         d.Get("isautodeployed").(string),
		Name:                   d.Get("name").(string),
		Resourceid:             d.Get("resourceid").(string),
		Ruletype:               d.Get("ruletype").(string),
		State:                  d.Get("state").(string),
		Trustedlearningclients: d.Get("trustedlearningclients").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_trustedlearningclients_binding.Type(), &appfwprofile_trustedlearningclients_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_trustedlearningclients_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_trustedlearningclients_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_trustedlearningclients_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_trustedlearningclients_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	trustedlearningclients := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_trustedlearningclients_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_trustedlearningclients_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_trustedlearningclients_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["trustedlearningclients"].(string) == trustedlearningclients {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_trustedlearningclients_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("trustedlearningclients", data["trustedlearningclients"])

	return nil

}

func deleteAppfwprofile_trustedlearningclients_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_trustedlearningclients_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	trustedlearningclients := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("trustedlearningclients:%s", url.QueryEscape(trustedlearningclients)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_trustedlearningclients_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
