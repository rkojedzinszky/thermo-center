import { usePersistentString } from '@/composables/usePersistentRef'

export type ProfileGroupAxis = 'sensor' | 'daytype'

const AXIS_KEY = 'profile_group_axis'

function isAxis(value: string): value is ProfileGroupAxis {
  return value === 'sensor' || value === 'daytype'
}

const { state: axis, set: setAxis } = usePersistentString<ProfileGroupAxis>(
  AXIS_KEY,
  'sensor',
  isAxis,
)

export function useProfileAxis() {
  return { axis, setAxis }
}
