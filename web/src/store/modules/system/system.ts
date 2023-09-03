import { defineStore } from 'pinia'

export const useChatStore = defineStore('system', {
  state: () => {
    return {
   
    }
  },

  actions: {
  },

  getters: {
  },

  persist: {
    key: 'system',
    paths: [],
    storage: localStorage,
  },
})
