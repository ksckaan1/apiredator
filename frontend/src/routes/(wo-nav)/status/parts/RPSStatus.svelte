<script lang="ts">
  export let latestRPS = 0;
  export let minRPS = 0;
  export let avgRPS = 0;
  export let maxRPS = 0;
</script>

<div class="tile">
  <h1>Requests per second</h1>
  <div class="mt-3 text-white/40 text-lg font-extralight">Latest</div>
  <span>
    {latestRPS}
  </span>
  <div class="flex gap-x-3 mt-3">
    <div class="flex-1">
      <div class="text-white/40 text-lg font-extralight">Min</div>
      <span>
        {minRPS}
      </span>
    </div>
    <div class="flex-1">
      <div class="text-white/40 text-lg font-extralight">Avg</div>
      <span>
        {avgRPS}
      </span>
    </div>
    <div class="flex-1">
      <div class="text-white/40 text-lg font-extralight">Max</div>
      <span>
        {maxRPS}
      </span>
    </div>
  </div>
</div>
