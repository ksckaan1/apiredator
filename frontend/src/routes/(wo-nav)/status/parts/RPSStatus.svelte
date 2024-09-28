<script lang="ts">
  interface Props {
    latestRPS?: number;
    minRPS?: number;
    avgRPS?: number;
    maxRPS?: number;
  }

  let {
    latestRPS = $bindable(0),
    minRPS = $bindable(0),
    avgRPS = $bindable(0),
    maxRPS = $bindable(0),
  }: Props = $props();
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
