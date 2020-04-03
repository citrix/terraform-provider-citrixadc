package citrixadc

import (
	"fmt"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"strings"
)

func testCiphersConfig(resourceTemplate string, ciphers []string) string {
	var ciphers_replacement string
	if ciphers == nil {
		ciphers_replacement = ""
	} else {
		quoted := make([]string, len(ciphers))
		for i, ciph := range ciphers {
			quoted[i] = fmt.Sprintf("\"%v\"", ciph)
		}
		ciphers_replacement = fmt.Sprintf("ciphers = [%v]", strings.Join(quoted, ", "))
	}
	//log.Printf("%v", fmt.Sprintf(resourceTemplate, ciphers_replacement))
	return fmt.Sprintf(resourceTemplate, ciphers_replacement)
}

func testCiphersuitesConfig(resourceTemplate string, ciphers []string) string {
	var ciphers_replacement string
	if ciphers == nil {
		ciphers_replacement = ""
	} else {
		quoted := make([]string, len(ciphers))
		for i, ciph := range ciphers {
			quoted[i] = fmt.Sprintf("\"%v\"", ciph)
		}
		ciphers_replacement = fmt.Sprintf("ciphersuites = [%v]", strings.Join(quoted, ", "))
	}
	//log.Printf("%v", fmt.Sprintf(resourceTemplate, ciphers_replacement))
	return fmt.Sprintf(resourceTemplate, ciphers_replacement)
}

func testCheckCiphersEqualToActual(ciphers []string, vserverName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		// Check existence of ciphers in correct order
		cipherBindings, err := nsClient.FindResourceArray(netscaler.Sslvserver_sslcipher_binding.Type(), vserverName)
		if err != nil && len(cipherBindings) != 0 {
			log.Printf("Error retrieving cipher resource array")
			return err
		}

		log.Printf("cipherBindings %v\n", cipherBindings)

		cipherList := make([]interface{}, 0, len(cipherBindings))
		for _, cipherBinding := range cipherBindings {
			cipherList = append(cipherList, cipherBinding["cipheraliasname"])
		}
		c := make([]interface{}, len(ciphers))
		for i, v := range ciphers {
			c[i] = v
		}
		if equal := slicesEqual(cipherList, c); !equal {
			return fmt.Errorf("Ciphers not equal. Check: %v. Actual: %v", ciphers, cipherList)
		}
		return nil
	}
}

func testCheckCiphersuitesEqualToActual(ciphersuites []string, vserverName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		// Check existence of ciphers in correct order
		ciphersuiteBindings, err := nsClient.FindResourceArray(netscaler.Sslvserver_sslciphersuite_binding.Type(), vserverName)
		if err != nil && len(ciphersuiteBindings) != 0 {
			log.Printf("Error retrieving ciphersuite resource array")
			return err
		}

		log.Printf("ciphersuiteBindings %v\n", ciphersuiteBindings)

		ciphersuiteList := make([]interface{}, 0, len(ciphersuiteBindings))
		for _, ciphersuiteBinding := range ciphersuiteBindings {
			ciphersuiteList = append(ciphersuiteList, ciphersuiteBinding["ciphername"])
		}
		c := make([]interface{}, len(ciphersuites))
		for i, v := range ciphersuites {
			c[i] = v
		}
		if equal := slicesEqual(ciphersuiteList, c); !equal {
			return fmt.Errorf("Ciphersuites not equal. Check: %v. Actual: %v", ciphersuites, ciphersuiteList)
		}
		return nil
	}
}

func testCheckCiphersConfiguredExpected(resourceString string, ciphers []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceString]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceString)
		}
		keys := make([]string, 0, len(rs.Primary.Attributes))
		for k, _ := range rs.Primary.Attributes {
			keys = append(keys, k)
		}
		// Verify correct order of ciphers
		for i, v := range ciphers {
			resource_key := fmt.Sprintf("ciphers.%v", i)
			resource_value := rs.Primary.Attributes[resource_key]
			if resource_value != v {
				return fmt.Errorf("Cipher at position %v differs from expected. Expected: %v Actual: %v", i, v, resource_value)
			}
		}
		//log.Printf("all keys %v", keys)
		log.Printf("name on resource %v", rs.Primary.Attributes["name"])
		return nil
	}
}

func testCheckCiphersuitesConfiguredExpected(resourceString string, ciphersuites []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceString]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceString)
		}
		keys := make([]string, 0, len(rs.Primary.Attributes))
		for k, _ := range rs.Primary.Attributes {
			keys = append(keys, k)
		}
		// Verify correct order of ciphers
		for i, v := range ciphersuites {
			resource_key := fmt.Sprintf("ciphersuites.%v", i)
			resource_value := rs.Primary.Attributes[resource_key]
			if resource_value != v {
				return fmt.Errorf("Ciphersuite at position %v differs from expected. Expected: %v Actual: %v", i, v, resource_value)
			}
		}
		log.Printf("name on resource %v", rs.Primary.Attributes["name"])
		return nil
	}
}
