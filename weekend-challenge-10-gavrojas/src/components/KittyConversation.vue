<script setup lang="ts">
  import type { IConversation } from '@/types/index';

  const props = defineProps<{
    conv: IConversation;
    index: number;
    isSelected: boolean;
  }>();

  const emit = defineEmits(['select']);
  const { conv, index, isSelected } = props;
</script>

<template>
  <li @click="emit('select', index)" :class="{ 'bg-turquoise text-blue': isSelected, 'bg-blue text-white': !isSelected }" class="list-none p-lg rounded-lg">
    {{ conv.name }}
  </li>
</template>
