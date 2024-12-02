<template>
  <div>
    <h1>Conversation - Kitty {{ $route.params.id }}</h1>
    <chat-box />
  </div>
</template>

<script setup lang="ts">
import ChatBox from '@/components/ChatBox.vue';
import { onMounted, watch } from 'vue';
import { useStore } from '@/stores/store';
import { useRoute } from 'vue-router';

const store = useStore();
const route = useRoute();

onMounted(() => {
  const conversationId = route.params.id;
  const conversation = store.conversations.find(conv => conv.id === conversationId);
  if (conversation) {
    store.selectedConversation = conversation;
  }
});

watch(
  () => route.params.id,
  (newId) => {
    if (newId) {
      const conversation = store.conversations.find(conv => conv.id === newId);
      if (conversation) {
        store.selectedConversation = conversation;
      }
    }
  }
);
</script>

<style scoped>
h1 {
  color: green;
}
</style>
