<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth';

const loggedUser = ref({ username: '', password: '' })
const registeredUser = ref({ username: '', password: '' })
const loginError = ref('')
const registerError = ref('')
const isRegistered = ref(false)
const router = useRouter()
const authStore = useAuthStore()

async function login() {
  try {
    const response = await fetch('http://localhost:3333/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(loggedUser.value),
    })
    const data = await response.json()
    authStore.setSession(data.token)
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(`${errorData.error}`)
    }
    router.push('/entries')
  } catch (error) {
    if (error instanceof Error) { 
      loginError.value = error.message 
    } else { 
      loginError.value = 'An unexpected error occurred' 
    } 
  }
}

async function register() {
  try {
    const response = await fetch('http://localhost:3333/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(registeredUser.value),
    })
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(`Registration failed ${errorData.error}`)
    }
    isRegistered.value = true

  } catch (error) {
    if (error instanceof Error) { 
      registerError.value = error.message 
    } else { 
      registerError.value = 'An unexpected error occurred' 
    } 
  }
}
</script>

<template>
  <div class="auth-container">
    <div class="auth-box">
      <div v-if="!isRegistered">
        <div class="register-form">
          <h2>Register</h2>
          <form @submit.prevent="register">
            <input v-model="registeredUser.username" type="text" placeholder="Username" />
            <input v-model="registeredUser.password" type="password" placeholder="Password" />
            <button type="submit">Register</button>
            <p v-if="registerError" class="error">{{ registerError }}</p>
          </form>
        </div>
      </div>
      <div v-else>
        <div class="login-form">
          <h2>Login</h2>
          <form @submit.prevent="login">
            <input v-model="loggedUser.username" type="text" placeholder="Username" />
            <input v-model="loggedUser.password" type="password" placeholder="Password" />
            <button type="submit">Login</button>
            <p v-if="loginError" class="error">{{ loginError }}</p>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.auth-box {
  display: flex;
  justify-content: space-around;
  width: 60%;
}
.auth-box form {
  display: flex;
  flex-direction: column;
  width: 45%;
}
.error {
  color: red;
}
</style>
