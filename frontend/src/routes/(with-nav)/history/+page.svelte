<script lang="ts">
  import DropdownSelect from "$components/ui/DropdownSelect.svelte";
  import TextInput from "$components/ui/TextInput.svelte";
  import type { models } from "$lib/wailsjs/go/models";
  import { flip } from "svelte/animate";
  import { fade } from "svelte/transition";
  import {
    GetAllBookmarks,
    GetAllTags,
  } from "$lib/wailsjs/go/service/AppService";

  let bookmarks: models.Bookmark[] = [];
  let bookmarkCount = 0;
  let searchValue = "";
  let tagValue = "";
  let allTags: any[] = [];

  let bouncedSearchValue = "";

  let timer: number;

  $: {
    if (searchValue != undefined) {
      clearTimeout(timer);
      timer = setTimeout(() => {
        bouncedSearchValue = searchValue;
      }, 750);
    }
  }

  const makeSearch = () => {
    GetAllBookmarks(bouncedSearchValue, tagValue, -1, 0).then((result) => {
      bookmarks = result.bookmarks;
      bookmarkCount = result.count;
      console.log(result);
    });
  };

  makeSearch();

  $: {
    if (bouncedSearchValue != undefined && tagValue != undefined) {
      makeSearch();
    }
  }

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
      color: "text-green-400",
    },
    {
      value: "POST",
      title: "POST",
      color: "text-amber-400",
    },
    {
      value: "PUT",
      title: "PUT",
      color: "text-sky-400",
    },
    {
      value: "PATCH",
      title: "PATCH",
      color: "text-purple-400",
    },
    {
      value: "DELETE",
      title: "DELETE",
      color: "text-red-400",
    },
    {
      value: "HEAD",
      title: "HEAD",
      color: "text-pink-400",
    },
    {
      value: "TRACE",
      title: "TRACE",
      color: "text-teal-400",
    },
    {
      value: "OPTIONS",
      title: "OPTIONS",
      color: "text-emerald-400",
    },
  ];
</script>

<div class="h-[calc(100vh-40px)] flex flex-col">
  <div class="flex-1 hide-scrollbar overflow-y-auto flex flex-col p-5">
    <div class="mx-auto w-full max-w-7xl flex flex-col gap-4">
      <div class="flex gap-2">
        <div class="flex-1">
          <TextInput bind:value={searchValue} label="Search" autoFocus />
        </div>
        <div class="w-48">
          <DropdownSelect items={allTags} bind:value={tagValue} />
        </div>
      </div>
      {#each bookmarks as bookmark (bookmark.id)}
        <div
          animate:flip={{ duration: 200 }}
          transition:fade={{ duration: 200 }}
          class="flex flex-col border border-white/20 rounded p-5 flex-shrink-0 hover:bg-white/5 cursor-pointer"
        >
          <div class="text-lg text-primary">
            {bookmark.title || "Untitled Bookmark"}
          </div>
          <div class="flex gap-3">
            <div
              class={requestMethods.find(
                (r) => r.value === bookmark.request.method,
              )?.color}
            >
              {bookmark.request.method}
            </div>
            <div>
              {bookmark.request.url}
            </div>
          </div>
          {#if bookmark.tags && bookmark.tags.length > 0}
            <div class="mt-1 flex flex-wrap gap-3">
              {#each bookmark.tags as tag}
                <button
                  on:click={() => {
                    tagValue = tag;
                  }}
                  class="bg-accent-bg text-white/50 px-2 py-1 rounded border border-white/20 transition-colors duration-200 hover:bg-primary hover:text-on-primary"
                >
                  {tag}
                </button>
              {/each}
            </div>
          {/if}
        </div>
      {/each}
    </div>
  </div>
</div>
