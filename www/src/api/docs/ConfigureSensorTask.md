
# ConfigureSensorTask


## Properties

Name | Type
------------ | -------------
`created` | Date
`started` | Date
`firstDiscovery` | Date
`lastDiscovery` | Date
`finished` | Date
`error` | string
`id` | number
`sensorId` | number
`sensorName` | string

## Example

```typescript
import type { ConfigureSensorTask } from ''

// TODO: Update the object below with actual values
const example = {
  "created": null,
  "started": null,
  "firstDiscovery": null,
  "lastDiscovery": null,
  "finished": null,
  "error": null,
  "id": null,
  "sensorId": null,
  "sensorName": null,
} satisfies ConfigureSensorTask

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ConfigureSensorTask
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


