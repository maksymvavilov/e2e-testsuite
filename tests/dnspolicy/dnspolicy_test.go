package dnspolicy

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{
	Output:      colors.Colored(os.Stdout),
	Concurrency: 4,
	Paths:       []string{"cases"},
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                 "DNSPolicy",
		Options:              &o,
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// We could create CRs here from fixtures
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Before each test case
		return ctx, nil
	})

	// shared
	ctx.Given(`^there is a lb DNSPolicy ([^"]*)$`, assertLBDNSPolicyExists)

	// loadbalanced.feature
	ctx.Given(`^there is a (DNSPolicy|Gateway) ([^"]*)$`, assertResourceExists)
	ctx.When(`^(DNSPolicy|HTTPRoute) ([^"]*) targets (Gateway|Secret) ([^"]*)$`, assertResourceTargets)
	ctx.Then(`^we should have ([^"]*) DNSRecord$`, assertDNSRecordExists)
	ctx.Then(`^we should have (\d+) DNSRecords$`, assertNumberOfDNSRecords)
	ctx.Step("endpoints should be traversable", allEndpointsTravertsable)

	// healthchecks.feature
	ctx.When(`DNSPolicy ([^\"]*) has healthchecks defined`, assertPolicyDefinesHealthchecks)
	ctx.Then(`we should have (DNSRecord|DNSPolicy|Gateway) ([^\"]*) healthy`, assertResourceIsHealthy)
	ctx.Then(`we should have (DNSRecord|DNSPolicy|Gateway) ([^\"]*) not healthy`, assertResourceIsNotHealthy)
	ctx.Then(`DNSHealthCheckProbe for DNSRecord ([^\"]*) should exist`, assertHealthcheckProbeExists)
	ctx.Then(`DNSHealthCheckProbe for DNSRecord ([^\"]*) should not exist`, assertHealthcheckProbeDoesNotExist)
}
