<script lang="ts">
  import { flip } from "svelte/animate";
  import { fade } from "svelte/transition";
  export let value: string = "";
  export let placeholder: string = "";
  export let label: string = "";
  export let autoComplete: string[] = [];
  export let autoFocus = false;

  let filteredAutoComplete = autoComplete;

  $: {
    if (value != undefined) {
      filteredAutoComplete = autoComplete.filter((ac) =>
        ac.toLowerCase().includes(value.toLowerCase()),
      );
    }
  }
  let showAutoCompletes = false;
  let inputElem: HTMLInputElement;
</script>

<div class="relative">
  <div class="wrapper">
    <slot name="head" />
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    {#if label}
      <div
        on:click={() => inputElem.focus()}
        class="flex items-center bg-default-bg justify-center h-full pl-2 text-white/60"
      >
        {label}
      </div>
    {/if}
    <input
      bind:this={inputElem}
      class="flex-1 w-full px-3 py-2 bg-default-bg focus:outline-none"
      type="text"
      autocorrect="off"
      autocapitalize="none"
      autofocus={autoFocus}
      on:keypress
      {placeholder}
      bind:value
      on:focus={() => (showAutoCompletes = true)}
      on:blur={() => {
        setTimeout(() => (showAutoCompletes = false), 100);
      }}
    />
    <slot name="tail" />
  </div>
  {#if showAutoCompletes && filteredAutoComplete.length > 0}
    <div
      class="absolute w-full flex flex-col bg-default-bg border border-white/20 rounded top-[calc(100%+.5rem)] max-h-48 overflow-y-auto hide-scrollbar"
    >
      {#each filteredAutoComplete as ac (ac)}
        <button
          animate:flip={{ duration: 200 }}
          transition:fade={{ duration: 200 }}
          class="py-2 px-3 hover:bg-white/5 text-left"
          on:click={() => {
            value = ac;
            showAutoCompletes = false;
          }}
        >
          {ac}
        </button>
      {/each}
    </div>
  {/if}
</div>

<style lang="postcss">
  .wrapper {
    @apply flex border rounded border-white/20 w-full h-10 overflow-hidden;
  }

  .wrapper:has(input:focus) {
    @apply border-primary;
  }
</style>
