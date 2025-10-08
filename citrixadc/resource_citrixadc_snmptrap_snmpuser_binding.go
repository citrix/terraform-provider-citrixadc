package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func resourceCitrixAdcSnmptrap_snmpuser_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmptrap_snmpuser_bindingFunc,
		ReadContext:   readSnmptrap_snmpuser_bindingFunc,
		DeleteContext: deleteSnmptrap_snmpuser_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"trapclass": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"trapdestination": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"securitylevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: false,
				ForceNew: true,
				Default:  0,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
				Default:  "V3",
			},
		},
	}
}

func createSnmptrap_snmpuser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmptrap_snmpuser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	snmptrapId := d.Get("trapclass").(string) + "," + d.Get("trapdestination").(string)
	username := d.Get("username").(string)
	log.Printf("%v", username)
	bindingId := fmt.Sprintf("%s,%s", snmptrapId, username)
	snmptrap_snmpuser_binding := snmp.Snmptrapsnmpuserbinding{
		Securitylevel:   d.Get("securitylevel").(string),
		Td:              d.Get("td").(int),
		Trapclass:       d.Get("trapclass").(string),
		Trapdestination: d.Get("trapdestination").(string),
		Username:        d.Get("username").(string),
		Version:         d.Get("version").(string),
	}
	log.Printf("%v", username)
	_, err := client.AddResource(service.Snmptrap_snmpuser_binding.Type(), bindingId, &snmptrap_snmpuser_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSnmptrap_snmpuser_bindingFunc(ctx, d, meta)
}

func readSnmptrap_snmpuser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmptrap_snmpuser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 3)

	trapclass := idSlice[0]
	trapdestination := idSlice[1]
	username := idSlice[2]

	args := make(map[string]string, 0)
	args["trapclass"] = trapclass
	args["trapdestination"] = trapdestination

	// Check if the `version` is given by user, if not then set the default value, because this is needed for GET api call in args
	if val, ok := d.GetOk("version"); ok {
		args["version"] = val.(string)
	} else {
		args["version"] = "V3"
	}

	// Check if the `td` is given by user, if not then set the default value, because this is needed for GET api call in args
	if val, ok := d.GetOk("td"); ok {
		args["td"] = strconv.Itoa(val.(int))
	} else {
		args["td"] = "0"
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading snmptrap_snmpuser_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "snmptrap_snmpuser_binding",
		ArgsMap:                  args,
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
		log.Printf("[WARN] citrixadc-provider: Clearing snmptrap_snmpuser_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["username"].(string) == username {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmptrap_snmpuser_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("securitylevel", data["securitylevel"])
	setToInt("td", d, data["td"])
	d.Set("trapclass", data["trapclass"])
	d.Set("trapdestination", data["trapdestination"])
	d.Set("username", data["username"])
	d.Set("version", data["version"])

	return nil

}

func deleteSnmptrap_snmpuser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmptrap_snmpuser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	trapdestination := idSlice[1]
	username := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("trapdestination:%s", trapdestination))
	args = append(args, fmt.Sprintf("username:%s", username))
	if val, ok := d.GetOk("version"); ok {
		args = append(args, fmt.Sprintf("version:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%v", val.(int)))
	}

	err := client.DeleteResourceWithArgs(service.Snmptrap_snmpuser_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
