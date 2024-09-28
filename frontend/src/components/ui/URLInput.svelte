<script lang="ts">
  import { activeEnvVars, type EnvVar } from "$stores/environment";
  import type { Snippet } from "svelte";

  interface Props {
    value?: string;
    placeholder?: string;
    label?: string;
    tail?: Snippet;
    head?: Snippet;
  }

  let {
    value = $bindable(""),
    placeholder = "",
    label = "",
    tail,
    head,
  }: Props = $props();

  let inputElem: HTMLInputElement;

  const focusInput = () => {
    inputElem.focus();
  };
</script>

<div class="wrapper" tabindex="-1">
  {@render head?.()}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  {#if label}
    <div
      onclick={focusInput}
      class="flex items-center justify-center h-full pl-2 text-white/60"
    >
      {label}
    </div>
  {/if}
  <input
    bind:this={inputElem}
    class="flex-1 w-full px-3 py-2 bg-accent-bg focus:outline-none"
    autocorrect="off"
    autocapitalize="none"
    type="url"
    {placeholder}
    bind:value
  />
  {@render tail?.()}
</div>

<style lang="postcss">
  .wrapper {
    @apply flex border rounded bg-accent-bg border-white/20 w-full h-10;
  }

  /* TODO: &:focus changed to :focus */
  .wrapper:has(input:focus, :focus) {
    @apply border-primary;
  }
</style>
