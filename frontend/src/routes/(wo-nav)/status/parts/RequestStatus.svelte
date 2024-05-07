<script lang="ts">
  export let testType = "count";
  export let sentCount = 0;
  export let numberOfRequests = 0;
  export let numberOfClients = 0;
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
