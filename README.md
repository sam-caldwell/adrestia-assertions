AdrestiaAssertions
==================

## Objectives:
    * Provide a simple assertion that can be used at runtime.
    
## Methods
```golang
AssertPanic(condition, response)
```
* Evaluates condition.
* If true, panic and write response string.
```golang
AssertError(condition, response)
```
* Evaluates condition.
* Returns error if true, otherwise nil.