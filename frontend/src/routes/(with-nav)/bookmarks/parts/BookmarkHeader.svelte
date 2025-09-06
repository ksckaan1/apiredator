<script lang="ts">
	import Button from "$components/ui/Button.svelte";
	import DropdownSelect from "$components/ui/DropdownSelect.svelte";
	import TextInput from "$components/ui/TextInput.svelte";
	import { fade } from "svelte/transition";

	interface Props {
		searchValue: string;
		tagValue: string;
		allTags: string[];
		selectedBookmarks: string[];
		showDeleteBookmarkModal: boolean;
	}

	let {
		searchValue = $bindable(""),
		tagValue = $bindable(""),
		allTags,
		selectedBookmarks = $bindable([]),
		showDeleteBookmarkModal = $bindable(false),
	}: Props = $props();

	const onDeselectAllButtonClicked = () => {
		selectedBookmarks = [];
	};

	const onDeleteBookmarksButtonClicked = () => {
		showDeleteBookmarkModal = true;
	};
</script>

<div class="flex gap-4">
	<div class="flex-1">
		<TextInput bind:value={searchValue} label="Search" autoFocus />
	</div>
	<div class="w-48">
		<DropdownSelect items={allTags} bind:value={tagValue} />
	</div>
</div>

<div
	class="sticky top-0 py-2 h-10 border-b flex justify-between items-center border-white/20 bg-default-bg z-40"
>
	<span>
		{selectedBookmarks.length} bookmark{#if selectedBookmarks.length > 1}s{/if}
		selected
	</span>
	{#if selectedBookmarks.length > 0}
		<div transition:fade={{ duration: 200 }} class="flex items-start gap-4">
			<Button
				onclick={onDeselectAllButtonClicked}
				variant="transparent"
				icon="bx:checkbox"
			>
				Deselect All
			</Button>
			<Button
				onclick={onDeleteBookmarksButtonClicked}
				variant="transparent"
				icon="ph:trash"
			>
				Delete Selected
			</Button>
		</div>
	{/if}
</div>
