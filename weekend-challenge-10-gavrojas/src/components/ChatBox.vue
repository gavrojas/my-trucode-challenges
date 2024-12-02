<script setup lang="ts">
  import { ref, onMounted, onUpdated, onBeforeUnmount } from 'vue';
  import { useStore } from '@/stores/store';
  import Message from '@/components/ConversationsMessages.vue';

  const localUserMessage = ref('');
  const store = useStore();
  const messagesContainer = ref<HTMLElement | null>(null);

  const handleScroll = () => {
    const container = messagesContainer.value;
    if (container) {
      if (container.scrollTop > 0) {
        container.classList.add('scrolling');
      } else {
        container.classList.remove('scrolling');
      }
    }
  };
  const sendLocalMessage = () => {
      if (localUserMessage.value.trim()) {
        store.handleSendMessage(localUserMessage.value);
        localUserMessage.value = '';
      }
    };
  const scrollToBottom =() => {
    const container = messagesContainer.value;
    if (container) {
      container.scrollTop = container.scrollHeight;
    }
  }
  onMounted(() => {
    const container = messagesContainer.value;
    if (container) {
      container.addEventListener('scroll', handleScroll);
      scrollToBottom();
    }
  });
  onUpdated(() => {
    scrollToBottom();
  });
  onBeforeUnmount(() => {
    const container = messagesContainer.value;
    if (container) {
      container.removeEventListener('scroll', handleScroll);
    }
  });
</script>

<template>
  <article class="chat-box bg-turquoise min-h-screen p-lg grid gap-lg content-between">
    <div class="messages max-h-[90vh] overflow-y-scroll grid gap-sm" ref="messagesContainer" @scroll="handleScroll">
      <message 
        v-for="msg in store.selectedConversation?.messages" 
        :key="msg.id"
        :msg="msg" 
        :botName="store.selectedConversation?.name ?? ''" 
      />
    </div>
    <div class="input flex gap-sm" v-if="store.selectedConversation?.messages.length !== 0">
      <input v-model="localUserMessage" @keyup.enter="sendLocalMessage" placeholder="Type your message here..." class="w-full rounded-lg border border-blue p-sm" />
      <button @click="sendLocalMessage" class="p-lg bg-blue rounded-lg text-white border-none">Send</button>
    </div>  
  </article>
</template>


