Feature: on-cluster healthchecks
  We should create a probe CR when we define heatcheck spec
  We should not have probes for wildcard listeners

  Scenario: Healtcheckprobes are created
    Given there is a DNSPolicy dnspolicy-sample
    And there is a Gateway test
    When DNSPolicy dnspolicy-sample targets Gateway test
    And DNSPolicy dnspolicy-sample has healthchecks defined
    Then we should have DNSRecord test-api not healthy
    And DNSHealthCheckProbe for DNSRecord test-api should exist
    And we should have DNSRecord test-app healthy
    And DNSHealthCheckProbe for DNSRecord test-app should not exist