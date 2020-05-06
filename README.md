AdrestiaAssertions
==================

## Objectives:
    * Provide a simple assertion that can be used at runtime.

## Status:
[![Build Status](https://dev.azure.com/asymmetric-effort/open%20source%20projects/_apis/build/status/sam-caldwell.adrestia-assertions?branchName=master)](https://dev.azure.com/asymmetric-effort/open%20source%20projects/_build/latest?definitionId=5&branchName=master)


## Methods
```golang
AdrestiaAssert.Panic(condition, response)
```
* Evaluates condition.
* If true, panic and write response string.

```golang
AdrestiaAssert.Error(condition, response)
```
* Evaluates condition.
* Returns error if true, otherwise nil.
