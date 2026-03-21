import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import Navigation from '@/components/Navigation.vue'
import { createRouter, createMemoryHistory } from 'vue-router'

describe('Navigation Dropdown Visibility', () => {
  let wrapper: ReturnType<typeof mount>

  beforeEach(() => {
    const router = createRouter({
      history: createMemoryHistory(),
      routes: [
        { path: '/overview', name: 'overview', component: { template: '<div>Overview</div>' } },
        { path: '/heating', name: 'heating', component: { template: '<div>Heating</div>' } },
      ],
    })

    wrapper = mount(Navigation, {
      props: { current: 'overview' },
      global: {
        plugins: [router],
      },
    })
  })

  it('should render menu button', () => {
    const button = wrapper.find('.menu-button')
    expect(button.exists()).toBe(true)
  })

  it('should hide dropdown initially', () => {
    const dropdown = document.body.querySelector('.menu-dropdown')
    expect(dropdown).toBeNull()
  })

  it('should show dropdown when button is clicked', async () => {
    const button = wrapper.find('.menu-button')
    await button.trigger('click')
    await nextTick()
    const dropdown = document.body.querySelector('.menu-dropdown')
    expect(dropdown).not.toBeNull()
  })

  it('dropdown should not be clipped by parent overflow', async () => {
    const button = wrapper.find('.menu-button')
    await button.trigger('click')
    await nextTick()

    const navMenu = wrapper.find('.nav-menu')
    const dropdown = document.body.querySelector('.menu-dropdown') as HTMLElement | null
    expect(dropdown).not.toBeNull()

    // Check computed styles
    const navMenuStyles = window.getComputedStyle(navMenu.element as HTMLElement)
    const dropdownStyles = window.getComputedStyle(dropdown as HTMLElement)

    console.log('nav-menu overflow:', navMenuStyles.overflow)
    console.log('nav-menu position:', navMenuStyles.position)
    console.log('dropdown z-index:', dropdownStyles.zIndex)
    console.log('dropdown position:', dropdownStyles.position)

    // Verify no overflow is clipping the dropdown
    expect(navMenuStyles.overflow).not.toBe('hidden')
  })

  it('dropdown should render into document body (teleport)', async () => {
    const button = wrapper.find('.menu-button')
    await button.trigger('click')
    await nextTick()

    const dropdown = document.body.querySelector('.menu-dropdown') as HTMLElement | null
    expect(dropdown).not.toBeNull()
    expect(dropdown?.parentElement).toBe(document.body)
  })

  it('should close dropdown when menu link is clicked', async () => {
    const button = wrapper.find('.menu-button')
    await button.trigger('click')
    await nextTick()

    let dropdown = document.body.querySelector('.menu-dropdown')
    expect(dropdown).not.toBeNull()

    const firstLink = document.body.querySelector('.menu-dropdown .menu-link') as HTMLElement | null
    expect(firstLink).not.toBeNull()
    firstLink?.click()
    await nextTick()
    dropdown = document.body.querySelector('.menu-dropdown')
    expect(dropdown).toBeNull()
  })
})
