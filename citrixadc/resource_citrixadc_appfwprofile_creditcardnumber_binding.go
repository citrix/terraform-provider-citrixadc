package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcAppfwprofile_creditcardnumber_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_creditcardnumber_bindingFunc,
		ReadContext:   readAppfwprofile_creditcardnumber_bindingFunc,
		DeleteContext: deleteAppfwprofile_creditcardnumber_bindingFunc,
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
			"creditcardnumber": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"creditcardnumberurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_creditcardnumber_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_creditcardnumber_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	creditcardnumber := d.Get("creditcardnumber")
	creditcardnumberurl := d.Get("creditcardnumberurl")
	bindingId := fmt.Sprintf("%s,%s,%s", name, creditcardnumber, creditcardnumberurl)
	appfwprofile_creditcardnumber_binding := appfw.Appfwprofilecreditcardnumberbinding{
		Alertonly:           d.Get("alertonly").(string),
		Comment:             d.Get("comment").(string),
		Creditcardnumber:    d.Get("creditcardnumber").(string),
		Creditcardnumberurl: d.Get("creditcardnumberurl").(string),
		Isautodeployed:      d.Get("isautodeployed").(string),
		Name:                d.Get("name").(string),
		Resourceid:          d.Get("resourceid").(string),
		Ruletype:            d.Get("ruletype").(string),
		State:               d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_creditcardnumber_binding.Type(), &appfwprofile_creditcardnumber_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_creditcardnumber_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_creditcardnumber_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_creditcardnumber_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	creditcardnumber := idSlice[1]
	creditcardnumberurl := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_creditcardnumber_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_creditcardnumber_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_creditcardnumber_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true
		if v["creditcardnumber"].(string) != creditcardnumber {
			match = false
		}
		if v["creditcardnumberurl"].(string) != creditcardnumberurl {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_creditcardnumber_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("creditcardnumber", data["creditcardnumber"])
	d.Set("creditcardnumberurl", data["creditcardnumberurl"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_creditcardnumber_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_creditcardnumber_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	creditcardnumber := idSlice[1]
	creditcardnumberurl := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("creditcardnumber:%s", creditcardnumber))
	args = append(args, fmt.Sprintf("creditcardnumberurl:%s", url.QueryEscape(creditcardnumberurl)))

	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_creditcardnumber_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
