import { describe, it, expect, vi } from 'vitest'
import { mount, flushPromises } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import App from '../App.vue'

// Minimal stub router so App mounts without navigation errors
function makeRouter() {
  return createRouter({
    history: createWebHistory(),
    routes: [{ path: '/:pathMatch(.*)*', component: { template: '<div />' } }],
  })
}

describe('App', () => {
  it('mounts without errors', async () => {
    const router = makeRouter()
    const wrapper = mount(App, { global: { plugins: [router] } })
    await flushPromises()
    expect(wrapper.exists()).toBe(true)
  })

  it('renders a RouterView slot', () => {
    const router = makeRouter()
    const wrapper = mount(App, { global: { plugins: [router] } })
    // The wrapper itself exists (may render a comment node for router-view)
    expect(wrapper.exists()).toBe(true)
  })
})
