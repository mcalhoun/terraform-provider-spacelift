package spacelift

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	. "github.com/spacelift-io/terraform-provider-spacelift/spacelift/internal/testhelpers"
)

func TestModuleData(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	testSteps(t, []resource.TestStep{{
		Config: fmt.Sprintf(`
			resource "spacelift_module" "test" {
                name            = "test-module-%s"
				administrative  = true
				branch          = "master"
				description     = "description"
				labels          = ["one", "two"]
				repository      = "terraform-bacon-tasty"
				shared_accounts = ["foo-subdomain", "bar-subdomain"]
			}
			data "spacelift_module" "test" {
				module_id = spacelift_module.test.id
			}
		`, randomID),
		Check: Resource(
			"data.spacelift_module.test",
			Attribute("id", Equals(fmt.Sprintf("test-module-%s", randomID))),
			Attribute("administrative", Equals("true")),
			Attribute("branch", Equals("master")),
			Attribute("description", Equals("description")),
			SetEquals("labels", "one", "two"),
			Attribute("name", Equals(fmt.Sprintf("test-module-%s", randomID))),
			Attribute("project_root", Equals("")),
			Attribute("repository", Equals("terraform-bacon-tasty")),
			SetEquals("shared_accounts", "bar-subdomain", "foo-subdomain"),
			Attribute("terraform_provider", Equals("default")),
		),
	}})
}

func TestModuleDataSpace(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	testSteps(t, []resource.TestStep{{
		Config: fmt.Sprintf(`
			resource "spacelift_module" "test" {
                name            = "test-module-%s"
				administrative  = true
				branch          = "master"
				description     = "description"
				labels          = ["one", "two"]
				repository      = "terraform-bacon-tasty"
				shared_accounts = ["foo-subdomain", "bar-subdomain"]
				space_id        = "root"
			}

			data "spacelift_module" "test" {
				module_id = spacelift_module.test.id
			}
		`, randomID),
		Check: Resource(
			"data.spacelift_module.test",
			Attribute("id", Equals(fmt.Sprintf("test-module-%s", randomID))),
			Attribute("administrative", Equals("true")),
			Attribute("branch", Equals("master")),
			Attribute("description", Equals("description")),
			SetEquals("labels", "one", "two"),
			Attribute("name", Equals(fmt.Sprintf("test-module-%s", randomID))),
			Attribute("project_root", Equals("")),
			Attribute("repository", Equals("terraform-bacon-tasty")),
			Attribute("space_id", Equals("root")),
			SetEquals("shared_accounts", "bar-subdomain", "foo-subdomain"),
			Attribute("terraform_provider", Equals("default")),
		),
	}})
}
