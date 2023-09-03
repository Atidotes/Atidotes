import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const pinia = createPinia()

/** 使用持久化存储 */
pinia.use(piniaPluginPersistedstate)

export default pinia