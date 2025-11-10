package citrixadc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"gopkg.in/yaml.v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcNintroResource() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNitroResourceFunc,
		ReadContext:   readNitroResourceFunc,
		UpdateContext: updateNitroResourceFunc,
		DeleteContext: deleteNitroResourceFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"workflows_file": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workflow": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attributes": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"non_updateable_attributes": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNitroResourceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In createNitroResourceFunc")
	//client := meta.(*NetScalerNitroClient).client

	workflow, err := readWorkflow(d)
	if err != nil {
		return diag.FromErr(err)
	}
	// log.Printf("workflows read %v", workflows)
	log.Printf("workflow read %v", workflow)

	switch workflow["lifecycle"] {
	case "object":
		err := createObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "non_updateable_object":
		err := createObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "binding":
		err := createBindingFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "object_by_args":
		err := createObjectByArgsFunc(d, meta, workflow)
		return diag.FromErr(err)
	default:
		return diag.Errorf("Lifecycle \"%v\" does not have a create function", workflow["lifecycle"])
	}
}

func readNitroResourceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readNitroResourceFunc")
	//client := meta.(*NetScalerNitroClient).client
	// id := d.Id()

	workflow, err := readWorkflow(d)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("workflow read %v", workflow)

	switch workflow["lifecycle"] {
	case "object":
		err := readObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "non_updateable_object":
		err := readObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "binding":
		err := readBindingFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "object_by_args":
		err := readObjectByArgsFunc(d, meta, workflow)
		return diag.FromErr(err)
	default:
		return diag.Errorf("Lifecycle \"%v\" does not have a read function", workflow["lifecycle"])
	}

}

func updateNitroResourceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In updateNitroResourceFunc")

	workflow, err := readWorkflow(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("workflow read %v", workflow)

	switch workflow["lifecycle"] {
	case "object":
		err := updateObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "object_by_args":
		err := updateObjectByArgsFunc(d, meta, workflow)
		return diag.FromErr(err)
	default:
		return diag.Errorf("Lifecycle \"%v\" does not have an update function", workflow["lifecycle"])
	}
}

func deleteNitroResourceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteServerFunc")

	workflow, err := readWorkflow(d)

	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("workflow read %v", workflow)

	switch workflow["lifecycle"] {
	case "object":
		err := deleteObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "non_updateable_object":
		err := deleteObjectFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "binding":
		err := deleteBindingFunc(d, meta, workflow)
		return diag.FromErr(err)
	case "object_by_args":
		err := deleteObjectByArgsFunc(d, meta, workflow)
		return diag.FromErr(err)
	default:
		return diag.Errorf("Lifecycle \"%v\" does not have a delete function", workflow["lifecycle"])
	}
}

func createObjectFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createObjectFunc")

	client := meta.(*NetScalerNitroClient).client

	primaryIdAttribute := workflow["primary_id_attribute"]

	primaryId := getConfiguredValue(d, primaryIdAttribute)

	if primaryId == nil {
		return fmt.Errorf("Configured object does not contain primary id attribute %v", primaryIdAttribute)
	}

	object := getConfiguredMap(d)

	_, err := client.AddResource(workflow["endpoint"].(string), primaryId.(string), &object)

	if err != nil {
		return err
	}

	d.SetId(primaryId.(string))

	err = readObjectFunc(d, meta, workflow)
	if err != nil {
		return fmt.Errorf("Error when reading created object %s", err)
	}

	return nil
}

func readObjectFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In readObjectFunc")
	client := meta.(*NetScalerNitroClient).client

	primaryId := d.Id()

	findParams := service.FindParams{
		ResourceType:             workflow["endpoint"].(string),
		ResourceName:             primaryId,
		ResourceMissingErrorCode: workflow["resource_missing_errorcode"].(int),
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing nitro resource state %s", primaryId)
		d.SetId("")
		return nil
	}

	if len(dataArr) > 1 {
		return fmt.Errorf("FindResourceArrayWithParams returned too many results")
	}

	data := dataArr[0]
	setConfiguredAttributes(d, data, workflow)

	return nil
}

func updateObjectFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateObjectFunc")

	client := meta.(*NetScalerNitroClient).client

	primaryIdAttribute := workflow["primary_id_attribute"].(string)

	primaryId := getConfiguredValue(d, primaryIdAttribute)

	if primaryId == nil {
		return fmt.Errorf("Configured object does not contain primary id attribute %v", primaryIdAttribute)
	}

	var object map[string]interface{}
	nitroObject := make(map[string]interface{})

	// The map copy will work for simple values
	if v, ok := d.GetOk("attributes"); ok {
		object = v.(map[string]interface{})
		for k, v := range object {
			nitroObject[k] = v
		}
	}

	// Add primary id even if defined in non_updateable_attributes
	if _, ok := nitroObject[primaryIdAttribute]; !ok {
		nitroObject[primaryIdAttribute] = primaryId
	}

	_, err := client.UpdateResource(workflow["endpoint"].(string), primaryId.(string), &nitroObject)
	if err != nil {
		return fmt.Errorf("Error updating object %s", err)
	}

	return readObjectFunc(d, meta, workflow)
}

func deleteObjectFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteObjectFunc")
	client := meta.(*NetScalerNitroClient).client
	primaryId := d.Id()

	err := client.DeleteResource(workflow["endpoint"].(string), primaryId)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func createBindingFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createBindingFunc")

	client := meta.(*NetScalerNitroClient).client

	primaryIdAttribute := workflow["primary_id_attribute"]
	primaryId := getConfiguredValue(d, primaryIdAttribute)
	if primaryId == nil {
		return fmt.Errorf("Configured binding does not contain primary id attribute %v", primaryIdAttribute)
	}

	secondaryIdAttribute := workflow["secondary_id_attribute"]
	secondaryId := getConfiguredValue(d, secondaryIdAttribute)
	if secondaryId == nil {
		return fmt.Errorf("Configured binding does not contain secondary id attribute %v", secondaryIdAttribute)
	}

	binding := getConfiguredMap(d)

	_, err := client.UpdateResource(workflow["endpoint"].(string), primaryId.(string), &binding)

	if err != nil {
		return err
	}
	bindingId := fmt.Sprintf("%v,%v", primaryId, secondaryId)
	d.SetId(bindingId)

	err = readBindingFunc(d, meta, workflow)
	if err != nil {
		return fmt.Errorf("Error when reading created binding %s", err)
	}
	return nil
}

func readBindingFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In readBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)

	primaryId := idSlice[0]
	secondaryId := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             workflow["endpoint"].(string),
		ResourceName:             primaryId,
		ResourceMissingErrorCode: workflow["bound_resource_missing_errorcode"].(int),
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing nitro resource state %s", primaryId)
		d.SetId("")
		return nil
	}
	secondaryIdAttribute := workflow["secondary_id_attribute"]
	foundIndex := -1
	for index, binding := range dataArr {
		if fmt.Sprintf("%v", binding[secondaryIdAttribute.(string)]) == secondaryId {
			foundIndex = index
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondary id attribute not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	setConfiguredAttributes(d, data, workflow)

	return nil
}

func deleteBindingFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteBindingFunc")

	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)

	primaryId := idSlice[0]
	secondaryId := idSlice[1]

	secondaryIdAttribute := workflow["secondary_id_attribute"]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("%v:%s", secondaryIdAttribute, secondaryId))

	err := client.DeleteResourceWithArgs(workflow["endpoint"].(string), primaryId, args)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func createObjectByArgsFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createObjectByArgsFunc")

	client := meta.(*NetScalerNitroClient).client

	deleteIdAttributes := workflow["delete_id_attributes"].([]interface{})
	idSlice := make([]string, 0, 0)

	for _, deleteIdAttribute := range deleteIdAttributes {
		attributeValue := getConfiguredValue(d, deleteIdAttribute)
		if attributeValue != nil {
			idItem := fmt.Sprintf("%s:%s", deleteIdAttribute, attributeValue)
			idSlice = append(idSlice, idItem)
		}
	}

	if len(idSlice) == 0 {
		return fmt.Errorf("Configured object does not contain any id attribute")
	}

	idString := strings.Join(idSlice, ",")

	object := getConfiguredMap(d)

	_, err := client.AddResource(workflow["endpoint"].(string), "", &object)

	if err != nil {
		return err
	}

	d.SetId(idString)

	err = readObjectByArgsFunc(d, meta, workflow)
	if err != nil {
		return fmt.Errorf("Error when reading created object %s", err)
	}

	return nil
}

func readObjectByArgsFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In readObjectByArgsFunc")
	client := meta.(*NetScalerNitroClient).client

	primaryId := d.Id()
	idItems := strings.Split(primaryId, ",")
	argsMap := make(map[string]string)

	for _, idItem := range idItems {
		idSlice := strings.Split(idItem, ":")
		key := url.QueryEscape(idSlice[0])
		value := url.QueryEscape(idSlice[1])
		argsMap[key] = value
	}

	findParams := service.FindParams{
		ResourceType:             workflow["endpoint"].(string),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: workflow["resource_missing_errorcode"].(int),
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing nitro resource state %s", primaryId)
		d.SetId("")
		return nil
	}

	if len(dataArr) > 1 {
		return fmt.Errorf("FindResourceArrayWithParams returned too many results")
	}

	data := dataArr[0]
	setConfiguredAttributes(d, data, workflow)

	return nil
}

func updateObjectByArgsFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateObjectByArgsFunc")

	client := meta.(*NetScalerNitroClient).client

	var object map[string]interface{}
	nitroObject := make(map[string]interface{})

	// The map copy will work for simple values
	if v, ok := d.GetOk("attributes"); ok {
		object = v.(map[string]interface{})
		for k, v := range object {
			nitroObject[k] = v
		}
	}

	for _, key := range workflow["delete_id_attributes"].([]interface{}) {
		value := getConfiguredValue(d, key)
		if value == nil {
			continue
		}
		// Add primary ids even if defined in non_updateable_attributes
		if _, ok := nitroObject[key.(string)]; !ok {
			nitroObject[key.(string)] = value
		}
	}

	err := client.UpdateUnnamedResource(workflow["endpoint"].(string), &nitroObject)
	if err != nil {
		return fmt.Errorf("Error updating object %s", err)
	}

	return readObjectByArgsFunc(d, meta, workflow)
}

func deleteObjectByArgsFunc(d *schema.ResourceData, meta interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteObjectByArgsFunc")
	client := meta.(*NetScalerNitroClient).client
	primaryId := d.Id()
	idItems := strings.Split(primaryId, ",")
	argsMap := make(map[string]string)

	for _, idItem := range idItems {
		idSlice := strings.Split(idItem, ":")
		key := url.QueryEscape(idSlice[0])
		value := url.QueryEscape(idSlice[1])
		argsMap[key] = value
	}

	err := client.DeleteResourceWithArgsMap(workflow["endpoint"].(string), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func readWorkflow(d *schema.ResourceData) (map[interface{}]interface{}, error) {
	log.Printf("[DEBUG]  netscaler-provider: In readWorkflow")
	yamlFileName := d.Get("workflows_file").(string)
	fileData, err := ioutil.ReadFile(yamlFileName)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = yaml.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	workflowsDict, ok := data.(map[interface{}]interface{})["workflow"]
	if !ok {
		return nil, fmt.Errorf("Top level workflow key not found in workflows yaml file")
	}

	specificWorkflow, ok := workflowsDict.(map[interface{}]interface{})[d.Get("workflow")]
	if !ok {
		return nil, fmt.Errorf("Key %v not found in workflows map", d.Get("workflow"))
	}

	return specificWorkflow.(map[interface{}]interface{}), nil
}

func getConfiguredValue(d *schema.ResourceData, key interface{}) interface{} {
	log.Printf("[DEBUG]  netscaler-provider: In getConfiguredValue")

	if v, ok := d.GetOk("attributes"); ok {
		retval, ok := v.(map[string]interface{})[key.(string)]
		if ok {
			return retval
		}
	}

	if v, ok := d.GetOk("non_updateable_attributes"); ok {
		retval, ok := v.(map[string]interface{})[key.(string)]
		if ok {
			return retval
		}
	}
	// Fallthrough
	return nil
}

func getConfiguredMap(d *schema.ResourceData) map[string]interface{} {
	log.Printf("[DEBUG]  netscaler-provider: In getConfiguredMap")
	retVal := make(map[string]interface{})
	if v, ok := d.GetOk("attributes"); ok {
		for k, v := range v.(map[string]interface{}) {
			retVal[k] = v
		}
	}

	if v, ok := d.GetOk("non_updateable_attributes"); ok {
		for k, v := range v.(map[string]interface{}) {
			retVal[k] = v
		}
	}

	return retVal
}

func setConfiguredAttributes(d *schema.ResourceData, data map[string]interface{}, workflow map[interface{}]interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In setConfiguredAttributes")

	attributesMap := make(map[string]interface{})
	nonUpdateableAttributesMap := make(map[string]interface{})

	for dataKey, dataValueRaw := range data {
		// Every value is converted to string
		dataValue := fmt.Sprintf("%v", dataValueRaw)
		log.Printf("[DEBUG]  netscaler-provider: exploring data key %v", dataKey)
		if v, ok := d.GetOk("attributes"); ok {
			log.Printf("[DEBUG] citrixadc-provider:  attributes %v", v)
			if _, ok := v.(map[string]interface{})[dataKey]; ok {
				log.Printf("[DEBUG]  netscaler-provider: Setting attribute %v:%v", dataKey, dataValue)
				attributesMap[dataKey] = dataValue
			}
		}

		if v, ok := d.GetOk("non_updateable_attributes"); ok {
			log.Printf("[DEBUG] citrixadc-provider:  non_updateable_attributes %v", v)
			if _, ok := v.(map[string]interface{})[dataKey]; ok {
				log.Printf("[DEBUG]  netscaler-provider: Setting non updateable attribute %v:%v", dataKey, dataValue)
				nonUpdateableAttributesMap[dataKey] = dataValue
			}
		}

	}

	d.Set("attributes", attributesMap)
	d.Set("non_updateable_attributes", nonUpdateableAttributesMap)

	return nil
}
