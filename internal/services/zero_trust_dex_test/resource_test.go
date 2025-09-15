package zero_trust_dex_test_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDeviceDexTest_Traceroute(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dex_test.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	sharedChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "interval", "0h30m0s"),
		resource.TestCheckResourceAttr(name, "enabled", "true"),
		resource.TestCheckResourceAttr(name, "data.kind", "traceroute"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create test
			{
				Config: testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd, "dash.cloudflare.com", "My Test"),
				Check: resource.ComposeTestCheckFunc(
					append(sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "My Test"),
						resource.TestCheckResourceAttr(name, "data.host", "dash.cloudflare.com"))...,
				),
			},
			// Update test in place
			{
				Config: testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd, "dash.cloudflare.com", "My Test Updated"),
				Check: resource.ComposeTestCheckFunc(
					append(
						sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "My Test Updated"),
						resource.TestCheckResourceAttr(name, "data.host", "dash.cloudflare.com"),
					)...,
				),
			},
			// Update test with replace
			{
				Config: testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd, "1.1.1.1", "My Test Updated"),
				Check: resource.ComposeTestCheckFunc(
					append(
						sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "My Test Updated"),
						resource.TestCheckResourceAttr(name, "data.host", "1.1.1.1"),
					)...,
				),
			},
			// import resource
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareDeviceDexTest_HTTP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dex_test.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	sharedChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "interval", "0h30m0s"),
		resource.TestCheckResourceAttr(name, "enabled", "true"),
		resource.TestCheckResourceAttr(name, "data.kind", "http"),
		resource.TestCheckResourceAttr(name, "data.method", "GET"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceDexTestsHttp(accountID, rnd, "https://dash.cloudflare.com/home", "My Test"),
				Check: resource.ComposeTestCheckFunc(
					append(sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "My Test"),
						resource.TestCheckResourceAttr(name, "data.host", "https://dash.cloudflare.com/home"))...,
				),
			},
			// Update test in place
			{
				Config: testAccCloudflareDeviceDexTestsHttp(accountID, rnd, "https://dash.cloudflare.com/home", "My Test Updated"),
				Check: resource.ComposeTestCheckFunc(
					append(
						sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "My Test Updated"),
						resource.TestCheckResourceAttr(name, "data.host", "https://dash.cloudflare.com/home"),
					)...,
				),
			},
			{
				Config: testAccCloudflareDeviceDexTestsHttp(accountID, rnd, "https://one.dash.cloudflare.com/home/quick-start", "My Test Updated"),
				Check: resource.ComposeTestCheckFunc(
					append(
						sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "My Test Updated"),
						resource.TestCheckResourceAttr(name, "data.host", "https://one.dash.cloudflare.com/home/quick-start"),
					)...,
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCloudflareDeviceDexTestsHttp(accountID, rnd, host, description string) string {
	return acctest.LoadTestCase("devicedextestshttp.tf", rnd, accountID, host, description)
}

func testAccCloudflareDeviceDexTestsTraceroute(accountID, rnd, target, description string) string {
	return acctest.LoadTestCase("devicedexteststraceroute.tf", rnd, accountID, target, description)
}
