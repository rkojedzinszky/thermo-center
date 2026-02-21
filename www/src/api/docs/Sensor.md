
# Sensor


## Properties

Name | Type
------------ | -------------
`valid` | boolean
`sensorResync` | string
`thsensor` | string
`id` | number
`vcc` | number
`rssi` | number
`lqi` | number
`interval` | number
`name` | string
`lastTsf` | number

## Example

```typescript
import type { Sensor } from ''

// TODO: Update the object below with actual values
const example = {
  "valid": null,
  "sensorResync": null,
  "thsensor": null,
  "id": null,
  "vcc": null,
  "rssi": null,
  "lqi": null,
  "interval": null,
  "name": null,
  "lastTsf": null,
} satisfies Sensor

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Sensor
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


