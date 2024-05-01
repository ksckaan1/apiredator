<script lang="ts">
  export let statusCodes: { [key: number]: number } = {};

  let keys: number[] = [];

  $: keys = Object.keys(statusCodes).map((sc) => Number(sc));

  let isCollapsed = true;
</script>

<div class="tile">
  <h1>Status Codes</h1>
  {#if statusCodes}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div
      class="flex flex-col gap-2 overflow-y-hidden cursor-pointer mt-3"
      class:max-h-32={isCollapsed}
      on:click={() => (isCollapsed = !isCollapsed)}
    >
      {#each keys as sc}
        <div class="flex justify-between items-center gap-2">
          <div
            class:text-sky-400={sc < 200}
            class:text-green-400={sc > 199 && sc < 300}
            class:text-orange-400={sc > 299 && sc < 400}
            class:text-red-400={sc > 399 && sc < 500}
            class:text-purple-400={sc > 500}
          >
            {sc}
          </div>
          <div
            class="flex-1 border-b border-dashed border-opacity-20"
            class:border-sky-400={sc < 200}
            class:border-green-400={sc > 199 && sc < 300}
            class:border-orange-400={sc > 299 && sc < 400}
            class:border-red-400={sc > 399 && sc < 500}
            class:border-purple-400={sc > 500}
          ></div>
          <div
            class:text-sky-400={sc < 200}
            class:text-green-400={sc > 199 && sc < 300}
            class:text-orange-400={sc > 299 && sc < 400}
            class:text-red-400={sc > 399 && sc < 500}
            class:text-purple-400={sc > 500}
          >
            {statusCodes[sc]}x
          </div>
        </div>
      {/each}
    </div>
    {#if isCollapsed && keys.length > 4}
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <!-- svelte-ignore a11y-no-static-element-interactions -->
      <span
        on:click={() => (isCollapsed = !isCollapsed)}
        class="text-white/40 cursor-pointer">click to see more...</span
      >
    {/if}
  {/if}
</div>
