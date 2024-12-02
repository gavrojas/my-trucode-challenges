<script lang="ts">
  import type { PropType } from 'vue';
  import type { IMessage } from '@/types/index';

  export default {
    props: {
      msg: {
        type: Object as PropType<IMessage>,
        default: () => ({}),
        required: true
      },
      botName: {
        type: String,
        required: true
      } 
    }
  }
</script>

<template>
  <span v-if="msg.author === 'bot'" class="bot">
    <h3>{{ botName }}:</h3> 
    <p>{{ msg.content }}</p>
  </span>
  <span v-if="msg.author === 'user'" class="user">
    <h3>Me:</h3> 
    <p>{{ msg.content }}</p>
  </span>
</template>

<style scoped>
  span{
    background-color: var(--white);
    border-radius: 16px;
    display: block;
    padding: 0.5rem var(--spacing-sm);
  }
</style>