<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'

const router = useRouter()
const { login, loading, error } = useAuth()

const username = ref('')
const password = ref('')

async function handleSubmit() {
  const ok = await login(username.value, password.value)
  if (ok) {
    router.push('/overview')
  }
}
</script>

<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-logo">
        <span class="login-icon">🌡️</span>
        <h1 class="login-title">Thermo Center</h1>
        <p class="login-subtitle">Sign in to continue</p>
      </div>

      <form class="login-form" @submit.prevent="handleSubmit">
        <div class="form-group">
          <label class="form-label" for="username">Username</label>
          <input
            id="username"
            v-model="username"
            class="form-input"
            type="text"
            autocomplete="username"
            placeholder="Enter username"
            required
          />
        </div>

        <div class="form-group">
          <label class="form-label" for="password">Password</label>
          <input
            id="password"
            v-model="password"
            class="form-input"
            type="password"
            autocomplete="current-password"
            placeholder="Enter password"
            required
          />
        </div>

        <div v-if="error" class="login-error" role="alert">
          {{ error }}
        </div>

        <button class="login-btn" type="submit" :disabled="loading">
          <span v-if="loading" class="spinner" />
          {{ loading ? 'Signing in…' : 'Sign In' }}
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  flex: 1;
  min-height: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow-y: auto;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%);
  padding: 1rem;
}

.login-card {
  background: rgba(255, 255, 255, 0.05);
  backdrop-filter: blur(16px);
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 1.5rem;
  padding: 2.5rem 2rem;
  width: 100%;
  max-width: 400px;
  box-shadow:
    0 8px 32px rgba(0, 0, 0, 0.4),
    0 0 0 1px rgba(255, 255, 255, 0.05);
}

.login-logo {
  text-align: center;
  margin-bottom: 2rem;
}

.login-icon {
  font-size: 3rem;
}

.login-title {
  color: #e2e8f0;
  font-size: 1.75rem;
  font-weight: 700;
  margin: 0.5rem 0 0.25rem;
  letter-spacing: -0.01em;
}

.login-subtitle {
  color: #94a3b8;
  font-size: 0.9rem;
  margin: 0;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.form-label {
  color: #cbd5e1;
  font-size: 0.85rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.form-input {
  background: rgba(255, 255, 255, 0.07);
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: 0.75rem;
  color: #e2e8f0;
  font-size: 1rem;
  padding: 0.75rem 1rem;
  outline: none;
  transition:
    border-color 0.2s,
    box-shadow 0.2s;
}

.form-input::placeholder {
  color: #475569;
}

.form-input:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.25);
}

.login-error {
  background: rgba(239, 68, 68, 0.15);
  border: 1px solid rgba(239, 68, 68, 0.4);
  border-radius: 0.75rem;
  color: #fca5a5;
  font-size: 0.9rem;
  padding: 0.75rem 1rem;
}

.login-btn {
  align-items: center;
  background: linear-gradient(135deg, #3b82f6, #2563eb);
  border: none;
  border-radius: 0.75rem;
  color: white;
  cursor: pointer;
  display: flex;
  font-size: 1rem;
  font-weight: 600;
  gap: 0.5rem;
  justify-content: center;
  padding: 0.85rem 1.5rem;
  transition:
    opacity 0.2s,
    transform 0.1s,
    box-shadow 0.2s;
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.35);
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.login-btn:not(:disabled):hover {
  opacity: 0.9;
  box-shadow: 0 4px 20px rgba(59, 130, 246, 0.5);
}

.login-btn:not(:disabled):active {
  transform: scale(0.98);
}

.spinner {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
  display: inline-block;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
