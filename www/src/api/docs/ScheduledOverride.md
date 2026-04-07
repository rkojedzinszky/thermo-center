
# ScheduledOverride


## Properties

Name | Type
------------ | -------------
`resourceUri` | string
`id` | number
`control` | string
`start` | Date
`end` | Date
`targetTemp` | number

## Example

```typescript
import type { ScheduledOverride } from ''

// TODO: Update the object below with actual values
const example = {
  "resourceUri": null,
  "id": null,
  "control": null,
  "start": null,
  "end": null,
  "targetTemp": null,
} satisfies ScheduledOverride

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ScheduledOverride
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


