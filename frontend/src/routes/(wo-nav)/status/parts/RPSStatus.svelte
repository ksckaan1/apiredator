<script lang="ts">
  export let latestRPS = 0;
  export let minRPS = 0;
  export let avgRPS = 0;
  export let maxRPS = 0;
</script>

<div class="tile">
  <h1>Requests per second</h1>
  <div class="grid grid-cols-2 gap-3 mt-3">
    <div class="bl">
      <h2>Latest</h2>
      <span>
        {latestRPS}
      </span>
    </div>
    <div class="bl">
      <h2>Avg</h2>
      <span>
        {avgRPS}
      </span>
    </div>
    <div class="bl">
      <h2>Min</h2>
      <span>
        {minRPS}
      </span>
    </div>
    <div class="bl">
      <h2>Max</h2>
      <span>
        {maxRPS}
      </span>
    </div>
  </div>
</div>
