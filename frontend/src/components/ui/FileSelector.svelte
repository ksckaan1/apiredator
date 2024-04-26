<script lang="ts">
  import { SelectFiles } from "$lib/wailsjs/go/service/AppService";

  export let files: string[] = [];
  export let is_multiple: boolean = true;

  const onSelectFiles = async () => {
    const selectedFiles = await SelectFiles(is_multiple);
    if (is_multiple) {
      selectedFiles.forEach((file) => {
        if (!files.includes(file)) {
          files = [...files, file];
        }
      });
      return;
    }
    files = selectedFiles;
  };

  const removeFile = (removedFile: string) => {
    files = files.filter((l) => l != removedFile);
  };
</script>

<div tabindex="-1" class="flex relative h-full group">
  <button class="text-left flex-1 px-3">
    {#if files.length > 1}
      <span>{files.length} files selected</span>
    {:else if files.length > 0}
      <span>1 file selected</span>
    {:else}
      <span>
        Select file{#if is_multiple}(s){/if}</span
      >
    {/if}
  </button>
  <div
    class="hidden group-focus:flex z-50 flex-col absolute top-0 px-3 pt-2 pb-3 left-0 bg-accent-bg rounded border border-white/20 w-full"
  >
    {#if files.length > 1}
      <span>{files.length} files selected</span>
    {:else if files.length > 0}
      <span>1 file selected</span>
    {:else}
      <span>
        Select file{#if is_multiple}(s){/if}</span
      >
    {/if}
    <button
      class="bg-white/5 rounded border border-white/20 px-3 py-2 my-3 hover:opacity-65"
      on:click={onSelectFiles}
    >
      Click to Add files
    </button>
    <div class="max-h-48 overflow-y-auto flex flex-col gap-y-3">
      {#each files as file}
        <div
          class="flex shrink-0 rounded items-baseline bg-white/5 px-3 py-2 border border-white/20"
        >
          <span class="flex-1">
            {file}
          </span>
          <i
            class="fa-regular fa-trash-can cursor-pointer"
            on:click={() => removeFile(file)}
          ></i>
        </div>
      {/each}
    </div>
  </div>
</div>
