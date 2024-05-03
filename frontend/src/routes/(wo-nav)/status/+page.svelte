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
  import { prettyTime } from "$utils/time";
  import RequestStatus from "./parts/RequestStatus.svelte";
  import DurationStatus from "./parts/DurationStatus.svelte";
  import RpsStatus from "./parts/RPSStatus.svelte";
  import TimeStatus from "./parts/TimeStatus.svelte";
  import StatusCodes from "./parts/StatusCodes.svelte";
  import StatusTiles from "./parts/StatusTiles.svelte";

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;

  // request infos
  let testType = "";
  let requestMethod = "";
  let requestURL = "";
  let numberOfClients = 0;
  let numberOfRequests = 0;
  let targetDuration = 0;

  // request statuses
  let isFinished = false;
  let sentCount = 0;
  let passedDuration = 0;
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
      setCurrentRequest(cr);
    }
    getStats();
  });

  const setCurrentRequest = (cr: domain.Data) => {
    testType = cr.options.test_type;
    requestURL = cr.request.url;
    requestMethod = cr.request.method;
    targetDuration =
      cr.options.duration.hours * 60 * 60 +
      cr.options.duration.minutes * 60 +
      cr.options.duration.seconds;
    numberOfClients = cr.options.number_of_clients;
    numberOfRequests = cr.options.number_of_requests;
  };

  const updateStats = (st: domain.Stat) => {
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
    passedDuration = Math.floor(st.passed_duration / Math.pow(10, 9));
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
      setCurrentRequest(cr);
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
    class="w-full p-5 mx-auto max-w-7xl flex-1 overflow-y-auto flex flex-col"
  >
    <div
      class="border border-white/20 rounded px-3 py-2 h-18 w-full flex gap-3"
    >
      <div class={requestMethods.find((r) => r.value === requestMethod)?.color}>
        {requestMethod}
      </div>
      <div>{requestURL}</div>
    </div>
    <StatusTiles
      {testType}
      {sentCount}
      {numberOfRequests}
      {numberOfClients}
      {passedDuration}
      {targetDuration}
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
