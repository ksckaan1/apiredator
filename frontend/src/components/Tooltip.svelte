<script lang="ts">
  import { fly } from "svelte/transition";
  let leftPos = -1;
  let topPos = -1;
  let ttPos = "bottom";
  let yFly = 0;
  let xFly = 0;
  let msg = "";

  const onMouseMove = (e: any) => {
    const target = document.elementFromPoint(e.clientX, e.clientY);

    if (!target?.getAttribute("data-tooltip-message")) {
      leftPos = -1;
      topPos = -1;
      return;
    }

    ttPos = target.getAttribute("data-tooltip-position") || "bottom";
    if (ttPos === "bottom") {
      leftPos = e.clientX;
      topPos = e.clientY + 20;
      yFly = 20;
      xFly = 0;
    } else if (ttPos === "top") {
      leftPos = e.clientX;
      topPos = e.clientY - 20;
      yFly = -20;
      xFly = 0;
    } else if (ttPos === "right") {
      leftPos = e.clientX + 20;
      topPos = e.clientY;
      yFly = 0;
      xFly = 20;
    } else if (ttPos === "left") {
      leftPos = e.clientX - 20;
      topPos = e.clientY;
      yFly = 0;
      xFly = -20;
    }
    msg = target.getAttribute("data-tooltip-message") || "";
  };

  document.addEventListener("mousemove", onMouseMove, { passive: true });
</script>

{#if leftPos >= 0 && topPos >= 0}
  <div
    transition:fly={{ duration: 200, y: yFly, x: xFly }}
    class:ttBottom={ttPos === "bottom"}
    class:ttTop={ttPos === "top"}
    class:ttRight={ttPos === "right"}
    class:ttLeft={ttPos === "left"}
    style="--left: {leftPos}px;--top: {topPos}px;"
    class="tooltip"
  >
    <div
      class="px-3 py-1 bg-amber-500 text-nowrap text-black/90 rounded border border-white/20 pointer-events-none"
    >
      {@html msg}
    </div>
  </div>
{/if}

<style lang="postcss">
  .tooltip {
    @apply w-0 h-0 fixed flex;
    left: var(--left);
    top: var(--top);
  }

  .ttBottom {
    @apply items-start justify-center;
  }

  .ttTop {
    @apply items-end justify-center;
  }

  .ttRight {
    @apply items-center justify-start;
  }

  .ttLeft {
    @apply items-center justify-end;
  }
</style>
