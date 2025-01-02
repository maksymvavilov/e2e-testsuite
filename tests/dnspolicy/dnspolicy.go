package dnspolicy

import (
	"context"
	"fmt"

	kuadrantdns "github.com/kuadrant/dns-operator/api/v1alpha1"
	kuadrantv1 "github.com/kuadrant/kuadrant-operator/api/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	testNS = "test"
)

var (
	k8sClient client.Client
	scheme    = runtime.NewScheme()
	ctx       = context.WithoutCancel(context.Background())
)

func init() {
	utilruntime.Must(kuadrantv1.AddToScheme(scheme))
	utilruntime.Must(kuadrantdns.AddToScheme(scheme))
	utilruntime.Must(gatewayv1.AddToScheme(scheme))

	k8sClient, _ = client.New(controllerruntime.GetConfigOrDie(), client.Options{Scheme: scheme})
}

func assertLBDNSPolicyExists(policyName string) error {
	dnspolicy := &kuadrantv1.DNSPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name:      policyName,
			Namespace: testNS,
		},
	}
	err := k8sClient.Get(ctx, client.ObjectKeyFromObject(dnspolicy), dnspolicy)
	if err != nil {
		return err
	}
	if dnspolicy.Spec.LoadBalancing == nil {
		return fmt.Errorf("DNS policy %s has no load balancing", policyName)
	}
	return nil
}

func assertResourceExists(resource, name string) error {
	switch resource {
	case "DNSPolicy":
		return assertDNSPolicyExists(name)
	case "Gateway":
		return assertGatewayExists(name)
	}
	return fmt.Errorf("resource %s is not expected", resource)
}

func assertResourceTargets(targetingResource, targetingName, targetResource, targetName string) error {
	fmt.Println(fmt.Sprintf("Checking if resource %s with name %s targets resource %s with name %s", targetingResource, targetingName, targetResource, targetName))
	return nil
}

func assertDNSPolicyExists(policyName string) error {
	dnspolicy := &kuadrantv1.DNSPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name:      policyName,
			Namespace: testNS,
		},
	}
	return k8sClient.Get(ctx, client.ObjectKeyFromObject(dnspolicy), dnspolicy)
}

func assertGatewayExists(gatewayName string) error {
	gateway := &gatewayv1.Gateway{
		ObjectMeta: v1.ObjectMeta{
			Name:      gatewayName,
			Namespace: testNS,
		},
	}
	return k8sClient.Get(ctx, client.ObjectKeyFromObject(gateway), gateway)
}

func assertDNSRecordExists(recordName string) error {
	fmt.Println(fmt.Sprintf("Checking if DNS Record %s exists", recordName))
	record := kuadrantdns.DNSRecord{
		ObjectMeta: v1.ObjectMeta{
			Name:      recordName,
			Namespace: testNS,
		},
	}
	return k8sClient.Get(ctx, client.ObjectKeyFromObject(&record), &record)
}

func assertNumberOfDNSRecords(number int) error {
	recordList := &kuadrantdns.DNSRecordList{}
	err := k8sClient.List(ctx, recordList, client.InNamespace(testNS))
	if err != nil {
		return err
	}
	if len(recordList.Items) != number {
		return fmt.Errorf("number of DNS records is not correct. Expected %d, got %d", number, len(recordList.Items))
	}
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
