<script lang="ts">
  import { ref } from 'vue';
  import ConversationList from '@/components/ConversationsList.vue';
  import ChatBox from '@/components/ChatBox.vue';
  import type { IConversation, IData } from '@/types/index';

  const BotMessages: string[] = [
    'Hello, Friend!',
    'How are you?',
    'What are you up to?',
    'Cool! I am working on a new app.',
    'Nice! Sorry for the hurry but I need to leave.',
    'Thanks for the conversation! See you later.'
  ]

  export default {
    data() {
      return {
        data: ref<IData>({ 
          conversations: [], 
          selectedConversation: null 
        }),
        conversationCount: ref(0),
      }
    },
    methods: {
      startNewConversation() {
        this.conversationCount++;
        const newConversation: IConversation = {
          id: `${this.conversationCount}`,
          name: `Kitty ${this.conversationCount}`,
          messages: [],
          botIndex: 0 
        };

        const initialBotMessage = this.getBotResponse(newConversation);
        if (initialBotMessage) {
          newConversation.messages.push({ id: `${Date.now()}`, author: 'bot', content: initialBotMessage });
        }

        this.data.conversations.unshift(newConversation);
        this.data.selectedConversation = newConversation;
      },

      selectConversation(index: number) {
        this.data.selectedConversation = this.data.conversations[index];
      },

      handleSendMessage(message: string) {
        if (this.data.selectedConversation && message.trim()) {
          const currentConversation = this.data.selectedConversation;
          currentConversation.messages.push({ id: `${Date.now()}`, author: 'user', content: message });

          const botResponse = this.getBotResponse(currentConversation);
          if (botResponse) {
            currentConversation.messages.push({ id: `${Date.now()}`, author: 'bot', content: botResponse });
          }
        }
      },

      getBotResponse(conversation: IConversation) {
        if (conversation.botIndex < BotMessages.length) {
          const response = BotMessages[conversation.botIndex];
          conversation.botIndex++;
          return response;
        }
        return null;
      },
    },
    computed: {},
    components: { ConversationList, ChatBox },
  }
</script>

<template>
  <main>
    <!-- conversations -->
    <!-- messages -->
    <conversation-list 
      :conversations="data.conversations" 
      @start-new="startNewConversation" 
      @select="selectConversation" 
      :selectedConversation="data.selectedConversation"
    />
    <chat-box 
      :messages="data.selectedConversation ? data.selectedConversation.messages : []" 
      :botName="data.selectedConversation ? data.selectedConversation.name : ''"
      @send="handleSendMessage"
    />

  </main>
</template>

<style scoped>
  main { 
    display: grid;
    grid-template-columns: 3fr 7fr;
    grid-template-rows: 1fr;
    height: 100vh;
    width: 100vw;
  }
</style>
