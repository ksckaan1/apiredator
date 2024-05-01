<script lang="ts">
  import DurationStatus from "./DurationStatus.svelte";
  import RequestStatus from "./RequestStatus.svelte";
  import TimeStatus from "./TimeStatus.svelte";
  import StatusCodes from "./StatusCodes.svelte";
  import RpsStatus from "./RPSStatus.svelte";

  export let testType = "";
  export let sentCount = 0;
  export let numberOfClients = 0;
  export let numberOfRequests = 0;
  export let passedDuration = 0;
  export let targetDuration = 0;
  export let latestRPS = 0;
  export let minRPS = 0;
  export let avgRPS = 0;
  export let maxRPS = 0;
  export let startedAt = "";
  export let endedAt = "";
  export let statusCodes: { [key: number]: number } = {};
</script>

<div class="grid grid-cols-3 gap-5 mt-5">
  <RequestStatus {testType} {sentCount} {numberOfClients} {numberOfRequests} />
  <DurationStatus
    {testType}
    {passedDuration}
    {numberOfClients}
    {targetDuration}
  />
  <RpsStatus {latestRPS} {minRPS} {avgRPS} {maxRPS} />
  <TimeStatus {startedAt} {endedAt} />
  <StatusCodes {statusCodes} />
</div>
