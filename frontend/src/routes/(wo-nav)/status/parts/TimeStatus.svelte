<script lang="ts">
  export let startedAt = "";
  export let endedAt = "";
</script>

<div class="tile">
  <h1>Time</h1>
  <div class="mt-3 text-lg text-white/40 font-extralight">Started at:</div>
  <span>
    {startedAt != "" ? startedAt : "not started yet"}
  </span>
  <div class="mt-3 text-lg text-white/40 font-extralight">Ended at:</div>
  <span class="mb-5">
    {endedAt != "" ? endedAt : "not ended yet"}
  </span>
</div>
