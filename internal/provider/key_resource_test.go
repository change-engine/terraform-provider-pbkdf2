package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccKeyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKeyResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbkdf2_key.test", "password", "one"),
					resource.TestCheckResourceAttr("pbkdf2_key.test", "iterations", "100000"),
				),
			},
			{
				Config: testAccKeyResourceConfig("two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pbkdf2_key.test", "password", "two"),
				),
			},
		},
	})
}

func testAccKeyResourceConfig(password string) string {
	return fmt.Sprintf(`
resource "pbkdf2_key" "test" {
  password = %[1]q
}
`, password)
}