package icinga2api

import (
	"fmt"
	"testing"

	"github.com/cnanaaron/terraform-provider-icinga2api/iapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCreateBasicHostGroup(t *testing.T) {

	var testAccCreateBasicHostGroup = fmt.Sprintf(`
		resource "icinga2_hostgroup" "tf-hg-1" {
		  name = "terraform-hostgroup-1"
		  display_name = "Terraform Test HostGroup"
	}`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateBasicHostGroup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists("icinga2_hostgroup.tf-hg-1"),
					testAccCheckResourceState("icinga2_hostgroup.tf-hg-1", "name", "terraform-hostgroup-1"),
					testAccCheckResourceState("icinga2_hostgroup.tf-hg-1", "display_name", "Terraform Test HostGroup"),
				),
			},
		},
	})
}

func testAccCheckHostgroupExists(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("Hostgroup resource not found: %s", rn)
		}

		if resource.Primary.ID == "" {
			return fmt.Errorf("Hostgroup resource id not set")
		}

		client := testAccProvider.Meta().(*iapi.Server)
		_, err := client.GetHostgroup(resource.Primary.ID)
		if err != nil {
			return fmt.Errorf("Error getting getting hostgroup: %s", err)
		}

		return nil
	}
}
