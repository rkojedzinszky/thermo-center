
# Control


## Properties

Name | Type
------------ | -------------
`resourceUri` | string
`sensorId` | number
`name` | string
`temperature` | number
`targetTemp` | number
`pidcontrol` | number
`age` | number
`id` | number
`kp` | number
`ki` | number
`kd` | number
`intabsmax` | number

## Example

```typescript
import type { Control } from ''

// TODO: Update the object below with actual values
const example = {
  "resourceUri": null,
  "sensorId": null,
  "name": null,
  "temperature": null,
  "targetTemp": null,
  "pidcontrol": null,
  "age": null,
  "id": null,
  "kp": null,
  "ki": null,
  "kd": null,
  "intabsmax": null,
} satisfies Control

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Control
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


