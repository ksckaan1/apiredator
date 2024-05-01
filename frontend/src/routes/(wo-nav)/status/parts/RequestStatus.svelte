<script lang="ts">
  export let testType = "count";
  export let sentCount = 0;
  export let numberOfRequests = 0;
  export let numberOfClients = 0;
</script>

<div class="tile">
  <h1>Requests</h1>
  <div class="flex gap-x-3 mt-3">
    <div class="flex-1">
      <h2 class="text-lg font-extralight text-white/30">Sent</h2>
      <span class="text-xl">
        {sentCount}
      </span>
    </div>
    {#if testType === "count"}
      <div class="flex-1">
        <h2 class="text-lg font-extralight text-white/30 px-3">Target</h2>
        <span class="text-xl px-3">
          {numberOfRequests * numberOfClients}
        </span>
      </div>
    {/if}
  </div>
  {#if testType === "count" && numberOfClients > 1}
    <div class="text-sm text-white/40">
      {numberOfClients} clients X {numberOfRequests} requests
    </div>
  {/if}
</div>
