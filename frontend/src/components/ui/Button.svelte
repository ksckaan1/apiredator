<script lang="ts">
    import type { Snippet } from "svelte";

  interface Props {
    variant?: "primary" | "secondary";
    onclick?: (e: MouseEvent) => void;
    children: Snippet;
  }

  let { variant = "primary", onclick,children }: Props = $props();
</script>

<button
  class:primary={variant === "primary"}
  class="h-10 px-3 border rounded border-white/20"
  {onclick}
>
  {@render children()}
</button>

<style lang="postcss">
  .primary {
    @apply bg-primary text-on-primary;
  }
</style>
