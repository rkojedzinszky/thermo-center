import { usePersistentString } from '@/composables/usePersistentRef'

export type ProfileEditMode = 'temperature' | 'structure'

const MODE_KEY = 'profile_edit_mode'

function isMode(value: string): value is ProfileEditMode {
  return value === 'temperature' || value === 'structure'
}

const { state: mode, set: setMode } = usePersistentString<ProfileEditMode>(
  MODE_KEY,
  'temperature',
  isMode,
)

export function useProfileEditMode() {
  return { mode, setMode }
}
