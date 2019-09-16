package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ssl"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func syncCiphers(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In syncCiphers")
	var ciphers interface{}
	var ok bool
	var err error
	var cipherBindings []map[string]interface{}
	if ciphers, ok = d.GetOk("ciphers"); ok {
		log.Printf("Configured ciphers %v", ciphers)
	}

	cipherBindings, err = client.FindResourceArray(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName)

	// When ciphers are not set and they do not exist there is no sync
	if !ok && len(cipherBindings) == 0 {
		return nil
	}

	// Catch error in Findresources
	if err != nil && len(cipherBindings) != 0 {
		log.Printf("Error finding ciphersuite bindings for vserver %v bindings %v with len %v ", vserverName, cipherBindings, len(cipherBindings))
		return err
	}

	// Evaluating the equality of resource and target ADC set cipher
	actualCiphers := make([]interface{}, 0, len(cipherBindings))
	for _, cipherBinding := range cipherBindings {
		actualCiphers = append(actualCiphers, cipherBinding["ciphername"])
	}

	// When existing and set are equal there is no sync
	if slicesEqual(ciphers.([]interface{}), actualCiphers) {
		log.Printf("Found actual cipher list to be equal to set cipher list")
		return nil
	}
	// Fallthrough

	// First delete all configured cipher bindings
	for _, cipherBinding := range cipherBindings {
		ciphername := cipherBinding["ciphername"].(string)
		argsMap := map[string]string{"ciphername": ciphername}
		log.Printf("Will delete ciphername %v from vserver %v", ciphername, vserverName)
		err := client.DeleteResourceWithArgsMap(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName, argsMap)
		if err != nil {
			log.Printf("Error deleting ciphername %v from vserver %v", ciphername, vserverName)
			return err
		}
	}

	// Then add all configured bindings
	cipherslice := ciphers.([]interface{})
	for _, ciphername := range cipherslice {
		resource := ssl.Sslvserversslciphersuitebinding{
			Vservername: vserverName,
			Ciphername:  ciphername.(string),
		}
		_, err = client.AddResource(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName, resource)
		if err != nil {
			log.Printf("Error binding cipher %v to vserver %v", ciphername, vserverName)
			return err
		}
	}

	return nil
}

func setCipherData(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In setCipherData")

	//exists := client.ResourceExists(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName)
	//cipherBindings := make([]map[string]interface{}, 0)
	cipherBindings, err := client.FindResourceArray(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName)
	if err != nil && len(cipherBindings) != 0 {
		log.Printf("Error retrieving cipher resource array")
		return err
	}
	log.Printf("cipherBindings %v\n", cipherBindings)
	cipherList := make([]interface{}, 0, len(cipherBindings))
	for _, cipherBinding := range cipherBindings {
		cipherList = append(cipherList, cipherBinding["ciphername"])
	}
	log.Printf("Setting ciphers to value %v", cipherList)
	d.Set("ciphers", cipherList)
	return nil
}

func slicesEqual(a, b []interface{}) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
