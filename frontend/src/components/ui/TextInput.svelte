<script lang="ts">
  import type { Snippet } from "svelte";
  import { flip } from "svelte/animate";
  import { fade } from "svelte/transition";

  interface Props {
    value?: string;
    placeholder?: string;
    autoComplete?: string[];
    autoFocus?: boolean;
    label?: string;
    head?: Snippet;
    tail?: Snippet;
    onkeypress?: (e: KeyboardEvent) => void;
    variant?: "primary" | "secondary";
  }

  let {
    value = $bindable(""),
    placeholder = "",
    autoComplete = [],
    autoFocus = false,
    label = "",
    variant = "primary",
    head,
    tail,
    onkeypress,
  }: Props = $props();

  let filteredAutoComplete = $state(autoComplete);

  $effect(() => {
    if (value != undefined) {
      filteredAutoComplete = autoComplete.filter((ac) =>
        ac.toLowerCase().includes(value.toLowerCase()),
      );
    }
  });

  let showAutoCompletes = $state(false);
  let inputElem: HTMLInputElement;
</script>

<div class="relative" class:bg-accent-bg={variant === "primary"}>
  <div class="wrapper"
  class:borderless={variant === "secondary"}
  >
    {@render head?.()}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    {#if label}
      <div
        onclick={() => inputElem.focus()}
        class="flex items-center  justify-center h-full pl-2 text-white/60"
      >
        {label}
      </div>
    {/if}
    <!-- svelte-ignore a11y_autofocus -->
    <input
      bind:this={inputElem}
      class="flex-1 w-full px-3 py-2 focus:outline-none"
      class:bg-accent-bg={variant === "primary"}
      class:bg-transparent={variant === "secondary"}
      type="text"
      autocorrect="off"
      autocapitalize="none"
      autofocus={autoFocus}
      {placeholder}
      onkeypress={onkeypress}
      bind:value
      onfocus={() => (showAutoCompletes = true)}
      onblur={() => {
        setTimeout(() => (showAutoCompletes = false), 100);
      }}
    />
    {@render tail?.()}
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
          onclick={() => {
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
    @apply flex rounded border border-white/20 w-full h-10 overflow-hidden;
  }

  .wrapper.borderless {
    @apply border-transparent border-none;
  }

  .wrapper:has(input:focus) {
    @apply border-primary;
  }
</style>
