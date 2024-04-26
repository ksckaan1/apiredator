<script lang="ts">
  import { activeEnvVars, type EnvVar } from "$stores/environment";

  export let value = "";
  export let previewVars = false;
  let isFocused = false;

  let inputElem: HTMLInputElement;

  const onClick = () => {
    isFocused = true;
    setTimeout(() => inputElem.focus(), 10);
  };

  const highlightText = (text: string): string => {
    if (!previewVars) return text;
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
      `<span class="text-red-400">$1</span>`,
    );
    return text;
  };

  let highlightedText = highlightText(value);

  $: {
    if ($activeEnvVars || value) highlightedText = highlightText(value);
  }
</script>

<div class="h-10 cursor-pointer overflow-hidden">
  {#if isFocused}
    <input
      class="w-full px-3 h-full border rounded bg-transparent focus:outline-none border-primary"
      bind:this={inputElem}
      on:blur={() => (isFocused = false)}
      type="text"
      autocorrect="off"
      autocapitalize="none"
      bind:value
    />
  {:else}
    <button
      class="px-3 w-full h-full text-left overflow-hidden text-nowrap relative truncate cursor-text"
      on:click={onClick}
    >
      {@html highlightedText}
    </button>
  {/if}
</div>
