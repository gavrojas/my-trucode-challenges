<script lang="ts">
  import Conversation from '@/components/KittyConversation.vue';
  import type { PropType } from 'vue';
  import type { IConversation } from '@/types/index';

  export default {
    props: {
      conversations: {
        type: Array as PropType<IConversation[]>,
        required: true
      },
      selectedConversation: {
        type: Object as PropType<IConversation | null>,
        required: true
      }
    },
    methods: {
      startNewConversation() {
        this.$emit('start-new');
      },
      selectConversation(index: number) {
        this.$emit('select', index);
      },
    },
    components: { Conversation },
  }
</script>

<template>
  <aside class="conversations">
    <p v-if="conversations.length === 0">No conversations</p>
    <ul>
      <conversation 
        v-for="(conv, index) in conversations" 
        :key="conv.id"
        :conv="conv" 
        :index="index" 
        :isSelected="conversations.length > 1 && selectedConversation !== null && selectedConversation.id === conv.id"
        @select="selectConversation"
      />
    </ul>
    <button @click="startNewConversation">Start a conversation</button>
  </aside>
</template>

<style scoped>
  .conversations {
    background-color: var(--lavender);
    min-height: 100vh;
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  ul{
    margin: 0;
    padding: 0;
    display: grid;
    gap: var(--spacing-lg)
  }

  button{
    padding: var(--spacing-lg);
    background-color: var(--blue);
    border-radius: var(--border-radius-lg);
    color: var(--white);
    border: none;
  }
</style>
