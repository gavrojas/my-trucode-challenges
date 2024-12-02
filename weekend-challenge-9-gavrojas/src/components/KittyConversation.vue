<script lang="ts">
import type { PropType } from 'vue';
import type { IConversation } from '@/types/index';

  export default {
    props: {
      conv: {
        type: Object as PropType<IConversation>,
        default: () => ({}),
        required: true,
      },
      index: {
        type: Number,
        required: true,
      },
      isSelected: {
        type: Boolean,
        required: true
      }
    }
  }
</script>

<template>
  <li @click="$emit('select', index)" :class="{ selected: isSelected }">
    {{ conv.name }}
  </li>
</template>

<style scoped>
  li{
    list-style: none;
    padding: var(--spacing-lg);
    background-color: var(--blue);
    border-radius: var(--border-radius-lg);
    color: var(--white);
  }

  li.selected{
    background-color: var(--turquoise);
    color: var(--blue);
  }
</style>
