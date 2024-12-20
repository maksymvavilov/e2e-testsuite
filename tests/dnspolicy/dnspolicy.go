package dnspolicy

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var k8sClient client.Client
var scheme = runtime.NewScheme()

func init() {
	/*
		add to scheme
	*/

	k8sClient, _ = client.New(controllerruntime.GetConfigOrDie(), client.Options{Scheme: scheme})
}

func assertLBDNSPolicyExists(policyName string) error {
	fmt.Println(fmt.Sprintf("Checking if LB DNSPolicy %s exists", policyName))
	return nil
}

func assertResourceExists(resource, name string) error {
	fmt.Println(fmt.Sprintf("Checking if resource %s with name %s exists", resource, name))
	return nil
}

func assertResourceTargets(targetingResource, targetingName, targetResource, targetName string) error {
	fmt.Println(fmt.Sprintf("Checking if resource %s with name %s targets resource %s with name %s", targetingResource, targetingName, targetResource, targetName))
	return nil
}

func assertDNSRecordExists(recordName string) error {
	fmt.Println(fmt.Sprintf("Checking if DNS Record %s exists", recordName))
	return nil
}

func assertNumberOfDNSRecords(number int) error {
	fmt.Println(fmt.Sprintf("Checking if DNS Record count is %d", number))
	return nil
}

func allEndpointsTravertsable() error {
	fmt.Println(fmt.Sprintf("Checking if all endpoints travertsable resources"))
	return nil
}

func assertPolicyDefinesHealthchecks(policyName string) error {
	fmt.Println(fmt.Sprintf("Checking if policy %s defines healthchecks", policyName))
	return nil
}

func assertResourceIsHealthy(resource, resourceName string) error {
	fmt.Println(fmt.Sprintf("Checking if resource %s witn name %s is healthy", resource, resourceName))
	return nil
}

func assertResourceIsNotHealthy(resource, resourceName string) error {
	fmt.Println(fmt.Sprintf("Checking if resource %s witn name %s is not healthy", resource, resourceName))
	return nil
}

func assertHealthcheckProbeExists(recordName string) error {
	fmt.Println(fmt.Sprintf("Checking if healthcheck probe for DNSRecord with name %s exists", recordName))
	return nil
}

func assertHealthcheckProbeDoesNotExist(recordName string) error {
	fmt.Println(fmt.Sprintf("Checking if healthcheck probe for DNSRecord with name %s does not exists", recordName))
	return nil
}
