import type { OptionsApiCall } from '@/types/index'
const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:3333'

type ApiCallOptions = OptionsApiCall | undefined

export async function apiCall(
  path: string,
  { method = 'GET', data, headers }: ApiCallOptions = {}
) {
  const sessionToken = sessionStorage.getItem('token')
  const response = await fetch(BASE_URL + path, {
    method,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${sessionToken}`,
      ...headers
    },
    body: JSON.stringify(data)
  })

  if (method === 'DELETE') {
    return response.ok
  }

  const jsonObj = await response.json()
  return jsonObj
}