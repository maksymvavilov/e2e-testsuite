# Inactive 
It was decided that at the moment such test suite is not needed. 


## e2e-testsuite
POC of text-driven test suite.
Uses [Cucumber](https://cucumber.io/docs/cucumber/) and it's implementation in GO - [Godog](https://github.com/cucumber/godog)

### Why
This allows for, hopefully, any person to understand what is being tested and what cases we have. 
This is a Proof Of a Concept of a test approach and not a design of the suite. So, please, ignore for a moment decisions like "how do we deploy workload" and "we don't change resources amid the test"

Here, the test declaration and test implementation are split. 
The declaration can be found in the `tests/dnspoliy/cases`.
The implementation is in the `tests/dnspolicy/*.go`

Note that the declaration is language agnostic and can be written in other languages
(such a python) even within the same test suite. 

### How
This suite expects you to have already kuadrant running. 
Refer to the [kuadrant-operator](https://github.com/Kuadrant/kuadrant-operator) on how to install. 
It uses Kustomize to deploy an initial set of resources on the cluster.
Such a way of deploying should make it more "readable" and ease the understanding of the initial setup. 
The test execution will firstly verify the installation and then will run tests against it. 

1. Run `make deploy-env` to deploy "starting" files 
2. Run `make test` to execute tests 

Note that the suite will not clean up after execution.
