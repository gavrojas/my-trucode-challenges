<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Entry } from '@/types/index';
import { useEntries } from '@/services/entries';

const { loadEntries, createEntry, deleteEntry, editEntry, updateEntry, entries, newEntry, selectedEntry } = useEntries()
const showCreateEntryForm = ref(false)

onMounted(loadEntries)
</script>
<template>
  <div>
    <h1>Entries</h1>
    <button @click="showCreateEntryForm = !showCreateEntryForm">
      {{ showCreateEntryForm ? 'Hide' : 'Create New Entry' }}
    </button>
    <div v-if="showCreateEntryForm">
      <form @submit.prevent="createEntry">
        <input v-model="newEntry.title" type="text" placeholder="Title" />
        <textarea v-model="newEntry.content" placeholder="Content"></textarea>
        <button type="submit">Create Entry</button>
      </form>
    </div>
    <ul>
      <li v-for="entry in entries" :key="entry.id">
        <h2>{{ entry.title }}</h2>
        <p>{{ entry.content }}</p>
        <button @click="editEntry(entry)">Edit</button>
        <button @click="deleteEntry(entry.id)">Delete</button>
      </li>
    </ul>
    <div v-if="selectedEntry">
      <h2>Edit Entry</h2>
      <form @submit.prevent="updateEntry">
        <input v-model="selectedEntry.title" type="text" placeholder="Title" />
        <textarea v-model="selectedEntry.content" placeholder="Content"></textarea>
        <button type="submit">Update Entry</button>
      </form>
    </div>
  </div>
</template>

