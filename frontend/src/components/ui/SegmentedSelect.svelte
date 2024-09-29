<script lang="ts">
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";

  interface Props {
    items: { title: string; value: string; icon?: string }[];
    activeItem: string;
  }

  let { items = $bindable([]), activeItem = $bindable("") }: Props = $props();

  let parentElem: HTMLDivElement;

  let indicatorLeft = $state(0);
  let indicatorWidth = $state(0);
  let isAnimActive = $state(false);

  const activateItem = (value: string) => {
    const activeItemIndex = items.findIndex((item) => item.value === value);
    if (activeItemIndex === -1) return;
    if (!parentElem) return;
    const target = parentElem.querySelectorAll("button")[activeItemIndex];
    if (!target) return;
    indicatorLeft =
      parentElem.querySelectorAll("button")[activeItemIndex].offsetLeft;
    indicatorWidth =
      parentElem.querySelectorAll("button")[activeItemIndex].offsetWidth;
  };

  $effect(() => {
    if (activeItem) {
      activateItem(activeItem);
    }
  });

  $effect(() => {
    activateItem(activeItem);
  });

  const selectItem = (val: string) => {
    isAnimActive = true;
    activeItem = val;
  };
</script>

<div
  bind:this={parentElem}
  class="rounded border border-white/20 bg-accent-bg font-extralight text-white/40 flex h-10 overflow-hidden relative"
>
  <div
    style="--indicator-left: {indicatorLeft}px; --indicator-width: {indicatorWidth}px"
    class="indicator"
    class:anim={isAnimActive}
  ></div>
  {#each items as item}
    <button
      class:active={item.value === activeItem}
      class="flex-1 z-20 px-4 flex gap-x-2 items-center"
      onclick={() => selectItem(item.value)}
    >
      {#if item.icon}
        <Icon icon={item.icon} width="24" />
      {/if}
      <span>
        {item.title}
      </span>
    </button>
  {/each}
</div>

<style lang="postcss">
  .active {
    @apply text-on-primary font-normal transition-all duration-200;
  }

  .anim {
    @apply transition-all duration-200;
  }

  .indicator {
    @apply absolute bg-primary h-full top-0  z-10;
    width: var(--indicator-width);
    left: var(--indicator-left);
  }
</style>
