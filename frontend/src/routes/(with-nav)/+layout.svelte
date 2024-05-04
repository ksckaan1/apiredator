<script>
  import { fade } from "svelte/transition";
  import MacHeader from "$components/MacHeader.svelte";
  import Nav from "$components/Nav.svelte";
</script>

<div in:fade={{ duration: 200, delay: 200 }} out:fade={{ duration: 200 }}>
  <MacHeader />
  <div class="flex">
    <Nav />
    <div class="flex flex-col w-full">
      <slot />
    </div>
  </div>
</div>
