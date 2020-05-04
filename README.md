AdrestiaAssertions
==================

## Objectives:
    * Provide a simple assertion that can be used at runtime.
    
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