<script lang="ts">
  import { flip } from "svelte/animate";
  import { fly } from "svelte/transition";
  import { showToast, toastStore } from "$stores/toast";
  import Icon from "@iconify/svelte";
</script>

<div
  class="fixed bottom-0 right-0 w-[40rem] p-5 z-50 flex flex-col gap-y-5 pointer-events-none"
>
  {#each $toastStore.list as toast, index (index)}
    <div
      class="rounded p-5 pl-3 border border-white/20 flex flex-row items-start gap-x-2 flex-shrink-0 bg-opacity-30 backdrop-blur-sm"
      class:bg-red-900={toast.type === "error"}
      class:bg-green-900={toast.type === "success"}
      class:bg-yellow-900={toast.type === "warning"}
      class:bg-blue-900={toast.type === "info"}
      animate:flip={{ duration: 200 }}
      in:fly={{ duration: 200, y: 100 }}
      out:fly={{ duration: 200, x: 400 }}
    >
      {#if toast.type === "error"}
        <Icon icon="fa6-regular:circle-xmark" width="24" class="text-red-800" />
      {:else if toast.type === "success"}
        <Icon icon="lets-icons:check-ring" width="24" class="text-green-800" />
      {:else if toast.type === "warning"}
        <Icon icon="ph:warning" width="24" class="text-yellow-800" />
      {:else if toast.type === "info"}
        <Icon
          icon="material-symbols:info-outline"
          width="24"
          class="text-blue-800"
        />
      {/if}
      <div>
        {toast.message}
      </div>
    </div>
  {/each}
</div>
