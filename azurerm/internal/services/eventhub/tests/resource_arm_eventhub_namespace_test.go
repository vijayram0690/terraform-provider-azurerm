package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMEventHubNamespace_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMEventHubNamespace_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			{
				Config:      testAccAzureRMEventHubNamespace_requiresImport(data),
				ExpectError: acceptance.RequiresImportError("azurerm_eventhub_namespace"),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_standard(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_standard(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMEventHubNamespace_networkrule_iprule(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_networkrule_iprule(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMEventHubNamespace_networkrule_vnet(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_networkrule_vnet(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMEventHubNamespace_readDefaultKeys(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestMatchResourceAttr(data.ResourceName, "default_primary_connection_string", regexp.MustCompile("Endpoint=.+")),
					resource.TestMatchResourceAttr(data.ResourceName, "default_secondary_connection_string", regexp.MustCompile("Endpoint=.+")),
					resource.TestMatchResourceAttr(data.ResourceName, "default_primary_key", regexp.MustCompile(".+")),
					resource.TestMatchResourceAttr(data.ResourceName, "default_secondary_key", regexp.MustCompile(".+")),
				),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_maximumThroughputUnits(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_maximumThroughputUnits(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMEventHubNamespace_NonStandardCasing(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	config := testAccAzureRMEventHubNamespaceNonStandardCasing(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists("azurerm_eventhub_namespace.test"),
				),
			},
			{
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_BasicWithTagsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")
	preConfig := testAccAzureRMEventHubNamespace_basic(data)
	postConfig := testAccAzureRMEventHubNamespace_basicWithTagsUpdate(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: preConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "1"),
				),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_BasicWithCapacity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")
	config := testAccAzureRMEventHubNamespace_capacity(data, 20)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "capacity", "20"),
				),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_BasicWithCapacityUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")
	preConfig := testAccAzureRMEventHubNamespace_capacity(data, 20)
	postConfig := testAccAzureRMEventHubNamespace_capacity(data, 2)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: preConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "capacity", "20"),
				),
			},
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "capacity", "2"),
				),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_BasicWithSkuUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")
	preConfig := testAccAzureRMEventHubNamespace_basic(data)
	postConfig := testAccAzureRMEventHubNamespace_standard(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: preConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "sku", "Basic"),
				),
			},
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "sku", "Standard"),
					resource.TestCheckResourceAttr(data.ResourceName, "capacity", "2"),
				),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_maximumThroughputUnitsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")
	preConfig := testAccAzureRMEventHubNamespace_maximumThroughputUnits(data)
	postConfig := testAccAzureRMEventHubNamespace_maximumThroughputUnitsUpdate(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: preConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "sku", "Standard"),
					resource.TestCheckResourceAttr(data.ResourceName, "capacity", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "maximum_throughput_units", "20"),
				),
			},
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "sku", "Standard"),
					resource.TestCheckResourceAttr(data.ResourceName, "capacity", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "maximum_throughput_units", "1"),
				),
			},
		},
	})
}

func TestAccAzureRMEventHubNamespace_autoInfalteDisabledWithAutoInflateUnits(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventhub_namespace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMEventHubNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventHubNamespace_autoInfalteDisabledWithAutoInflateUnits(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventHubNamespaceExists(data.ResourceName),
				),
			},
		},
	})
}

func testCheckAzureRMEventHubNamespaceDestroy(s *terraform.State) error {
	conn := acceptance.AzureProvider.Meta().(*clients.Client).Eventhub.NamespacesClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_eventhub_namespace" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		resp, err := conn.Get(ctx, resourceGroup, name)

		if err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return err
			}
		}
	}

	return nil
}

func testCheckAzureRMEventHubNamespaceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		namespaceName := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Event Hub Namespace: %s", namespaceName)
		}

		conn := acceptance.AzureProvider.Meta().(*clients.Client).Eventhub.NamespacesClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		resp, err := conn.Get(ctx, resourceGroup, namespaceName)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Event Hub Namespace %q (resource group: %q) does not exist", namespaceName, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on eventHubNamespacesClient: %+v", err)
		}

		return nil
	}
}

func testAccAzureRMEventHubNamespace_basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Basic"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMEventHubNamespace_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_eventhub_namespace" "import" {
  name                = "${azurerm_eventhub_namespace.test.name}"
  location            = "${azurerm_eventhub_namespace.test.location}"
  resource_group_name = "${azurerm_eventhub_namespace.test.resource_group_name}"
  sku                 = "${azurerm_eventhub_namespace.test.sku}"
}
`, template)
}

func testAccAzureRMEventHubNamespace_standard(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Standard"
  capacity            = "2"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_networkrule_iprule(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Standard"
  capacity            = "2"

  network_rulesets {
    default_action = "Deny"
    ip_rule {
      ip_mask = "10.0.0.0/16"
    }
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_networkrule_vnet(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctvn-%[1]d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctsub-%[1]d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%[1]d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Standard"
  capacity            = "2"

  network_rulesets {
    default_action = "Deny"
    virtual_network_rule {
      subnet_id = "${azurerm_subnet.test.id}"

      ignore_missing_virtual_network_service_endpoint = true
    }
  }
}
`, data.RandomInteger, data.Locations.Primary)
}

func testAccAzureRMEventHubNamespaceNonStandardCasing(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "basic"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_maximumThroughputUnits(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                     = "acctesteventhubnamespace-%d"
  location                 = "${azurerm_resource_group.test.location}"
  resource_group_name      = "${azurerm_resource_group.test.name}"
  sku                      = "Standard"
  capacity                 = "2"
  auto_inflate_enabled     = true
  maximum_throughput_units = 20
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_basicWithTagsUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Basic"

  tags = {
    environment = "Production"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_capacity(data acceptance.TestData, capacity int) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Basic"
  capacity            = %d
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, capacity)
}

func testAccAzureRMEventHubNamespace_maximumThroughputUnitsUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                     = "acctesteventhubnamespace-%d"
  location                 = "${azurerm_resource_group.test.location}"
  resource_group_name      = "${azurerm_resource_group.test.name}"
  sku                      = "Standard"
  capacity                 = 1
  auto_inflate_enabled     = true
  maximum_throughput_units = 1
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMEventHubNamespace_autoInfalteDisabledWithAutoInflateUnits(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                     = "acctesteventhubnamespace-%d"
  location                 = "${azurerm_resource_group.test.location}"
  resource_group_name      = "${azurerm_resource_group.test.name}"
  sku                      = "Standard"
  capacity                 = 1
  auto_inflate_enabled     = false
  maximum_throughput_units = 0
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}
