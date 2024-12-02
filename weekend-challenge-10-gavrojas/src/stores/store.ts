import { defineStore } from 'pinia';
import type { IMessage, IConversation } from '@/types/index'
import { useRouter } from 'vue-router'

interface State {
  messages: IMessage[];
  botName: string;
  conversations: IConversation[];
  selectedConversation: IConversation | null;
  conversationCount: number;
}

const BotMessages: string[] = [
  'Hello, Friend!',
  'How are you?',
  'What are you up to?',
  'Cool! I am working on a new app.',
  'Nice! Sorry for the hurry but I need to leave.',
  'Thanks for the conversation! See you later.'
]

export const useStore = defineStore('main', {
  state: (): State => ({
    messages: [],
    botName: 'Bot',
    conversations: [],
    selectedConversation: null,
    conversationCount: 0,
  }),
  actions: {
    sendMessage(message: string) {
      const newMessage: IMessage = {
        id: Date.now().toString(),
        content: message,
        author: 'user',
      };
      this.selectedConversation?.messages.push(newMessage);
    },
    startNewConversation() {
      this.conversationCount++;
      const newConversation: IConversation = {
        id: `${this.conversationCount}`,
        name: `Kitty ${this.conversationCount}`,
        messages: [],
        botIndex: 0,
      };

      const initialBotMessage = this.getBotResponse(newConversation);
      if (initialBotMessage) {
        newConversation.messages.push({ id: `${Date.now()}`, author: 'bot', content: initialBotMessage });
      }
      
      this.conversations.unshift(newConversation);
      this.selectedConversation = newConversation;

      // Navegación - router
      const router = useRouter();
      router.push({ name: 'Chat', params: { id: newConversation.id } });
    },
    selectConversation(index: number) {
      this.selectedConversation = this.conversations[index];

      // Navegación - router
      const selectedConversation = this.selectedConversation;
      if (selectedConversation) {
        const router = useRouter();
        router.push({ name: 'Chat', params: { id: selectedConversation.id } });
      }
    },
    getBotResponse(conversation: IConversation): string | null {
      if (conversation.botIndex < BotMessages.length) {
        const response = BotMessages[conversation.botIndex];
        conversation.botIndex++;
        return response;
      }
      return null;
    },
    handleSendMessage(message: string) {
      if (this.selectedConversation && message.trim()) {
        const currentConversation = this.selectedConversation;
        currentConversation.messages.push({
          id: `${Date.now()}`,
          author: 'user',
          content: message,
        });
        const botResponse = this.getBotResponse(currentConversation);
        if (botResponse) {
          currentConversation.messages.push({
            id: `${Date.now()}`,
            author: 'bot',
            content: botResponse,
          });
        }
      }
    },
  },
});