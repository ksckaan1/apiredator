<script lang="ts">
  import { slide } from "svelte/transition";
  interface Props {
    startedAt?: string;
    endedAt?: string;
  }
  let { startedAt = $bindable(""), endedAt = $bindable("") }: Props = $props();
</script>

<div class="tile">
  <h1>Time</h1>
  <div class="grid gap-3 mt-3">
    <div class="bl">
      <h2>Started at</h2>
      <span>
        {startedAt != "" ? startedAt : "not started yet"}
      </span>
    </div>
    {#if endedAt !== ""}
      <div transition:slide class="bl">
        <h2>Ended at</h2>
        <span>
          {endedAt}
        </span>
      </div>
    {/if}
  </div>
</div>
