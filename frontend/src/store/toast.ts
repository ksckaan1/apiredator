import { writable } from "svelte/store";
import { uniqueId } from "lodash"


export type Toast = {
  id?: string
  message: string
  type: "success" | "error" | "warning" | "info"
  duration?: number
}

interface ToastStore {
  list: Toast[];
}

export const toastStore = writable<ToastStore>({
  list: []
})

export const showToast = (t: Toast) => {
  t.id = uniqueId()

  if (!t.duration) {
    t.duration = 3000
  }

  toastStore.update(store => {
    store.list.push(t);
    return store;
  })

  setTimeout(() => {
    toastStore.update(store => {
      store.list = store.list.filter(item => item.id !== t.id)
      return store;
    })
    console.log("removed", t);

  }, t.duration)

  console.log(t);
}