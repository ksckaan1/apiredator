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
  import { goto } from "$app/navigation";
  import Button from "$components/ui/Button.svelte";
  import type { models } from "$lib/wailsjs/go/models";
  import { prettyTime } from "$utils/time";
  import StatusTiles from "./parts/StatusTiles.svelte";
  import UrlInfo from "./parts/URLInfo.svelte";
  import PinButton from "./parts/PinButton.svelte";
  import { showToast } from "$stores/toast";
  import AddBookmarkModal from "./parts/AddBookmarkModal.svelte";
    import ChartView from "$components/ui/ChartView.svelte";

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;

  let pageScroll = $state(0);
  let isPinned = $state(true);
  let showAddBookmarkModal = $state(false);

  // request infos
  let testType = $state("");
  let requestMethod = $state("");
  let requestURL = $state("");
  let numberOfClients = $state(0);
  let numberOfRequests = $state(0);
  let testDuration = $state("");

  // request statuses
  let chartData: number[] = $state([]);
  let isFinished = $state(false);
  let sentCount = $state(0);
  let passedDuration = $state("0s");
  let latestRPS: number = $state(0);
  let minRPS: number = $state(0);
  let avgRPS: number = $state(0);
  let maxRPS: number = $state(0);
  let statusCodes: any = $state({});
  let startedAt = $state("");
  let endedAt = $state("");

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

  $effect(() => {
    (async () => {
            let cr = await GetCurrentRequest();
      if (cr) {
        setCurrentRequest(cr);
      }
      getStats();
    })();
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
    chartData = st.rps.list;
    sentCount = st.sent_count;
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
    chartData = [];
    sentCount = 0;
    passedDuration = "0s";
    latestRPS = 0;
    minRPS = 0;
    avgRPS = 0;
    maxRPS = 0;
    statusCodes = {};
    startedAt = "";
    endedAt = "";
  };

  const onRetryButtonClicked = async () => {
    resetStatuses();
    await StartCurrentRequest();
    let cr = await GetCurrentRequest();
    if (cr) {
      setCurrentRequest(cr);
    }
    chartData = [];
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

  const addBookmark = async (bookmarkTitle: string, bookmarkTags: string[]) => {
    await AddToBookmark(bookmarkTitle, bookmarkTags)
      .then(() => {
        showToast({
          message: "Bookmark created",
          type: "success",
        })        
      })
      .catch((e: Error) => {
        showToast({
          message: e.message,
          type: "error",
        }) 
      });
  }

  const onBookmarkButtonClicked = async () => {
    showAddBookmarkModal = true;
  }
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
        <Button
          variant="secondary"
          onclick={onBackButtonClicked}
          icon="ic:outline-arrow-back"
        >
          Back to Request
        </Button>
        <div class="flex items-center gap-3">
          {#if !isFinished}
            <Button
            variant="danger"
              icon="carbon:stop-outline"
              onclick={onStopBtnClicked}
            >
              Stop
            </Button>
          {:else}
            <Button
              variant="transparent"
              icon="bi:bookmark-plus"
              onclick={onBookmarkButtonClicked}
            />
            <Button
            icon="ic:round-refresh"
            variant="success"
            onclick={onRetryButtonClicked}
            >
              Retry
            </Button>
          {/if}
          <Button 
          onclick={onNewRequestButtonClicked} 
          icon="mdi:globe-plus"
          >New Request</Button>
        </div>
      </div>
    </div>
  </div>

  <div
    in:fly={{ delay: 400, duration: 200, y: -50 }}
    class="w-screen fixed top-24 h-80"
    style="--wails-draggable:drag"
  >
    {#if chartData.length > 0}
      <ChartView data={chartData} />
    {/if}
  </div>

  <div
    class="pointer-events-none w-full relative z-50 flex-1 overflow-y-auto flex flex-col hide-scrollbar"
    onscroll={(e: any) => {
      pageScroll = e.target.scrollTop || 0;
    }}
  >
    <div class="w-full">
      <div class="h-80" style="--wails-draggable:drag"></div>
      <div
        class="w-full top-0 content-side pointer-events-auto"
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
      <div class="w-screen bg-default-bg pointer-events-auto">
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


<AddBookmarkModal
  showModal={showAddBookmarkModal}
  {addBookmark}
/>

<style lang="postcss">
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
