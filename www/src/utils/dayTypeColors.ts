export interface DayTypeColor {
  accent: string
  soft: string
}

const PALETTE: DayTypeColor[] = [
  { accent: '#6366f1', soft: 'rgba(99, 102, 241, 0.18)' },
  { accent: '#f59e0b', soft: 'rgba(245, 158, 11, 0.18)' },
  { accent: '#ef4444', soft: 'rgba(239, 68, 68, 0.18)' },
  { accent: '#10b981', soft: 'rgba(16, 185, 129, 0.18)' },
  { accent: '#06b6d4', soft: 'rgba(6, 182, 212, 0.18)' },
  { accent: '#a855f7', soft: 'rgba(168, 85, 247, 0.18)' },
  { accent: '#ec4899', soft: 'rgba(236, 72, 153, 0.18)' },
  { accent: '#84cc16', soft: 'rgba(132, 204, 22, 0.18)' },
]

/** Stable color for a day type, keyed by its position in the day-type list. */
export function dayTypeColor(index: number): DayTypeColor {
  const first = PALETTE[0]!
  if (index < 0) return first
  return PALETTE[index % PALETTE.length] ?? first
}
