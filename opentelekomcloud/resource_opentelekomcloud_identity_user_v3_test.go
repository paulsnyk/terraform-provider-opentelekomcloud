package opentelekomcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/users"
)

func TestAccIdentityV3User_basic(t *testing.T) {
	var user users.User
	var userName = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAdminOnly(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIdentityV3UserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityV3User_basic(userName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityV3UserExists("opentelekomcloud_identity_user_v3.user_1", &user),
					resource.TestCheckResourceAttrPtr(
						"opentelekomcloud_identity_user_v3.user_1", "name", &user.Name),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_identity_user_v3.user_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_identity_user_v3.user_1", "email", "test@acme.org"),
				),
			},
			{
				Config: testAccIdentityV3User_update(userName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityV3UserExists("opentelekomcloud_identity_user_v3.user_1", &user),
					resource.TestCheckResourceAttrPtr(
						"opentelekomcloud_identity_user_v3.user_1", "name", &user.Name),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_identity_user_v3.user_1", "enabled", "false"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_identity_user_v3.user_1", "email", "test2@acme.org"),
				),
			},
		},
	})
}

func testAccCheckIdentityV3UserDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	identityClient, err := config.identityV3Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating Opentelekomcloud identity client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_identity_user_v3" {
			continue
		}

		_, err := users.Get(identityClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("User still exists")
		}
	}

	return nil
}

func testAccCheckIdentityV3UserExists(n string, user *users.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		identityClient, err := config.identityV3Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating Opentelekomcloud identity client: %s", err)
		}

		found, err := users.Get(identityClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("User not found")
		}

		*user = *found

		return nil
	}
}

func testAccIdentityV3User_basic(userName string) string {
	return fmt.Sprintf(`
    resource "opentelekomcloud_identity_user_v3" "user_1" {
      name = "%s"
      password = "password123@!"
      enabled = true
      email = "test@acme.org"
    }
  `, userName)
}

func testAccIdentityV3User_update(userName string) string {
	return fmt.Sprintf(`
    resource "opentelekomcloud_identity_user_v3" "user_1" {
      name = "%s"
      enabled = false
      password = "password123@!"
      email = "tEst2@acme.org"
    }
  `, userName)
}
