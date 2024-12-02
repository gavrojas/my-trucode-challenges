<script setup lang="ts">
  import { useStore } from '@/stores/store';
  import { onMounted } from 'vue';
  import { useRouter } from 'vue-router'
  import ConversationList from '@/components/ConversationsList.vue';
  import ChatBox from '@/components/ChatBox.vue';

  const store = useStore();

  const startNewConversation = () => {
    store.startNewConversation();
  };

  const selectConversation = (index: number) => {
    store.selectConversation(index);
  };

  const handleSendMessage = (message: string) => {
    store.handleSendMessage(message);
  };

  onMounted(() => {
  const router = useRouter();
  const conversationId = router.currentRoute.value.params.id;
  if (conversationId) {
    const conversation = store.conversations.find(conv => conv.id === conversationId);
    if (conversation) {
      store.selectedConversation = conversation;
    }
  }
});
</script>

<template>
  <main class="grid grid-cols-10 grid-rows-1 h-screen w-screen">
    <!-- conversations -->
    <conversation-list class="col-span-3"
      :conversations="store.conversations" 
      @start-new="startNewConversation" 
      @select="selectConversation" 
      :selectedConversation="store.selectedConversation"
    />
    <!-- messages -->
    <chat-box class="col-span-7"
      @send="handleSendMessage"
      :messages="store.selectedConversation ? store.selectedConversation.messages : []" 
      :botName="store.selectedConversation ? store.selectedConversation.name : ''"
    />
  </main>
</template>
