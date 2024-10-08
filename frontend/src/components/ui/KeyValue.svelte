<script lang="ts">
  import type { KeyValueData } from "$types/custom";
  import TextInput from "./TextInput.svelte";
  import Switch from "./Switch.svelte";

  interface Props {
    rows?: KeyValueData[];
  }

  let { rows = $bindable([]) }: Props = $props();

  const deleteRow = (i: number) => {
    rows = rows.filter((_, index) => i !== index);
  };

  let keyInput: HTMLInputElement;
  let valueInput: HTMLInputElement;

  let newKey = $state("");
  let newValue = $state("");

  const onKeyInput = (e: any) => {
    if (e.key !== "Enter") return;

    const keyTrimmed = newKey.trim();
    if (keyTrimmed === "") return;

    const valueTrimmed = newValue.trim();
    if (valueTrimmed === "") {
      valueInput.focus();
      return;
    }

    addNewRow(keyTrimmed, valueTrimmed);
  };

  const onValueInput = (e: any) => {
    if (e.key !== "Enter") return;

    const keyTrimmed = newKey.trim();
    if (keyTrimmed === "") {
      keyInput.focus();
      return;
    }

    const valueTrimmed = newValue.trim();

    addNewRow(keyTrimmed, valueTrimmed);
  };

  const addNewRow = (key: string, value: string) => {
    rows = [
      ...rows,
      {
        key,
        value,
        is_active: true,
      },
    ];

    newKey = "";
    newValue = "";

    keyInput.focus();
  };
</script>

<div>
  <div class="grid grid-cols-[8rem,1fr,1fr,3rem] border-b border-white/20">
    <div class="font-bold my-2 px-3">Using?</div>
    <div class="font-bold my-2 px-6">Key</div>
    <div class="font-bold my-2 px-3">Value</div>
  </div>
  {#each rows as row, i}
    <div
      class="rw grid w-full grid-cols-[8rem,1fr,1fr,3rem] border-b px-3 border-white/20 items-center"
    >
      <Switch bind:value={row.is_active} />
      <TextInput bind:value={row.key} variant="secondary" />
      <TextInput bind:value={row.value} variant="secondary" />
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button onclick={() => deleteRow(i)}>
        <i class="fa-regular fa-trash-can"></i>
      </button>
    </div>
  {/each}
  <div
    class="grid w-full grid-cols-[8rem,1fr,1fr,3rem] border-b px-3 border-white/20 items-center"
  >
    <span></span>
    <input
      bind:this={keyInput}
      bind:value={newKey}
      autocorrect="off"
      autocapitalize="none"
      class="h-10 bg-transparent placeholder:text-white/20 border border-transparent rounded px-3 focus:outline-none focus:border-primary"
      type="text"
      placeholder="key"
      onkeydown={onKeyInput}
    />
    <input
      bind:this={valueInput}
      bind:value={newValue}
      autocorrect="off"
      autocapitalize="none"
      class="h-10 bg-transparent placeholder:text-white/20 border border-transparent rounded px-3 focus:outline-none focus:border-primary"
      type="text"
      placeholder="value"
      onkeydown={onValueInput}
    />
    <span></span>
  </div>
</div>

<style lang="postcss">
  .rw:nth-child(odd) {
    @apply bg-default-bg;
  }

  .rw:nth-child(odd) :global(.elipsis) {
    @apply bg-gradient-to-r from-default-bg/0 to-default-bg pointer-events-none;
  }
  .rw:nth-child(even) {
    @apply bg-accent-bg;
  }
  .rw:nth-child(even) :global(.elipsis) {
    @apply bg-gradient-to-r from-accent-bg/0 to-accent-bg pointer-events-none;
  }
</style>
