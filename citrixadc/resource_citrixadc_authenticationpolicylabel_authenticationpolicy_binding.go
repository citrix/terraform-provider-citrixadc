package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcAuthenticationpolicylabel_authenticationpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationpolicylabel_authenticationpolicy_bindingFunc,
		Read:          readAuthenticationpolicylabel_authenticationpolicy_bindingFunc,
		Delete:        deleteAuthenticationpolicylabel_authenticationpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nextfactor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAuthenticationpolicylabel_authenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationpolicylabel_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	labelname := d.Get("labelname")
	policyname := d.Get("policyname")
	bindingId := fmt.Sprintf("%s,%s", labelname, policyname)
	authenticationpolicylabel_authenticationpolicy_binding := authentication.Authenticationpolicylabelauthenticationpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Labelname:              d.Get("labelname").(string),
		Nextfactor:             d.Get("nextfactor").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), &authenticationpolicylabel_authenticationpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAuthenticationpolicylabel_authenticationpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationpolicylabel_authenticationpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAuthenticationpolicylabel_authenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationpolicylabel_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	labelname := idSlice[0]
	policyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationpolicylabel_authenticationpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "authenticationpolicylabel_authenticationpolicy_binding",
		ResourceName:             labelname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationpolicylabel_authenticationpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationpolicylabel_authenticationpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("labelname", data["labelname"])
	d.Set("nextfactor", data["nextfactor"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteAuthenticationpolicylabel_authenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationpolicylabel_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	labelname := idSlice[0]
	policyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	
	if v, ok := d.GetOk("priority"); ok {
		args = append(args, fmt.Sprintf("priority:%s", url.QueryEscape(fmt.Sprintf("%v", v))))
	}

	err := client.DeleteResourceWithArgs(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), labelname, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
