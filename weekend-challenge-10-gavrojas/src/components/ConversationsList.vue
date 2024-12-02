<script setup lang="ts">
  import {computed } from 'vue';
  import { useStore } from '@/stores/store';
  import Conversation from '@/components/KittyConversation.vue';

  const store = useStore();

  const conversations = computed(() => store.conversations);
  const selectedConversation = computed(() => store.selectedConversation);

  const startNewConversation = () => {
    store.startNewConversation();
  };

  const selectConversation = (index: number) => {
    store.selectConversation(index);
  };

</script>

<template>
  <aside class="conversations bg-lavender min-h-screen p-lg flex flex-col justify-between">
    <p v-if="conversations.length === 0">No conversations</p>
    <ul class="m-0 p-0 grid gap-lg">
      <conversation 
        v-for="(conv, index) in conversations" 
        :key="conv.id"
        :conv="conv" 
        :index="index" 
        :isSelected="conversations.length > 1 && selectedConversation !== null && selectedConversation.id === conv.id"
        @select="selectConversation"
      />
    </ul>
    <button @click="startNewConversation" class="p-lg bg-blue rounded-lg text-white border-none">Start a conversation</button>
  </aside>
</template>

