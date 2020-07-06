package citrixadc

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/chiradeep/go-nitro/config/ssl"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSslcipher() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcipherFunc,
		Read:          readSslcipherFunc,
		// Update:        updateSslcipherFunc, // All fields are ForceNew or Computed w/out Optional, Update is superfluous
		Delete: deleteSslcipherFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ciphergroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// sslcipher_sslciphersuite_binidng is MANDATORY attribute
			"ciphersuitebinding": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      sslcipherCipherSuitebindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ciphername": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cipherpriority": &schema.Schema{
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

func createSslcipherFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcipherFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcipherGroupName := d.Get("ciphergroupname").(string)

	sslcipher := ssl.Sslcipher{
		Ciphergroupname: sslcipherGroupName,
	}

	_, err := client.AddResource(netscaler.Sslcipher.Type(), sslcipherGroupName, &sslcipher)
	if err != nil {
		return err
	}

	d.SetId(sslcipherGroupName)

	err = updateSslCipherCipherSuiteBindings(d, meta)
	if err != nil {
		return err
	}

	err = readSslcipherFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslcipher but we can't read it ?? %s", sslcipherGroupName)
		return err
	}
	return nil
}

func readSslcipherFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcipherFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcipherGroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcipher state %s", sslcipherGroupName)
	data, err := client.FindResource(netscaler.Sslcipher.Type(), sslcipherGroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcipher state %s", sslcipherGroupName)
		d.SetId("")
		return nil
	}

	err = readSslCipherCipherSuitebindings(d, meta)
	if err != nil {
		return err
	}

	d.Set("ciphergroupname", data["ciphergroupname"])

	return nil

}

func deleteSslcipherFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcipherFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcipherGroupName := d.Id()
	err := client.DeleteResource(netscaler.Sslcipher.Type(), sslcipherGroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

// sslcipher_sslciphersuite_binding

type cipherPriority struct {
	cipherName     string
	cipherPriority int
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
		cipherpriority := val["cipherpriority"].(int)
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

	if err := client.DeleteResourceWithArgs(netscaler.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting sslciphersuite binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleSslCipherCipherSuiteBinding(d *schema.ResourceData, meta interface{}, binding cipherPriority) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSslCipherCipherSuiteBinding")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("Adding binding %v", binding)

	bindingStruct := ssl.Sslciphersslciphersuitebinding{}
	bindingStruct.Ciphergroupname = d.Get("ciphergroupname").(string)
	bindingStruct.Ciphername = binding.cipherName
	bindingStruct.Cipherpriority = binding.cipherPriority

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource(netscaler.Sslcipher_sslciphersuite_binding.Type(), bindingStruct.Ciphergroupname, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateSslCipherCipherSuiteBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslCipherCipherSuiteBindings")
	oldSet, newSet := d.GetChange("ciphersuitebinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))

	for _, binding := range remove.List() {
		if err := deleteSingleSslCipherCipherSuiteBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

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
	return hashcode.String(buf.String())
}

func readSslCipherCipherSuitebindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSslCipherCipherSuitebindings")
	client := meta.(*NetScalerNitroClient).client
	ciphergroupname := d.Get("ciphergroupname").(string)
	bindings, _ := client.FindResourceArray(netscaler.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i, _ := range bindings {
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
