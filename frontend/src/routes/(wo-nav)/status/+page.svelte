<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import Chart from "chart.js/auto";
  import {
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
  import type { domain } from "$lib/wailsjs/go/models";
  import { prettyTime } from "$utils/time";
  import StatusTiles from "./parts/StatusTiles.svelte";
  import UrlInfo from "./parts/URLInfo.svelte";

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;

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

    gradient.addColorStop(0, "rgba(225, 175, 75, .6)"); // Başlangıç rengi
    gradient.addColorStop(1, "rgb(13, 15, 24)"); // Bitiş rengi

    ch.data.datasets[0].backgroundColor = gradient;
    ch.resize();
  };

  onMount(async () => {
    ch = new Chart(chartElem, {
      type: "line",
      options: {
        aspectRatio: 16 / 4,
        responsive: true,
        layout: {
          padding: {
            top: 100,
          },
        },
        scales: {
          x: {
            display: false, // X eksenini gizle
          },
          y: {
            display: false, // Y eksenini gizle
          },
        },
        plugins: {
          legend: {
            display: false, // Açıklamayı gizle
          },
          tooltip: {
            enabled: false, // Araç ipuçlarını kapat
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

    gradient.addColorStop(0, "rgba(225, 175, 75, .6)"); // Başlangıç rengi
    gradient.addColorStop(1, "rgb(13, 15, 24)"); // Bitiş rengi

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

  const setCurrentRequest = (cr: domain.Data) => {
    testType = cr.options.test_type;
    requestURL = cr.request.url;
    requestMethod = cr.request.method;
    testDuration = cr.options.test_duration;
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
</script>

<div
  in:fade={{ delay: 200, duration: 200 }}
  out:fade={{ duration: 200 }}
  class="flex flex-col h-screen"
>
  <div class="bg-accent-bg" style="--wails-draggable:drag">
    <div class="w-full p-5 pt-10 mx-auto max-w-7xl flex-1 flex flex-col">
      <div
        in:fly={{ delay: 200, duration: 300, y: -50 }}
        class="flex items-center justify-between"
      >
        <button
          on:click={onBackButtonClicked}
          class="bg-default-bg rounded px-5 py-2 border border-white/20"
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
    </div>
  </div>

  <div
    in:fly={{ delay: 400, duration: 200, y: -50 }}
    class="bg-accent-bg w-screen"
    style="--wails-draggable:drag"
  >
    <canvas
      class="w-full bg-gradient-to-b from-accent-bg to-default-bg"
      bind:this={chartElem}
    ></canvas>
  </div>

  <div
    class="w-full p-5 mx-auto max-w-7xl flex-1 overflow-y-auto flex flex-col hide-scrollbar"
  >
    <div in:fly={{ duration: 200, delay: 600, y: -50 }}>
      <UrlInfo {requestMethod} {requestURL} />
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

<style>
  .hide-scrollbar::-webkit-scrollbar {
    width: 0px;
    display: none;
  }
</style>
