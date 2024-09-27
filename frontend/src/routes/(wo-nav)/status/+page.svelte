<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import { flip } from "svelte/animate";
  import Chart from "chart.js/auto";
  import {
    AddToBookmark,
    GetCurrentRequest,
    GetStats,
    ResetCurrentRequest,
    StartCurrentRequest,
    StopWork,
    WaitWork,
  } from "$lib/wailsjs/go/service/AppService";
  import { onDestroy, onMount } from "svelte";
  import { goto } from "$app/navigation";
  import Button from "$components/ui/Button.svelte";
  import type { models } from "$lib/wailsjs/go/models";
  import { prettyTime } from "$utils/time";
  import StatusTiles from "./parts/StatusTiles.svelte";
  import UrlInfo from "./parts/URLInfo.svelte";
  import PinButton from "./parts/PinButton.svelte";
  import BookmarkButton from "./parts/BookmarkButton.svelte";
  import TextInput from "$components/ui/TextInput.svelte";

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;

  let pageScroll = 0;
  let isPinned = true;
  let isBookmarkCreated = false;
  let isBookmarkTitleModalShown = false;
  let bookmarkTitle = "";
  let bookmarkTag = "";
  let bookmarkTags: string[] = [
    "job",
    "custom api",
    "dyb api",
    "appman",
    "test1",
    "test6",
    "test7",
    "test8",
    "test9",
    "test10",
  ];

  let allTags = [
    "Example",
    "Naber",
    "Nasılsın",
    "How are you?",
    "Ne var yu?",
    "test1",
    "test6",
    "test7",
    "test8",
    "test9",
    "test10",
    "test11",
    "test12",
  ];

  // request infos
  let testType = "";
  let requestMethod = "";
  let requestURL = "";
  let numberOfClients = 0;
  let numberOfRequests = 0;
  let testDuration = "";

  // request statuses
  let isFinished = false;
  let sentCount = 0;
  let passedDuration = "0s";
  let rpsValues: number[] = [];
  let latestRPS: number = 0;
  let minRPS: number = 0;
  let avgRPS: number = 0;
  let maxRPS: number = 0;
  let statusCodes: any = {};
  let startedAt = "";
  let endedAt = "";

  let datasets = [
    {
      label: "Request per second",
      data: [],
      fill: true,
      borderColor: "rgb(225, 175, 75)",
      borderWidth: 0,
      pointRadius: 0,
      pointBorderWidth: 0,
      tension: 0.5,
    },
  ];

  const removeTag = (tag: string) => {
    bookmarkTags = bookmarkTags.filter((t) => t !== tag);
  };

  const getStats = async () => {
    let interval = setInterval(async () => {
      stats = await GetStats();
      updateStats(stats);
    }, 1000);
    await WaitWork();
    clearInterval(interval);
    stats = await GetStats();
    updateStats(stats);
    isFinished = true;
  };

  const resizeEvent = () => {
    var gradient = chartElem
      .getContext("2d")!
      .createLinearGradient(0, 0, 0, chartElem.height);

    gradient.addColorStop(0, "rgba(225, 175, 75, .6)");
    gradient.addColorStop(1, "rgb(13, 15, 24)");

    ch.data.datasets[0].backgroundColor = gradient;
    ch.resize();
  };

  onMount(async () => {
    ch = new Chart(chartElem, {
      type: "line",
      options: {
        responsive: true,
        maintainAspectRatio: false,
        layout: {
          padding: {
            top: 100,
          },
        },
        scales: {
          x: {
            display: false,
          },
          y: {
            display: false,
          },
        },
        plugins: {
          legend: {
            display: false,
          },
          tooltip: {
            enabled: false,
          },
        },
      },
      data: {
        labels: [],
        datasets: datasets,
      },
    });

    var gradient = chartElem
      .getContext("2d")!
      .createLinearGradient(0, 0, 0, chartElem.height);

    gradient.addColorStop(0, "rgba(225, 175, 75, .6)");
    gradient.addColorStop(1, "rgb(13, 15, 24)");

    ch.data.datasets[0].backgroundColor = gradient;
    ch.update();

    window.addEventListener("resize", resizeEvent);

    let cr = await GetCurrentRequest();
    if (cr) {
      setCurrentRequest(cr);
    }
    getStats();
  });

  onDestroy(() => {
    window.removeEventListener("resize", resizeEvent);
  });

  const setCurrentRequest = (cr: models.Data) => {
    testType = cr.options.test_type;
    requestURL = cr.request.url;
    requestMethod = cr.request.method;
    testDuration = cr.options.test_duration;
    numberOfClients = cr.options.number_of_clients;
    numberOfRequests = cr.options.number_of_requests;
  };

  const updateStats = (st: models.Stat) => {
    sentCount = st.sent_count;
    let rps = st.rps.list;
    rpsValues = rps;
    ch.data.datasets[0].data = rpsValues;
    ch.data.labels = rps.map((r) => "");
    ch.update();
    latestRPS = st.rps.latest;
    minRPS = st.rps.min;
    avgRPS = Number(st.rps.avg.toFixed(2));
    maxRPS = st.rps.max;
    passedDuration = st.passed_duration;
    startedAt = prettyTime(st.started_at);
    endedAt = prettyTime(st.ended_at);
    statusCodes = st.status_codes;
  };

  const onStopBtnClicked = () => {
    StopWork();
  };

  const onBackButtonClicked = async () => {
    if (!isFinished) {
      await StopWork();
    }

    goto("/new-request", {
      replaceState: true,
    });
  };

  const resetStatuses = () => {
    isBookmarkCreated = false;
    sentCount = 0;
    passedDuration = "0s";
    rpsValues = [];
    latestRPS = 0;
    minRPS = 0;
    avgRPS = 0;
    maxRPS = 0;
    statusCodes = {};
    startedAt = "";
    endedAt = "";
    ch.data.datasets[0].data = [];
    ch.data.labels = [];
    ch.update();
  };

  const onRetryButtonClicked = async () => {
    resetStatuses();
    await StartCurrentRequest();
    let cr = await GetCurrentRequest();
    if (cr) {
      setCurrentRequest(cr);
    }
    rpsValues = [];
    ch.data.labels = [];
    isFinished = false;
    getStats();
  };

  const onBookmarkButtonClicked = async () => {
    if (!isBookmarkCreated) {
      isBookmarkTitleModalShown = true;
    }
  };

  const onNewRequestButtonClicked = async () => {
    if (!isFinished) {
      await StopWork();
    }
    await ResetCurrentRequest();
    goto("/new-request", {
      replaceState: true,
    });
  };
</script>

<div
  in:fade={{ delay: 200, duration: 200 }}
  out:fade={{ duration: 200 }}
  class="flex flex-col h-screen"
>
  <div class="bg-default-bg" style="--wails-draggable:drag">
    <div class="w-full p-5 pt-10 mx-auto max-w-7xl flex-1 flex flex-col">
      <div
        in:fly={{ delay: 200, duration: 300, y: -50 }}
        class="flex items-center justify-between"
      >
        <button
          on:click={onBackButtonClicked}
          class="bg-accent-bg rounded px-5 py-2 border border-white/20"
        >
          &larr; Back to request
        </button>
        <div class="flex items-center gap-3">
          {#if !isFinished}
            <button
              on:click={onStopBtnClicked}
              class="bg-red-900 rounded px-5 py-2 border border-white/20"
            >
              Stop
            </button>
          {:else}
            <BookmarkButton
              isActive={isBookmarkCreated}
              on:click={onBookmarkButtonClicked}
            />
            <button
              on:click={onRetryButtonClicked}
              class="bg-green-900 rounded px-5 py-2 border border-white/20"
            >
              <i class="fa-solid fa-rotate-right"></i>
              Retry
            </button>
          {/if}
          <Button on:click={onNewRequestButtonClicked}>New Request</Button>
        </div>
      </div>
    </div>
  </div>

  <div
    in:fly={{ delay: 400, duration: 200, y: -50 }}
    class="w-screen fixed top-24 h-80"
    style="--wails-draggable:drag"
  >
    <canvas class="w-screen bg-default-bg" bind:this={chartElem}></canvas>
  </div>

  <div
    class="w-full relative z-50 flex-1 overflow-y-auto flex flex-col hide-scrollbar"
    on:scroll={(e) => {
      pageScroll = e.target.scrollTop || 0;
    }}
  >
    <div class="w-full">
      <div class="h-80" style="--wails-draggable:drag"></div>
      <div
        class="w-full top-0 content-side"
        class:content-side-solid={pageScroll > 320}
        class:sticky={isPinned}
        class:static={!isPinned}
      >
        <div class="w-full mx-auto max-w-7xl px-5">
          <div in:fly={{ duration: 200, delay: 600, y: -50 }}>
            <div class="flex gap-2 items-center">
              <UrlInfo {requestMethod} {requestURL} />
              <PinButton bind:isPinned />
            </div>
          </div>
          <StatusTiles
            {testType}
            {sentCount}
            {numberOfRequests}
            {numberOfClients}
            {passedDuration}
            {testDuration}
            {latestRPS}
            {minRPS}
            {avgRPS}
            {maxRPS}
            {startedAt}
            {endedAt}
            {statusCodes}
          />
        </div>
      </div>
      <div class="w-screen bg-default-bg">
        <div class="max-w-7xl mx-auto w-full flex-shrink-0 px-5">
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
          <p>asdasd</p>
        </div>
      </div>
    </div>
  </div>
</div>

{#if isBookmarkTitleModalShown}
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/50 z-50 flex items-start justify-center"
    on:click|self={() => (isBookmarkTitleModalShown = false)}
  >
    <div
      class="bg-accent-bg border border-white/20 rounded p-5 mt-20 w-full max-w-xl"
    >
      <h2 class="text-2xl mb-2">Save Bookmark</h2>
      <TextInput label="Insert Title" bind:value={bookmarkTitle} />
      {#if bookmarkTags.length > 0}
        <div class="flex flex-wrap gap-2 py-5">
          {#each bookmarkTags as tag (tag)}
            <button
              animate:flip={{ duration: 200 }}
              transition:fade={{ duration: 200 }}
              on:click={() => removeTag(tag)}
              class="bg-primary text-on-primary px-2 py-1 rounded"
            >
              {tag}
            </button>
          {/each}
        </div>
      {:else}
        <div class="py-6">No tags added</div>
      {/if}
      <TextInput
        label="Add Tag"
        autoComplete={allTags}
        bind:value={bookmarkTag}
        on:keypress={(e) => {
          if (
            e.key === "Enter" &&
            !bookmarkTags.includes(bookmarkTag) &&
            bookmarkTag.trim() != ""
          ) {
            bookmarkTags = [...bookmarkTags, bookmarkTag];
            bookmarkTag = "";
          }
        }}
      />
      <div class="flex justify-end gap-3 items-center mt-5">
        <button
          on:click={() => {
            isBookmarkTitleModalShown = false;
            bookmarkTitle = "";
            bookmarkTags = [];
          }}
          class="text-white/70">Cancel</button
        >
        <Button
          on:click={async () => {
            await AddToBookmark(bookmarkTitle, bookmarkTags)
              .then(() => {
                isBookmarkCreated = true;
                isBookmarkTitleModalShown = false;
                bookmarkTitle = "";
                bookmarkTags = [];
              })
              .catch((e) => {
                console.error(e);
              });
          }}
        >
          <div class="flex gap-1 items-center justify-center w-24">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              width="1.5rem"
              ><path
                fill="currentColor"
                d="M11 13H6q-.425 0-.712-.288T5 12t.288-.712T6 11h5V6q0-.425.288-.712T12 5t.713.288T13 6v5h5q.425 0 .713.288T19 12t-.288.713T18 13h-5v5q0 .425-.288.713T12 19t-.712-.288T11 18z"
              /></svg
            >
            <span>Add</span>
          </div></Button
        >
      </div>
    </div>
  </div>
{/if}

<style>
  .hide-scrollbar::-webkit-scrollbar {
    width: 0px;
    display: none;
  }

  .content-side {
    background: linear-gradient(
      to bottom,
      rgba(13, 15, 24, 0) 0px,
      rgba(13, 15, 24, 1) 100px,
      rgba(13, 15, 24, 1) 100px
    );
  }

  .content-side-solid {
    @apply bg-default-bg;
  }
</style>
