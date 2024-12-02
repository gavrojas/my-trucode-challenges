<script lang="ts">
  import Message from '@/components/ConversationsMessages.vue';
  import type { PropType } from 'vue';
  import type { IMessage } from '@/types/index';

  export default {
    props: {
      messages: {
        type: Array as PropType<IMessage[]>,
        required: true
      },
      botName: {
        type: String,
        required: true
      },
    },

    data() {
      return {
        localUserMessage: ''
      };
    },

    methods: {
      handleScroll() {
        const container = this.$refs.messagesContainer as HTMLElement;
        if (container.scrollTop > 0) {
          container.classList.add('scrolling');
        } else {
          container.classList.remove('scrolling');
        }
      },
      sendLocalMessage() {
        if (this.localUserMessage.trim()) {
          this.$emit('send', this.localUserMessage);
          this.localUserMessage = '';
        }
      },
      scrollToBottom() {
        const container = this.$refs.messagesContainer as HTMLElement;
        container.scrollTop = container.scrollHeight;
      }
    },
    mounted() {
      const container = this.$refs.messagesContainer as HTMLElement;
      container.addEventListener('scroll', this.handleScroll);
      this.scrollToBottom(); // Scroll a la parte inferior cuando el componente se monta
    },

    updated() {
      this.scrollToBottom(); // Scroll a la parte inferior cuando el componente se actualiza
    },

    beforeUnmount() {
      const container = this.$refs.messagesContainer as HTMLElement;
      container.removeEventListener('scroll', this.handleScroll);
    },
    components: { Message },
  };
</script>

<template>
  <article class="chat-box">
    <div class="messages" ref="messagesContainer" @scroll="handleScroll">
      <message 
        v-for="msg in messages" 
        :key="msg.id"
        :msg="msg" 
        :botName="botName" 
      />
    </div>
    <div class="input" v-if="messages.length !== 0">
      <input v-model="localUserMessage" @keyup.enter="sendLocalMessage" placeholder="Type your message here..." />
      <button @click="sendLocalMessage">Send</button>
    </div>  
  </article>
</template>

<style scoped>
  .chat-box {
    background-color: var(--turquoise);
    min-height: 100vh;
    padding: var(--spacing-lg);
    display: grid;
    gap: var(--spacing-lg);
    align-content: space-between;
  }

  .messages {
    max-height: 90vh;
    overflow-y: scroll;
    scrollbar-width: thin;
    scrollbar-color: transparent transparent;
    display: grid;
    gap: var(--spacing-sm);
  }

  .messages.scrolling {
    scrollbar-color: var(--black25) transparent; /* Para Firefox */
  }

  input{
    width: 100%;
    border-radius: var(--border-radius-lg);
    border: solid 1px var(--blue);
  }

  .input{
    gap: 10px;
    display: flex;
  }

  button{
    padding: var(--spacing-lg);
    background-color: var(--blue);
    border-radius: var(--border-radius-lg);
    color: var(--white);
    border: none;
  }
</style>
