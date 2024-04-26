<script lang="ts">
  export let variant: "primary" | "secondary" = "primary";
</script>

<button
  class:primary={variant === "primary"}
  class="h-10 px-3 border rounded border-white/20"
  on:click
>
  <slot />
</button>

<style lang="postcss">
  .primary {
    @apply bg-primary text-on-primary;
  }
</style>
