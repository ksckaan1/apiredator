<script lang="ts">
  import { convertSeconds } from "$utils/time";
  export let testType = "duration";
  export let passedDuration = "0s";
  export let testDuration = "0s";
  export let numberOfClients = 1;
</script>

<div class="tile">
  <h1>Duration</h1>
  <div class="grid gap-3 mt-3">
    <div class="bl">
      <h2>Passed</h2>
      <span>
        {passedDuration}
      </span>
    </div>
    {#if testType === "duration"}
      <div class="bl">
        <h2>Target</h2>
        <span>
          {testDuration}
        </span>
      </div>
    {/if}
  </div>
  {#if testType === "duration" && numberOfClients > 1}
    <div class="text-sm text-white/40 mt-3">
      {numberOfClients} clients X {testDuration}
    </div>
  {/if}
</div>
