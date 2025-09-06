<script lang="ts">
	import DropdownSelect from "$components/ui/DropdownSelect.svelte";
	import TextInput from "$components/ui/TextInput.svelte";
	import type { models } from "$lib/wailsjs/go/models";
	import { fade, slide } from "svelte/transition";
	import {
		DeleteBookmarks,
		GetAllBookmarks,
		GetAllTags,
	} from "$lib/wailsjs/go/service/AppService";
	import { showToast } from "$stores/toast";
	import DeleteBookmarksModal from "./parts/DeleteBookmarksModal.svelte";
	import Button from "$components/ui/Button.svelte";
	import { goto } from "$app/navigation";
	import BookmarkCard from "./parts/BookmarkCard.svelte";

	let bookmarks: models.Bookmark[] = $state([]);
	let searchValue = $state("");
	let tagValue = $state("");
	let allTags: any[] = $state([]);

	let bouncedSearchValue = $state("");
	let selectedBookmarks: string[] = $state([]);

	let showDeleteBookmarkModal = $state(false);

	let timer: number;

	$effect(() => {
		if (searchValue != undefined) {
			clearTimeout(timer);
			timer = setTimeout(() => {
				bouncedSearchValue = searchValue;
			}, 750);
		}
	});

	const makeSearch = () => {
		GetAllBookmarks(bouncedSearchValue, tagValue, -1, 0).then((result) => {
			bookmarks = result.bookmarks;
			console.log(result);
		});
	};

	makeSearch();

	$effect(() => {
		if (bouncedSearchValue != undefined && tagValue != undefined) {
			makeSearch();
		}
	});

	GetAllTags().then((result) => {
		allTags = [
			{ title: "All Tags", value: "" },
			...result.map((t) => ({ value: t, title: t })),
		];
	});

	const onDeselectAllButtonClicked = () => {
		selectedBookmarks = [];
	};

	const onDeleteBookmarksButtonClicked = () => {
		showDeleteBookmarkModal = true;
	};

	const deleteBookmarks = () => {
		DeleteBookmarks(selectedBookmarks)
			.then(() => {
				showToast({
					type: "success",
					message: `${selectedBookmarks.length} bookmark(s) deleted`,
				});
				selectedBookmarks = [];
				showDeleteBookmarkModal = false;
				makeSearch();
			})
			.catch((e: Error) => {
				showToast({
					type: "error",
					message: e.message,
				});
				selectedBookmarks = [];
				showDeleteBookmarkModal = false;
				makeSearch();
			});
	};

	const onClickBookmark = (id: string) => {
		goto(`/inspect?id=${id}`);
	};

	const onBookmarkSelected = (id: string, selected: boolean) => {
		if (selected && !selectedBookmarks.includes(id)) {
			selectedBookmarks.push(id);
			return;
		}

		selectedBookmarks = selectedBookmarks.filter(
			(bookmarkId) => bookmarkId !== id,
		);
	};
</script>

<div
	in:fade={{ duration: 200, delay: 200 }}
	out:fade={{ duration: 200 }}
	class="h-[calc(100vh-40px)] flex flex-col"
>
	<div class="flex-1 hide-scrollbar overflow-y-auto flex flex-col px-5 pt-5">
		<div class="mx-auto w-full max-w-7xl flex flex-col gap-4">
			<h1 class="text-3xl">Bookmarks</h1>
			<div class="flex gap-4">
				<div class="flex-1">
					<TextInput
						bind:value={searchValue}
						label="Search"
						autoFocus
					/>
				</div>
				<div class="w-48">
					<DropdownSelect items={allTags} bind:value={tagValue} />
				</div>
			</div>
			{#if bookmarks && bookmarks.length > 0}
				<div
					class="flex flex-col pb-5 pr-[1px] gap-4 h-[calc(100vh-10.625rem)] hide-scrollbar overflow-y-scroll"
				>
					{#if selectedBookmarks.length > 0}
						<div
							class="sticky top-0 py-2 border-b flex justify-between items-center border-white/20 bg-default-bg z-40"
							transition:slide={{ duration: 200 }}
						>
							<span>
								{selectedBookmarks.length} bookmark(s) selected
							</span>
							<div class="flex items-start gap-4">
								<Button
									onclick={onDeselectAllButtonClicked}
									variant="outlined"
									icon="bx:checkbox"
								>
									Deselect All
								</Button>
								<Button
									onclick={onDeleteBookmarksButtonClicked}
									variant="outlined"
									icon="ph:trash"
								>
									Delete Selected Ones
								</Button>
							</div>
						</div>
					{/if}
					{#each bookmarks as bookmark (bookmark.id)}
						<BookmarkCard
							{bookmark}
							{onClickBookmark}
							bind:tagValue
							selected={selectedBookmarks.includes(bookmark.id)}
							onBookmarkSelected={(selected) =>
								onBookmarkSelected(bookmark.id, selected)}
						/>
					{/each}
				</div>
			{:else}
				<div
					class="text-center flex flex-col justify-center items-center text-white/50"
				>
					There is no bookmark added!
				</div>
			{/if}
		</div>
	</div>
</div>

<DeleteBookmarksModal
	bind:showModal={showDeleteBookmarkModal}
	{selectedBookmarks}
	{deleteBookmarks}
/>
