<script lang="ts">
  import { WindowToggleMaximise } from "$lib/wailsjs/runtime/runtime";

  const onTitleBarDoubleClick = (e: MouseEvent) => {
    if (e.target !== e.currentTarget) return;
    WindowToggleMaximise();
  };
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
  style="--wails-draggable:drag"
  class="flex items-center h-10 pl-20 border-b bg-default-bg border-white/20"
  ondblclick={onTitleBarDoubleClick}
>
  <h1 class="cursor-default select-none">
    <span class="text-primary">API</span>redator
  </h1>
</div>
