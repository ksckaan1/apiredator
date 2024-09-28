<script lang="ts">
  interface Props {
    value?: boolean;
  }

  let { value = $bindable(false) }: Props = $props();
</script>

<!-- svelte-ignore a11y_consider_explicit_label -->
<button
  class="relative flex bg-accent-bg flex-shrink-0 items-center transition-all duration-200 border rounded-full cursor-pointer h-7 w-14 border-white/20"
  class:active={value}
  onclick={() => (value = !value)}
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
