<script lang="ts">
  import Button from "$components/ui/Button.svelte";
  import TextInput from "$components/ui/TextInput.svelte";
  import { GetAllTags } from "$lib/wailsjs/go/service/AppService";
  import { showToast } from "$stores/toast";
  import { flip } from "svelte/animate";
  import { fly, fade } from "svelte/transition";

  interface Props {
    showModal: boolean;
    addBookmark: (bookmarkTitle: string, bookmarkTags: string[]) => void;
  }

  let { showModal = $bindable(false), addBookmark }: Props = $props();

  let bookmarkTitle = $state("");
  let bookmarkTags: string[] = $state([]);

  let bookmarkTag = $state("");
  let allTags: string[] = $state([]);

  $effect(() => {
    if (showModal) {
      GetAllTags().then((tags) => {
        allTags = tags;
      });
    }
  });

  const onClickOutside = (e: MouseEvent) => {
    if (e.target !== e.currentTarget) return;
    showModal = false;
  };

  const removeTag = (tag: string) => {
    bookmarkTags = bookmarkTags.filter((t) => t !== tag);
  };

  const onAddButtonClicked = () => {
    if (!/.{1,}/gm.test(bookmarkTitle)) {
      showToast({
        message: "Invalid bookmark title",
        type: "error",
      });
      return;
    }

    addBookmark(bookmarkTitle, bookmarkTags);

    bookmarkTitle = "";
    bookmarkTags = [];
    bookmarkTag = "";
    showModal = false;
  };

  const onCancelAddBookmarkButtonClicked = () => {
    bookmarkTitle = "";
    bookmarkTags = [];
    bookmarkTag = "";
    showModal = false;
  };
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->

{#if showModal}
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/50 z-50 flex items-start justify-center"
    onclick={onClickOutside}
  >
    <div
      class="bg-accent-bg border border-white/20 rounded p-5 mt-20 w-full max-w-xl"
      transition:fly={{ duration: 200, y: -200 }}
    >
      <h2 class="text-2xl mb-2">Save Bookmark</h2>
      <TextInput label="Insert Title" bind:value={bookmarkTitle} />
      {#if bookmarkTags.length > 0}
        <div class="flex flex-wrap gap-2 py-5">
          {#each bookmarkTags as tag (tag)}
            <button
              animate:flip={{ duration: 200 }}
              transition:fade={{ duration: 200 }}
              onclick={() => removeTag(tag)}
              class="bg-primary text-on-primary px-2 py-1 rounded"
            >
              {tag}
            </button>
          {/each}
        </div>
      {:else}
        <div class="py-6">No tags added</div>
      {/if}
      <TextInput
        label="Add Tag"
        autoComplete={allTags}
        bind:value={bookmarkTag}
        onkeypress={(e) => {
          if (
            e.key === "Enter" &&
            !bookmarkTags.includes(bookmarkTag) &&
            bookmarkTag.trim() != ""
          ) {
            bookmarkTags = [...bookmarkTags, bookmarkTag];
            bookmarkTag = "";
          }
        }}
      />
      <div class="flex justify-end gap-3 items-center mt-5">
        <Button
          onclick={onCancelAddBookmarkButtonClicked}
          variant="transparent"
        >
          Cancel
        </Button>
        <Button onclick={onAddButtonClicked} icon="ic:round-add">Add</Button>
      </div>
    </div>
  </div>
{/if}
