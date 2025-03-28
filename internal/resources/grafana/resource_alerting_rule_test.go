package grafana_test

import (
	"testing"

	"github.com/grafana/grafana-openapi-client-go/models"
	"github.com/grafana/terraform-provider-grafana/v3/internal/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAlertRuleSeparate_compound(t *testing.T) {
	testutils.CheckOSSTestsEnabled(t, ">=9.1.0")

	var alertRule models.ProvisionedAlertRule

	resource.ParallelTest(t, resource.TestCase{
		ProtoV5ProviderFactories: testutils.ProtoV5ProviderFactories,
		// Implicitly tests deletion.
		CheckDestroy: alertingRuleCheckExists.destroyed(&alertRule, nil),
		Steps: []resource.TestStep{
			// Test creation.
			{
				Config: testutils.TestAccExample(t, "resources/grafana_rule/resource.tf"),
				Check: resource.ComposeTestCheckFunc(
					alertingRuleCheckExists.exists("grafana_rule.test_rule", &alertRule),
					resource.TestCheckResourceAttr("grafana_rule.test_rule", "name", "My Alert Rule"),
					testutils.CheckLister("grafana_rule.test_rule"),
				),
			},
			// Test update
			{
				Config: testutils.TestAccExampleWithReplace(t, "resources/grafana_rule/resource.tf", map[string]string{
					"My Alert Rule": "Our Alert Rule",
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("grafana_rule.test_rule", "name", "Our Alert Rule"),
				),
			},
		},
	})
}
