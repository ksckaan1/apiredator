<script lang="ts">
  import { onMount } from "svelte";

  interface Props {
    tabs?: any[];
    value?: string;
  }

  let {
    tabs = [
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
    ],
    value = $bindable(""),
  }: Props = $props();

  let wrapper: HTMLDivElement;

  let indicatorLeft = $state(0);
  let indicatorWidth = $state(0);

  $effect(() => {
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
      onclick={() => onTabClicked(tab.value)}
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
