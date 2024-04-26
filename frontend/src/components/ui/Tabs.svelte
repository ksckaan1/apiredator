<script lang="ts">
  import { onMount } from "svelte";

  export let tabs: any[] = [
    {
      title: "Options",
      value: "options",
    },
    {
      title: "Header",
      value: "header",
    },
    {
      title: "Body",
      value: "body",
    },
  ];

  export let value: string = "";

  let wrapper: HTMLDivElement;

  let indicatorLeft = 0;
  let indicatorWidth = 0;

  onMount(() => {
    if (value) {
      const [left, width] = getTabOffset(value);
      indicatorLeft = left;
      indicatorWidth = width;
    }
  });

  const getTabOffset = (tabValue: string) => {
    let activeIndex = tabs.findIndex((tab) => tab.value == tabValue);
    if (activeIndex >= 0) {
      return [
        wrapper.querySelectorAll("button")[activeIndex].offsetLeft,
        wrapper.querySelectorAll("button")[activeIndex].offsetWidth,
      ];
    }
    return [0, 0];
  };

  const onTabClicked = (tabValue: string) => {
    const [left, width] = getTabOffset(tabValue);
    indicatorLeft = left;
    indicatorWidth = width;
    value = tabValue;
  };
</script>

<div bind:this={wrapper} class="relative border-b border-white/20">
  {#if indicatorWidth > 0}
    <div
      style="--left: {indicatorLeft}px; --width: {indicatorWidth}px"
      class="indicator"
    ></div>
  {/if}
  {#each tabs as tab}
    <button
      class="px-4 py-2 hover:bg-white/5"
      on:click={() => onTabClicked(tab.value)}
    >
      {tab.title}
    </button>
  {/each}
</div>

<style lang="postcss">
  .indicator {
    @apply absolute h-0.5 bg-primary bottom-0 transition-all duration-200;
    left: var(--left);
    width: var(--width);
  }
</style>
