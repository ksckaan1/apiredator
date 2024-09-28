<script lang="ts">
  interface Props {
    testType?: string;
    sentCount?: number;
    numberOfRequests?: number;
    numberOfClients?: number;
  }

  let {
    testType = $bindable("count"),
    sentCount = $bindable(0),
    numberOfRequests = $bindable(0),
    numberOfClients = $bindable(0),
  }: Props = $props();
</script>

<div class="tile">
  <h1>Requests</h1>
  <div class="grid gap-3 mt-3">
    <div class="bl">
      <h2>Sent</h2>
      <span>
        {sentCount}
      </span>
    </div>
    {#if testType === "count"}
      <div class="bl">
        <h2>Target</h2>
        <span>
          {numberOfRequests * numberOfClients}
        </span>
      </div>
    {/if}
  </div>
  {#if testType === "count" && numberOfClients > 1}
    <div class="text-sm text-white/40 mt-3">
      {numberOfClients} clients X {numberOfRequests} requests
    </div>
  {/if}
</div>
