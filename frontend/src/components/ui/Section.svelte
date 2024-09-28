<script lang="ts">
  import { slide } from "svelte/transition";
  import Switch from "./Switch.svelte";
  import type { Snippet } from "svelte";

  interface Props {
    title?: string;
    subtitle?: string;
    isActive?: boolean | null;
    children?: Snippet;
    tail?: Snippet;
  }

  let { title = "", subtitle = "", isActive = null, children, tail }: Props = $props();
</script>

<div class="bg-default-bg" tabindex="-1">
  <div class="flex items-center justify-between">
    <div>
      <h2 class="text-lg text-bold">{title}</h2>
      {#if subtitle}
        <p class="mb-3 text-sm text-white/70">{subtitle}</p>
      {/if}
    </div>
    {#if isActive != null}
      <Switch bind:value={isActive} />
    {/if}
    {@render tail?.()}
  </div>
  {#if isActive == null || isActive == true}
    <div transition:slide={{ duration: 200 }}>{@render children?.()}</div>
  {/if}
</div>
