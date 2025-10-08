package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslcipher() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcipherFunc,
		ReadContext:   readSslcipherFunc,
		UpdateContext: updateSslcipherFunc,
		DeleteContext: deleteSslcipherFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ciphergroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// sslcipher_sslciphersuite_binidng is MANDATORY attribute
			"ciphersuitebinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      sslcipherCipherSuitebindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ciphername": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cipherpriority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createSslcipherFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcipherFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcipherGroupName := d.Get("ciphergroupname").(string)

	sslcipher := ssl.Sslcipher{
		Ciphergroupname: sslcipherGroupName,
	}

	_, err := client.AddResource(service.Sslcipher.Type(), sslcipherGroupName, &sslcipher)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslcipherGroupName)

	// Ignore bindings unless there is an explicit configuration for it
	if _, ok := d.GetOk("ciphersuitebinding"); ok {
		err = updateSslCipherCipherSuiteBindings(d, meta)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return readSslcipherFunc(ctx, d, meta)
}

func updateSslcipherFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslcipherFunc")
	sslcipherGroupName := d.Get("ciphergroupname").(string)
	hasChange := false
	if d.HasChange("ciphersuitebinding") {
		log.Printf("[DEBUG]  citrixadc-provider: Ciphersuitebinding has changed for sslcipherGroupName %s, starting update", sslcipherGroupName)
		hasChange = true
	}
	if hasChange {
		err := updateSslCipherCipherSuiteBindings(d, meta)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func readSslcipherFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcipherFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcipherGroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcipher state %s", sslcipherGroupName)
	dataArr, err := client.FindAllResources(service.Sslcipher.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcipher state %s", sslcipherGroupName)
		d.SetId("")
		return nil
	}
	// for some of the NetScaler version Get was not working so using GetAll
	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: Sslcipher does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["ciphergroupname"].(string) == sslcipherGroupName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllresources Sslcipher not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing Sslcipher state %s", sslcipherGroupName)
		d.SetId("")
		return nil
	}
	data := dataArr[foundIndex]

	// Ignore bindings unless there is an explicit configuration for it
	if _, ok := d.GetOk("ciphersuitebinding"); ok {
		err = readSslCipherCipherSuitebindings(d, meta)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.Set("ciphergroupname", data["ciphergroupname"])

	return nil

}

func deleteSslcipherFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcipherFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcipherGroupName := d.Id()
	err := client.DeleteResource(service.Sslcipher.Type(), sslcipherGroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

type cipherPriority struct {
	cipherName     string
	cipherPriority uint32
}

type cipherPriorities []cipherPriority

func (n cipherPriorities) Len() int {
	return len(n)
}

func getSortedCipherBindigs(unsortedCipherBindings *schema.Set) cipherPriorities {
	log.Printf("[DEBUG]  citrixadc-provider: In getSortedCipherBindigs")
	sortedciphers := make(cipherPriorities, 0, unsortedCipherBindings.Len())

	for _, v := range unsortedCipherBindings.List() {
		val := v.(map[string]interface{})
		ciphername := val["ciphername"].(string)
		cipherpriority := uint32(val["cipherpriority"].(int))
		cipher := cipherPriority{
			cipherName:     ciphername,
			cipherPriority: cipherpriority,
		}

		sortedciphers = append(sortedciphers, cipher)
	}
	sort.Slice(sortedciphers, func(i, j int) bool {
		return sortedciphers[i].cipherPriority < sortedciphers[j].cipherPriority
	})
	return sortedciphers
}

func deleteSingleSslCipherCipherSuiteBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSslCipherCipherSuiteBinding")
	client := meta.(*NetScalerNitroClient).client

	ciphergroupname := d.Get("ciphergroupname").(string)
	// Construct args from binding data
	args := make([]string, 0, 1)

	if d, ok := binding["ciphername"]; ok {
		s := fmt.Sprintf("ciphername:%v", d.(string))
		args = append(args, s)
	}

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting sslciphersuite binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleSslCipherCipherSuiteBinding(d *schema.ResourceData, meta interface{}, binding cipherPriority) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSslCipherCipherSuiteBinding")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("Adding binding %v", binding)

	bindingStruct := ssl.Sslcipherciphersuitebinding{}
	bindingStruct.Ciphergroupname = d.Get("ciphergroupname").(string)
	bindingStruct.Ciphername = binding.cipherName
	bindingStruct.Cipherpriority = binding.cipherPriority

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource(service.Sslcipher_sslciphersuite_binding.Type(), bindingStruct.Ciphergroupname, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateSslCipherCipherSuiteBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslCipherCipherSuiteBindings")
	client := meta.(*NetScalerNitroClient).client

	findParams := service.FindParams{
		ResourceType: "sslcipher_sslciphersuite_binding",
		ResourceName: d.Get("ciphergroupname").(string),
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}

	// We need to do this since adding a ciphersuite with lower priority than an existing one
	// will bump the existing priority by one.
	// Delete all existing bindings
	for _, data := range dataArr {
		binding := make(map[string]interface{})
		binding["ciphername"] = data["ciphername"]
		binding["cipherpriority"] = data["cipherpriority"]

		if err := deleteSingleSslCipherCipherSuiteBinding(d, meta, binding); err != nil {
			return err
		}
	}
	// Add all configured bindings
	add := d.Get("ciphersuitebinding").(*schema.Set)
	for _, binding := range getSortedCipherBindigs(add) {
		if err := addSingleSslCipherCipherSuiteBinding(d, meta, binding); err != nil {
			return err
		}
	}
	return nil
}

func sslcipherCipherSuitebindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In sslcipherCipherSuitebindingMappingHash")
	var buf bytes.Buffer

	m := v.(map[string]interface{})
	if d, ok := m["ciphername"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["cipherpriority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}
	return hashString(buf.String())
}

func readSslCipherCipherSuitebindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSslCipherCipherSuitebindings")
	client := meta.(*NetScalerNitroClient).client
	ciphergroupname := d.Get("ciphergroupname").(string)
	bindings, _ := client.FindResourceArray(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i := range bindings {
		processedBindings[i] = make(map[string]interface{})
		processedBindings[i].(map[string]interface{})["ciphername"] = bindings[i]["ciphername"].(string)
		processedBindings[i].(map[string]interface{})["cipherpriority"], _ = strconv.Atoi(bindings[i]["cipherpriority"].(string))
	}

	updatedSet := schema.NewSet(sslcipherCipherSuitebindingMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("ciphersuitebinding", updatedSet); err != nil {
		return err
	}
	return nil
}
