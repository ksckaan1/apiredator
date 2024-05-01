<script lang="ts">
  import { convertSeconds } from "$utils/time";
  export let testType = "duration";
  export let passedDuration = 0;
  export let targetDuration = 0;
  export let numberOfClients = 1;
</script>

<div class="tile">
  <h1>Duration</h1>
  <div class="flex gap-x-3 mt-3">
    <div class="flex-1">
      <h2 class="text-lg font-extralight text-white/30">Passed</h2>
      <span class="text-xl">
        {convertSeconds(passedDuration)}
      </span>
    </div>
    {#if testType === "duration"}
      <div class="flex-1">
        <h2 class="text-lg font-extralight text-white/30 px-3">Target</h2>
        <span class="text-xl px-3">
          {convertSeconds(targetDuration)}
        </span>
      </div>
    {/if}
  </div>
  {#if testType === "duration" && numberOfClients > 1}
    <div class="text-sm text-white/40">
      {numberOfClients} clients X {convertSeconds(targetDuration)}
    </div>
  {/if}
</div>
