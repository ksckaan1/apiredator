<script>
  import MacHeader from "$components/MacHeader.svelte";
  import Nav from "$components/Nav.svelte";
</script>

<MacHeader />
<div class="flex">
  <Nav />
  <div class="flex flex-col w-full">
    <slot />
  </div>
</div>


