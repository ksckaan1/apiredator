<script lang="ts">
  import type { FormData } from "$types/custom";
  import DropdownSelect from "./DropdownSelect.svelte";
  import EditableInput from "./EditableInput.svelte";
  import FileSelector from "./FileSelector.svelte";
  import Switch from "./Switch.svelte";

  export let rows: FormData[] = [
    {
      key: "file",
      text_value: "",
      file_value: ["abc.jpg", "def.jpg"],
      row_type: "file",
      is_active: true,
    },
    {
      key: "text",
      text_value: "value",
      file_value: [],
      row_type: "text",
      is_active: true,
    },
  ];

  const deleteRow = (i: number) => {
    rows = rows.filter((_, index) => i !== index);
  };

  let newKey = "";
  let newTextValue: string = "";
  let newFileValue: string[] = [];
  let newRowType: "text" | "file" = "text";

  const addRow = () => {
    if (newKey.trim() == "") return;

    rows = [
      ...rows,
      {
        key: newKey,
        text_value: newTextValue,
        file_value: newFileValue,
        row_type: newRowType,
        is_active: true,
      },
    ];

    newKey = "";
    newFileValue = [];
    newTextValue = "";
  };
</script>

<div>
  <div class="grid grid-cols-[8rem,1fr,7rem,1fr,3rem] border-b border-white/20">
    <div class="font-bold my-2 px-3">Using?</div>
    <div class="font-bold my-2 px-6">Key</div>
    <div class="font-bold my-2 px-6">Type</div>
    <div class="font-bold my-2 px-3">Value</div>
  </div>
  {#each rows as row, i}
    <div
      class="rw grid w-full grid-cols-[8rem,1fr,7rem,1fr,3rem] border-b px-3 border-white/20 items-center"
    >
      <Switch bind:value={row.is_active} />
      <EditableInput bind:value={row.key} />
      <div class="px-4">
        <DropdownSelect
          bind:value={row.row_type}
          items={[
            {
              title: "File",
              value: "file",
            },
            {
              title: "Text",
              value: "text",
            },
          ]}
        />
      </div>
      {#if row.row_type === "file"}
        <FileSelector bind:files={row.file_value} />
      {:else}
        <EditableInput bind:value={row.text_value} />
      {/if}
      <button on:click={() => deleteRow(i)}>
        <i class="fa-regular fa-trash-can"></i>
      </button>
    </div>
  {/each}
  <div
    class="grid w-full grid-cols-[8rem,1fr,7rem,1fr,3rem] border-b px-3 border-white/20 items-center"
  >
    <span></span>
    <input
      bind:value={newKey}
      autocorrect="off"
      autocapitalize="none"
      class="h-10 bg-transparent placeholder:text-white/20 border border-transparent rounded px-3 focus:outline-none focus:border-primary"
      type="text"
      placeholder="key"
    />
    <div class="px-4">
      <DropdownSelect
        bind:value={newRowType}
        items={[
          {
            title: "File",
            value: "file",
          },
          {
            title: "Text",
            value: "text",
          },
        ]}
      />
    </div>
    {#if newRowType === "text"}
      <input
        bind:value={newTextValue}
        autocorrect="off"
        autocapitalize="none"
        class="h-10 bg-transparent placeholder:text-white/20 border border-transparent rounded px-3 focus:outline-none focus:border-primary"
        type="text"
        placeholder="value"
      />
    {:else}
      <FileSelector bind:files={newFileValue} />
    {/if}
    <button on:click={addRow}>
      <i class="fa-solid fa-plus"></i>
    </button>
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
