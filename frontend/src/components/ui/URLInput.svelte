<script lang="ts">
  import { activeEnvVars, type EnvVar } from "$stores/environment";

  export let value: string = "";
  export let placeholder: string = "";
  export let label: string = "";

  let inputElem: HTMLInputElement;

  const highlightText = (text: string): string => {
    $activeEnvVars.forEach((item: EnvVar) => {
      text = text.replaceAll(
        new RegExp(`{{${item.key}}}`, "gi"),
        `<div class="inline text-primary" data-tooltip-message="<b>env:</b> ${item.key}" data-tooltip-position="bottom">
          ${item.value}
        </div>`,
      );
    });

    text = text.replaceAll(
      /{{([a-z0-9_\-]+)}}/gi,
      `<span class="text-red-400">{{$1}}</span>`,
    );
    return text;
  };

  let highlightedText = highlightText(value);

  $: if ($activeEnvVars || value) highlightedText = highlightText(value);

  const focusInput = () => {
    inputElem.focus();
  };
</script>

<div class="wrapper" tabindex="-1">
  <slot name="head" />
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  {#if label}
    <div
      on:click={focusInput}
      class="flex items-center justify-center h-full pl-2 text-white/60"
    >
      {label}
    </div>
  {/if}
  <input
    bind:this={inputElem}
    class="flex-1 w-full px-3 py-2 bg-default-bg focus:outline-none"
    autocorrect="off"
    autocapitalize="none"
    type="url"
    {placeholder}
    bind:value
  />
  <slot name="tail" />
</div>

<style lang="postcss">
  .wrapper {
    @apply flex border rounded border-white/20 w-full h-10;
  }

  .wrapper:has(input:focus, &:focus) {
    @apply border-primary;
  }
</style>
