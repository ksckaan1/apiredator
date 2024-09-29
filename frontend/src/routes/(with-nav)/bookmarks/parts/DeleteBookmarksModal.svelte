<script lang="ts">
  import Button from "$components/ui/Button.svelte";
  import { fly } from "svelte/transition";

  interface Props {
    showModal: boolean;
    deleteBookmarks: () => void;
    selectedBookmarks: string[];
  }

  let {
    showModal = $bindable(false),
    selectedBookmarks = [],
    deleteBookmarks,
  }: Props = $props();

  const onCancelDeleteBookmarkButtonClicked = () => {
    showModal = false;
  }
</script>

{#if showModal}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/50 z-50 flex items-start justify-center"
    onclick={(e) => {
      if (e.target !== e.currentTarget) return;
      showModal = false;
    }}
  >
    <div
      transition:fly={{ duration: 200, y: -200 }}
      class="bg-accent-bg border border-white/20 rounded p-5 mt-20 w-full max-w-xl"
    >
      <h2 class="text-2xl mb-2">Delete Bookmark(s)</h2>
      <p>
        Are you sure you want to delete the {selectedBookmarks.length} selected bookmark(s)?
      </p>
      <div class="flex justify-end gap-3 items-center mt-5">
        <Button
          onclick={onCancelDeleteBookmarkButtonClicked}
          variant="transparent"
        >
          Cancel
        </Button>
        <Button
          onclick={deleteBookmarks}
          variant="danger"
          icon="ph:trash"
        >
          Delete
        </Button>
      </div>
    </div>
  </div>
{/if}
