<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Entry } from '@/types/index';

const entries = ref<Entry[]>([])
const newEntry = ref({ title: '', content: '' })
const selectedEntry = ref<Entry | null>(null)
const showCreateEntryForm = ref(false)

async function loadEntries() {
  const response = await fetch('http://localhost:3333/entries', {
    credentials: 'include',
  })
  const data = await response.json()
  entries.value = data
}

async function createEntry() {
  await fetch('http://localhost:3333/entries', {
    credentials: 'include',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(newEntry.value)
  })
  loadEntries()
  newEntry.value = { title: '', content: '' }
}

async function deleteEntry(id: number) {
  await fetch(`http://localhost:3333/entries/${id}`, {
    credentials: 'include',
    method: 'DELETE',
  })
  loadEntries()
}

async function editEntry(entry: Entry) {
  selectedEntry.value = { ...entry }
}

async function updateEntry() {
  await fetch(`http://localhost:3333/entries/${selectedEntry.value!.id}`, {
    credentials: 'include',
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(selectedEntry.value)
  })
  loadEntries()
  selectedEntry.value = null
}

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

