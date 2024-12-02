export interface Entry {
  id: number
  title: string
  content: string
}

export interface JWTPayload {
  MapClaims: {
    eat: number
    iat: number
  }
  session: string
}

export interface OptionsApiCall {
  method?: string
  data?: any
  headers?: any
}