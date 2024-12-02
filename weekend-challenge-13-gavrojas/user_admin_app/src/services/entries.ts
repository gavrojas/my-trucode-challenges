import { apiCall } from '@/services/utils';
import { ref } from 'vue';
import type { Entry } from '@/types/index';



export function useEntries(){
  const entries = ref<Entry[]>([])

  async function loadEntries() {
    const data = await apiCall('/entries')
    entries.value = data as Entry[]
  }

  const selectedEntry = ref<Entry | null>(null)
  const newEntry = ref({ title: '', content: '' })


  async function createEntry() {
    await apiCall('/entries', {
      method: 'POST',
      data: newEntry.value
    })
    loadEntries()
    newEntry.value = { title: '', content: '' }
  }

  async function deleteEntry(id: number) {
    await apiCall('/entries', {
      method: 'DELETE',
    })
    loadEntries()
  }
  
  async function editEntry(entry: Entry) {
    selectedEntry.value = { ...entry }
  }
  
  async function updateEntry() {
    await apiCall(`/entries/${selectedEntry.value!.id}`, {
      method: 'PUT',
      data: selectedEntry.value
    })
    loadEntries()
    selectedEntry.value = null
  }

  return {loadEntries, createEntry, deleteEntry, editEntry, updateEntry, entries, newEntry, selectedEntry}
}


