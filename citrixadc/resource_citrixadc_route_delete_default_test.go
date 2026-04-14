package citrixadc

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Test 1: Positive case - add route with delete_default_route=true, verify default route is deleted
func TestAccRoute_delete_default_route(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRouteDeleteDefaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute_delete_default_true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "delete_default_route", "true"),
					resource.TestCheckResourceAttrSet("citrixadc_route.test_route", "original_default_gateway"),
					testAccCheckDefaultRouteAbsent(),
				),
			},
		},
	})
}

// Test 2: Destroy restores default route - add route with delete_default_route=true,
// then destroy and verify default route is restored
func TestAccRoute_delete_default_route_restore_on_destroy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRouteDeleteDefaultRestored,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute_delete_default_true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "delete_default_route", "true"),
					testAccCheckDefaultRouteAbsent(),
				),
			},
		},
	})
}

// Test 3: Update mutable attribute (distance) without affecting delete_default_route
func TestAccRoute_update_distance_preserves_delete_default(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRouteDeleteDefaultRestored,
		Steps: []resource.TestStep{
			// Step 1: Create route with delete_default_route=true
			{
				Config: testAccRoute_delete_default_true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "delete_default_route", "true"),
					testAccCheckDefaultRouteAbsent(),
				),
			},
			// Step 2: Update distance to 3, delete_default_route stays true
			{
				Config: testAccRoute_delete_default_true_update_distance,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "distance", "3"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "delete_default_route", "true"),
					resource.TestCheckResourceAttrSet("citrixadc_route.test_route", "original_default_gateway"),
					testAccCheckDefaultRouteAbsent(),
				),
			},
		},
	})
}

// Test 4: Route with delete_default_route=false (default behavior, no default route deletion)
func TestAccRoute_delete_default_route_false(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRouteDestroyedGeneric,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute_delete_default_false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route_no_delete"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route_no_delete", "delete_default_route", "false"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route_no_delete", "original_default_gateway", ""),
					testAccCheckDefaultRoutePresent(),
				),
			},
		},
	})
}

// Test 5: Route without specifying delete_default_route (defaults to false)
func TestAccRoute_delete_default_route_omitted(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRouteDestroyedGeneric,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute_delete_default_omitted,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route_omitted"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route_omitted", "delete_default_route", "false"),
					testAccCheckDefaultRoutePresent(),
				),
			},
		},
	})
}

// Test 6: Changing delete_default_route from true to false forces recreate
func TestAccRoute_delete_default_route_force_new(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRouteDeleteDefaultRestored,
		Steps: []resource.TestStep{
			// Step 1: Create with delete_default_route=true
			{
				Config: testAccRoute_delete_default_true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "delete_default_route", "true"),
					testAccCheckDefaultRouteAbsent(),
				),
			},
			// Step 2: Change delete_default_route to false - should force destroy+recreate
			// The destroy will restore the default route, then create without deleting it
			{
				Config: testAccRoute_delete_default_changed_to_false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.test_route"),
					resource.TestCheckResourceAttr("citrixadc_route.test_route", "delete_default_route", "false"),
					testAccCheckDefaultRoutePresent(),
				),
			},
		},
	})
}

// --- Test Configs ---

const testAccRoute_delete_default_true = `
resource "citrixadc_route" "test_route" {
  network              = "192.168.10.0"
  netmask              = "255.255.255.0"
  gateway              = "10.102.126.1"
  delete_default_route = true
}
`

const testAccRoute_delete_default_true_update_distance = `
resource "citrixadc_route" "test_route" {
  network              = "192.168.10.0"
  netmask              = "255.255.255.0"
  gateway              = "10.102.126.1"
  distance             = 3
  delete_default_route = true
}
`

const testAccRoute_delete_default_false = `
resource "citrixadc_route" "test_route_no_delete" {
  network              = "192.168.20.0"
  netmask              = "255.255.255.0"
  gateway              = "10.102.126.1"
  delete_default_route = false
}
`

const testAccRoute_delete_default_omitted = `
resource "citrixadc_route" "test_route_omitted" {
  network = "192.168.30.0"
  netmask = "255.255.255.0"
  gateway = "10.102.126.1"
}
`

const testAccRoute_delete_default_changed_to_false = `
resource "citrixadc_route" "test_route" {
  network              = "192.168.10.0"
  netmask              = "255.255.255.0"
  gateway              = "10.102.126.1"
  delete_default_route = false
}
`

// --- Check Functions ---

// testAccCheckRouteExist verifies the route exists on the NetScaler
func testAccCheckRouteExist(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("route resource not found in state: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("route resource ID is not set")
		}

		client, err := testAccGetClient()
		if err != nil {
			return err
		}

		findParams := service.FindParams{
			ResourceType: service.Route.Type(),
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return fmt.Errorf("error fetching routes: %s", err.Error())
		}

		network := rs.Primary.Attributes["network"]
		netmask := rs.Primary.Attributes["netmask"]
		gateway := rs.Primary.Attributes["gateway"]

		for _, r := range dataArray {
			if r["network"] == network && r["netmask"] == netmask && r["gateway"] == gateway {
				return nil
			}
		}
		return fmt.Errorf("route %s/%s via %s not found on NetScaler", network, netmask, gateway)
	}
}

// testAccCheckDefaultRouteAbsent verifies the default route (0.0.0.0/0.0.0.0) does NOT exist
func testAccCheckDefaultRouteAbsent() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetClient()
		if err != nil {
			return err
		}

		findParams := service.FindParams{
			ResourceType: service.Route.Type(),
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return fmt.Errorf("error fetching routes: %s", err.Error())
		}

		for _, r := range dataArray {
			if r["network"] == "0.0.0.0" && r["netmask"] == "0.0.0.0" {
				return fmt.Errorf("default route (0.0.0.0/0.0.0.0) still exists with gateway %s, expected it to be deleted", r["gateway"])
			}
		}
		return nil
	}
}

// testAccCheckDefaultRoutePresent verifies the default route (0.0.0.0/0.0.0.0) exists
func testAccCheckDefaultRoutePresent() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetClient()
		if err != nil {
			return err
		}

		findParams := service.FindParams{
			ResourceType: service.Route.Type(),
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return fmt.Errorf("error fetching routes: %s", err.Error())
		}

		for _, r := range dataArray {
			if r["network"] == "0.0.0.0" && r["netmask"] == "0.0.0.0" {
				return nil
			}
		}
		return fmt.Errorf("default route (0.0.0.0/0.0.0.0) not found, expected it to be present")
	}
}

// testAccCheckRouteDeleteDefaultDestroy verifies route is deleted after destroy
func testAccCheckRouteDeleteDefaultDestroy(s *terraform.State) error {
	client, err := testAccGetClient()
	if err != nil {
		return err
	}

	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return fmt.Errorf("error fetching routes: %s", err.Error())
	}

	for _, r := range dataArray {
		if r["network"] == "192.168.10.0" && r["netmask"] == "255.255.255.0" {
			return fmt.Errorf("route 192.168.10.0/255.255.255.0 still exists after destroy")
		}
	}
	return nil
}

// testAccCheckRouteDeleteDefaultRestored verifies route is deleted AND default route is restored after destroy
func testAccCheckRouteDeleteDefaultRestored(s *terraform.State) error {
	client, err := testAccGetClient()
	if err != nil {
		return err
	}

	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return fmt.Errorf("error fetching routes: %s", err.Error())
	}

	// Verify the managed route is gone
	for _, r := range dataArray {
		if r["network"] == "192.168.10.0" && r["netmask"] == "255.255.255.0" {
			return fmt.Errorf("route 192.168.10.0/255.255.255.0 still exists after destroy")
		}
	}

	// Verify the default route was restored
	for _, r := range dataArray {
		if r["network"] == "0.0.0.0" && r["netmask"] == "0.0.0.0" {
			return nil
		}
	}
	return fmt.Errorf("default route (0.0.0.0/0.0.0.0) was not restored after destroy")
}

// testAccCheckRouteDestroyed returns a CheckDestroy func for a specific resource
func testAccCheckRouteDestroyedGeneric(s *terraform.State) error {
	client, err := testAccGetClient()
	if err != nil {
		return err
	}

	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return fmt.Errorf("error fetching routes: %s", err.Error())
	}

	// Look up the network from the last known state
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_route" {
			continue
		}
		network := rs.Primary.Attributes["network"]
		netmask := rs.Primary.Attributes["netmask"]
		for _, r := range dataArray {
			if r["network"] == network && r["netmask"] == netmask {
				return fmt.Errorf("route %s/%s still exists after destroy", network, netmask)
			}
		}
	}
	return nil
}
