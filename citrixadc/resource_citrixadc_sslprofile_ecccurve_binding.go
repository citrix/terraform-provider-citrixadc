package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslprofile_ecccurve_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslprofile_ecccurve_bindingFunc,
		Read:          readSslprofile_ecccurve_bindingFunc,
		Delete:        deleteSslprofile_ecccurve_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"remove_existing_ecccurve_binding": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ecccurvename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslprofile_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslprofile_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	ecccurvename := d.Get("ecccurvename")
	bindingId := fmt.Sprintf("%s,%s", name, ecccurvename)
	sslprofile_ecccurve_binding := ssl.Sslprofileecccurvebinding{
		Ecccurvename: d.Get("ecccurvename").(string),
		Name:         d.Get("name").(string),
	}

	if val, ok := d.GetOk("remove_existing_ecccurve_binding"); ok && val.(bool) {
		log.Printf("[DEBUG]  citrixadc-provider: Removing all sslprofile_ecccurve_binding from %s", name)
		defaultEcccurves, err := getDefault_SslprofileEcccurveBindings(d, meta)
		log.Printf("[DEBUG] citrixadc-provider: defaultSslprofileEcccurveBindings: %v", defaultEcccurves)
		if err != nil {
			return err
		}
		for _, ecccurvename := range defaultEcccurves {
			deleteSingleSslprofileEcccurveBindings(d, meta, ecccurvename)
		}
	}

	_, err := client.AddResource(service.Sslprofile_ecccurve_binding.Type(), bindingId, &sslprofile_ecccurve_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslprofile_ecccurve_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslprofile_ecccurve_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func getDefault_SslprofileEcccurveBindings(d *schema.ResourceData, meta interface{}) ([]string, error) {
	log.Printf("[DEBUG]  citrixadc-provider: In getDefaultSslprofileEcccurveBindings")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)
	bindings, _ := client.FindResourceArray(service.Sslprofile_ecccurve_binding.Type(), sslprofileName)
	log.Printf("bindings %v\n", bindings)

	defaultSslprofileEcccurveBindings := make([]string, len(bindings))

	for i, val := range bindings {
		defaultSslprofileEcccurveBindings[i] = val["ecccurvename"].(string)
	}

	return defaultSslprofileEcccurveBindings, nil
}

func deleteSingleSslprofileEcccurveBindings(d *schema.ResourceData, meta interface{}, ecccurvename string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSslprofileEcccurveBinding")
	client := meta.(*NetScalerNitroClient).client

	sslprofileName := d.Get("name").(string)
	args := make([]string, 0, 1)

	s := fmt.Sprintf("ecccurvename:%s", ecccurvename)
	args = append(args, s)

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs(service.Sslprofile_ecccurve_binding.Type(), sslprofileName, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting EccCurve binding %v\n", sslprofileName)
		return err
	}

	return nil
}

func readSslprofile_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslprofile_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ecccurvename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslprofile_ecccurve_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslprofile_ecccurve_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile_ecccurve_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ecccurvename"].(string) == ecccurvename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile_ecccurve_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ecccurvename", data["ecccurvename"])
	d.Set("name", data["name"])

	return nil

}

func deleteSslprofile_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslprofile_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ecccurvename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ecccurvename:%s", ecccurvename))

	err := client.DeleteResourceWithArgs(service.Sslprofile_ecccurve_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
