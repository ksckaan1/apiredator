<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { Snippet } from "svelte";

  interface Props {
    variant?: "primary" | "secondary" | "danger" | "transparent" | "outlined" | "success";
    onclick?: (e: MouseEvent) => void;
    icon?: string;
    iconSize?: number;
    children?: Snippet;
  }

  let { variant = "primary", icon, iconSize = 26, onclick, children }: Props = $props();
</script>

<button
  class="h-10 px-3 border rounded border-white/20 flex gap-x-2 items-center hover:opacity-70"
  class:primary={variant === "primary"}
  class:success={variant === "success"}
  class:danger={variant === "danger"}
  class:outlined={variant === "outlined"}
  class:transparent={variant === "transparent"}
  {onclick}
>
  {#if icon}
    <Icon icon={icon} width={iconSize} />
  {/if}
  {#if children}
    <span>
      {@render children()}
    </span>
  {/if}
</button>

<style lang="postcss">
  .primary {
    @apply bg-primary text-on-primary;
  }

  .success{
    @apply bg-green-900 text-green-300;
  }

  .danger {
    @apply bg-red-900 text-red-300;
  }

  .outlined{
    @apply bg-transparent text-white/70 hover:text-primary border border-white/20 hover:border-primary;
  }

  .transparent {
    @apply border-none bg-transparent text-white/80 hover:text-primary;
  }
</style>
