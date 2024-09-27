<script lang="ts">
  export let value = true;
</script>

<button
  class="relative flex flex-shrink-0 items-center transition-all duration-200 border rounded-full cursor-pointer h-7 w-14 border-white/20"
  class:active={value}
  on:click={() => (value = !value)}
>
  <div
    class="absolute transition-all duration-200 rounded-full bg-white/20 size-5 left-1"
  ></div>
</button>

<style lang="postcss">
  button.active {
    @apply border-primary;
  }

  button.active div {
    @apply bg-primary left-8;
  }
</style>
