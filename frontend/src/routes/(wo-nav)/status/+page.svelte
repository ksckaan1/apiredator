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

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;
  let currentRequest: domain.Data;
  let isFinished = false;
  let completedRequest = 0;
  let passedDuration = 0;
  let rpsValues: number[] = [];

  let datasets = [
    {
      label: "Request per second",
      data: [],
      fill: true,
      borderColor: "rgb(75, 192, 192)",
      tension: 0.5,
    },
  ];

  const getStats = async () => {
    let last = 0;
    let interval = setInterval(async () => {
      stats = await GetStats();
      completedRequest = stats.completed;
      let rps = stats.completed - last;
      rpsValues = [...rpsValues, rps];
      ch.data.datasets[0].data = rpsValues;
      ch.data.labels?.push("");
      ch.update();
      passedDuration++;
      last = stats.completed;
    }, 1000);
    await WaitWork();
    isFinished = true;
    clearInterval(interval);
    stats = await GetStats();
    completedRequest = stats.completed;
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
    <div class="bg-default-bg w-full rounded-lg mt-5 p-5 chartWrapper">
      <canvas class="w-full" bind:this={chartElem}></canvas>
    </div>
  </div>
</div>

<div class="w-full p-5 mx-auto max-w-7xl flex-1 flex flex-col">
  <div class="border border-white/20 rounded px-3 py-2 h-18 w-full flex gap-3">
    <div
      class={requestMethods.find(
        (r) => r.value === currentRequest?.request.method,
      )?.color}
    >
      {currentRequest?.request.method}
    </div>
    <div>{currentRequest?.request.url}</div>
  </div>
  <div class="grid grid-cols-3 gap-5">
    <div>
      {completedRequest}/{currentRequest?.options.number_of_requests ?? "0"}
    </div>
    <div>
      {rpsValues.length > 0 ? rpsValues[rpsValues.length - 1] : "0"} rps
    </div>
    <div></div>
  </div>
  <pre>{JSON.stringify(stats, null, "\t")}</pre>
</div>
