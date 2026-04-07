import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { InstantProfile, PatchInstantProfileRequest } from '@/api'

const profilesMap = ref<Map<number, InstantProfile>>(new Map())

const instantProfiles = computed<InstantProfile[]>(() => {
  return Array.from(profilesMap.value.values())
})

export function useInstantProfiles() {
  async function loadInstantProfiles() {
    try {
      const result = await api.listInstantProfile()
      const items: InstantProfile[] = result.objects ?? []
      const newMap = new Map<number, InstantProfile>()
      for (const profile of items) {
        newMap.set(profile.id, profile)
      }
      profilesMap.value = newMap
    } catch (error) {
      console.error('Failed to load instant profiles:', error)
    }
  }

  function updateInstantProfileDirect(updatedProfile: InstantProfile) {
    profilesMap.value.set(updatedProfile.id, updatedProfile)
  }

  async function toggleProfile(id: number) {
    try {
      const profile = profilesMap.value.get(id)
      if (!profile) return

      const patchData: PatchInstantProfileRequest = {
        active: !profile.active,
      }

      // Send patch request
      await api.patchInstantProfile({
        id: id,
        patchInstantProfileRequest: patchData,
      })

      // Query the profile again to verify success and get real state
      const verifiedProfile: InstantProfile = await api.getInstantProfile({ id })
      updateInstantProfileDirect(verifiedProfile)
    } catch (error) {
      console.error('Failed to toggle instant profile:', error)
    }
  }

  return {
    instantProfiles,
    loadInstantProfiles,
    updateInstantProfileDirect,
    toggleProfile,
  }
}
