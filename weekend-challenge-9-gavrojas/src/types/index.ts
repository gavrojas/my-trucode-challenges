export interface IConversation {
  id: string
  name: string
  messages: IMessage[]
  botIndex: number
}

export interface IMessage {
  id: string
  author: string
  content: string
}

export interface IData {
  conversations: IConversation[]
  selectedConversation: IConversation | null
}
