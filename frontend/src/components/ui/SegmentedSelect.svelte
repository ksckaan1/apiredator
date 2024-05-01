<script lang="ts">
  import { onMount } from "svelte";

  export let items: { title: string; value: string }[] = [
    {
      title: "None",
      value: "none",
    },
    {
      title: "Form Data",
      value: "format-data",
    },
    {
      title: "Raw",
      value: "raw",
    },
    {
      title: "Binary",
      value: "binary",
    },
  ];
  export let activeItem = "none";

  let parentElem: HTMLDivElement;

  let indicatorLeft = 0;
  let indicatorWidth = 0;

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

  $: {
    if (activeItem) {
      activateItem(activeItem);
    }
  }

  onMount(() => {
    activateItem(activeItem);
  });
</script>

<div
  bind:this={parentElem}
  class="rounded border border-white/20 flex h-10 overflow-hidden relative"
>
  <div
    style="--indicator-left: {indicatorLeft}px; --indicator-width: {indicatorWidth}px"
    class=" indicator"
  ></div>
  {#each items as item}
    <button
      class:active={item.value === activeItem}
      class="flex-1 z-20 px-4"
      on:click={() => (activeItem = item.value)}
    >
      {item.title}
    </button>
  {/each}
</div>

<style lang="postcss">
  .active {
    @apply text-default-bg transition-all duration-200;
  }

  .indicator {
    @apply absolute bg-primary h-full top-0 transition-all duration-200 z-10;
    width: var(--indicator-width);
    left: var(--indicator-left);
  }
</style>
