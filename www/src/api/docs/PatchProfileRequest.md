
# PatchProfileRequest


## Properties

Name | Type
------------ | -------------
`control` | string
`daytype` | string
`start` | string
`targetTemp` | number

## Example

```typescript
import type { PatchProfileRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "control": null,
  "daytype": null,
  "start": null,
  "targetTemp": null,
} satisfies PatchProfileRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PatchProfileRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


