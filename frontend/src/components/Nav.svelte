<script lang="ts">
  import { goto, onNavigate } from "$app/navigation";
  import { page } from "$app/stores";
  import { screenStore } from "../store/screen";
  import Icon from "@iconify/svelte";

  let activeNavBarItemIndex = $state(-1);
  let showIndicator = $state(false);

  $effect(() => {
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
      icon: "mdi:globe",
      rgx: /\/new\-request/,
    },
    {
      id: "bookmarks",
      title: "Bookmarks",
      icon: "material-symbols:bookmark-outline",
      rgx: /\/bookmarks/,
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
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="z-20 flex items-center justify-center w-full cursor-pointer aspect-square"
      onclick={() => {
        $screenStore.active = item.id;
        goto(`/${item.id}`);
      }}
    >
      <Icon
        class={`${activeNavBarItemIndex == index ? "text-primary" : "text-white/30"}`}
        icon={item.icon}
        width="40"
      />
    </button>
  {/each}
</nav>

<style>
  .indicator {
    top: var(--bar-top);
    transition: top 0.2s;
  }
</style>
