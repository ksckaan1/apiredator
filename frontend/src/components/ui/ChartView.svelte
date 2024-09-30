<script lang="ts">
  import Chart from "chart.js/auto";
  import { onDestroy, onMount } from "svelte";

  interface Props {
    data: number[];
  }

  let { data = [] }: Props = $props();

  let chartElem: HTMLCanvasElement;
  let ch: Chart;
  let showHud = $state(false);

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
            grid: {
              color: "rgba(255, 255, 255, 0.02)",
            },
          },
          y: {
            display: false,
            grid: {
              color: "rgba(255, 255, 255, 0.02)",
            },
          },
        },
        plugins: {
          legend: {
            display: false,
          },
          tooltip: {
            enabled: false,
            intersect: false,
            mode: "nearest",
            borderWidth: 2,
            borderColor: "rgb(225, 175, 75)",
          },
        },
      },
      data: {
        labels: data.map(() => ""),
        datasets: [
          {
            label: "Request per second",
            data: [...data],
            fill: true,
            borderColor: "rgb(225, 175, 75)",
            borderWidth: 0,
            pointRadius: 0,
            pointBorderWidth: 0,
            tension: 0.5,
          },
        ],
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
  });

  onDestroy(() => {
    window.removeEventListener("resize", resizeEvent);
  });

  $effect(() => {
    if (ch === undefined || ch.data === undefined) return;
    ch.data.datasets[0].data = data.map((v: number) => v);
    ch.data.labels = ch.data.datasets[0].data.map((_, index: number) =>
      (index + 1).toString(),
    );
    ch.update();
  });

  $effect(() => {
    var gradient = chartElem
      .getContext("2d")!
      .createLinearGradient(0, 0, 0, chartElem.height);

    gradient.addColorStop(0, "rgba(225, 175, 75, .6)");
    gradient.addColorStop(1, "rgb(13, 15, 24)");

    if (showHud) {
      ch.data.datasets[0].borderWidth = 2;
      ch.data.datasets[0].pointRadius = 2;
      ch.data.datasets[0].pointBorderWidth = 2;
      ch.options.scales!.y!.display = true;
      ch.options.scales!.x!.display = true;
      ch.options.plugins!.tooltip!.enabled = true;

      ch.data.datasets[0].backgroundColor = gradient;

      ch.update();
      return;
    }

    ch.data.datasets[0].borderWidth = 0;
    ch.data.datasets[0].pointRadius = 0;
    ch.data.datasets[0].pointBorderWidth = 0;
    ch.options.scales!.y!.display = false;
    ch.options.scales!.x!.display = false;
    ch.options.plugins!.tooltip!.enabled = false;

    ch.data.datasets[0].backgroundColor = gradient;

    ch.update();
  });
</script>

<canvas
  onmouseenter={() => (showHud = true)}
  onmouseleave={() => (showHud = false)}
  class="w-screen bg-default-bg"
  bind:this={chartElem}
></canvas>
