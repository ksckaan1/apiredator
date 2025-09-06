<script lang="ts">
	import Button from "$components/ui/Button.svelte";
	import DropdownSelect from "$components/ui/DropdownSelect.svelte";
	import TextInput from "$components/ui/TextInput.svelte";
	import { fade } from "svelte/transition";

	interface Props {
		searchValue: string;
		tagValue: string;
		allTags: string[];
		allIDs: string[];
		selectedBookmarks: string[];
		showDeleteBookmarkModal: boolean;
	}

	let {
		searchValue = $bindable(""),
		tagValue = $bindable(""),
		allTags,
		allIDs = [],
		selectedBookmarks = $bindable([]),
		showDeleteBookmarkModal = $bindable(false),
	}: Props = $props();

	const onSelectAllButtonClicked = () => {
		selectedBookmarks = allIDs;
	};

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
	<div class="flex">
		<Button
			onclick={onSelectAllButtonClicked}
			variant="transparent"
			icon="bxs:checkbox"
			disabled={selectedBookmarks.length === allIDs.length}
		>
			Select All
		</Button>
		<Button
			onclick={onDeselectAllButtonClicked}
			variant="transparent"
			icon="bx:checkbox"
			disabled={selectedBookmarks.length === 0}
		>
			Deselect All
		</Button>
		<Button
			onclick={onDeleteBookmarksButtonClicked}
			variant="transparent"
			icon="ph:trash"
			disabled={selectedBookmarks.length === 0}
		>
			Delete Selected
		</Button>
	</div>
</div>
