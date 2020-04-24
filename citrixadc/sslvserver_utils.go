package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ssl"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func syncCiphersuites(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In syncCiphersuites")
	var ciphersuites interface{}
	var ok bool
	var err error
	var ciphersuiteBindings []map[string]interface{}
	if ciphersuites, ok = d.GetOk("ciphersuites"); ok {
		log.Printf("Configured ciphersuites %v", ciphersuites)
	}

	ciphersuiteBindings, err = client.FindResourceArray(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName)

	// When cipher suites are not set and they do not exist there is no sync
	if !ok && len(ciphersuiteBindings) == 0 {
		return nil
	}

	// Catch error in Findresources
	if err != nil && len(ciphersuiteBindings) != 0 {
		log.Printf("Error finding ciphersuite bindings for vserver %v bindings %v with len %v ", vserverName, ciphersuiteBindings, len(ciphersuiteBindings))
		return err
	}

	// Evaluating the equality of resource and target ADC set cipher
	actualCiphersuites := make([]interface{}, 0, len(ciphersuiteBindings))
	for _, ciphersuiteBinding := range ciphersuiteBindings {
		actualCiphersuites = append(actualCiphersuites, ciphersuiteBinding["ciphername"])
	}

	// When existing and set are equal there is no sync
	if slicesEqual(ciphersuites.([]interface{}), actualCiphersuites) {
		log.Printf("Found actual ciphersuites list to be equal to set cipher list")
		return nil
	}
	// Fallthrough

	// First delete all configured cipher bindings
	for _, ciphersuiteBinding := range ciphersuiteBindings {
		ciphername := ciphersuiteBinding["ciphername"].(string)
		argsMap := map[string]string{"ciphername": ciphername}
		log.Printf("Will delete ciphername %v from vserver %v", ciphername, vserverName)
		err := client.DeleteResourceWithArgsMap(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName, argsMap)
		if err != nil {
			log.Printf("Error deleting ciphersuite %v from vserver %v", ciphername, vserverName)
			return err
		}
	}

	// Then add all configured bindings
	ciphersuiteslice := ciphersuites.([]interface{})
	for _, ciphername := range ciphersuiteslice {
		resource := ssl.Sslvserversslciphersuitebinding{
			Vservername: vserverName,
			Ciphername:  ciphername.(string),
		}
		_, err = client.AddResource(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName, resource)
		if err != nil {
			log.Printf("Error binding ciphersuite %v to vserver %v", ciphername, vserverName)
			return err
		}
	}

	return nil
}

func setCiphersuiteData(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In setCiphersuiteData")

	ciphersuiteBindings, err := client.FindResourceArray(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName)
	if err != nil && len(ciphersuiteBindings) != 0 {
		log.Printf("Error retrieving ciphersuite resource array")
		return err
	}
	log.Printf("ciphersuiteBindings %v\n", ciphersuiteBindings)
	ciphersuiteList := make([]interface{}, 0, len(ciphersuiteBindings))
	for _, ciphersuiteBinding := range ciphersuiteBindings {
		ciphersuiteList = append(ciphersuiteList, ciphersuiteBinding["ciphername"])
	}
	log.Printf("Setting ciphersuites to value %v", ciphersuiteList)
	d.Set("ciphersuites", ciphersuiteList)
	return nil
}

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

	findParams := netscaler.FindParams{
		ResourceType:             "sslvserver_sslcipher_binding",
		ResourceName:             vserverName,
		ResourceMissingErrorCode: 258,
	}
	cipherBindings, err = client.FindResourceArrayWithParams(findParams)

	if err != nil {
		return err
	}

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
		actualCiphers = append(actualCiphers, cipherBinding["cipheraliasname"])
	}

	// When existing and set are equal there is no sync
	if slicesEqual(ciphers.([]interface{}), actualCiphers) {
		log.Printf("Found actual cipher list to be equal to set cipher list")
		return nil
	}
	// Fallthrough

	// First delete all configured cipher bindings
	for _, cipherBinding := range cipherBindings {
		cipheraliasname := cipherBinding["cipheraliasname"].(string)
		argsMap := map[string]string{"ciphername": cipheraliasname}
		log.Printf("Will delete cipheraliasname %v from vserver %v", cipheraliasname, vserverName)
		err := client.DeleteResourceWithArgsMap(netscaler.Sslvserver_sslcipher_binding.Type(), vserverName, argsMap)
		if err != nil {
			log.Printf("Error deleting cipheraliasname %v from vserver %v", cipheraliasname, vserverName)
			return err
		}
	}

	// Then add all configured bindings
	cipherslice := ciphers.([]interface{})
	for _, cipheraliasname := range cipherslice {
		resource := ssl.Sslvserversslciphersuitebinding{
			Vservername: vserverName,
			Ciphername:  cipheraliasname.(string),
		}
		_, err = client.AddResource(netscaler.Sslvserver_sslcipher_binding.Type(), vserverName, resource)
		if err != nil {
			log.Printf("Error binding cipher %v to vserver %v", cipheraliasname, vserverName)
			return err
		}
	}

	return nil
}

func setCipherData(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In setCipherData")

	cipherBindings, err := client.FindResourceArray(netscaler.Sslvserver_sslcipher_binding.Type(), vserverName)
	if err != nil && len(cipherBindings) != 0 {
		log.Printf("Error retrieving cipher resource array")
		return err
	}
	log.Printf("cipherBindings %v\n", cipherBindings)
	cipherList := make([]interface{}, 0, len(cipherBindings))
	for _, cipherBinding := range cipherBindings {
		cipherList = append(cipherList, cipherBinding["cipheraliasname"])
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

func syncSnisslcert(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In syncSnisslcert")
	old, new := d.GetChange("snisslcertkeys")

	var oldlist, newlist []interface{}
	if old == nil {
		oldlist = make([]interface{}, 0)
	} else {
		oldlist = old.(*schema.Set).List()
	}

	newlist = new.(*schema.Set).List()

	var toadd, todelete []string
	toadd = make([]string, 0, len(newlist))

	if old == nil {
		todelete = make([]string, 0)
	} else {
		todelete = make([]string, 0, len(oldlist))
	}

	// certificate key bindings that exist only in the old data will be deleted
	for _, oldcertkey := range oldlist {
		exists := false
		for _, newcertkey := range newlist {
			if oldcertkey.(string) == newcertkey.(string) {
				exists = true
				break
			}
		}
		if !exists {
			todelete = append(todelete, oldcertkey.(string))
		}
	}
	log.Printf("[DEBUG] The following sni certificate key bindings are marked for deletion %v", todelete)

	// certificate key bindings that exist only in the new data will be created
	for _, newcertkey := range newlist {
		exists := false
		for _, oldcertkey := range oldlist {
			if oldcertkey.(string) == newcertkey.(string) {
				exists = true
				break
			}
		}
		if !exists {
			toadd = append(toadd, newcertkey.(string))
		}
	}
	log.Printf("[DEBUG] The following sni certificate key bindings are marked for addition %v", toadd)

	// Do the unbindings first
	for _, snisslcertkey := range todelete {

		args := map[string]string{"certkeyname": snisslcertkey, "snicert": "true"}
		err := client.DeleteResourceWithArgsMap(netscaler.Sslvserver_sslcertkey_binding.Type(), sslvserverName, args)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding sni sslcertkey from sslvserver %s", snisslcertkey)
		}
		log.Printf("[DEBUG] netscaler-provider: sni sslcertkey has been unbound from sslvserver for sslcertkey %s ", snisslcertkey)
	}

	// Do the bindings
	for _, snisslcertkey := range toadd {
		binding := ssl.Sslvserversslcertkeybinding{
			Vservername: sslvserverName,
			Certkeyname: snisslcertkey,
			Snicert:     true,
		}
		log.Printf("[INFO] netscaler-provider:  Binding sni ssl cert %s to sslvserver %s", snisslcertkey, sslvserverName)
		err := client.BindResource(netscaler.Sslvserver.Type(), sslvserverName, netscaler.Sslcertkey.Type(), snisslcertkey, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind sni ssl cert %s to sslvserver %s", snisslcertkey, sslvserverName)
			err2 := client.DeleteResource(netscaler.Lbvserver.Type(), sslvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete sslvserver %s after bind to sni ssl cert failed", sslvserverName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete sslvserver %s after bind to sni ssl cert failed", sslvserverName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind sni ssl cert %s to sslvserver %s", snisslcertkey, sslvserverName)
		}
	}

	return nil
}

func snisslcertkeysExist(snisslcertkeys, meta interface{}) error {
	log.Printf("[DEBUG] In snisslcertkeysExist")
	client := meta.(*NetScalerNitroClient).client
	allkeys := snisslcertkeys.(*schema.Set).List()
	missingKeys := make([]string, 0, len(allkeys))
	for _, certkey := range allkeys {
		log.Printf("[DEBUG] checking existence of sslcertkey %v", certkey)
		exists := client.ResourceExists(netscaler.Sslcertkey.Type(), certkey.(string))
		if !exists {
			missingKeys = append(missingKeys, certkey.(string))
		}
	}
	if len(missingKeys) > 0 {
		return fmt.Errorf("The following ssl certificate keys do not exist on target ADC %v", missingKeys)
	} else {
		return nil
	}
}

func readSslcerts(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcerts")
	client := meta.(*NetScalerNitroClient).client
	bindings, err := client.FindAllBoundResources(netscaler.Sslvserver.Type(), sslvserverName, netscaler.Sslcertkey.Type())
	if err != nil {
		log.Printf("[WARN] netscaler-provider: sslvserver binding to ssl error %s", sslvserverName)
		return nil
	}
	var boundCert string
	snicerts := make([]string, 0, len(bindings))
	for _, binding := range bindings {
		cert, ok := binding["certkeyname"]
		snicert, ok2 := binding["snicert"]
		log.Printf("Reading ssl binding certkeyname %v, %v", cert, ok)
		log.Printf("Reading ssl binding snicert %v, %v", snicert, ok2)
		if ok && ok2 && snicert == false {
			boundCert = cert.(string)
		} else if ok && ok2 && snicert == true {
			snicerts = append(snicerts, cert.(string))
		}
	}
	d.Set("sslcertkey", boundCert)
	d.Set("snisslcertkeys", snicerts)
	return nil
}
