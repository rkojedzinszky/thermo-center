import type { Calendar, Control, DayType, Profile } from '../api'

export const CONTROL_URI = '/api/v1/control/1/'
export const DAYTYPE_URI = '/api/v1/daytype/1/'

export function makeCalendar(id: number, day: string, daytypeUri: string = DAYTYPE_URI): Calendar {
  return {
    id,
    resourceUri: `/api/v1/calendar/${id}/`,
    day: new Date(`${day}T00:00:00.000Z`),
    daytype: daytypeUri,
  }
}

export function makeProfile(
  id: number,
  start: string,
  targetTemp: number | null = 20,
  overrides: Partial<Profile> = {},
): Profile {
  return {
    resourceUri: `/api/v1/profile/${id}/`,
    id,
    control: CONTROL_URI,
    daytype: DAYTYPE_URI,
    start,
    targetTemp,
    ...overrides,
  }
}

export function makeControl(id: number, name: string, overrides: Partial<Control> = {}): Control {
  return {
    resourceUri: `/api/v1/control/${id}/`,
    sensorId: id,
    name,
    id,
    kp: 1,
    ki: 0,
    kd: 0,
    temperature: 21.5,
    targetTemp: 22,
    ...overrides,
  }
}

export function makeDayType(id: number, name: string): DayType {
  return { resourceUri: `/api/v1/daytype/${id}/`, id, name }
}
