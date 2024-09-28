<script lang="ts">
  import DropdownSelect from "$components/ui/DropdownSelect.svelte";
  import TextInput from "$components/ui/TextInput.svelte";
  import type { models } from "$lib/wailsjs/go/models";
  import { flip } from "svelte/animate";
  import { fade, slide } from "svelte/transition";
  import {
    GetAllBookmarks,
    GetAllTags,
  } from "$lib/wailsjs/go/service/AppService";
  import Icon from "@iconify/svelte";

  let bookmarks: models.Bookmark[] = $state([]);
  let bookmarkCount = $state(0);
  let searchValue = $state("");
  let tagValue = $state("");
  let allTags: any[] = $state([]);

  let bouncedSearchValue = $state("");
  let selectedBookmarks: string[] = $state([]);

  let timer: number;

  $effect(() => {
    if (searchValue != undefined) {
      clearTimeout(timer);
      timer = setTimeout(() => {
        bouncedSearchValue = searchValue;
      }, 750);
    }
  });

  const makeSearch = () => {
    GetAllBookmarks(bouncedSearchValue, tagValue, -1, 0).then((result) => {
      bookmarks = result.bookmarks;
      bookmarkCount = result.count;
      console.log(result);
    });
  };

  makeSearch();

  $effect(() => {
    if (bouncedSearchValue != undefined && tagValue != undefined) {
      makeSearch();
    }
  });

  GetAllTags().then((result) => {
    allTags = [
      { title: "All Tags", value: "" },
      ...result.map((t) => ({ value: t, title: t })),
    ];
  });

  let requestMethods = [
    {
      value: "GET",
      title: "GET",
      color: "bg-green-700",
    },
    {
      value: "POST",
      title: "POST",
      color: "bg-amber-700",
    },
    {
      value: "PUT",
      title: "PUT",
      color: "bg-sky-700",
    },
    {
      value: "PATCH",
      title: "PATCH",
      color: "bg-purple-700",
    },
    {
      value: "DELETE",
      title: "DELETE",
      color: "bg-red-700",
    },
    {
      value: "HEAD",
      title: "HEAD",
      color: "bg-pink-700",
    },
    {
      value: "TRACE",
      title: "TRACE",
      color: "bg-teal-700",
    },
    {
      value: "OPTIONS",
      title: "OPTIONS",
      color: "bg-emerald-700",
    },
  ];
</script>

<div class="h-[calc(100vh-40px)] flex flex-col">
  <div class="flex-1 hide-scrollbar overflow-y-auto flex flex-col px-5 pt-5">
    <div class="mx-auto w-full max-w-7xl flex flex-col gap-4">
      <h1 class="text-3xl">Bookmarks</h1>
      <div class="flex gap-2">
        <div class="flex-1">
          <TextInput bind:value={searchValue} label="Search" autoFocus />
        </div>
        <div class="w-48">
          <DropdownSelect items={allTags} bind:value={tagValue} />
        </div>
      </div>
      {#if bookmarks && bookmarks.length > 0}
        <div
          class="flex flex-col pb-5 pr-1 gap-4 h-[calc(100vh-10.625rem)] hide-scrollbar overflow-y-scroll"
        >
          {#if selectedBookmarks.length > 0}
            <div
              class="sticky top-0 py-2 border-b flex justify-between items-center border-white/20 bg-default-bg z-40"
              transition:slide={{ duration: 200 }}
            >
              <span>
                {selectedBookmarks.length} bookmark(s) selected
              </span>
              <div class="flex items-start gap-5">
                <button
                  onclick={() => {
                    selectedBookmarks = [];
                  }}
                  class="flex gap-1"
                >
                  <Icon icon="bx:checkbox" width="24" />
                  <span> Deselect All </span>
                </button>
                <button
                  onclick={() => {
                    selectedBookmarks = [];
                  }}
                  class="flex gap-1"
                >
                  <Icon icon="ph:trash" width="24" />
                  <span> Delete Selected Ones </span>
                </button>
              </div>
            </div>
          {/if}
          {#each bookmarks as bookmark (bookmark.id)}
            <div
              animate:flip={{ duration: 200 }}
              transition:fade={{ duration: 200 }}
              class="flex border overflow-hidden bg-accent-bg border-white/20 rounded flex-shrink-0 hover:bg-white/5 cursor-pointer"
            >
              <div
                class={`${
                  requestMethods.find(
                    (r) => r.value === bookmark.request.method,
                  )?.color
                } flex w-8 pl-1 items-center flex-shrink-0`}
              >
                <div class="[writing-mode:vertical-lr] rotate-180">
                  {bookmark.request.method}
                </div>
              </div>
              <div class="p-5 flex-1">
                <div class="text-lg text-primary">
                  {bookmark.title || "Untitled Bookmark"}
                </div>
                <div class="flex gap-3">
                  <div>
                    {bookmark.request.url}
                  </div>
                </div>
                {#if bookmark.tags && bookmark.tags.length > 0}
                  <div class="mt-1 flex flex-wrap gap-3 items-center">
                    <Icon icon="mdi:tags" class="text-white/50" />
                    {#each bookmark.tags as tag}
                      <button
                        onclick={() => {
                          tagValue = tag;
                        }}
                        class=" text-white/50 transition-colors duration-200 hover:text-primary"
                      >
                        #{tag}
                      </button>
                    {/each}
                  </div>
                {/if}
              </div>
              <div class="flex flex-col justify-center pr-3 flex-shrink-0">
                <input
                  type="checkbox"
                  bind:group={selectedBookmarks}
                  id={bookmark.id}
                  value={bookmark.id}
                  class="hidden"
                />
                <label for={bookmark.id} class="cursor-pointer">
                  {#if selectedBookmarks.includes(bookmark.id)}
                    <Icon icon="bxs:checkbox" width="40" class="text-primary" />
                  {:else}
                    <Icon icon="bx:checkbox" width="40" class="text-white/30" />
                  {/if}
                </label>
              </div>
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </div>
</div>
