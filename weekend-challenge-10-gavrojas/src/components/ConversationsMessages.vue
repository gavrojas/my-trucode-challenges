<script setup lang="ts">
  import type { IMessage } from '@/types/index';

  const props = defineProps<{
    msg: IMessage;
    botName: string;
  }>();

  const { msg, botName } = props;
</script>

<template>
  <span v-if="msg.author === 'bot'" class="bot bg-white rounded-xlg grid gap-sm p-sm">
    <h3 class="text-xl font-bold">{{ botName }}:</h3> 
    <p>{{ msg.content }}</p>
  </span>
  <span v-if="msg.author === 'user'" class="user bg-white rounded-xlg grid gap-sm p-sm">
    <h3 class="text-xl font-bold">Me:</h3> 
    <p>{{ msg.content }}</p>
  </span>
</template>
