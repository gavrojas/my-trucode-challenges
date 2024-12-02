// composable, los demás componentes van a usar este store
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { jwtDecode } from 'jwt-decode'
import type { JWTPayload } from '@/types/index'



export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const session = ref('')

  const isLoggedIn = computed(() => !!session.value ) 

  async function init() {
    const tokenStr = sessionStorage.getItem('token')
    if (tokenStr) {
      setSession(tokenStr)
    }
  }

  function setSession(tokenStr: string) {
    const payload = jwtDecode(tokenStr) as JWTPayload
    console.log(payload);
    // La diferencia entre sessionStorage y localStorage es que sessionStorage se limpia cuando se cierra el navegador, mientras que localStorage no, se va a quedar ahi poara siempre hasta que el sitio borre eso. 
    const now = new Date()
    const diff = payload.MapClaims.eat * 1000 - now.getTime()

    sessionStorage.setItem('token', tokenStr)
    session.value = tokenStr

    setTimeout(()=>{
      clearSession()
    }, diff)

    router.push('/entries')
  }

  function clearSession() { 
    session.value = ''
    router.push('/')
  }

  return { isLoggedIn, init, setSession, clearSession }
})

