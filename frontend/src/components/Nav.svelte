<script lang="ts">
  import { goto, onNavigate } from "$app/navigation";
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import { screenStore } from "../store/screen";

  let activeNavBarItemIndex = -1;
  let showIndicator = false;

  onMount(() => {
    setNavItem($page.route.id ?? "");
  });

  onNavigate((nav) => {
    setNavItem(nav.to?.route.id ?? $page.route.id ?? "");
  });

  const setNavItem = (target: string) => {
    activeNavBarItemIndex = navBarItems.findIndex((i) => i.rgx.test(target));
    showIndicator = activeNavBarItemIndex != -1;
  };

  const navBarItems = [
    {
      id: "new-request",
      title: "New Request",
      icon: "fa-solid fa-globe",
      rgx: /\/new\-request/,
    },
    {
      id: "history",
      title: "History",
      icon: "fa-solid fa-clock-rotate-left",
      rgx: /\/history/,
    },
    {
      id: "environment",
      title: "Environments",
      icon: "fa-solid fa-recycle",
      rgx: /\/environment/,
    },
  ];
</script>

<nav
  class="relative w-16 h-[calc(100vh-2.5rem)] border-r bg-default-bg border-white/20"
>
  {#if showIndicator}
    <div
      style="--bar-top: {activeNavBarItemIndex * 60}px"
      class="absolute left-0 z-10 flex w-[2px] h-[60px] bg-primary pointer-events-none indicator aspect-square"
    ></div>
  {/if}
  {#each navBarItems as item, index}
    <button
      class="z-20 flex items-center justify-center w-full cursor-pointer aspect-square"
      on:click={() => {
        $screenStore.active = item.id;
        goto(`/${item.id}`);
      }}
      data-tooltip-message={item.title}
      data-tooltip-position="right"
    >
      <i
        class="{item.icon} {index == activeNavBarItemIndex
          ? 'text-primary'
          : 'text-zinc-500'} text-3xl transition-colors duration-200"
        data-tooltip-message={item.title}
        data-tooltip-position="right"
      ></i>
    </button>
  {/each}
</nav>

<style>
  .indicator {
    top: var(--bar-top);
    transition: top 0.2s;
  }
</style>
