<script lang="ts">
  import { fade } from "svelte/transition";
  import MacHeader from "$components/MacHeader.svelte";
  import Nav from "$components/Nav.svelte";
  import type { Snippet } from "svelte";

  interface Props {
    children?: Snippet;
  }

  let { children }: Props = $props();
</script>

<div in:fade={{ duration: 200, delay: 200 }} out:fade={{ duration: 200 }}>
  <MacHeader />
  <div class="flex">
    <Nav />
    <div class="flex flex-col w-full">
      {@render children?.()}
    </div>
  </div>
</div>
