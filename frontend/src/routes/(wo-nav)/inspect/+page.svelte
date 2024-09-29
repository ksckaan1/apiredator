<script lang="ts">
  import Button from "$components/ui/Button.svelte";
  import { models } from "$lib/wailsjs/go/models";
  import { fade, fly } from "svelte/transition";
  import UrlInfo from "../status/parts/URLInfo.svelte";
  import PinButton from "../status/parts/PinButton.svelte";
  import StatusTiles from "../status/parts/StatusTiles.svelte";
  import { Chart } from "chart.js/auto";
  import { goto } from "$app/navigation";
  import { DeleteBookmarks, GetBookmarkByID, SetCurrentRequest, StartCurrentRequest } from "$lib/wailsjs/go/service/AppService";
  import { prettyTime } from "$utils/time";
  import Icon from "@iconify/svelte";
  import { showToast } from "$stores/toast";
  import DeleteBookmarkModal from "./parts/DeleteBookmarkModal.svelte";

  const usp = new URLSearchParams(window.location.search);
  const bookmarkID: string = usp.get("id") ?? "";

  let showDeleteModal = $state(false);

  let bookmarkData: models.Bookmark | null = $state(null);

  let chartElem: HTMLCanvasElement;
  let ch: Chart;

  let pageScroll = $state(0);
  let isPinned = $state(true);

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

  const resizeEvent = () => {
    var gradient = chartElem
      .getContext("2d")!
      .createLinearGradient(0, 0, 0, chartElem.height);

    gradient.addColorStop(0, "rgba(225, 175, 75, .6)");
    gradient.addColorStop(1, "rgb(13, 15, 24)");

    ch.data.datasets[0].backgroundColor = gradient;
    ch.resize();
  };

  $effect(() => {
    (async () => {
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

      GetBookmarkByID(bookmarkID).then((result) => {
        bookmarkData = result
        setChartData(result);
      });
    })();

    return () => {
      window.removeEventListener("resize", resizeEvent);
    };
  });

  const setChartData = (b: models.Bookmark) => {
    ch.data.datasets[0].data = b.stat.rps.list;
    ch.data.labels = ch.data.datasets[0].data.map((r) => "");
    ch.update();
  }

  const onBackButtonClicked = ()=>{
    goto('/bookmarks')
  }

  const onRetryButtonClicked = ()=>{
    SetCurrentRequest(new models.Data({
        request: bookmarkData?.request,
        options: bookmarkData?.options
    })).then(()=>{
      StartCurrentRequest()
      .then(()=>goto('/status'))
      .catch((e: Error)=>{
        showToast({
          message: e.message,
          type: "error"
        })
      })
    }).catch((e: Error)=>{
      showToast({
        message: e.message,
        type: "error"
      })
    })
  }

  const onNewRequestButtonClicked = ()=>{
    SetCurrentRequest(new models.Data({
        request: bookmarkData?.request,
        options: bookmarkData?.options
    })).then(()=>{
      goto('/new-request')
    }).catch((e: Error)=>{
      showToast({
        message: e.message,
        type: "error"
      })
    })
  }

  const onDeleteBookmarkButtonClicked = ()=>{
    showDeleteModal = true
  }

  const deleteBookmark = ()=>{
    DeleteBookmarks([bookmarkID]).then(()=>{
      showDeleteModal = false
      goto('/bookmarks')
      showToast({
        message: "Bookmark deleted",
        type: "success"
      })
    }).catch((e: Error)=>{
      showToast({
        message: e.message,
        type: "error"
      })
    })
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
          icon="ic:outline-arrow-back"
          onclick={onBackButtonClicked}
        >
          Back to Bookmarks
        </Button>
        <div class="flex items-center gap-3">
          <Button
            icon="ic:round-refresh"
            variant="success"
            onclick={onRetryButtonClicked}
            >
              Retry
            </Button>
          <Button 
          onclick={onNewRequestButtonClicked} 
          icon="fluent:globe-arrow-up-20-regular"
          >Use as New Request</Button>
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

{#if bookmarkData != null}
  <div
    class="w-full relative z-50 flex-1 overflow-y-auto flex flex-col hide-scrollbar"
    onscroll={(e: any) => {
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
            <div class="flex justify-between">
              <div>
                <div class="text-3xl text-white/70 mb-2 font-light">
                  {bookmarkData?.title}
                </div>
                <div class="mt-2 text-white/30 font-light">
                  {new Date(bookmarkData?.create_at).toString()}
                </div>
                <div class="mb-3 flex flex-wrap gap-3 items-center">
                  <Icon icon="mdi:tags" width="20" class="text-white/50" />
                  {#if bookmarkData?.tags && bookmarkData?.tags.length > 0}
                    {#each bookmarkData?.tags as tag}
                      <div class="text-white/50">
                        #{tag}
                      </div>
                    {/each}
                  {:else}
                    <div class="text-white/50">No Tag</div>
                  {/if}
                </div>
              </div>
              <div class=" flex gap-x-2 items-start">
                <Button
                  variant="outlined"
                    icon="ic:outline-edit"
                >
                  Edit
                </Button>
                <Button
                  variant="outlined-danger"
                  icon="ic:outline-delete"
                  onclick={onDeleteBookmarkButtonClicked}
                >
                  Delete
                </Button>
              </div>
            </div>
            <div class="flex gap-2 items-center">
              <UrlInfo requestMethod={bookmarkData?.request.method} requestURL={bookmarkData?.request.url} />
              <PinButton bind:isPinned />
            </div>
          </div>
          <StatusTiles
            testType={bookmarkData?.options.test_type}
            sentCount={bookmarkData?.stat.sent_count}
            numberOfRequests={bookmarkData?.options.number_of_requests}
            numberOfClients={bookmarkData?.options.number_of_clients}
            passedDuration={bookmarkData?.stat.passed_duration}
            testDuration={bookmarkData?.options.test_duration}
            latestRPS={bookmarkData?.stat.rps.latest}
            minRPS={bookmarkData?.stat.rps.min}
            avgRPS={Number(bookmarkData?.stat.rps.avg.toFixed(2))}
            maxRPS={bookmarkData?.stat.rps.max}
            startedAt={prettyTime(bookmarkData?.create_at)}
            endedAt={prettyTime(bookmarkData?.stat.ended_at)}
            statusCodes={bookmarkData?.stat.status_codes}
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
  {/if}
</div>

<DeleteBookmarkModal
  bind:showModal={showDeleteModal}
  selectedBookmark={bookmarkID}
  {deleteBookmark}
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

