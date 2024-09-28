<script lang="ts">
  import { fly } from "svelte/transition";
  import DurationStatus from "./DurationStatus.svelte";
  import RequestStatus from "./RequestStatus.svelte";
  import TimeStatus from "./TimeStatus.svelte";
  import StatusCodes from "./StatusCodes.svelte";
  import RpsStatus from "./RPSStatus.svelte";

  interface Props {
    testType?: string;
    sentCount?: number;
    numberOfClients?: number;
    numberOfRequests?: number;
    passedDuration?: string;
    testDuration?: string;
    latestRPS?: number;
    minRPS?: number;
    avgRPS?: number;
    maxRPS?: number;
    startedAt?: string;
    endedAt?: string;
    statusCodes?: { [key: number]: number };
  }

  let {
    testType = $bindable("count"),
    sentCount = $bindable(0),
    numberOfClients = $bindable(0),
    numberOfRequests = $bindable(0),
    passedDuration = $bindable("0s"),
    testDuration = $bindable("0s"),
    latestRPS = $bindable(0),
    minRPS = $bindable(0),
    avgRPS = $bindable(0),
    maxRPS = $bindable(0),
    startedAt = $bindable(""),
    endedAt = $bindable(""),
    statusCodes = $bindable({}),
  }: Props = $props();
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
