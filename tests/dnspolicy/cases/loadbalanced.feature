Feature: loadbalanced routing strategy
  DNSPolicy with loadbalanced strategy
  Should resul in the creation of DNSRecords per listener
  Endpoints on DNSRecords should be traversable

  Scenario: LB strategy creates DNSRecords
    Given there is a lb DNSPolicy dnspolicy-sample
    And there is a Gateway test
    When DNSPolicy dnspolicy-sample targets Gateway test
    Then we should have test-api DNSRecord
    And we should have test-app DNSRecord

  Scenario: All endpoints should be traversable
    Given there is a lb DNSPolicy dnspolicy-sample
    And there is a Gateway test
    When DNSPolicy dnspolicy-sample targets Gateway test
    Then we should have 2 DNSRecords
    And endpoints should be traversable
