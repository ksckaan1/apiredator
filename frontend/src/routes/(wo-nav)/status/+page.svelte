<script lang="ts">
  import Chart from "chart.js/auto";
  import {
    GetStats,
    StopWork,
    WaitWork,
  } from "$lib/wailsjs/go/service/AppService";
  import { onMount } from "svelte";

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let stats: any;
  let isFinished = false;
  let completed = 0;
  let total = 0;
  let duration = 0;
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
      completed = stats.completed;
      let rps = stats.completed - last;
      rpsValues = [...rpsValues, rps];
      ch.data.datasets[0].data = rpsValues;
      ch.data.labels?.push("");
      ch.update();
      console.log({
        completed: stats.completed,
        rps: stats.completed - last,
      });
      duration++;
      last = stats.completed;
    }, 1000);
    await WaitWork();
    isFinished = true;
    clearInterval(interval);
    stats = await GetStats();
  };

  getStats();

  onMount(() => {
    ch = new Chart(chartElem, {
      type: "line",
      options: {
        aspectRatio: 16 / 4,
        // animation: {
        //   duration: 0,
        // },
      },
      data: {
        labels: [],
        datasets: datasets,
      },
    });
  });

  const onStopBtnClicked = () => {
    StopWork();
  };
</script>

<div class="bg-accent-bg" style="--wails-draggable:drag">
  <div class="w-full p-5 pt-10 mx-auto max-w-7xl flex-1 flex flex-col">
    <div class="flex items-center">
      <button>&larr; back to request</button>
      <div class="mx-auto text-center">
        {duration}s
      </div>
      {#if !isFinished}
        <button
          on:click={onStopBtnClicked}
          class="bg-red-900 rounded px-5 py-2"
        >
          Stop
        </button>
      {/if}
    </div>
    <div class="bg-default-bg w-full rounded-lg mt-5 p-5 chartWrapper">
      <canvas class="w-full" bind:this={chartElem}></canvas>
    </div>
  </div>
</div>

<div class="w-full p-5 pt-10 mx-auto max-w-7xl flex-1 flex flex-col">
  <div class="grid grid-cols-3 gap-5">
    <div>{completed}/{total}</div>
    <div>
      {rpsValues.length > 0 ? rpsValues[rpsValues.length - 1] : "0"} rps
    </div>
    <div></div>
  </div>
  <pre>{JSON.stringify(stats, null, "\t")}</pre>
</div>
