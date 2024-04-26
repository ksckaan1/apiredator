<script lang="ts">
  export let value: string = "";
  export let placeholder: string = "";
  export let label: string = "";
  let inputElem: HTMLInputElement;
</script>

<div class="wrapper">
  <slot name="head" />
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  {#if label}
    <div
      on:click={() => inputElem.focus()}
      class="flex items-center justify-center h-full pl-2 text-white/60"
    >
      URL
    </div>
  {/if}
  <input
    bind:this={inputElem}
    class="flex-1 w-full px-3 py-2 bg-default-bg focus:outline-none"
    type="text"
    autocorrect="off"
    autocapitalize="none"
    {placeholder}
    bind:value
  />
  <slot name="tail" />
</div>

<style lang="postcss">
  .wrapper {
    @apply flex border rounded border-white/20 w-full h-10 overflow-hidden;
  }

  .wrapper:has(input:focus) {
    @apply border-primary;
  }
</style>
