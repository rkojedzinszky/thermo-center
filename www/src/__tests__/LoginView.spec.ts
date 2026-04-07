import { describe, it, expect, vi, beforeEach } from 'vitest'
import { ref } from 'vue'
import { mount, flushPromises } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'

// Mock useAuth composable
const mockLogin = vi.fn()
const mockLoading = ref(false)
const mockError = ref<string | null>(null)

vi.mock('../composables/useAuth', () => ({
  useAuth: () => ({
    login: mockLogin,
    loading: mockLoading,
    error: mockError,
    session: ref(null),
    isLoggedIn: ref(false),
    checkSession: vi.fn(),
  }),
}))

function createTestRouter() {
  return createRouter({
    history: createWebHistory(),
    routes: [
      { path: '/login', component: LoginView },
      { path: '/overview', component: { template: '<div>Overview</div>' } },
    ],
  })
}

describe('LoginView', () => {
  beforeEach(() => {
    mockLogin.mockReset()
    mockError.value = null as string | null
    mockLoading.value = false
  })

  describe('layout & structure', () => {
    it('renders a login form', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('form').exists()).toBe(true)
    })

    it('renders username input', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('input[type="text"]').exists()).toBe(true)
    })

    it('renders password input', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('input[type="password"]').exists()).toBe(true)
    })

    it('renders sign in button', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      const btn = wrapper.find('button[type="submit"]')
      expect(btn.exists()).toBe(true)
      expect(btn.text()).toContain('Sign In')
    })

    it('renders Thermo Center title', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('.login-title').text()).toBe('Thermo Center')
    })
  })

  describe('readability', () => {
    it('has a visible subtitle for context', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('.login-subtitle').text()).toBeTruthy()
    })

    it('has labeled username input', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      const label = wrapper.find('label[for="username"]')
      expect(label.exists()).toBe(true)
      expect(label.text().toLowerCase()).toContain('username')
    })

    it('has labeled password input', () => {
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      const label = wrapper.find('label[for="password"]')
      expect(label.exists()).toBe(true)
      expect(label.text().toLowerCase()).toContain('password')
    })
  })

  describe('interaction', () => {
    it('calls login with entered credentials on submit', async () => {
      mockLogin.mockResolvedValue(false)
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })

      await wrapper.find('input[type="text"]').setValue('admin')
      await wrapper.find('input[type="password"]').setValue('secret')
      await wrapper.find('form').trigger('submit')
      await flushPromises()

      expect(mockLogin).toHaveBeenCalledWith('admin', 'secret')
    })

    it('shows error message when login fails', async () => {
      mockError.value = 'Invalid credentials'
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })

      expect(wrapper.find('.login-error').exists()).toBe(true)
      expect(wrapper.find('.login-error').text()).toContain('Invalid credentials')
    })

    it('does NOT show error element when no error', () => {
      mockError.value = null
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('.login-error').exists()).toBe(false)
    })

    it('disables submit button while loading', async () => {
      mockLoading.value = true
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      const btn = wrapper.find('button[type="submit"]')
      expect(btn.attributes('disabled')).toBeDefined()
    })

    it('shows loading text on button when loading', () => {
      mockLoading.value = true
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('button[type="submit"]').text()).toContain('Signing in')
    })

    it('shows spinner element while loading', () => {
      mockLoading.value = true
      const wrapper = mount(LoginView, { global: { plugins: [createTestRouter()] } })
      expect(wrapper.find('.spinner').exists()).toBe(true)
    })
  })
})
