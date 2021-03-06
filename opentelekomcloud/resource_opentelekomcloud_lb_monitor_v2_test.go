package opentelekomcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/monitors"
)

func TestAccLBV2Monitor_basic(t *testing.T) {
	var monitor monitors.Monitor

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV2MonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: TestAccLBV2MonitorConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBV2MonitorExists("opentelekomcloud_lb_monitor_v2.monitor_1", &monitor),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "monitor_port", "112"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "delay", "20"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "timeout", "10"),
				),
			},
			{
				Config: TestAccLBV2MonitorConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "name", "monitor_1_updated"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "delay", "30"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "timeout", "15"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "monitor_port", "120"),
				),
			},
		},
	})
}

func TestAccLBV2Monitor_minConfig(t *testing.T) {
	var monitor monitors.Monitor

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV2MonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: TestAccLBV2MonitorConfig_minConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBV2MonitorExists("opentelekomcloud_lb_monitor_v2.monitor_1", &monitor),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "delay", "20"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "timeout", "10"),
				),
			},
			{
				Config: TestAccLBV2MonitorConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "name", "monitor_1_updated"),
					resource.TestCheckResourceAttr("opentelekomcloud_lb_monitor_v2.monitor_1", "monitor_port", "120"),
				),
			},
		},
	})
}

func testAccCheckLBV2MonitorDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkingClient, err := config.networkingV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating OpenTelekomCloud networking client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_lb_monitor_v2" {
			continue
		}

		_, err := monitors.Get(networkingClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("monitor still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckLBV2MonitorExists(n string, monitor *monitors.Monitor) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkingClient, err := config.networkingV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating OpenTelekomCloud networking client: %s", err)
		}

		found, err := monitors.Get(networkingClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("monitor not found")
		}

		*monitor = *found

		return nil
	}
}

var TestAccLBV2MonitorConfig_basic = fmt.Sprintf(`
resource "opentelekomcloud_lb_loadbalancer_v2" "loadbalancer_1" {
  name          = "loadbalancer_1"
  vip_subnet_id = "%s"
}

resource "opentelekomcloud_lb_listener_v2" "listener_1" {
  name            = "listener_1"
  protocol        = "HTTP"
  protocol_port   = 8080
  loadbalancer_id = opentelekomcloud_lb_loadbalancer_v2.loadbalancer_1.id
}

resource "opentelekomcloud_lb_pool_v2" "pool_1" {
  name        = "pool_1"
  protocol    = "HTTP"
  lb_method   = "ROUND_ROBIN"
  listener_id = opentelekomcloud_lb_listener_v2.listener_1.id
}

resource "opentelekomcloud_lb_monitor_v2" "monitor_1" {
  name         = "monitor_1"
  type         = "TCP"
  delay        = 20
  timeout      = 10
  max_retries  = 5
  pool_id      = opentelekomcloud_lb_pool_v2.pool_1.id
  monitor_port = 112

  timeouts {
    create = "5m"
    update = "5m"
    delete = "5m"
  }
}
`, OS_SUBNET_ID)

var TestAccLBV2MonitorConfig_update = fmt.Sprintf(`
resource "opentelekomcloud_lb_loadbalancer_v2" "loadbalancer_1" {
  name          = "loadbalancer_1"
  vip_subnet_id = "%s"
}

resource "opentelekomcloud_lb_listener_v2" "listener_1" {
  name            = "listener_1"
  protocol        = "HTTP"
  protocol_port   = 8080
  loadbalancer_id = opentelekomcloud_lb_loadbalancer_v2.loadbalancer_1.id
}

resource "opentelekomcloud_lb_pool_v2" "pool_1" {
  name        = "pool_1"
  protocol    = "HTTP"
  lb_method   = "ROUND_ROBIN"
  listener_id = opentelekomcloud_lb_listener_v2.listener_1.id
}

resource "opentelekomcloud_lb_monitor_v2" "monitor_1" {
  name           = "monitor_1_updated"
  type           = "TCP"
  delay          = 30
  timeout        = 15
  max_retries    = 10
  admin_state_up = "true"
  pool_id        = opentelekomcloud_lb_pool_v2.pool_1.id
  monitor_port   = 120

  timeouts {
    create = "5m"
    update = "5m"
    delete = "5m"
  }
}
`, OS_SUBNET_ID)

var TestAccLBV2MonitorConfig_minConfig = fmt.Sprintf(`
resource "opentelekomcloud_lb_loadbalancer_v2" "loadbalancer_1" {
  name          = "loadbalancer_1"
  vip_subnet_id = "%s"
}

resource "opentelekomcloud_lb_listener_v2" "listener_1" {
  name            = "listener_1"
  protocol        = "HTTP"
  protocol_port   = 8080
  loadbalancer_id = opentelekomcloud_lb_loadbalancer_v2.loadbalancer_1.id
}

resource "opentelekomcloud_lb_pool_v2" "pool_1" {
  name        = "pool_1"
  protocol    = "HTTP"
  lb_method   = "ROUND_ROBIN"
  listener_id = opentelekomcloud_lb_listener_v2.listener_1.id
}

resource "opentelekomcloud_lb_monitor_v2" "monitor_1" {
  type         = "TCP"
  delay        = 20
  timeout      = 10
  max_retries  = 5
  pool_id      = opentelekomcloud_lb_pool_v2.pool_1.id
}
`, OS_SUBNET_ID)
