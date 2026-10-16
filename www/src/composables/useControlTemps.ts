import { computed, type MaybeRefOrGetter, toValue } from 'vue'
import type { Control } from '@/api'
import { formatDecimal2 } from '@/composables/useSensorFormatting'

function formatTemp(value: number | null | undefined): string {
  return value != null ? `${formatDecimal2(value)} °C` : '—'
}

/**
 * Shared current/target temperature display strings for a control.
 * Replaces duplicated `formatDecimal2` snippets in profile views.
 */
export function useControlTemps(control: MaybeRefOrGetter<Control | null | undefined>) {
  const currentTemp = computed(() => formatTemp(toValue(control)?.temperature))
  const targetNow = computed(() => formatTemp(toValue(control)?.targetTemp))
  return { currentTemp, targetNow }
}
