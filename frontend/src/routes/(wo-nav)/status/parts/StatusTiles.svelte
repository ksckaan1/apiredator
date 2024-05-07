<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import DurationStatus from "./DurationStatus.svelte";
  import RequestStatus from "./RequestStatus.svelte";
  import TimeStatus from "./TimeStatus.svelte";
  import StatusCodes from "./StatusCodes.svelte";
  import RpsStatus from "./RPSStatus.svelte";

  export let testType = "";
  export let sentCount = 0;
  export let numberOfClients = 0;
  export let numberOfRequests = 0;
  export let passedDuration = "0s";
  export let testDuration = "0s";
  export let latestRPS = 0;
  export let minRPS = 0;
  export let avgRPS = 0;
  export let maxRPS = 0;
  export let startedAt = "";
  export let endedAt = "";
  export let statusCodes: { [key: number]: number } = {};
</script>

<div class="statuses mt-5">
  <div in:fly={{ delay: 600, duration: 200, y: 50 }}>
    <RequestStatus
      {testType}
      {sentCount}
      {numberOfClients}
      {numberOfRequests}
    />
  </div>
  <div in:fly={{ delay: 750, duration: 200, y: 50 }}>
    <DurationStatus
      {testType}
      {passedDuration}
      {numberOfClients}
      {testDuration}
    />
  </div>
  <div in:fly={{ delay: 900, duration: 200, y: 50 }}>
    <RpsStatus {latestRPS} {minRPS} {avgRPS} {maxRPS} />
  </div>
  <div in:fly={{ delay: 1050, duration: 200, y: 50 }}>
    <TimeStatus {startedAt} {endedAt} />
  </div>
  <div in:fly={{ delay: 1200, duration: 200, y: 50 }}>
    <StatusCodes {statusCodes} />
  </div>
</div>

<style class="postcss">
  .statuses {
    @apply grid gap-5 w-full;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  }
</style>
