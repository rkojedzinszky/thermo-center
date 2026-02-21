
# PatchSensorRequest


## Properties

Name | Type
------------ | -------------
`vcc` | number
`rssi` | number
`lqi` | number
`interval` | number
`name` | string
`lastTsf` | number

## Example

```typescript
import type { PatchSensorRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "vcc": null,
  "rssi": null,
  "lqi": null,
  "interval": null,
  "name": null,
  "lastTsf": null,
} satisfies PatchSensorRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PatchSensorRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


