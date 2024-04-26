<script lang="ts">
  import { WindowToggleMaximise } from "$lib/wailsjs/runtime/runtime";

  const onTitleBarDoubleClick = () => {
    WindowToggleMaximise();
  };
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
  style="--wails-draggable:drag"
  class="flex items-center h-10 pl-20 border-b bg-default-bg border-white/20"
  on:dblclick|self={onTitleBarDoubleClick}
>
  <h1 class="cursor-default select-none">Bomber</h1>
</div>
