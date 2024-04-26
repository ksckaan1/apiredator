<script lang="ts">
  import { slide } from "svelte/transition";
  import Switch from "./Switch.svelte";

  export let title: string = "";
  export let subtitle: string = "";
  export let isActive: boolean | null = null;
</script>

<div class="" tabindex="-1">
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
  </div>
  {#if isActive == null || isActive == true}
    <div transition:slide={{ duration: 200 }}><slot /></div>
  {/if}
</div>
