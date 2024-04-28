<script lang="ts">
  import Chart from "chart.js/auto";
  import {
    GetCurrentRequest,
    GetStats,
    ResetCurrentRequest,
    StartCurrentRequest,
    StopWork,
    WaitWork,
  } from "$lib/wailsjs/go/service/AppService";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import Button from "$components/ui/Button.svelte";
  import type { domain } from "$lib/wailsjs/go/models";
  import { convertSeconds, prettyTime } from "$utils/time";

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;
  let currentRequest: domain.Data;
  let isFinished = false;
  let completedRequest = 0;
  let passedDuration = 0;
  let rpsValues: number[] = [];
  let statusCodes: any = {};
  let startedAt = "";
  let endedAt = "";

  let datasets = [
    {
      label: "Request per second",
      data: [],
      fill: true,
      borderColor: "rgb(225, 175, 75)",
      tension: 0.5,
    },
  ];

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

  onMount(async () => {
    ch = new Chart(chartElem, {
      type: "line",
      options: {
        aspectRatio: 16 / 4,
      },
      data: {
        labels: [],
        datasets: datasets,
      },
    });
    let cr = await GetCurrentRequest();
    if (cr) {
      currentRequest = cr;
    }
    getStats();
  });

  const updateStats = (st: domain.Stat) => {
    completedRequest = st.completed;
    let rps = st.request_per_second;
    rpsValues = rps;
    ch.data.datasets[0].data = rpsValues;
    ch.data.labels = rps.map((r) => "");
    ch.update();
    passedDuration = Math.floor(st.duration / Math.pow(10, 9));
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

  const onRetryButtonClicked = async () => {
    await StartCurrentRequest();
    let cr = await GetCurrentRequest();
    if (cr) {
      currentRequest = cr;
    }
    rpsValues = [];
    ch.data.labels = [];
    isFinished = false;
    getStats();
  };

  const onBookmarkButtonClicked = async () => {};

  const onNewRequestButtonClicked = async () => {
    if (!isFinished) {
      await StopWork();
    }
    await ResetCurrentRequest();
    goto("/new-request", {
      replaceState: true,
    });
  };

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

<div class="flex flex-col h-screen">
  <div class="bg-accent-bg" style="--wails-draggable:drag">
    <div class="w-full p-5 pt-10 mx-auto max-w-7xl flex-1 flex flex-col">
      <div class="flex items-center justify-between">
        <button on:click={onBackButtonClicked}>&larr; back to request</button>
        <div class="flex items-center gap-3">
          {#if !isFinished}
            <button
              on:click={onStopBtnClicked}
              class="bg-red-900 rounded px-5 py-2 border border-white/20"
            >
              Stop
            </button>
          {:else}
            <i class="fa-regular fa-bookmark text-2xl"></i>
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
      <div class="bg-default-bg w-full rounded-lg mt-5 pl-2 chartWrapper">
        <canvas class="w-full" bind:this={chartElem}></canvas>
      </div>
    </div>
  </div>

  <div
    class="w-full p-5 mx-auto max-w-7xl flex-1 overflow-y-scroll flex flex-col"
  >
    <div
      class="border border-white/20 rounded px-3 py-2 h-18 w-full flex gap-3"
    >
      <div
        class={requestMethods.find(
          (r) => r.value === currentRequest?.request.method,
        )?.color}
      >
        {currentRequest?.request.method}
      </div>
      <div>{currentRequest?.request.url}</div>
    </div>
    <div class="grid grid-cols-3 gap-5 mt-5">
      <div class="tile">
        <h1>Requests</h1>
        <span>
          {completedRequest}
          {#if !currentRequest?.options.duration.is_duration_active}
            / {currentRequest?.options.number_of_requests *
              currentRequest?.options.number_of_clients}
          {/if}
        </span>
        {#if !currentRequest?.options.duration.is_duration_active && currentRequest?.options.number_of_clients > 1}
          <div class="text-sm text-white/50 p-3 pt-0">
            {currentRequest?.options.number_of_clients} clients X {currentRequest
              ?.options.number_of_requests} requests
          </div>
        {/if}
      </div>
      <div class="tile">
        <h1>Duration</h1>
        <span>
          {convertSeconds(passedDuration)}
          {#if currentRequest?.options.duration.is_duration_active}
            / {convertSeconds(
              currentRequest.options.duration.hours * 60 * 60 +
                currentRequest.options.duration.minutes * 60 +
                currentRequest.options.duration.seconds,
            )}
          {/if}
        </span>
        {#if currentRequest?.options.duration.is_duration_active && currentRequest?.options.number_of_clients > 1}
          <div class="text-sm text-white/50 p-3 pt-0">
            {currentRequest?.options.number_of_clients} clients X {convertSeconds(
              currentRequest.options.duration.hours * 60 * 60 +
                currentRequest.options.duration.minutes * 60 +
                currentRequest.options.duration.seconds,
            )}
          </div>
        {/if}
      </div>
      <div class="tile">
        <h1>Requests per second</h1>
        <span>
          {rpsValues.length > 0 ? rpsValues[rpsValues.length - 1] : "0"} rps
        </span>
      </div>
      <div class="tile">
        <h1>Time</h1>
        <div class="px-3 mt-3 text-white/50">Started at:</div>
        <span>
          {startedAt}
        </span>
        <div class="px-3 mt-3 text-white/50">Ended at:</div>
        <span class="mb-5">
          {endedAt != "" ? endedAt : "not ended yet"}
        </span>
      </div>
      <div class="tile">
        <h1>Status Codes</h1>
        {#if statusCodes}
          <div class="flex flex-col gap-2 px-3">
            {#each Object.keys(statusCodes) as sc}
              <div class="flex justify-between items-center gap-2">
                <div>
                  {sc}
                </div>
                <div
                  class="flex-1 border-b border-dashed border-white/20"
                ></div>
                <div>
                  {statusCodes[sc]}x
                </div>
              </div>
            {/each}
          </div>
        {/if}
      </div>
    </div>
  </div>
</div>

<style lang="postcss">
  .tile {
    @apply border border-white/20 rounded bg-accent-bg;
  }

  .tile h1 {
    @apply text-xl opacity-50 p-3 pb-0 font-extralight;
  }

  .tile span {
    @apply text-2xl font-light px-3 block;
  }
</style>
